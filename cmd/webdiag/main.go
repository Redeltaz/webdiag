package main

import (
	"log"

	"github.com/webdiag/pkg/cli"
	"github.com/webdiag/pkg/request"
)

func main() {
  cli := cli.NewCLI()

  err := cli.Init()
  if err != nil {
    log.Fatal(err)
  }

  if err := request.CheckReachability(cli.Url); err != nil  {
    log.Fatal(err)
  }
}
