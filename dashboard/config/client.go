package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Headers struct {
	Key   string
	Value string
}

func PostHttp(payload interface{}, path string, headers []Headers) (*http.Response, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}
	var url = fmt.Sprintf("%s:%s", os.Getenv("APIGW_HOST"), os.Getenv("APIGW_PORT"))
	url = url + "/" + path
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	for _, v := range headers {
		req.Header.Add(v.Key, v.Value)
	}
	client := &http.Client{}
	return client.Do(req)
}

func PatchHttp(payload interface{}, path string, headers []Headers) (*http.Response, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}
	var url = fmt.Sprintf("%s:%s", os.Getenv("APIGW_HOST"), os.Getenv("APIGW_PORT"))
	url = url + "/" + path
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	for _, v := range headers {
		req.Header.Add(v.Key, v.Value)
	}
	client := &http.Client{}
	return client.Do(req)
}

func GetHttp(path string, headers []Headers) (*http.Response, error) {
	var url = fmt.Sprintf("%s:%s", os.Getenv("APIGW_HOST"), os.Getenv("APIGW_PORT"))
	url = url + "/" + path
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	for _, v := range headers {
		req.Header.Add(v.Key, v.Value)
	}
	client := &http.Client{}
	return client.Do(req)
}

func DelHttp(path string, headers []Headers) (*http.Response, error) {
	var url = fmt.Sprintf("%s:%s", os.Getenv("APIGW_HOST"), os.Getenv("APIGW_PORT"))
	url = url + "/" + path
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	for _, v := range headers {
		req.Header.Add(v.Key, v.Value)
	}
	client := &http.Client{}
	return client.Do(req)
}
