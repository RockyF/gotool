package gotool

import (
	"io/ioutil"
	"net/http"
	"strings"
	"net/url"
	"encoding/json"
)

var UnSerialize func(bytes []byte, result interface{}) (interface{}, error)

func HttpGet(url string, options ...interface{}) (interface{}, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	return readAndUnSerialize(resp, options)
}

func HttpPost(url string, args string, options ...interface{}) (interface{}, error) {
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(args))
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	return readAndUnSerialize(resp, options)
}

func HttpPostForm(url string, args url.Values, options ...interface{}) (interface{}, error) {
	resp, err := http.PostForm(url, args)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	return readAndUnSerialize(resp, options)
}

func readAndUnSerialize(resp *http.Response, options []interface{}) (interface{}, error) {
	if bytes, err := ioutil.ReadAll(resp.Body); err == nil {
		if UnSerialize == nil {
			return bytes, nil
		} else {
			target := options[0]
			return UnSerialize(bytes, target)
		}
	} else {
		return nil, err
	}
}

func HttpJsonSerializefunc(bytes []byte, result interface{}) (interface{}, error) {
	var data interface{}

	if result == nil{
		data = make(map[string]interface{})
	}else{
		data = result
	}

	if err := json.Unmarshal(bytes, data); err == nil{
		return data, nil
	}else{
		return nil, err
	}
}
