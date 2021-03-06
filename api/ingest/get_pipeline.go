// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package ingest

import (
	"fmt"
	"net/http"
	"net/url"
)

// GetPipeline - see https://www.elastic.co/guide/en/elasticsearch/plugins/5.x/ingest.html for more info.
//
// id: comma separated list of pipeline ids. Wildcards supported.
//
// options: optional parameters. Supports the following functional options: WithID, WithMasterTimeout, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Ingest) GetPipeline(id string, options ...*Option) (*GetPipelineResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: i.transport.URL.Scheme,
			Host:   i.transport.URL.Host,
		},
		Method: "GET",
	}
	methodOptions := supportedOptions["GetPipeline"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := i.transport.Do(req)
	return &GetPipelineResponse{resp}, err
}

// GetPipelineResponse is the response for GetPipeline
type GetPipelineResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}
