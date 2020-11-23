package wrtc

import (
	"bufio"
	"fmt"
	"github.com/at-wat/ebml-go/webm"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/pion/rtp"
	"github.com/pion/rtp/codecs"
	"github.com/pion/webrtc/v3/pkg/media/samplebuilder"
)

var (
	videoPath = "videos/"
	chunkPath = "chunks/"
)

type VideoRecorder struct {
	audioWriter, videoWriter   webm.BlockWriteCloser
	audioBuilder, videoBuilder *samplebuilder.SampleBuilder

	audioChunker, videoChunker webm.BlockWriteCloser

	audioTimestamp, videoTimestamp uint32
	Broadcaster                    *Broadcaster
	path                           string
	name                           string
}

func newVideoRecorder() *VideoRecorder {
	return &VideoRecorder{
		audioBuilder: samplebuilder.New(10, &codecs.OpusPacket{}),
		videoBuilder: samplebuilder.New(10, &codecs.VP8Packet{}),
	}
}

func (s *VideoRecorder) Close() {

	log(s.Broadcaster.Uid, fmt.Sprintf("Recording finished - %s", s.Broadcaster.Recorder.name))
	if s.audioWriter != nil {
		if err := s.audioWriter.Close(); err != nil {
			// panic(err)
		}
	}
	s.audioWriter = nil
	if s.videoWriter != nil {
		if err := s.videoWriter.Close(); err != nil {
			// panic(err)
		}
	}
	s.videoWriter = nil
}
func (s *VideoRecorder) PushOpus(rtpPacket *rtp.Packet) {
	s.audioBuilder.Push(rtpPacket)

	for {
		sample := s.audioBuilder.Pop()
		if sample == nil {
			return
		}
		if s.audioWriter != nil {
			s.audioTimestamp += sample.Samples
			t := s.audioTimestamp / 48
			if _, err := s.audioWriter.Write(true, int64(t), sample.Data); err != nil {
				panic(err)
			}
			if _, err := s.audioChunker.Write(true, int64(t), sample.Data); err != nil {
				panic(err)
			}
		}
	}
}
func (s *VideoRecorder) PushVP8(rtpPacket *rtp.Packet) {
	s.videoBuilder.Push(rtpPacket)

	for {
		sample := s.videoBuilder.Pop()
		if sample == nil {
			return
		}
		// Read VP8 header.
		videoKeyframe := (sample.Data[0]&0x1 == 0)
		if videoKeyframe {
			// Keyframe has frame information.
			raw := uint(sample.Data[6]) | uint(sample.Data[7])<<8 | uint(sample.Data[8])<<16 | uint(sample.Data[9])<<24
			width := int(raw & 0x3FFF)
			height := int((raw >> 16) & 0x3FFF)

			if s.videoWriter == nil || s.audioWriter == nil {
				s.InitChunker(width, height)
				s.InitWriter(width, height)

			}
		}
		if s.videoWriter != nil {
			s.videoTimestamp += sample.Samples
			t := s.videoTimestamp / 90
			if _, err := s.videoWriter.Write(videoKeyframe, int64(t), sample.Data); err != nil {
				panic(err)
			}
			if _, err := s.videoChunker.Write(videoKeyframe, int64(t), sample.Data); err != nil {
				panic(err)
			}
		}
	}
}
func (s *VideoRecorder) InitWriter(width, height int) {

	uid := s.Broadcaster.Uid.String()
	now := time.Now().Format("2006-01-02_15-04-05")
	filename := uid + ".mp4"
	filepath := videoPath + now + "_" + filename
	s.path = filepath
	s.name = filename

	ffmpeg := exec.Command("ffmpeg", "-re", "-i", "pipe:0", "-c:v", "libx264", "-loglevel", "panic", filepath)

	ffmpegIn, _ := ffmpeg.StdinPipe()
	ffmpegErr, _ := ffmpeg.StderrPipe()

	if err := ffmpeg.Start(); err != nil {
		panic(err)
	}

	go func() {
		scanner := bufio.NewScanner(ffmpegErr)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	ws, err := GetVideoBlockWriter(ffmpegIn, width, height)
	if err != nil {
		panic(err)
	}

	log(s.Broadcaster.Uid, fmt.Sprintf("Record starting - video width=%d, height=%d", width, height))

	s.audioWriter = ws[0]
	s.videoWriter = ws[1]
}

func (s *VideoRecorder) InitChunker(width, height int) {

	uid := s.Broadcaster.Uid.String()
	now := time.Now().Format("2006-01-02_15-04-05")
	filename := uid + ".mp4"
	filepath := videoPath + now + "_" + filename
	s.path = filepath
	s.name = filename

	dirPath := chunkPath + s.Broadcaster.BroadcastId

	// Delete all videos in directory
	if _, err := os.Stat(chunkPath); os.IsNotExist(err) {
		os.Mkdir(chunkPath, os.ModePerm)
	}

	os.RemoveAll(dirPath)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.Mkdir(dirPath, os.ModePerm)
	}

	ffmpeg := exec.Command("ffmpeg", "-re", "-i", "pipe:0", "-loglevel", "panic", "-c:v", "libx264", "-map", "0", "-segment_time", "1", "-f", "segment", "-reset_timestamps", "1", "-vf", "fps=60", dirPath+"/%d.mp4") //nolint

	ffmpegIn, _ := ffmpeg.StdinPipe()
	ffmpegErr, _ := ffmpeg.StderrPipe()

	if err := ffmpeg.Start(); err != nil {
		panic(err)
	}

	go func() {
		scanner := bufio.NewScanner(ffmpegErr)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	ws, err := GetVideoBlockWriter(ffmpegIn, width, height)
	if err != nil {
		panic(err)
	}

	log(s.Broadcaster.Uid, fmt.Sprintf("Chunker starting - video width=%d, height=%d", width, height))

	s.audioChunker = ws[0]
	s.videoChunker = ws[1]
}

func GetVideoBlockWriter(ffmpegIn io.WriteCloser, width, height int) ([]webm.BlockWriteCloser, error) {
	ws, err := webm.NewSimpleBlockWriter(ffmpegIn,
		[]webm.TrackEntry{
			{
				Name:            "Audio",
				TrackNumber:     1,
				TrackUID:        12345,
				CodecID:         "A_OPUS",
				TrackType:       2,
				DefaultDuration: 20000000,
				Audio: &webm.Audio{
					SamplingFrequency: 48000.0,
					Channels:          2,
				},
			}, {
				Name:            "Video",
				TrackNumber:     2,
				TrackUID:        67890,
				CodecID:         "V_VP8",
				TrackType:       1,
				DefaultDuration: 33333333,
				Video: &webm.Video{
					PixelWidth:  uint64(width),
					PixelHeight: uint64(height),
				},
			},
		})
	return ws, err
}
