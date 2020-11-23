package rpcware

import (
	"context"
	"github.com/bic4907/webrtc/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"log"
)

type RpcClient struct {
	connection    *grpc.ClientConn
	cci           protobuf.ServiceClient
	SendChunk     chan *protobuf.VideoChunk
	connectedHost string
}

var instance *RpcClient

func GetRpcInstance() *RpcClient {
	if instance == nil {
		instance = new(RpcClient)

		instance.SendChunk = make(chan *protobuf.VideoChunk)

	}
	return instance
}

func (c *RpcClient) Connect(remoteHost string) {
	conn, e := grpc.Dial(remoteHost, grpc.WithInsecure())
	if e != nil {
		log.Println("Error while connecting gRPC")
		return
	} else {
	}
	c.connectedHost = remoteHost
	c.connection = conn
	cci := protobuf.NewServiceClient(conn)
	c.cci = cci

	go c.Listen()
}

func (c *RpcClient) Disconnect() {
	c.connection.Close()
}

func (c *RpcClient) Reconnect() {
	c.Disconnect()
	c.Connect(c.connectedHost)
}

func (c *RpcClient) Listen() {
	log.Println("gRPC Listening...")

	vStream, _ := protobuf.ServiceClient.StreamVideo(c.cci, context.TODO())
	//aStream, _ := protobuf.ServiceClient.StreamAudio(c.cci, context.Background())

	for {
		if c.connection.GetState() != connectivity.Ready {
			go c.Reconnect()
			break
		}
		select {
		case chunk := <-c.SendChunk:

			err := vStream.Send(chunk)
			if err != nil {
				log.Println("gRPC connection not established")
				break
			}
		}
	}
}
