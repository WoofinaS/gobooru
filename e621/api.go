package e621

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	postSearchPrefix = "https://e621.net/posts.json?"
)

var (
	client = &http.Client{}
)

func NewClient(userName string) Client {
	return Client("gobooru/1.0 " + userName)
}

func (c Client) SearchPost(f PostFilter) (result *postSearchResult, err error) {
	err = request(postSearchPrefix, string(c), &result)
	return
}

func request(url, agent string, v interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", agent)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, v)
}
