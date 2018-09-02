package main // import "github.com/golangci/go-tools/cmd/errcheck-ng"

import (
	"os"

	"github.com/golangci/go-tools/errcheck"
	"github.com/golangci/go-tools/lint/lintutil"
)

func main() {
	c := lintutil.CheckerConfig{
		Checker:     errcheck.NewChecker(),
		ExitNonZero: true,
	}
	lintutil.ProcessArgs("errcheck-ng", []lintutil.CheckerConfig{c}, os.Args[1:])
}
