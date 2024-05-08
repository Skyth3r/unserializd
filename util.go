package unserializd

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
