package main

import (
	"fmt"
	"os"

	"github.com/matzhouse/causal/pkg/causal"
)

func main() {

	agent, err := causal.New()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = agent.Start()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
