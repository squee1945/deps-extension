package main

import (
	"context"
	"log"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	ctx := context.Background()

	mcpServer := mcp.NewServer(&mcp.Implementation{Name: "deps-extension", Version: "v1.0.0"}, nil)

	s := newServer()

	mcp.AddTool(mcpServer, &mcp.Tool{Name: "list_dependents", Description: "This tool uses deps.dev to find dependents for a given package"}, s.ListDependents)

	if err := mcpServer.Run(ctx, &mcp.StdioTransport{}); err != nil {
		exit("server.Run(): %v", err)
	}
}

func exit(f string, args ...any) {
	log.Printf("Error: "+f+"\n", args...)
	os.Exit(1)
}
