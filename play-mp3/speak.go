package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"cloud.google.com/go/texttospeech/apiv1"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

type SpeakInput struct {
	Text string `json:"text"`
}

type SpeakVoice struct {
	LanguageCode string `json:"languageCode"`
	Name         string `json:"name"`
	SsmlGender   string `json:"ssmlGender"`
}

type SpeakAudioConfig struct {
	AudioEncoding string `json:"audioEncoding"`
}

type SpeakRequest struct {
	Input       SpeakInput       `json:"input"`
	Voice       SpeakVoice       `json:"voice"`
	AudioConfig SpeakAudioConfig `json:"audioConfig"`
}

var (
	// SSML Request Example
	speakRequest = SpeakRequest{
		Input: SpeakInput{
			Text: "Hello, my friend",
		},
		Voice: SpeakVoice{
			LanguageCode: "en-gb",
			Name:         "en-GB-Standard-D",
			SsmlGender:   "MALE",
		},
		AudioConfig: SpeakAudioConfig{
			AudioEncoding: "MP3",
		},
	}
)

func speak(arguments []string) {

	_ = arguments

	// Instantiates a client.
	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Perform the text-to-speech request on the text input with the selected
	// voice parameters and audio file type.
	req := texttospeechpb.SynthesizeSpeechRequest{
		// Set the text input to be synthesized.
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: strings.Join(arguments, " ")},
		},
		// Build the voice request, select the language code ("en-US") and the SSML
		// voice gender ("neutral").
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "en-US",
			SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
		},
		// Select the type of audio file you want returned.
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}

	// Append EOF to get back to the prompt
	resp.AudioContent = append(resp.AudioContent, byte(0))

	// The resp's AudioContent is binary.
	filename := "output.mp3"
	err = ioutil.WriteFile(filename, resp.AudioContent, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Audio content written to file: %v\n", filename)

	fmt.Printf("speakRequest: %s\n", speakRequest)

	play([]string{"output.mp3"})

	//requestJson, err := json.MarshalIndent(speakRequest, "", " ")
	//if err != nil {
	//	fmt.Printf("failed to marshal response: %v\n", err)
	//	return
	//}
	//fmt.Printf("%s\n", requestJson)
	//
	//requestJson, err = json.Marshal(speakRequest)
	//if err != nil {
	//	fmt.Printf("failed to marshal response: %v\n", err)
	//	return
	//}
	//
	//
	//
	//// Send request to service
	//res, err := http.Post("https://texttospeech.googleapis.com/v1/text:synthesize",
	//	"application/json",
	//	bytes.NewBuffer(requestJson))
	//if err != nil {
	//	fmt.Printf("failed to send request: %v\n", err)
	//	return
	//}
	//if res.StatusCode != http.StatusOK {
	//	b, _ := ioutil.ReadAll(res.Body)
	//	fmt.Printf("Received no \"200 OK\" from request: %q\n", strings.TrimSuffix(string(b), "\n"))
	//	return
	//}
	//fmt.Printf("Received reply from request: %v\n", res.Status)

}

/*

 */

/*
curl -s -X POST -H "Authorization: Bearer "$(gcloud auth application-default print-access-token) \
-H "Content-Type: application/json; charset=utf-8" \
-d @1.json \
https://texttospeech.googleapis.com/v1/text:synthesize | jq .audioContent | tr -d '"' > 1.txt && \
base64 1.txt --decode > 1.mp3 && \
echo EOF >> 1.mp3 && echo OK
*/
