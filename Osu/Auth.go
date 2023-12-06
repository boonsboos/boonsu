package osu

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	util "boonsboos.nl/boonsu/Util"
)

var osuAuthToken string

type authResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func Authorize() {

	form := url.Values{
		"client_id":     {util.Options.OsuClientID},
		"client_secret": {util.Options.OsuToken},
		"grant_type":    {"client_credentials"},
		"scope":         {"public"},
	}

	req, err := http.NewRequest("POST",
		"https://osu.ppy.sh/oauth/token",
		strings.NewReader(form.Encode()),
	)

	if err != nil {
		log.Println("auth request bad!")
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("failed to authorize with osu api: " + err.Error())
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("something went wrong while reading auth response:", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("failed to authorize with osu api: " + string(bytes))
	}

	var a authResponse
	json.Unmarshal(bytes, &a)

	osuAuthToken = a.AccessToken
	refresher = a.ExpiresIn
}

var refresher int = 0

func AutoReAuth() {
	for refresher > 10 {
		time.Sleep(1 * time.Second)
		refresher--
	}

	go Authorize()
}
