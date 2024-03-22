package main

import (
	"context"
	"runtime/trace"
	"testing"
)

func TestTask(t *testing.T) {
	ctx := context.Background()
	c1, task := trace.NewTask(ctx, "A")
	defer task.End()
	c2, task := trace.NewTask(c1, "A1")
	defer task.End()
	_ = c2
}
