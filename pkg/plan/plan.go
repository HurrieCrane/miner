package plan

type Plan interface {
	// Executes the current plan
	// and deploys the relevant
	// infrastructure
	Execute()
}
