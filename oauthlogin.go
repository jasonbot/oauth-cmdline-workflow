package main

import "time"
import "fmt"
import "os"

import "./oauthworkflow"

func main() {
	var port uint32 = 8327

	timeout := 5 * time.Second
	agoflow := oauthworkflow.MakeAGOFlow("", "")

	success, error := oauthworkflow.FullOAuthHandshake(agoflow, timeout,
		port)

	if success != "" {
		fmt.Print(success)
	} else {
		fmt.Fprintln(os.Stderr, error)
	}
}
