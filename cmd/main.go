package main

import (
	"fmt"
	"log"
	"miner/pkg/cli"
	"miner/pkg/plan"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error: %v\n", r)
		}
	}()

	args, err := cli.ParseArgs().Validate()
	if err != nil {
		log.Fatalf("%s", err)
	}

	plan, err := plan.LoadPlan(args.PlanPath)
	if err != nil {
		log.Fatalf("%s", err)
	}

	plan.Execute()
}
