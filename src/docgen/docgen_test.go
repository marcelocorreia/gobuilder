package docgen_test

import (
	"testing"
	"docgen"
)

func TestSTS(t *testing.T) {
	dg := docgen.DocGen{}
	dg.ProcessMarkdown("/home/marcelo/Documents/da-service/README.md")

}