package main

import "net/http"

type OAuthFlow interface {
	InitializeOAuthFlow(port uint32, success chan string, error chan string)
	FirstURL() string
	http.Handler
}
