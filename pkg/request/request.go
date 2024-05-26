package request

import (
	"fmt"
	"net/http"
	"net/url"
)

func CheckReachability(url *url.URL) error {
  res, err := http.Get(url.String())
  if err != nil {
    return err
  }
  defer res.Body.Close()
  
  if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
    return fmt.Errorf("URL not reachable, returned %d as status code", res.StatusCode)
  }

  return nil
}
