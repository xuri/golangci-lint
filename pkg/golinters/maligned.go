package golinters

import (
	"context"
	"fmt"

	"github.com/golangci/golangci-lint/pkg/result"
	malignedAPI "github.com/golangci/maligned"
)

type Maligned struct{}

func (Maligned) Name() string {
	return "maligned"
}

func (m Maligned) Run(ctx context.Context, lintCtx *Context) (*result.Result, error) {
	issues := malignedAPI.Run(lintCtx.Program)

	res := &result.Result{}
	for _, i := range issues {
		text := fmt.Sprintf("struct of size %d bytes could be of size %d bytes", i.OldSize, i.NewSize)
		if lintCtx.RunCfg().Maligned.SuggestNewOrder {
			text += fmt.Sprintf(":\n%s", formatCodeBlock(i.NewStructDef, lintCtx.RunCfg()))
		}
		res.Issues = append(res.Issues, result.Issue{
			File:       i.Pos.Filename,
			LineNumber: i.Pos.Line,
			Text:       text,
			FromLinter: m.Name(),
		})
	}
	return res, nil
}