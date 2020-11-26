package plan

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"miner/internal/utils"
)

type Plan interface {
	// Executes the current plan
	// and deploys the relevant
	// infrastructure
	Execute()
}

// LoadPlan loads a plan into memory
func LoadPlan(uri string) (Plan, error) {
	if utils.IsHTTPURL(uri) {
		return nil, errors.New("Downloading Plans over HTTPS isn't implemented yet")
	}

	rawPlan, err := loadPlanFromDisk(uri)
	if err != nil {
		return nil, err
	}

	plan := make(map[string]interface{})

	jsonErr := json.Unmarshal(rawPlan, &plan)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if plan["Provider"] == "" {
		return nil, ErrMissingProviderValue
	}

	return nil, nil
}

func loadPlanFromDisk(fileURI string) ([]byte, error) {
	content, err := ioutil.ReadFile(fileURI)
	if err != nil {
		return nil, err
	}
	return content, err
}
