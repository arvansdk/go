package arvan

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type APIClient struct {
	initialized	bool
	httpClient	*http.Client
	apiUrl		string
	apiKey		string
}

func NewClient() *APIClient {
	client := &APIClient{}
	client.init()
	return client
}

func (client *APIClient) init() {
	if client.initialized{
		return
	}
	godotenv.Load()
	client.initialized = true
	client.httpClient = &http.Client{}
	client.apiUrl = os.Getenv("ARVAN_API_URL")
	client.apiKey = os.Getenv("ARVAN_API_KEY")
}

func MapToQueryString(queryParams map[string]string) string {
	queryString := ""
	for key, value := range queryParams {
		queryString += fmt.Sprintf("%s=%s&", key, value)
	}
	return queryString
}

func (client *APIClient) CurlGet(url string, queryParams map[string]string) []byte {
	client.init()
	log.Println("Getting:", url)
	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/%s?%s", client.apiUrl, url, MapToQueryString(queryParams)),
		nil,
	)
	req.Header.Set("authorization", client.apiKey)
	req.Header.Set("accept", "application/json, text/plain, */*")
	res, err := client.httpClient.Do(req)
	if err != nil {
		log.Println("HTTP request failed!")
		return []byte("")
	}
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Reading response body failed!")
		return []byte("")
	}
	return resBody
}

func (client *APIClient) CurlPost(url string, data interface{}) []byte {
	client.init()
	jsonData, _ := json.Marshal(data)
	body := strings.NewReader(string(jsonData))
	log.Println("Posting:", url)
	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", client.apiUrl, url),
		body,
	)
	req.Header.Set("authorization", client.apiKey)
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("content-type", "application/json")
	res, err := client.httpClient.Do(req)
	if err != nil {
		log.Println("HTTP request failed!")
		return []byte("")
	}
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Reading response body failed!")
		return []byte("")
	}
	log.Printf("%s\n", resBody)
	return resBody
}

func (client *APIClient) CurlPatch(url string, data interface{}) []byte {
	client.init()
	jsonData, _ := json.Marshal(data)
	body := strings.NewReader(string(jsonData))
	log.Println("Patching:", url)
	req, _ := http.NewRequest(
		"PATCH",
		fmt.Sprintf("%s/%s", client.apiUrl, url),
		body,
	)
	req.Header.Set("authorization", client.apiKey)
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("content-type", "application/json")
	res, err := client.httpClient.Do(req)
	if err != nil {
		log.Println("HTTP request failed!")
		return []byte("")
	}
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Reading response body failed!")
		return []byte("")
	}
	log.Printf("%s\n", resBody)
	return resBody
}