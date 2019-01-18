package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cinic0101/go-esconn/api/template"
	"io/ioutil"
	"net/http"
)

const contentTypeJson = "application/json"

func (a *API) Search(index string, body map[string]interface{}) (*SearchResponse, error) {
	jsonValue, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s/_search", a.Config.URL, index)
	resp, err := http.Post(url, contentTypeJson, bytes.NewBuffer(jsonValue))
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	return &SearchResponse{
		StatusCode: resp.StatusCode,
		Resp: bodyString,
	}, nil
}

func (a *API) SearchTemplate(index string, template *template.SearchTemplate) (*SearchResponse, error) {
	return a.Search(index, template.Body)
}

type SearchResponse struct {
	StatusCode int
	Resp string
} 