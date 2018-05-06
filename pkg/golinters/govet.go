package golinters

import (
	"context"

	"github.com/golangci/golangci-lint/pkg/result"
	govetAPI "github.com/golangci/govet"
)

type Govet struct{}

func (Govet) Name() string {
	return "govet"
}

func (g Govet) Run(ctx context.Context, lintCtx *Context) (*result.Result, error) {
	issues, err := govetAPI.Run(lintCtx.Paths.MixedPaths(), lintCtx.RunCfg().BuildTags, lintCtx.RunCfg().Govet.CheckShadowing)
	if err != nil {
		return nil, err
	}

	res := &result.Result{}
	for _, i := range issues {
		res.Issues = append(res.Issues, result.Issue{
			File:       i.Pos.Filename,
			LineNumber: i.Pos.Line,
			Text:       i.Message,
			FromLinter: g.Name(),
		})
	}
	return res, nil
}