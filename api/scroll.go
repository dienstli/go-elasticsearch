// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"
)

// Scroll - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/search-request-scroll.html for more info.
//
// options: optional parameters. Supports the following functional options: WithScrollID, WithScroll, WithScrollIDParam, WithBody, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (a *API) Scroll(options ...*Option) (*ScrollResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: a.transport.URL.Scheme,
			Host:   a.transport.URL.Host,
		},
		Method: "GET",
	}
	methodOptions := supportedOptions["Scroll"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := a.transport.Do(req)
	return &ScrollResponse{resp}, err
}

// ScrollResponse is the response for Scroll
type ScrollResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}
