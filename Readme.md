## Realtime Youtube Subs monitor using youtube data API and Websockets 

Inspired by `TutotialEdge` [youtube video](https://www.youtube.com/watch?v=n3BQLHtsrkM)

[Youtube Data API Documentation](https://developers.google.com/youtube/v3/docs/channels/list?hl=en_GB)

### Libraries used 

- `net/http` for http routing
- `gorilla/websocket` for websocket implementation
- `godotenv` to read environment variables from `.env` file, used hide the app secrets and API keys used in the app.

### Run the app 

> App was built on go v1.16 on Apple M1

Clone the repository and copy it to your `GO_PATH` and run the below command

Run the app
```
$ go run main.go
```
> If it asks for downloading the missing packages, follow the command-line instructions and get the packages

> Make sure to create .env file in the root folder and create and set YOUTUBE_KEY and CHANNEL_ID to youtube API to work properly, or replace the those keys directly with the actual code in `/youtube/youtube.go` file

Open any browser and accesss the URL `http://localhost:3030` and you are good to go, wait for 5 seconds to websocket to pull the data from youtube data API.