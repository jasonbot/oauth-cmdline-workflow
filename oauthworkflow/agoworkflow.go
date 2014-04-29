package oauthworkflow

import "fmt"
import "io/ioutil"
import "net/http"
import "net/url"

type AGOLogin struct {
	APPID     string
	APPSECRET string
	Port      uint32
	Success   chan string
	Error     chan string
}

func MakeAGOFlow(APPID, APPSECRET string) OAuthFlow {
	flow := AGOLogin{APPID: APPID, APPSECRET: APPSECRET}
	return flow
}

func (self AGOLogin) InitializeOAuthFlow(port uint32, success chan string,
	error chan string) {
	self.Port = port
	self.Success = success
	self.Error = error
}

func (self AGOLogin) FirstURL() string {
	redirect_uri := url.QueryEscape(fmt.Sprintf("http://127.0.0.1:%v/gotLogin",
		self.Port))
	url := fmt.Sprintf("https://www.arcgis.com/sharing/oauth2/authorize?"+
		"client_id=%v&response_type=code&redirect_uri=%v",
		self.APPID, redirect_uri)

	return url
}

func (self AGOLogin) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	error_string := ""
	if req.URL.Path == "/gotLogin" {
		code := req.URL.Query().Get("code")
		if code != "" {
			resp, post_err := http.PostForm("https://www.arcgis.com/sharing/"+
				"oauth2/token",
				url.Values{
					"client_id":     {self.APPID},
					"client_secret": {self.APPSECRET},
					"grant_type":    {"authorization_code"},
					"code":          {code}})

			if post_err != nil {
				auth_code, newerror := ioutil.ReadAll(resp.Body)

				if newerror != nil {
					error_string = newerror.Error()
				} else {
					headers := writer.Header()
					headers.Set("Content-Type", "text/plain")
					response := "You are logged in. You can close this window."
					writer.Write([]byte(response))
					self.Success <- string(auth_code)

					return
				}
			} else {
				error_string = post_err.Error()
			}
		}

		if error_string == "" {
			error_string = req.URL.Query().Get("error")
		}

		if error_string != "" {
			headers := writer.Header()
			headers.Set("Content-Type", "text/plain")
			response := fmt.Sprintf("Error logging in: %v.", error_string)
			writer.Write([]byte(response))

			self.Error <- error_string
		} else {
			http.Redirect(writer, req, self.FirstURL(), http.StatusSeeOther)
		}
	} else {
		http.Redirect(writer, req, self.FirstURL(), http.StatusSeeOther)
	}
}
