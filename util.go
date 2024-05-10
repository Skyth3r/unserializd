package unserializd

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
)

type SortingOption int

const (
	DateAddedDesc SortingOption = iota
	DateAddedAsc
	NameAsc
	NameDesc
	ReleaseDateDesc
	ReleaseDateAsc
	LastAiredDesc
	LastAiredAsc
)

var sortingSuffixes = map[SortingOption]string{
	DateAddedDesc:   "date_added_desc",
	DateAddedAsc:    "date_added_asc",
	NameAsc:         "name_asc",
	NameDesc:        "name_desc",
	ReleaseDateDesc: "release_date_desc",
	ReleaseDateAsc:  "release_date_asc",
	LastAiredDesc:   "last_aired_desc",
	LastAiredAsc:    "last_aired_asc",
}

func decodeResponse(rsp *http.Response) ([]byte, error) {
	var reader io.Reader

	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %v", rsp.StatusCode)
	}

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

	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return b, nil

}
