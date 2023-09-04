# concurrent-map
A concurrent map implementation in Go inspired by Python's `concurrent.futures` functionality (order is maintained). 

## Example
```go
// function
repeat := func(x ...interface{}) string {
	out := ""
	for i := 0; i < x[1].(int); i++ {
		out += fmt.Sprintf("%s! ", x[0])
	}
	return out
}
// inputs 
inputs := [][]interface{}{
	{"foo", 1}, {"bar", 2}, {"baz", 3},
}

// send inputs to args channel
args := make(chan []interface{})
go func() {
	for _, x := range inputs {
		args <- x
	}
	close(args)
}()

// concurrently eval function over input args (w/ default number of threads)
out := cmap.ConcurrentMap(repeat, args)

// consume output
for x := range out {
	fmt.Println(x)
}
```
