package wrtc

import (
	"encoding/json"
	"fmt"
	"github.com/bic4907/webrtc/common"
	"github.com/bic4907/webrtc/protobuf"
	"github.com/bic4907/webrtc/rpcware"
	"github.com/google/uuid"
	"github.com/pion/rtcp"
	"github.com/pion/webrtc/v3"
	"gopkg.in/djherbis/times.v1"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type Broadcaster struct {
	Pc             *webrtc.PeerConnection
	Ws             *websocket.Conn
	MessageChannel chan []byte

	Uid      uuid.UUID
	Recorder *VideoRecorder

	LastHit int64

	UserId      string
	RoomId      string
	BroadcastId string

	VideoTrack *webrtc.Track
	AudioTrack *webrtc.Track

	BroadcastChannel chan common.BroadcastChunk
	IsBroadcasted    bool
}

func MakeBroadcasterPeerConnection(description webrtc.SessionDescription, broadcaster *Broadcaster) *webrtc.PeerConnection {

	m := webrtc.MediaEngine{}
	m.RegisterCodec(webrtc.NewRTPVP8Codec(webrtc.DefaultPayloadTypeVP8, 90000))
	m.RegisterCodec(webrtc.NewRTPOpusCodec(webrtc.DefaultPayloadTypeOpus, 48000))

	api := webrtc.NewAPI(webrtc.WithMediaEngine(m))

	pc, err := api.NewPeerConnection(WebRTCConfig)
	if err != nil {
		panic(err)
	}

	if _, err = pc.AddTransceiverFromKind(webrtc.RTPCodecTypeAudio); err != nil {
		panic(err)
	} else if _, err = pc.AddTransceiverFromKind(webrtc.RTPCodecTypeVideo); err != nil {
		panic(err)
	}

	recorder := newVideoRecorder()
	broadcaster.Recorder = recorder
	broadcaster.Pc = pc
	recorder.Broadcaster = broadcaster

	var videoCheckerChannel *webrtc.DataChannel = nil

	pc.OnTrack(func(track *webrtc.Track, receiver *webrtc.RTPReceiver) {
		if track.Kind() == webrtc.RTPCodecTypeVideo {

			broadcaster.VideoTrack = track

			go func() {
				ticker := time.NewTicker(time.Second * 1)
				for range ticker.C {
					errSend := pc.WriteRTCP([]rtcp.Packet{&rtcp.PictureLossIndication{MediaSSRC: track.SSRC()}})
					if errSend != nil {
						break
					}
				}
			}()

			go func() {
				ticker := time.NewTicker(time.Second * 1)
				for range ticker.C {
					if broadcaster.LastHit != -1 && makeTimestamp()-broadcaster.LastHit > 3000 {
						log(broadcaster.Uid, fmt.Sprintf("Closed with Time-out"))

						broadcaster.Recorder.Close()
						broadcaster.Pc.Close()
						if broadcaster.Ws != nil {
							broadcaster.Ws.Close()
						}
						return
					}

					if videoCheckerChannel != nil {
						videoCheckerChannel.SendText("video-ok")
					}
				}
			}()

			broadcaster.StartChunkSender()

		} else {
			broadcaster.AudioTrack = track
		}

		for {

			rtp, err := track.ReadRTP()
			if err != nil {
				if err == io.EOF {
					return
				}
				panic(err)
			}

			switch track.Kind() {
			case webrtc.RTPCodecTypeAudio:
				broadcaster.Recorder.PushOpus(rtp)

				chunk := common.BroadcastChunk{BroadcastId: broadcaster.BroadcastId, Chunk: rtp, CodecType: webrtc.RTPCodecTypeAudio}
				broadcaster.BroadcastChannel <- chunk

			case webrtc.RTPCodecTypeVideo:
				broadcaster.Recorder.PushVP8(rtp)

				chunk := common.BroadcastChunk{BroadcastId: broadcaster.BroadcastId, Chunk: rtp, CodecType: webrtc.RTPCodecTypeVideo}
				broadcaster.BroadcastChannel <- chunk
			}
		}
	})

	pc.OnICECandidate(func(c *webrtc.ICECandidate) {
		if c == nil {
			return
		}

		b, err := json.Marshal(c.ToJSON())
		if err != nil {
			panic(err)
		}
		actualSerialized := string(b)

		payload := make(map[string]interface{})
		payload["type"] = "iceCandidate"
		payload["message"] = actualSerialized
		message, _ := json.Marshal(payload)

		defer func() {
			recover()
		}()
		broadcaster.Ws.WriteMessage(1, message)

	})

	pc.OnDataChannel(func(d *webrtc.DataChannel) {
		if d.Label() == "health-check" {
			d.OnMessage(func(msg webrtc.DataChannelMessage) {
				arr := strings.Split(string(msg.Data), "-")
				d.SendText("pong-" + arr[1])
				broadcaster.LastHit = makeTimestamp()
			})
		}

		if d.Label() == "video-check" {
			videoCheckerChannel = d
		}
	})

	err = pc.SetRemoteDescription(description)
	if err != nil {
		panic(err)
	}

	answer, err := pc.CreateAnswer(nil)
	if err != nil {
		panic(err)
	}

	err = pc.SetLocalDescription(answer)
	if err != nil {
		panic(err)
	}

	return pc
}

func (b *Broadcaster) StartChunkSender() {

	dirPath := chunkPath + b.BroadcastId
	var current = 0

	chunkSender := rpcware.GetRpcInstance()

	go func() {
		defer func() {
			recover()
		}()

		for {

			curFileName := dirPath + "/" + strconv.Itoa(current) + ".mp4"
			nextFileName := dirPath + "/" + strconv.Itoa(current+1) + ".mp4"

			if _, err := os.Stat(curFileName); os.IsNotExist(err) {
				time.Sleep(100)
				continue
			}

			if _, err := os.Stat(nextFileName); os.IsNotExist(err) {
				time.Sleep(100)
				continue
			}
			data, err := ioutil.ReadFile(curFileName)
			if err == nil {

				var ctime int64
				tStat, _ := times.Stat(curFileName)
				if tStat.HasBirthTime() == false {
					ctime = time.Now().UnixNano() - int64(1000000000*4.5)
				} else {
					tStat, _ := times.Stat(curFileName)
					ctime = tStat.BirthTime().UnixNano()
				}

				vChunk := &protobuf.VideoChunk{RoomId: b.RoomId, UserId: b.UserId, Chunk: data, CreatedAt: ctime}
				chunkSender.SendChunk <- vChunk
				// If file writing is completed
				err = os.Remove(curFileName)
				if err != nil {
				}

				current += 1
			}
		}
	}()

}
