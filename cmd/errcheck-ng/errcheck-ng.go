package main // import "github.com/golangci/go-tools/cmd/errcheck-ng"

import (
	"os"

	"github.com/golangci/go-tools/errcheck"
	"github.com/golangci/go-tools/lint"
	"github.com/golangci/go-tools/lint/lintutil"
)

func main() {
	lintutil.ProcessArgs("errcheck-ng", []lint.Checker{errcheck.NewChecker()}, os.Args[1:])
}
