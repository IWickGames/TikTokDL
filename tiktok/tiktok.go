package tiktok

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func GetVideos(url string) ([]string, *TikTokRequest) {
	fmt.Println("[TikTok] Downloading webpage")
	request := webGet(url, nil)
	if request.err != nil {
		fmt.Println("ERROR: Failed to fetch TicTok url")
		fmt.Println(request.err.Error())
		os.Exit(1)
	}

	fmt.Println("[TikTok] Parcing webpage")
	return parce(request.body), request
}

func DownloadVideo(url string, ttr TikTokRequest) io.Reader {
	fmt.Println("[Download] Downloading video")
	fmt.Println("[Download] Using: " + url)
	media, _, err := mediaGet(url, ttr)
	if err != nil {
		fmt.Println("ERROR: Failed to download video from TikTok CDN")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return bytes.NewReader(media)
}
