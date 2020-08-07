package main

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

}

/*
{
  "input":{
    "text":"Android is a mobile operating system developed by Google, based on the Linux kernel and designed primarily for touchscreen mobile devices such as smartphones and tablets."
  },
  "voice":{
    "languageCode":"en-gb",
    "name":"en-GB-Standard-D",
    "ssmlGender":"MALE"
  },
  "audioConfig":{
    "audioEncoding":"MP3"
  }
}
*/

/*

 */
