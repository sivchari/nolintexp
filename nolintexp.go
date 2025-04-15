package nolintexp

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
	"time"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "nolintexp diagnostics that the nolint directive specified expiration date exceeds the current date"

// Analyzer is an analyzer that checks for the expiration date of the nolint directive
var Analyzer = &analysis.Analyzer{
	Name: "nolintexp",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var expirationDate string

func init() {
	Analyzer.Flags.StringVar(&expirationDate, "nolintexp", "", "nolint expiration date")
}

func run(pass *analysis.Pass) (any, error) {
	if expirationDate == "" {
		expirationDate = time.Now().Format("2006-01-02")
	}
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CommentGroup)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		commentGroup, ok := n.(*ast.CommentGroup)
		if !ok {
			return
		}
		checkCommentGroup(pass, commentGroup, expirationDate)
	})

	return nil, nil
}

func checkCommentGroup(pass *analysis.Pass, commentGroup *ast.CommentGroup, expirationDate string) {
	var nolint string
	var pos token.Pos
	for _, comment := range commentGroup.List {
		if strings.Contains(comment.Text, "nolint") {
			nolint, _, _ = strings.Cut(strings.TrimPrefix(comment.Text, "//nolint:"), " ")
			pos = comment.Pos()
			break
		}
	}
	var nolintexp string
	for _, comment := range commentGroup.List {
		if strings.Contains(comment.Text, "nolintexp") {
			nolintexp, _, _ = strings.Cut(strings.TrimPrefix(comment.Text, "//nolintexp:"), " ")
			break
		}
	}
	if nolint == "" || nolintexp == "" {
		return
	}

	nolintexpDate, err := time.Parse("2006-01-02", nolintexp)
	if err != nil {
		panic(fmt.Sprintf("invalid nolintexp date: %s", nolintexp))
	}

	expiration, err := time.Parse("2006-01-02", expirationDate)
	if err != nil {
		panic(fmt.Sprintf("invalid expiration date: %s", expirationDate))
	}

	if nolintexpDate.Before(expiration) {
		pass.Reportf(pos, "nolint directive %q expired on %s", nolint, nolintexpDate.Format("2006-01-02"))
		return
	}
}
