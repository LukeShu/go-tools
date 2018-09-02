package simple

import (
	"testing"

	"github.com/golangci/go-tools/lint/testutil"
)

func TestAll(t *testing.T) {
	testutil.TestAll(t, NewChecker(), "")
}
