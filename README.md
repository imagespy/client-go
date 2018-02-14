# client-go

## Usage

```go
package main

import (
  "log"

  "github.com/imagespy/client-go"
)

func main() {
  client := imagespy.NewClientV2()
  img, err := client.Image.Get("golang:1.9.1")
  if err != nil {
    log.Fatal(err)
  }

  log.Println(img.LatestImage.Name)
  log.Println(img.LatestImage.Digest)
  log.Println(img.LatestImage.Tags)
}
```
