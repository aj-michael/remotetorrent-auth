package main

import (
	"fmt"
	"github.com/pusher/pusher-http-go"
	"io/ioutil"
	"net/http"
	"os"
)

func pusherAuth(res http.ResponseWriter, req *http.Request) {
	params, _ := ioutil.ReadAll(req.Body)
	appid := os.Getenv("RT_PUSHER_APP_ID")
	appkey := os.Getenv("RT_PUSHER_APP_KEY")
	appsecret := os.Getenv("RT_PUSHER_APP_SECRET")
	client := pusher.Client{
		AppId:  appid,
		Key:    appkey,
		Secret: appsecret,
	}
	response, err := client.AuthenticatePrivateChannel(params)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(res, string(response))
}

func main() {
	http.HandleFunc("/auth", pusherAuth)
	http.ListenAndServe(":5000", nil)
}
