## Playground For Translation Text Via Google Cloud Functions

We use JSON as a wire format to get some text translated between different languages.

<br>

Request, e.g.

```json
{
 "clientVersion": "0.0.1",
 "clientId": "beab10c6-deee-4843-9757-719566214526",
 "text": "This is a test",
 "sourceLanguage": "en",
 "targetLanguage": "fr"
}
```
<br>

Response, e.g.

```json
{
 "taskId": "dbeba1c5-0f2d-487d-9131-67b5d4020830",
 "translatedText": "C'est un test",
 "loadCommands": [
  "firebase database:get --pretty --instance migros-showcase --project hybrid-cloud-22365 /translations_v0_0_1/beab10c6-deee-4843-9757-719566214526/dbeba1c5-0f2d-487d-9131-67b5d4020830",
  "bq query 'SELECT * FROM migros_showcase.translations_v0_0_1 WHERE taskId = \"dbeba1c5-0f2d-487d-9131-67b5d4020830\"'",
  "gsutil cat gs://hybrid-cloud-22365.appspot.com/0.0.1/beab10c6-deee-4843-9757-719566214526/dbeba1c5-0f2d-487d-9131-67b5d4020830 | jq",
  "gcloud logging read 'resource.type=cloud_function resource.labels.region=europe-west1 textPayload=dbeba1c5-0f2d-487d-9131-67b5d4020830'"
 ]
}
```
<br>

HTTPS POST Call

`https://europe-west1-hybrid-cloud-22365.cloudfunctions.net/Translation`



