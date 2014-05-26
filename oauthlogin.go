// Command-line app that uses a hard-coded API ID/SECRET
package main

import "time"
import "fmt"
import "os"

import "./oauthworkflow"

func main() {
	var port uint32 = 8327
	APPID, APPSECRET := "", ""

	timeout := 5 * time.Second
	agoflow := oauthworkflow.MakeAGOFlow(APPID, APPSECRET, port)

	success, error := oauthworkflow.FullOAuthHandshake(agoflow, timeout,
		port)

	if success != "" {
		fmt.Print(success)
	} else if error != "" {
		fmt.Fprintln(os.Stderr, error)
	} else {
		fmt.Println(os.Stderr, "Unspecified error")
	}
}
