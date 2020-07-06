// Package switcheroo executes an anonymous function based on certain string patterns.
package switcheroo

import (
	"context"
	"sync"

	"github.com/didip/switcheroo/libstring"
)

// New is the constructor of Router
func New(ctx context.Context) *Router {
	r := &Router{
		ctx:      ctx,
		mtx:      &sync.Mutex{},
		routeMap: make(map[string]RouterFunc),
	}
	return r
}

// RouterFunc is the matched handler.
type RouterFunc func(context.Context, map[string]string)

// Router multiplexes based on string pattern.
// The string pattern is {name} based. Similar to most HTTP routers pattern.
type Router struct {
	ctx      context.Context
	mtx      *sync.Mutex
	routeMap map[string]RouterFunc
}

// Add a new pattern and handler into the router.
func (r *Router) Add(curlyPattern string, fn RouterFunc) {
	r.mtx.Lock()
	r.routeMap[curlyPattern] = fn
	r.mtx.Unlock()
}

// Executes all the patterns against the incoming string.
func (r *Router) Run(in string) {
	for curlyPattern, routerFunc := range r.routeMap {
		isMatched, params, err := libstring.Match(curlyPattern, in)

		if isMatched && err == nil {
			routerFunc(r.ctx, params)
		}
	}
}
