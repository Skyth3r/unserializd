package unserializd

import "net/http"

type Client struct {
	httpClient *http.Client
}

type SerializdDiary struct {
	Reviews      []DiaryReviews `json:"reviews"`
	TotalPages   int            `json:"totalPages"`
	TotalReviews int            `json:"totalReviews"`
}

type DiaryReviews struct {
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
	Items      []WatchingItem `json:"items"`
	TotalItems int            `json:"totalItems"`
	TotalPages int            `json:"totalPages"`
}

type WatchingItem struct {
	ShowID      int    `json:"showId"`
	DateAdded   string `json:"dateAdded"`
	ShowName    string `json:"showName"`
	BannerImage string `json:"bannerImage"`
}

type Watched struct {
	Items           []WatchedItem `json:"items"`
	TotalPages      int           `json:"totalPages"`
	NumberOfShows   int           `json:"numberOfShows"`
	NumberOfSeasons int           `json:"numberOfSeasons"`
}

type WatchedItem struct {
	ShowID           int    `json:"showId"`
	SeasonIDs        []int  `json:"seasonIds"`
	DateAdded        string `json:"dateAdded"`
	ShowName         string `json:"showName"`
	BannerImage      string `json:"bannerImage"`
	NumberOfSeasons  int    `json:"numSeasons"`
	NumberOfEpisodes int    `json:"numEpisodes"`
	PremierDate      string `json:"premierDate"`
}

type Watchlist struct {
	WatchlistItems  []WatchlistItem `json:"items"`
	TotalPages      int             `json:"totalPages"`
	NumberOfShows   int             `json:"numberOfShows"`
	NumberOfSeasons int             `json:"numberOfSeasons"`
}

type WatchlistItem struct {
	ShowID           int    `json:"showId"`
	SeasonIDs        []int  `json:"seasonIds"`
	DateAdded        string `json:"dateAdded"`
	ShowName         string `json:"showName"`
	BannerImage      string `json:"bannerImage"`
	NumberOfSeasons  int    `json:"numSeasons"`
	NumberOfEpisodes int    `json:"numEpisodes"`
}

type Paused struct {
	PausedItems []PausedItem `json:"items"`
	TotalItems  int          `json:"totalItems"`
	TotalPages  int          `json:"totalPages"`
}

type PausedItem struct {
	ShowID      int    `json:"showId"`
	ShowName    string `json:"showName"`
	DateAdded   string `json:"dateAdded"`
	BannerImage string `json:"bannerImage"`
}

type Dropped struct {
	DroppedItems []DroppedItem `json:"items"`
	TotalItems   int           `json:"totalItems"`
	TotalPages   int           `json:"totalPages"`
}

type DroppedItem struct {
	ShowID      int    `json:"showId"`
	ShowName    string `json:"showName"`
	DateAdded   string `json:"dateAdded"`
	BannerImage string `json:"bannerImage"`
}

type Reviews struct {
	ReviewItems []ReviewItem `json:"items"`
}

type PinnedReviews struct {
	PinnedReviewItems []ReviewItem `json:"reviews"`
}

type ReviewItem struct {
	DateAdded           string       `json:"dateAdded"`
	Rating              int          `json:"rating"`
	Like                bool         `json:"like"`
	ID                  int          `json:"id"`
	ReviewText          string       `json:"reviewText"`
	SeasonID            int          `json:"seasonId"`
	ShowID              int          `json:"showId"`
	Author              string       `json:"author"`
	AuthorImageUrl      string       `json:"authorImageUrl"`
	ContainsSpoiler     bool         `json:"containsSpoiler"`
	BackDate            string       `json:"backdate"`
	ShowName            string       `json:"showName"`
	ShowBannerImage     string       `json:"showBannerImage"`
	ShowSeasons         []ShowSeason `json:"showSeasons"`
	ShowPremiereDate    string       `json:"showPremiereDate"`
	NumberOfComments    int          `json:"numComments"`
	NumberOfLikes       int          `json:"numLikes"`
	DoesUserHaveComment bool         `json:"doesUserHaveComment"`
	DoesUserHaveLike    bool         `json:"doesUserHaveLike"`
	IsRewatch           bool         `json:"isRewatch"`
	IsLog               bool         `json:"isLog"`
	EpisodeNumber       int          `json:"episodeNumber"`
	EpisodeName         string       `json:"episodeName"`
	Tags                []string     `json:"tags"`
	LikeCount           int          `json:"likeCount"`
}
