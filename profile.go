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

func (c *Client) Watchlist(username string, s SortingOption) (*Watchlist, error) {

	var w Watchlist

	url := baseUrl + username + "/watchlistpage_v2/1?sort_by="

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
		tempUrl := baseUrl + username + "/watchlistpage_v2/" + fmt.Sprint(i) + "?sort_by=" + suffix

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

		var temp Watchlist
		if err := json.Unmarshal(body, &temp); err != nil {
			return nil, err
		}

		w.WatchlistItems = append(w.WatchlistItems, temp.WatchlistItems...)
	}

	return &w, nil
}

func (c *Client) Paused(username string, s SortingOption) (*Paused, error) {

	var p Paused

	url := baseUrl + username + "/paused_shows_page/1?sort_by="

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

	if err := json.Unmarshal(body, &p); err != nil {
		return nil, err
	}

	if p.TotalPages == 1 {
		return &p, nil
	}

	for i := 2; i <= p.TotalPages; i++ {
		tempUrl := baseUrl + username + "/paused_shows_page/" + fmt.Sprint(i) + "?sort_by=" + suffix

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

		var temp Paused
		if err := json.Unmarshal(body, &temp); err != nil {
			return nil, err
		}

		p.PausedItems = append(p.PausedItems, temp.PausedItems...)
	}

	return &p, nil
}

func (c *Client) Dropped(username string, s SortingOption) (*Dropped, error) {

	var d Dropped

	url := baseUrl + username + "/dropped_shows_page/1?sort_by="

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

	if err := json.Unmarshal(body, &d); err != nil {
		return nil, err
	}

	if d.TotalPages == 1 {
		return &d, nil
	}

	for i := 2; i <= d.TotalPages; i++ {
		tempUrl := baseUrl + username + "/dropped_shows_page/" + fmt.Sprint(i) + "?sort_by=" + suffix

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

		var temp Dropped
		if err := json.Unmarshal(body, &temp); err != nil {
			return nil, err
		}

		d.DroppedItems = append(d.DroppedItems, temp.DroppedItems...)
	}

	return &d, nil
}
