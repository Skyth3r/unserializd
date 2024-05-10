package unserializd

import (
	"encoding/json"
	"fmt"
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

	body, err := decodeResponse(rsp)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &d); err != nil {
		return nil, err
	}

	return &d, nil
}

func (c *Client) Watching(username string, s SortingOption) (*CurrentlyWatching, error) {

	var w CurrentlyWatching

	url := baseUrl + username + "/currently_watching_page/1?sort_by="

	suffix, found := sortingSuffixes[s]
	if !found {
		return nil, fmt.Errorf("invalid sorting option: %v", s)
	}
	url += suffix

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	rsp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	body, err := decodeResponse(rsp)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &w); err != nil {
		return nil, err
	}

	if w.TotalPages == 1 {
		return &w, nil
	}

	for i := 2; i <= w.TotalPages; i++ {
		tempUrl := baseUrl + username + "/currently_watching_page/" + fmt.Sprint(i) + "?sort_by=" + suffix

		req, err := http.NewRequest("GET", tempUrl, nil)
		if err != nil {
			return nil, err
		}

		rsp, err := c.Do(req)
		if err != nil {
			return nil, err
		}
		defer rsp.Body.Close()

		body, err := decodeResponse(rsp)
		if err != nil {
			return nil, err
		}

		var temp CurrentlyWatching
		if err := json.Unmarshal(body, &temp); err != nil {
			return nil, err
		}

		w.Items = append(w.Items, temp.Items...)
	}

	return &w, nil
}

func (c *Client) Watched(username string, s SortingOption) (*Watched, error) {

	var w Watched

	url := baseUrl + username + "/watchedpage_v2/1?sort_by="

	suffix, found := sortingSuffixes[s]
	if !found {
		return nil, fmt.Errorf("invalid sorting option: %v", s)
	}
	url += suffix

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	rsp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	body, err := decodeResponse(rsp)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &w); err != nil {
		return nil, err
	}

	if w.TotalPages == 1 {
		return &w, nil
	}

	for i := 2; i <= w.TotalPages; i++ {
		tempUrl := baseUrl + username + "/watchedpage_v2/" + fmt.Sprint(i) + "?sort_by=" + suffix

		req, err := http.NewRequest("GET", tempUrl, nil)
		if err != nil {
			return nil, err
		}

		rsp, err := c.Do(req)
		if err != nil {
			return nil, err
		}
		defer rsp.Body.Close()

		body, err := decodeResponse(rsp)
		if err != nil {
			return nil, err
		}

		var temp Watched
		if err := json.Unmarshal(body, &temp); err != nil {
			return nil, err
		}

		w.Items = append(w.Items, temp.Items...)
	}

	return &w, nil
}
