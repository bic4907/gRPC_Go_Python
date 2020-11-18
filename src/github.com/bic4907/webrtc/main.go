package main

import (
	"context"
	"fmt"
	"github.com/bic4907/webrtc/protobuf"
	"github.com/bic4907/webrtc/web"
	"google.golang.org/grpc"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {

	go web.StartWebService()

	conn, e := grpc.Dial("localhost:10002", grpc.WithInsecure())
	if e != nil {
		log.Println("Error while connecting grpc")
		return
	}
	defer conn.Close()

	c := protobuf.NewServiceClient(conn)
	go func() {
		for {
			result, err := c.SendMessage(context.Background(), &protobuf.ReqMessage{Content: "Hi"})

			if err != nil {
				log.Println(err)
			} else {
				fmt.Println(result)
			}

			time.Sleep(1 * time.Second)
		}
	}()

	closed := make(chan os.Signal, 1)
	signal.Notify(closed, os.Interrupt)
	<-closed

	// if err := peerConnection.Close(); err != nil {
	//	panic(err)
	//}
	// saver.Close()
}
