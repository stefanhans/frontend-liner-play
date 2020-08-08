package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Request struct {
	ClientVersion  string `json:"clientVersion"`
	ClientId       string `json:"clientId"`
	Text           string `json:"text"`
	SourceLanguage string `json:"sourceLanguage"`
	TargetLanguage string `json:"targetLanguage"`
}

type Response struct {
	TaskId         string   `json:"taskId"`
	TranslatedText string   `json:"translatedText"`
	LoadCommands   []string `json:"loadCommands"`
}

var (
	request = Request{
		ClientVersion: "0.0.1",
		ClientId:      "beab10c6-deee-4843-9757-719566214526",
	}
)

func translate(arguments []string) {

	_ = arguments

	request.SourceLanguage = "de"
	request.Text = strings.Join(arguments, " ")
	request.TargetLanguage = "en"

	fmt.Printf("request: %s\n", request)

	requestJson, err := json.MarshalIndent(request, "", " ")
	if err != nil {
		fmt.Printf("failed to marshal response: %v\n", err)
		return
	}
	fmt.Printf("%s\n", requestJson)

	requestJson, err = json.Marshal(request)
	if err != nil {
		fmt.Printf("failed to marshal response: %v\n", err)
		return
	}

	// Send request to service
	res, err := http.Post("https://europe-west1-hybrid-cloud-22365.cloudfunctions.net/Translation",
		"application/json",
		bytes.NewBuffer(requestJson))
	if err != nil {
		fmt.Printf("failed to send request: %v\n", err)
		return
	}
	if res.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(res.Body)
		fmt.Printf("Received no \"200 OK\" from request: %q\n", strings.TrimSuffix(string(b), "\n"))
		return
	}
	fmt.Printf("Received reply from request: %v\n", res.Status)

	// Read response body in JSON
	body, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		fmt.Printf("failed to read response of request: %v\n", err)
		return
	}

	// Unmarshall response data
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("failed to unmarshall response of request (%v): %v\n", res.Proto, err)
		return
	}
	fmt.Printf("response: %v\n", response)

	responseJson, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		fmt.Printf("failed to marshal response: %v\n", err)
		return
	}
	fmt.Printf("%s\n", responseJson)

	speak([]string{response.TranslatedText})
}

//{
//"clientVersion": "0.0.1",
//"clientId": "beab10c6-deee-4843-9757-719566214526",
//"text": "Today is Wednesday",
//"sourceLanguage": "en",
//"targetLanguage": "fr"
//}
//
//{
//"taskId": "4d71c2a3-e6e1-4efd-bb1f-082227cfb0a5",
//"translatedText": "Aujourd'hui nous sommes mercredi",
//"loadCommands": [
//"firebase database:get --pretty --instance migros-showcase --project hybrid-cloud-22365 /translations_v0_0_1/beab10c6-deee-4843-9757-719566214526/4d71c2a3-e6e1-4efd-bb1f-082227cfb0a5",
//"bq query 'SELECT * FROM migros_showcase.translations_v0_0_1 WHERE taskId = \"4d71c2a3-e6e1-4efd-bb1f-082227cfb0a5\"'",
//"gsutil cat gs://hybrid-cloud-22365.appspot.com/0.0.1/beab10c6-deee-4843-9757-719566214526/4d71c2a3-e6e1-4efd-bb1f-082227cfb0a5 | jq",
//"gcloud logging read 'resource.type=cloud_function resource.labels.region=europe-west1 textPayload=4d71c2a3-e6e1-4efd-bb1f-082227cfb0a5'"
//]
//}

/*
curl -s -X POST https://europe-west1-hybrid-cloud-22365.cloudfunctions.net/Translation -d '
{
"clientVersion": "0.0.1",
"clientId": "beab10c6-deee-4843-9757-719566214526",
"text": "Here, we see a schema",  "sourceLanguage": "en",  "targetLanguage": "fr"
}
'
*/
