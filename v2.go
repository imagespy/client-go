package imagespy

import (
	"fmt"
	"net/http"
)

// ImageV2 is an Image of the V2 API.
type ImageV2 struct {
	Digest      string         `json:"digest"`
	LatestImage *LatestImageV2 `json:"latest_image"`
	Name        string         `json:"name"`
	Tags        []string       `json:"tags"`
}

// LatestImageV2 is the latest image of an image.
type LatestImageV2 struct {
	Digest string   `json:"digest"`
	Name   string   `json:"name"`
	Tags   []string `json:"tags"`
}

// LayerV2 is a Layer of the V2 API.
type LayerV2 struct {
	Digest       string     `json:"digest"`
	SourceImages []*ImageV2 `json:"source_images"`
}

// ImageServiceV2 exposes the /v2/images API.
type ImageServiceV2 struct {
	registryWhitelist map[string]struct{}
	requester         *requester
}

// Get returns an Image.
func (i *ImageServiceV2) Get(ref string) (*ImageV2, error) {
	resp, err := i.requester.readAsJSON("/v2/images/" + ref)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		image := &ImageV2{}
		err = i.requester.parseJSON(resp.Body, image)
		if err != nil {
			return nil, err
		}

		return image, nil
	case http.StatusNotFound:
		// TODO: This is only necessary because httpcache only caches when the body is read. Another solution possible?
		i.requester.parseJSON(resp.Body, struct{}{})
		return nil, &NotFoundError{message: fmt.Sprintf("Error retrieving ImageSpy: API returned status code %d", resp.StatusCode)}
	default:
		return nil, fmt.Errorf("Error retrieving ImageSpy: API returned status code %d", resp.StatusCode)
	}
}

// LayerServiceV2 exposes the /v2/images API.
type LayerServiceV2 struct {
	requester *requester
}

// Get returns a Layer.
func (l *LayerServiceV2) Get(digest string) (*LayerV2, error) {
	resp, err := l.requester.readAsJSON("/v2/layers/" + digest)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		layer := &LayerV2{}
		err = l.requester.parseJSON(resp.Body, layer)
		if err != nil {
			return nil, err
		}

		return layer, nil

	case http.StatusNotFound:
		// TODO: This is only necessary because httpcache only caches when the body is read. Another solution possible?
		l.requester.parseJSON(resp.Body, struct{}{})
		return nil, &NotFoundError{message: fmt.Sprintf("Error retrieving Layer: No Layer with digest %s found", digest)}

	default:
		return nil, fmt.Errorf("Error retrieving Layer: API returned status code %d", resp.StatusCode)
	}
}
