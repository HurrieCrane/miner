package main

import (
	"fmt"
	"miner/pkg/cli"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error: %v\n", r)
		}
	}()

	args, err := cli.ParseArgs().Validate()
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf(args.PlanPath)
}
