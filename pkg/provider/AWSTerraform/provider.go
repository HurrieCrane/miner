package AWSTerraform

import (
	"miner/pkg/plan"
	"miner/pkg/provider"
)

type prov struct{}

func NewProvider() provider.Cloud {
	return &prov{}
}

func (p prov) CreatePlan() plan.Plan {
	return nil
}
