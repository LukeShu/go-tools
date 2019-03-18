// gosimple detects code that could be rewritten in a simpler way.
package main // import "github.com/golangci/go-tools/cmd/gosimple"
import (
	"fmt"
	"os"

	"github.com/golangci/go-tools/lint"
	"github.com/golangci/go-tools/lint/lintutil"
	"github.com/golangci/go-tools/simple"
)

func main() {
	fmt.Fprintln(os.Stderr, "Gosimple has been deprecated. Please use staticcheck instead.")
	fs := lintutil.FlagSet("gosimple")
	gen := fs.Bool("generated", false, "Check generated code")
	fs.Parse(os.Args[1:])
	c := simple.NewChecker()
	c.CheckGenerated = *gen
	lintutil.ProcessFlagSet([]lint.Checker{c}, fs)
}
