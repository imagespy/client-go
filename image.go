package imagespy

import (
	"fmt"
	"net/http"
	"time"
)

// Image is a Docker image.
type Image struct {
	Created time.Time `json:"created"`
	Digest  string    `json:"digest"`
	Name    string    `json:"name"`
	Tag     string    `json:"tag"`
}

// ImageSpy is a ImageSpy.
type ImageSpy struct {
	CurrentImage *Image `json:"current_image"`
	LatestImage  *Image `json:"latest_image"`
	Name         string `json:"name"`
}

// ImageSpyService handles interactions.
type ImageSpyService struct {
	cacheUnknownImages bool
	requester          *requester
}

// Get retrieves an ImageSpy.
// Creates the ImageSpy if it does not exist.
func (is *ImageSpyService) Get(name string) (*ImageSpy, error) {
	resp, err := is.requester.readAsJSON("/v1/images/" + name)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		imageSpy := &ImageSpy{}
		err = is.requester.parseJSON(resp.Body, imageSpy)
		if err != nil {
			return nil, err
		}

		return imageSpy, nil
	default:
		return nil, fmt.Errorf("Error retrieving ImageSpy: API returned status code %d", resp.StatusCode)
	}
}
