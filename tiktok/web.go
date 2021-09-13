package tiktok

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type TikTokRequest struct {
	body    string
	cookies []*http.Cookie
	err     error
}

/*
This function will take in a TikTok url and output the raw data from it
It also will automaticly pass the verification cookie system that TikTok has
*/
func webGet(url string, cook *http.Cookie) *TikTokRequest {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}

	if cook != nil { // Add cookies to request if supplied
		req.AddCookie(cook)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close() // Make sure to close stream

	// Reads the data from the body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	// Creates a TikTokRequest type and fills in the values
	ttr := TikTokRequest{
		body:    string(body),
		cookies: resp.Cookies(),
		err:     err,
	}

	// TikTok will return only two cookies if it is asking for verification
	// This will then need to request again with the new verification cookies
	// To then reach the real page
	if len(resp.Cookies()) == 2 {
		// Verification is needed so request again with verification cookie
		verifiedReq := webGet(url, resp.Cookies()[0])

		// Build new type now that the verification is complete
		ttr = TikTokRequest{
			body:    verifiedReq.body,
			cookies: verifiedReq.cookies,
			err:     err,
		}
	}

	return &ttr
}

/*
This function will grab the media from a TikTok media url
The cookies where already gotten from webGet and need to be passed into here
to make sure we don't get an "Access Denied" message
*/
func mediaGet(url string, ttReq TikTokRequest) ([]byte, []*http.Cookie, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	// Add all the supplied cookies
	for _, co := range ttReq.cookies {
		req.AddCookie(co)
	}

	// Set some header information to make request more believeable
	req.Header.Set("Host", "v16-web.tiktok.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
	req.Header.Set("Accept", "video/webm,video/ogg,video/*;q=0.9,application/ogg;q=0.7,audio/*;q=0.6,*/*;q=0.5")
	req.Header.Set("Referer", "https://www.tiktok.com/")

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close() // Make sure to close stream

	// Read all the body information
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return body, resp.Cookies(), nil
}

/*
This function will take in a block of html data requested from the website
and find any v16-web.tiktok.com media links embeded in to and return a list
of them. These urls can then be fed into mediaGet to download them
*/
func parce(data string) []string {
	var links []string

	splits := strings.Split(data, "https://")
	for _, value := range splits {
		value, _ = url.QueryUnescape(value)                 // Clean up any encoded characters
		value = strings.ReplaceAll(value, "&amp;", "&")     // Remove html excape characters
		if strings.HasPrefix(value, "v16-web.tiktok.com") { // Makes sure the URL is of the correct type
			links = append(links, "https://"+strings.Split(value, "\"")[0])
		}
	}

	return links
}
