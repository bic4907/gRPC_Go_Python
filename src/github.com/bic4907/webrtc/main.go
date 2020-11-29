package main

import (
	"github.com/bic4907/webrtc/archive"
	"github.com/bic4907/webrtc/rpcware"
	"github.com/bic4907/webrtc/web"
	"os"
	"os/signal"
)

func main() {

	go web.StartWebService()

	rpcInstance := rpcware.GetRpcInstance()
	rpcInstance.Connect(RpcRemoteHost)

	if AccessKey != "" && AccessSecret != "" && Bucket != "" {
		fileUploaderInstance := archive.GetBucketInstance()
		fileUploaderInstance.Connect(Bucket, AccessKey, AccessSecret)

		fileUploaderInstance.CallbackUrl = Video_Register_API
	}

	closed := make(chan os.Signal, 1)
	signal.Notify(closed, os.Interrupt)
	<-closed

}
