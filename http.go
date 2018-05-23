package gotool

import (
	"io/ioutil"
	"net/http"
	"strings"
	"net/url"
)

func HttpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func HttpPost(url string, args string) ([]byte, error) {
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(args))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func HttpPostForm(url string, args url.Values) ([]byte, error) {
	resp, err := http.PostForm(url, args)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
