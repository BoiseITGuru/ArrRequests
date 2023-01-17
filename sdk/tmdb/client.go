package tmdb

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL  *url.URL
	Trending *TrendingService

	httpClient *http.Client
}

func NewClient(httpClient *http.Client, apiKey string) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	//Create Base URL
	baseURL, err := url.Parse("https://api.themoviedb.org/3")
	if err != nil {
		return nil, err
	}

	client := &Client{
		BaseURL:    baseURL,
		httpClient: httpClient,
	}

	client.Trending = &TrendingService{client: client}

	return client, nil
}

func (c *Client) NewRequest(method string, path *url.URL, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, path.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
