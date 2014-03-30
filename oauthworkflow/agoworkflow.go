package oauthworkflow

import "fmt"
import "net/http"

type AGOLogin struct {
	port    uint32
	success chan string
	error   chan string
}

func (self AGOLogin) InitializeOAuthFlow(port uint32, success chan string,
	error chan string) {
	self.port = port
	self.success = success
	self.error = error
}

func (self AGOLogin) FirstURL() string {
	return fmt.Sprintf("http://127.0.0.1:%v", self.port)
}

func (self AGOLogin) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {

	} else if req.URL.Path == "/gotLogin" {
		code := req.URL.Query().Get("code")
		if code != "" {
			response := "You are now logged in."
			writer.Write([]byte(response))

			self.success <- code
		}

		error := req.URL.Query().Get("error")

		if error != "" {
			response := fmt.Sprintf("Error logging in: %v.", error)
			writer.Write([]byte(response))

			self.error <- error
		}

		http.Redirect(writer, req, "/", 303)

	} else {
		http.Redirect(writer, req, "/", 303)
	}
}
