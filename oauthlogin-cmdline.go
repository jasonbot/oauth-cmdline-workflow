// Command-line AGOL login app that has everything set as options
package main

import "flag"
import "time"
import "fmt"
import "os"

import "./oauthworkflow"

func main() {
	var port, timeout_in_seconds uint64
	var APPID, APPSECRET string

	flag.Uint64Var(&port, "port", 8327, "Port to bind to locally")
	flag.Uint64Var(&timeout_in_seconds, "timeout", 30, "Timeout (in seconds)")
	flag.StringVar(&APPID, "appid", "<APPID>", "App ID")
	flag.StringVar(&APPSECRET, "appsecret", "<APPSECRET>", "App Secret")
	flag.Parse()

	timeout := time.Duration(timeout_in_seconds) * time.Second
	agoflow := oauthworkflow.MakeAGOFlow(APPID, APPSECRET, uint32(port))

	success, error := oauthworkflow.FullOAuthHandshake(agoflow, timeout,
		uint32(port))

	if success != "" {
		fmt.Print(success)
	} else if error != "" {
		fmt.Fprintln(os.Stderr, error)
	} else {
		fmt.Println(os.Stderr, "Unspecified error")
	}
}
