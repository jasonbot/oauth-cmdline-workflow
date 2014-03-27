package main

import "./oauthworkflow"

import "fmt"
import "time"

func main() {
	oauth_token_channel := make(chan string)
	failure_channel := make(chan string)

	addr := fmt.Sprintf("http://127.0.0.1:%v", 8088)
	timeout := 5 * time.Second

	oauthworkflow.StartWebServer(oauth_token_channel, failure_channel, 8088)
	oauthworkflow.OpenBrowser(addr, failure_channel)
	oauthworkflow.WaitForToken(oauth_token_channel, failure_channel, timeout)
}
