package opa

import (
	"io"
	"time"

	"github.com/kroekle/opa/ast"
	"github.com/kroekle/opa/metrics"
	"github.com/kroekle/opa/topdown/cache"
	"github.com/kroekle/opa/topdown/print"
)

// Result holds the evaluation result.
type Result struct {
	Result []byte
}

// EvalOpts define options for performing an evaluation.
type EvalOpts struct {
	Input                  *interface{}
	Metrics                metrics.Metrics
	Entrypoint             int32
	Time                   time.Time
	Seed                   io.Reader
	InterQueryBuiltinCache cache.InterQueryCache
	PrintHook              print.Hook
	Capabilities           *ast.Capabilities
}
