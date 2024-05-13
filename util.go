package unserializd

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
)

type SortingOption int
type ReviewSortingOption int

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

const (
	WhenWatchedDesc ReviewSortingOption = iota
	WhenWatchedAsc
	WhenReviewedDesc
	WhenReviewedAsc
	LowestRatingFirst
	HighestRatingFirst
	MostLikedFirst
	LeastLikedFirst
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

var reviewSortingSuffixes = map[ReviewSortingOption]string{
	WhenWatchedDesc:    "backdate_desc",
	WhenWatchedAsc:     "backdate_asc",
	WhenReviewedDesc:   "date_added_desc",
	WhenReviewedAsc:    "date_added_asc",
	LowestRatingFirst:  "rating_asc",
	HighestRatingFirst: "rating_desc",
	MostLikedFirst:     "like_desc",
	LeastLikedFirst:    "like_asc",
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
