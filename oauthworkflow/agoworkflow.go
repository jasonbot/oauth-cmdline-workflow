package oauthworkflow

import "fmt"
import "net/http"

type AGOLogin struct {
	Port uint32
}

func (self *AGOLogin) FirstURL() string {
	return fmt.Sprintf("http://127.0.0.1:%v", self.Port)
}

func (self *AGOLogin) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	response := "Hello"
	writer.Write([]byte(response))
	fmt.Println("HI", req.URL.Query().Get("code"))
}
