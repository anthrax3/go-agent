package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/yhat/go-agent"
	"github.com/yhat/go-agent/internal/utilization"
)

func main() {
	util := utilization.Gather(utilization.Config{
		DetectAWS:    true,
		DetectDocker: true,
	}, newrelic.NewDebugLogger(os.Stdout))

	js, err := json.MarshalIndent(util, "", "\t")
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Printf("%s\n", js)
	}
}
