package main

import (
	"github.com/hoddy3190/casecheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(casecheck.Analyzer) }