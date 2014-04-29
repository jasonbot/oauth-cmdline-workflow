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

func (self AGOLogin) InitializeOAuthFlow(port uint32, success chan string,
	error chan string) {
	self.Port = port
	self.Success = success
	self.Error = error
}

func (self AGOLogin) FirstURL() string {
	redirect_uri := url.QueryEscape(fmt.Sprintf("http://127.0.0.1:%v/gotLogin", self.Port))
	url := fmt.Sprintf("https://www.arcgis.com/sharing/oauth2/authorize?client_id=%v&response_type=code&redirect_uri=%v", self.APPID, redirect_uri)

	return url
}

func (self AGOLogin) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/gotLogin" {
		code := req.URL.Query().Get("code")
		if code != "" {
			// TODO
			/* Access token

			Once the authorization code has been obtained, the app needs to exchange it for an access token.

			The actual request is a POST request to the /token endpoint for the portal, shown here for arcgis.com:

			https://www.arcgis.com/sharing/oauth2/token
			All the parameters (in the following example) must be sent in the request body and not as part of the query component of the URI.

			client_id=APPID&
			client_secret=APPSECRET&
			grant_type=authorization_code&
			code=CODE_OBTAINED_IN_THE_PREVIOUS_STEP */
			resp, err := http.PostForm("https://www.arcgis.com/sharing/oauth2/token",
				url.Values{
					"client_id":     {self.APPID},
					"client_secret": {self.APPSECRET},
					"grant_type":    {"authorization_code"},
					"code":          {code}})

			if err != nil {
				auth_code, _ := ioutil.ReadAll(resp.Body)
				response := "You are now logged in. You can close this window."
				writer.Write([]byte(response))
				self.Success <- string(auth_code)

				return
			}

		}

		error := req.URL.Query().Get("error")

		if error != "" {
			response := fmt.Sprintf("Error logging in: %v.", error)
			writer.Write([]byte(response))

			self.Error <- error
		}

		http.Redirect(writer, req, self.FirstURL(), 303)

	} else {
		http.Redirect(writer, req, self.FirstURL(), 303)
	}
}
