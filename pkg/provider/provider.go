package provider

import "miner/pkg/plan"

// Cloud provides an interface for
// executing plans for specific cloud
// providers
type Cloud interface {

	// CreatePlan creates the plan
	// for building the relevant Cloud
	// infrastructure
	CreatePlan() plan.Plan
}
