package rpcware

import (
	"context"
	"github.com/bic4907/webrtc/protobuf"
	"google.golang.org/grpc"
	"log"
)

type RpcClient struct {
	connection *grpc.ClientConn
	cci        protobuf.ServiceClient
	SendChunk  chan []byte
}

var instance *RpcClient

func GetRpcInstance() *RpcClient {
	if instance == nil {
		instance = new(RpcClient)

		instance.SendChunk = make(chan []byte)

	}
	return instance
}

func (c *RpcClient) Connect(remoteHost string) {
	conn, e := grpc.Dial(remoteHost, grpc.WithInsecure())
	if e != nil {
		log.Println("Error while connecting gRPC")
		return
	} else {
		log.Println("Successfully connected to gRPC")
	}
	c.connection = conn
	cci := protobuf.NewServiceClient(conn)
	c.cci = cci
}

func (c *RpcClient) Disconnect() {
	c.connection.Close()
}

func (c *RpcClient) Listen() {
	log.Println("gRPC Listening...")

	vStream, _ := protobuf.ServiceClient.StreamVideo(c.cci, context.TODO())
	//aStream, _ := protobuf.ServiceClient.StreamAudio(c.cci, context.Background())

	for {
		select {
		case rtp := <-c.SendChunk:
			//fmt.Println("Stream")

			vChunk := &protobuf.VideoChunk{RoomId: "1", UserId: "1", Rtp: rtp}
			vStream.Send(vChunk)

			//fmt.Println("Send Packet")
			//vMessage := &protobuf.ReqMessage{Content: "HI"}
			//_, _ = c.cci.SendMessage(context.Background(), vMessage)

			//log.Println("Chunk Received")
			//log.Println(chunk)
		}
	}
}
