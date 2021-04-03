package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Response struct {
	Kind  string `json:"kind"`
	Items []Item `json:"items"`
}

type Item struct {
	Kind  string `json:"kind"`
	Id    string `json:"id"`
	Stats Stats  `json:"statistics"`
}

type Stats struct {
	Views       string `json:"viewCount"`
	Subscribers string `json:"subscriberCount"`
	Videos      string `json:"videoCount"`
}

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}

// API Documentation, https://developers.google.com/youtube/v3/docs/channels/list?hl=en_GB
func GetSubscribers() (Item, error) {
	req, err := http.NewRequest("GET", "https://youtube.googleapis.com/youtube/v3/channels", nil)

	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}

	// here we define the query parameters and their respective values
	q := req.URL.Query()
	q.Add("key", getEnvVariable("YOUTUBE_KEY"))
	q.Add("id", getEnvVariable("CHANNEL_ID"))
	q.Add("part", "statistics")
	req.URL.RawQuery = q.Encode()

	// make the request to the URL that we have just constructed
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}

	defer resp.Body.Close()

	// fmt.Println("Response:", resp.Status)
	// we then read in all of the body of the
	// JSON response
	body, _ := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}

	var response Response
	// and finally unmarshal it into an Response struct
	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}

	return response.Items[0], nil
}
