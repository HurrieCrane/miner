package cli

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

/* ===== ParseArgs tests ===== */

func TestParseArgsForDisplayHelp(t *testing.T) {
	_ = flag.Set("help", "true")
	_ = flag.Set("providers", "false")
	args := ParseArgs()

	assert.True(t, args.DisplayHelp)
	assert.False(t, args.ListProviders)
	assert.Equal(t, "", args.PlanPath)
}

func TestParseArgsForListProviders(t *testing.T) {
	_ = flag.Set("help", "false")
	_ = flag.Set("providers", "true")
	args := ParseArgs()

	assert.False(t, args.DisplayHelp)
	assert.True(t, args.ListProviders)
	assert.Equal(t, "", args.PlanPath)
}

func TestParseArgsForPlanPath(t *testing.T) {
	_ = flag.Set("help", "false")
	_ = flag.Set("providers", "false")
	args = Arguments{}

	sut := "./go.mod"
	os.Args = append(os.Args, sut)
	args := ParseArgs()

	assert.False(t, args.DisplayHelp)
	assert.False(t, args.ListProviders)
	assert.Equal(t, sut, args.PlanPath)
}

func TestParseArgsForLen1Args(t *testing.T) {
	_ = flag.Set("help", "false")
	_ = flag.Set("providers", "false")
	args = Arguments{}

	os.Args = []string{""}
	args := ParseArgs()

	assert.Equal(t, 1, len(os.Args))

	assert.False(t, args.DisplayHelp)
	assert.False(t, args.ListProviders)
	assert.Equal(t, "", args.PlanPath)
}

/* ===== Validate tests ===== */
func TestValidateReturnsArgumentsForValidArguments(t *testing.T) {
	sut := Arguments{
		PlanPath:      "./go.mod",
		DisplayHelp:   false,
		ListProviders: false,
	}
	args, err := sut.Validate()
	assert.Nil(t, err)
	assert.Equal(t, &sut, args)
}

func TestValidateReturnsArgumentsForEmptyPathAndHelp(t *testing.T) {
	sut := Arguments{
		PlanPath:      "",
		DisplayHelp:   true,
		ListProviders: false,
	}
	args, err := sut.Validate()
	assert.Nil(t, err)
	assert.Equal(t, &sut, args)
}

func TestValidateReturnsArgumentsForEmptyPathAndProviders(t *testing.T) {
	sut := Arguments{
		PlanPath:      "",
		DisplayHelp:   false,
		ListProviders: true,
	}
	args, err := sut.Validate()
	assert.Nil(t, err)
	assert.Equal(t, &sut, args)
}

func TestValidateReturnsArgumentsForEmptyPathAndProvidersAndHelp(t *testing.T) {
	sut := Arguments{
		PlanPath:      "",
		DisplayHelp:   true,
		ListProviders: true,
	}
	args, err := sut.Validate()
	assert.Nil(t, err)
	assert.Equal(t, &sut, args)
}

func TestValidateReturnsMissingPlanForEmptyPathAndNoHelpOrProviders(t *testing.T) {
	sut := Arguments{
		PlanPath:      "",
		DisplayHelp:   false,
		ListProviders: false,
	}
	args, err := sut.Validate()
	assert.Equal(t, ErrMissingPlanPath, err)
	assert.Nil(t, args)
}
