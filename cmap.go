//go:build go1.18
// +build go1.18

package cmap

import "runtime"

// ConcurrentMap is a function that takes a channel of function arguments (slices of type X),
// and a function that takes function arguments and returns a channel of returns (type Y).
func ConcurrentMap[X, Y any](f func(...X) Y, args chan []X, nThreads ...int) chan Y {
	// nThreads is optional
	if len(nThreads) == 0 {
		nThreads = []int{runtime.NumCPU()}
	}

	out := make(chan Y)
	queue := make(chan chan Y, nThreads[0])
	go func() {
		for x := range args {
			queue <- func(x ...X) chan Y {
				tmp := make(chan Y)
				go func() { tmp <- f(x...) }()
				return tmp
			}(x...)
		}
		close(queue)
	}()
	go func() {
		for tmp := range queue {
			out <- <-tmp
		}
		close(out)
	}()
	return out
}
