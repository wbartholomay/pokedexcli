package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)


func makeCachedRequest(client *Client, url string) ([]byte, error){
	data, ok := client.cache.Get(url)
	if ok { return data, nil }

	//only supporting get requests for now
	return makeGetRequest(client, url)
}

func makeGetRequest(client *Client, url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return []byte{}, err
	} 
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("HTTP call failed, Status Code: %v Status: %v", resp.StatusCode, resp.Status)
	}
	

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}