package unserializd

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Diary(username string) (*SerializdDiary, error) {

	var d SerializdDiary

	url := baseUrl + username + "/diary"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	rsp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %v", rsp.StatusCode)
	}

	var reader io.Reader
	if rsp.Header.Get("Content-Encoding") == "gzip" {
		gz, err := gzip.NewReader(rsp.Body)
		if err != nil {
			return nil, err
		}
		defer gz.Close()
		reader = gz
	} else {
		reader = rsp.Body
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &d); err != nil {
		return nil, err
	}

	return &d, nil
}
