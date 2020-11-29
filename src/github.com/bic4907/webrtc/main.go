package main

import (
	"encoding/json"
	"github.com/bic4907/webrtc/archive"
	"github.com/bic4907/webrtc/rpcware"
	"github.com/bic4907/webrtc/web"
	"os"
	"os/signal"
)

var (
	AccessKey    = "AKIAJ6HEDAAEYCRTIVMQ"
	AccessSecret = "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"
	Bucket       = "storage.aisupervisor"

	RpcRemoteHost = "test.inchang.dev:10002"

	Video_Register_API = "https://test.inchang.dev:9000/room/video"
)

type Configuration struct {
	AccessKey          string
	AccessSecret       string
	Bucket             string
	RpcRemoteHost      string
	Video_Register_API string
}

func main() {

	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Configuration{}
	err := decoder.Decode(&config)
	if err != nil {
		panic("Error while reading config file")
	}
	file.Close()

	go web.StartWebService()

	rpcInstance := rpcware.GetRpcInstance()
	rpcInstance.Connect(RpcRemoteHost)

	if config.AccessKey != "" && config.AccessSecret != "" && config.Bucket != "" {
		fileUploaderInstance := archive.GetBucketInstance()
		fileUploaderInstance.Connect(config.Bucket, config.AccessKey, config.AccessSecret)

		fileUploaderInstance.CallbackUrl = config.Video_Register_API
	}

	closed := make(chan os.Signal, 1)
	signal.Notify(closed, os.Interrupt)
	<-closed

}
