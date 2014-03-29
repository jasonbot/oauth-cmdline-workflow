package oauthworkflow

type AGOLogin struct {
	Addr string
}

func (self *AGOLogin) FirstURL() string {
	return "http://127.0.0.1:8088"
}

func (self *AGOLogin) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	response := "Hello"
	writer.Write([]byte(response))
	fmt.Println("HI", req.URL.Query().Get("code"))
}
