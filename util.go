package unserializd

import "net/http"

func setHeaders(r *http.Request, s string) *http.Request {
	r.Header.Set("Accept", "application/json, text/plain, */*")
	r.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	r.Header.Set("Accept-Language", "en-US,en;q=0.9")
	r.Header.Set("Dnt", "1")
	r.Header.Set("Referer", s)
	r.Header.Set("Sec-Ch-Ua", `"Chromium";v="123", "Not:A-Brand";v="8"`)
	r.Header.Set("Sec-Ch-Ua-Mobile", "?1")
	r.Header.Set("Sec-Ch-Ua-Platform", `"Android"`)
	r.Header.Set("Sec-Fetch-Dest", "empty")
	r.Header.Set("Sec-Fetch-Mode", "cors")
	r.Header.Set("Sec-Fetch-Site", "same-origin")
	r.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Mobile Safari/537.36")
	r.Header.Set("X-Requested-With", "serializd_vercel")

	return r
}
