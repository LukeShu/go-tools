package errcheck

import (
	"testing"

	"github.com/golangci/go-tools/lint/testutil"
)

func TestAll(t *testing.T) {
	c := NewChecker()
	testutil.TestAll(t, c, "")
}
