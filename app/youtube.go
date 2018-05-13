// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"golang.org/x/net/context"
	youtube "google.golang.org/api/youtube/v3"
	"log"
	"net/http"
	"os"
)

func uploadVideo(filename string) *youtube.Video {
	ctx := context.Background()

	client, err := buildOAuthHTTPClient(ctx, []string{youtube.YoutubeUploadScope, youtube.YoutubeScope})
	if nil != err {
		log.Fatalf("Error authenticating: %v", err)
	}
	return youtubeMain(client, filename)
}

// youtubeMain is an example that demonstrates calling the YouTube API.
// It is similar to the sample found on the Google Developers website:
// https://developers.google.com/youtube/v3/docs/videos/insert
// but has been modified slightly to fit into the examples framework.
//
// Example usage:
//   go build -o go-api-demo
//   go-api-demo -clientid="my-clientid" -secret="my-secret" youtube filename
func youtubeMain(client *http.Client, filename string) *youtube.Video {

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Unable to create YouTube service: %v", err)
	}

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       "Test Title",
			Description: "Test Description", // can not use non-alpha-numeric characters
			CategoryId:  "22",
		},
		Status: &youtube.VideoStatus{PrivacyStatus: "unlisted"},
	}

	// The API returns a 400 Bad Request response if tags is an empty string.
	upload.Snippet.Tags = []string{"test", "upload", "api"}

	call := service.Videos.Insert("snippet,status", upload)

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("Error opening %v: %v", filename, err)
	}

	response, err := call.Media(file).Do()
	if err != nil {
		log.Fatalf("Error making YouTube API call: %v", err)
	}
	fmt.Printf("Upload successful! Video ID: %v\n", response.Id)

	return response
}

func videoInfo(id string) *youtube.Video {

	ctx := context.Background()

	client, err := buildOAuthHTTPClient(ctx, []string{youtube.YoutubeUploadScope, youtube.YoutubeScope})
	if nil != err {
		log.Fatalf("Error authenticating: %v", err)
	}
	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Unable to create YouTube service: %v", err)
	}
	responseList, err := service.Videos.List("contentDetails, fileDetails, player").Id(id).Do()
	if err != nil {
		log.Fatalf("Error making YouTube API call: %v", err)
	}

	return responseList.Items[0]
}
