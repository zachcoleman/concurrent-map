package cmap

import "testing"

func TestConcurrentMap(t *testing.T) {
	in := make(chan []int)
	out := ConcurrentMap(
		func(x ...int) int {
			return x[0] + x[1]
		}, // f
		in, // args
	)

	in <- []int{1, 2}
	in <- []int{3, 4}
	close(in)

	if <-out != 3 {
		t.Fail()
	}
	if <-out != 7 {
		t.Fail()
	}
	if _, ok := <-out; ok {
		t.Fail()
	}
}
