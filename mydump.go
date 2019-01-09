package mydump

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/go-test/deep"
)

// Dump does spew.Dump ignoring stringer.
func Dump(a ...interface{}) {
	scs := spew.Config
	scs.Indent = "  "
	scs.ContinueOnMethod = true
	scs.DisableMethods = true
	scs.MaxDepth = 10
	scs.Dump(a...)
}

// Diff does deep.Equal.
func Diff(a, b interface{}) []string {
	return deep.Equal(a, b)
}
