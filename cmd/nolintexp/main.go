package main

import (
	"github.com/sivchari/nolintexp"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(nolintexp.Analyzer) }
