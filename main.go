package main

import (
	"TikTok-DL/tiktok"
	"fmt"
	"io"
	"os"
)

func main() {
	// Validate that all arguments are given
	if len(os.Args) != 3 {
		fmt.Println("ERROR: Invalid arguments provided")
		fmt.Println("  Usage: tiktokdl <tiktokurl> <output (no file extenchion)>")
		os.Exit(1)
	}

	// Get the list of videos on the TikTok url
	links, ttrequest := tiktok.GetVideos(os.Args[1])
	if len(links) == 0 {
		fmt.Println("ERROR: Found no media urls in the provided link")
		os.Exit(1)
	}

	// Download the raw video data
	rawData := tiktok.DownloadVideo(links[0], *ttrequest)

	// Create the output file
	file, err := os.Create(os.Args[2] + ".mp4")
	if err != nil {
		fmt.Println("ERROR: Failed to generate output file (permission?)")
		os.Exit(1)
	}
	defer file.Close() // Make sure to close the stream

	// Copy the bytes into the output file
	_, err = io.Copy(file, rawData)
	if err != nil {
		fmt.Println("ERROR: Failed to write to output file (permission?)")
		os.Exit(1)
	}

	fmt.Println("Saved output media as", os.Args[2]+".mp4")
}
