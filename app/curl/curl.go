package curl

import (
	"io/ioutil"
	"net/http"
	"net/url"
)


func Get(apiURL string, params url.Values) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		return nil, err
	}
	Url.RawQuery = params.Encode()
	resp, err := http.Get(Url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
