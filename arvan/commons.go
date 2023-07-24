package arvan

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type APIClient struct {
	initialized	bool
	httpClient	*http.Client
	apiUrl		string
	auth		string
	Service		string
}

func (cc *APIClient) init() {
	if cc.initialized {
		return
	}
	cc.initialized = true
	cc.httpClient = &http.Client{}
	cc.apiUrl = os.Getenv("API_URL")
	cc.auth = os.Getenv("API_KEY")
}

func mapToQueryString(queryParams map[string]string) string {
	queryString := ""
	for key, value := range queryParams {
		queryString += fmt.Sprintf("%s=%s&", key, value)
	}
	return queryString
}

func (cc *APIClient) CurlGet(url string, queryParams map[string]string) []byte {
	cc.init()
	log.Println("Getting:", url)
	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/%s?%s", cc.apiUrl, url, mapToQueryString(queryParams)),
		nil,
	)
	req.Header.Set("authorization", cc.auth)
	req.Header.Set("accept", "application/json, text/plain, */*")
	res, err := cc.httpClient.Do(req)
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

func (cc *APIClient) CurlPost(url string, data interface{}) []byte {
	cc.init()
	jsonData, _ := json.Marshal(data)
	body := strings.NewReader(string(jsonData))
	log.Println("Posting:", url)
	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", cc.apiUrl, url),
		body,
	)
	req.Header.Set("authorization", cc.auth)
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("content-type", "application/json")
	res, err := cc.httpClient.Do(req)
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

func (cc *APIClient) CurlPatch(url string, data interface{}) []byte {
	cc.init()
	jsonData, _ := json.Marshal(data)
	body := strings.NewReader(string(jsonData))
	log.Println("Patching:", url)
	req, _ := http.NewRequest(
		"PATCH",
		fmt.Sprintf("%s/%s", cc.apiUrl, url),
		body,
	)
	req.Header.Set("authorization", cc.auth)
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("content-type", "application/json")
	res, err := cc.httpClient.Do(req)
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