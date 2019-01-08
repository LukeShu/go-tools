// staticcheck analyses Go code and makes it better.
package main // import "github.com/golangci/go-tools/cmd/staticcheck"

import (
	"os"

	"github.com/golangci/go-tools/lint"
	"github.com/golangci/go-tools/lint/lintutil"
	"github.com/golangci/go-tools/simple"
	"github.com/golangci/go-tools/staticcheck"
	"github.com/golangci/go-tools/stylecheck"
	"github.com/golangci/go-tools/unused"
)

func main() {
	fs := lintutil.FlagSet("staticcheck")
	fs.Parse(os.Args[1:])

	checkers := []lint.Checker{
		simple.NewChecker(),
		staticcheck.NewChecker(),
		stylecheck.NewChecker(),
	}

	uc := unused.NewChecker(unused.CheckAll)
	uc.ConsiderReflection = true
	checkers = append(checkers, unused.NewLintChecker(uc))

	lintutil.ProcessFlagSet(checkers, fs)
}
