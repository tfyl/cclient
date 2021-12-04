package cclient

import (
	"bytes"
	"encoding/json"
	http "github.com/useflyent/fhttp"
	"io/ioutil"
	"net/url"
	"strings"
)

// ParseJSON  closes body and decodes resp body to pointer
func ParseJSON(resp *http.Response, dataOut interface{}) error {
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	ioreader := bytes.NewReader(data)

	err = json.NewDecoder(ioreader).Decode(dataOut)
	if err != nil {
		return err
	}

	return nil
}

// NewPostFormData creates new post request using string map
// TODO: make it ordered
func NewPostFormData(urlstring string, formdata map[string]string) (*http.Request, error) {
	val := &url.Values{}
	for key, data := range formdata {
		val.Add(key, data)
	}
	return http.NewRequest("POST", urlstring, strings.NewReader(val.Encode()))
}

// NewPostJson creates new post request and marshals the data as a JSON
func NewPostJson(urlstring string, data interface{}) (*http.Request, error) {
	byteBuffer, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(byteBuffer)
	return http.NewRequest("POST", urlstring, body)
}

// NewPostString creates new post request and adds string as payload
func NewPostString(urlstring string, data string) (*http.Request, error) {
	body := strings.NewReader(data)
	return http.NewRequest("POST", urlstring, body)
}
