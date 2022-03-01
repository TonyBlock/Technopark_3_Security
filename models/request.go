package models

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

type Request struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Url      string `json:"url" gorm:"not null"`
	Method   string `json:"method" gorm:"not null"`
	Protocol string `json:"protocol" gorm:"not null"`
	Headers  string `json:"headers"`
	Body     string `json:"body"`
}

func (request *Request) GetHTTPRequest() (*http.Request, error) {
	httpReq, err := http.NewRequest(request.Method, request.Url, strings.NewReader(request.Body))
	if err != nil {
		return nil, err
	}
	httpReq.Proto = request.Protocol
	jsonMap := make(map[string][]string)
	if unmarshalErr := json.Unmarshal([]byte(request.Headers), &jsonMap); unmarshalErr != nil {
		return nil, unmarshalErr
	}
	for key, values := range jsonMap {
		for _, value := range values {
			httpReq.Header.Add(key, value)
		}
	}
	return httpReq, nil
}

func (request *Request) SetHTTPRequest(httpReq http.Request) error {
	request.Url = httpReq.RequestURI
	request.Method = httpReq.Method
	request.Protocol = httpReq.Proto
	tmp, err := json.Marshal(httpReq.Header)
	if err != nil {
		return err
	}
	request.Headers = string(tmp)
	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(httpReq.Body); err != nil {
		return err
	}
	request.Body = buf.String()
	return nil
}
