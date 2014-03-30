package oauthworkflow

import "fmt"
import "net/http"
import "os"
import "os/exec"
import "runtime"
import "strings"
import "time"

type OAuthFlow interface {
	InitializeOAuthFlow(port uint32, success chan string, error chan string)
	FirstURL() string
	http.Handler
}

func _webServer(token_chan, error_chan chan string,
	port uint32, flow OAuthFlow) {
	addr := fmt.Sprintf("127.0.0.1:%v", port)

	server := &http.Server{
		Addr:           addr,
		Handler:        flow,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		error_chan <- err.Error()
	}
}

func StartWebServer(token_chan, error_chan chan string, port uint32,
	flow OAuthFlow) {
	go _webServer(token_chan, error_chan, port, flow)
}

func OpenBrowser(url string, error_channel chan string) {
	var commandline, args string

	if runtime.GOOS == "windows" {
		// Windows
		commandline = "cmd.exe"
		args = fmt.Sprintf("/c start %v", strings.Replace(url, "&", "^&", -1))
	} else if runtime.GOOS == "darwin" {
		// OSX
		commandline = "open"
		args = strings.Replace(url, "&", "\\&", -1)
	} else {
		// Default: assume Linuxlike with a Freedesktop-compliant env running
		commandline = "xdg-open"
		args = strings.Replace(url, "&", "\\&", -1)
	}

	_, err := exec.Command(commandline, args).Output()

	if err != nil {
		error_channel <- fmt.Sprintf("Error opening browser: %v", err.Error())
	}
}

func OpenLocalhostBrowser(port uint32, error_chan chan string) {
	addr := fmt.Sprintf("http://127.0.0.1:%v", port)

	OpenBrowser(addr, error_chan)
}

func WaitForToken(token_chan, error_chan chan string, timeout time.Duration) {
	select {
	case v := <-token_chan:
		// Web server successfully got a token response
		fmt.Print(v)
	case err := <-error_chan:
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Error: %v", err))
	case <-time.After(timeout):
		// Timed out (default)
		fmt.Fprintln(os.Stderr, "Error: OAuth handshake timed out")
	}
}
