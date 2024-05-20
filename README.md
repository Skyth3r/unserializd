![unserializd header image](/header.png)

# unserializd
An unofficial Go package for accessing the [Serializd](https://www.serializd.com/) API

**Status:** Work in Progress

## Getting Started

### Installing
Use `go get` to retrieve the package to add it to your project's Go module dependencies.

	go get github.com/skyth3r/unserializd

To update the package use `go get -u` to retrieve the latest version of the package.

	go get -u github.com/skyth3r/unserializd

## Docs

🔜

## Quick Examples

```Go
// Create new client
client := unserializd.NewClient()

// Get user's diary
diary, err := client.Diary("USERNAME")
if err != nil {
    fmt.Println(err)
}
fmt.Println(diary)

// Get user's watched TV Shows
watched, err := client.Watched("USERNAME", 0)
if err != nil {
    fmt.Println(err)
}
fmt.Println(watched)

// Get TV shows on user's watchlist
watchlist, err := client.Watchlist("USERNAME", 0)
if err != nil {
    fmt.Println(err)
}
fmt.Println(watchlist.NumberOfShows)
for _, item := range watchlist.WatchlistItems {
    fmt.Println(item.ShowName)
}
```