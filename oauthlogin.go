package main

import "time"

import "./oauthworkflow"

func main() {
	oauth_token_channel := make(chan string)
	failure_channel := make(chan string)

	var port uint32 = 8327

	timeout := 5 * time.Second

	agoflow := oauthworkflow.AGOLogin{}
	oauthworkflow.StartWebServer(oauth_token_channel, failure_channel, port,
		agoflow)
	oauthworkflow.OpenLocalhostBrowser(agoflow.FirstURL(), failure_channel)
	oauthworkflow.WaitForToken(oauth_token_channel, failure_channel, timeout)
}
