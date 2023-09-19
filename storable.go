package storage

import "net/url"

// use url.JoinPath to implement Storable URL
var BaseURL *url.URL

type Storable interface {
	// Url base of of BseURL
	ToKey() *url.URL
}

func init() {
	var err error
	BaseURL, err = url.Parse("/")
	if err != nil {
		panic(err)
	}
}
