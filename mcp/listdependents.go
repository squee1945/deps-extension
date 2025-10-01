package main

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type ListDependentsInput struct {
	Name string `json:"name" jsonschema:"the name of the package to list dependents for"`
}

type ListDependentsOutput struct {
	Dependents []string `json:"deps" jsonschema:"the dependents of the input dependency package"`
}

func (s *server) ListDependents(ctx context.Context, req *mcp.CallToolRequest, in ListDependentsInput) (*mcp.CallToolResult, ListDependentsOutput, error) {
	out := ListDependentsOutput{
		Dependents: []string{"foo", "bar"},
	}
	return nil, out, nil
}
