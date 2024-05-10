package unserializd

import "net/http"

type Client struct {
	httpClient *http.Client
}

type SerializdDiary struct {
	Reviews      []Reviews `json:"reviews"`
	TotalPages   int       `json:"totalPages"`
	TotalReviews int       `json:"totalReviews"`
}

type Reviews struct {
	ID               int          `json:"id"`
	ShowID           int          `json:"showId"`
	SeasonID         int          `json:"seasonId"`
	SeasonName       string       `json:"seasonName"`
	DateAdded        string       `json:"dateAdded"`
	Rating           int          `json:"rating"`
	ReviewText       string       `json:"reviewText"`
	Author           string       `json:"author"`
	AuthorImageUrl   string       `json:"authorImageUrl"`
	ContainsSpoiler  bool         `json:"containsSpoilers"`
	BackDate         string       `json:"backdate"`
	ShowName         string       `json:"showName"`
	ShowBannerImage  string       `json:"showBannerImage"`
	ShowSeasons      []ShowSeason `json:"showSeasons"`
	ShowPremiereDate string       `json:"showPremiereDate"`
	IsRewatched      bool         `json:"isRewatched"`
	IsLogged         bool         `json:"isLogged"`
	EpisodeNumber    int          `json:"episodeNumber"`
}

type ShowSeason struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	SeasonNumber int    `json:"seasonNumber"`
	PosterPath   string `json:"posterPath"`
}

type CurrentlyWatching struct {
	Items      []Items `json:"items"`
	TotalItems int     `json:"totalItems"`
	TotalPages int     `json:"totalPages"`
}

type Items struct {
	ShowID      int    `json:"showId"`
	DateAdded   string `json:"dateAdded"`
	ShowName    string `json:"showName"`
	BannerImage string `json:"bannerImage"`
}

type Watched struct {
	Items           []WatchedItems `json:"items"`
	TotalPages      int            `json:"totalPages"`
	NumberOfShows   int            `json:"numberOfShows"`
	NumberOfSeasons int            `json:"numberOfSeasons"`
}

type WatchedItems struct {
	ShowID           int    `json:"showId"`
	SeasonIDs        []int  `json:"seasonIds"`
	DateAdded        string `json:"dateAdded"`
	ShowName         string `json:"showName"`
	BannerImage      string `json:"bannerImage"`
	NumberOfSeasons  int    `json:"numSeasons"`
	NumberOfEpisodes int    `json:"numEpisodes"`
	PremierDate      string `json:"premierDate"`
}
