package main

import (
	"github.com/bic4907/webrtc/rpcware"
	"github.com/bic4907/webrtc/web"
	"os"
	"os/signal"
)

func main() {

	go web.StartWebService()

	rpcInstance := rpcware.GetRpcInstance()
	rpcInstance.Connect("localhost:10002")
	//go rpcInstance.Listen()

	closed := make(chan os.Signal, 1)
	signal.Notify(closed, os.Interrupt)
	<-closed

	//rpcInstance := rpcware.GetRpcInstance()
	//rpcInstance.Connect("127.0.0.1:10002")
	//go rpcInstance.Listen()
	// if err := peerConnection.Close(); err != nil {
	//	panic(err)
	//}
	// saver.Close()
}
