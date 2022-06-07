package util

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) (map[string]interface{}, error) {
	res := make(map[string]interface{})

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}
