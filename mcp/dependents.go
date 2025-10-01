package main

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type ListDependentsInput struct {
	Package string `json:"package" jsonschema:"the name of the package to list dependents for"`
}

type ListDependentsOutput struct {
	Dependents []string `json:"dependents" jsonschema:"the dependents of the input dependency package"`
}

func (s *server) ListDependents(ctx context.Context, req *mcp.CallToolRequest, in ListDependentsInput) (*mcp.CallToolResult, ListDependentsOutput, error) {
	out := ListDependentsOutput{
		Dependents: []string{"foo", "bar"},
	}
	return nil, out, nil
}
