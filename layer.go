package imagespy

import (
	"fmt"
	"net/http"
)

// Layer is the digest of a layer and a list of images which created the layer.
type Layer struct {
	Digest       string   `json:"digest"`
	SourceImages []*Image `json:"source_images"`
}

// LayerService exposes the Layer API of the Image Spy API.
type LayerService struct {
	requester *requester
}

// Get retrieves a Layer.
func (ls *LayerService) Get(digest string) (*Layer, error) {
	resp, err := ls.requester.readAsJSON("/v1/layers/" + digest)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		layer := &Layer{}
		err = ls.requester.parseJSON(resp.Body, layer)
		if err != nil {
			return nil, err
		}

		return layer, nil

	case http.StatusNotFound:
		// TODO: This is only necessary because httpcache only caches when the body is read. Another solution possible?
		ls.requester.parseJSON(resp.Body, struct{}{})
		return nil, fmt.Errorf("Error retrieving Layer: No Layer with digest %s exists", digest)

	default:
		return nil, fmt.Errorf("Error retrieving Layer: API returned status code %d", resp.StatusCode)
	}
}
