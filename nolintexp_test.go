package nolintexp_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/sivchari/nolintexp"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	analyzer := nolintexp.Analyzer
	analyzer.Flags.Set("nolintexp", "2025-01-01")
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, analyzer, "a")
}
