package cli

import (
	"flag"
	"net/url"
)

type CLI struct {
  Url *url.URL
}

func NewCLI() *CLI {
  return &CLI{}
}

func (c *CLI) Init() error {
  var urlFlag = flag.String("u", "", "URL to diagnosis")
  flag.Parse()

  parsedUrlFlag, err := url.ParseRequestURI(*urlFlag) 
  if err != nil {
    return err
  }

  c.Url = parsedUrlFlag

  return nil
}
