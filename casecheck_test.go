package casecheck_test

import (
	"testing"

	"github.com/hoddy3190/casecheck"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, casecheck.Analyzer, "a")
}