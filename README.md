# gassert

`gassert` is Go package that provides function like `assert` of Python or C++.

With gassert, you can check validation of parameters or values more easily and cleaner.

## Usage

- #### Basic

Since Golang don't provide `assert` function by default, if you want to check validation of params, you should write like below. 

```go
func MyFunc(n int, s string, arr []string, p *MyStruct) {
    if n == 0 || len(s) == 0 || arr == nil || p == nil {
        // error
    }
}
```

With `gassert`, you can write that:

```go
func MyFunc(n int, s string, arr []string, p *MyStruct) {
  	// raise panic if any of them has zero-value
	gassert.Zeros(n, s, arr, p)
}
```

And you can check other conditions besides zero-value.

```go
func MyFunc(n int, s string, arr []string, p *MyStruct) {
  	gassert.NumLess(n, 100)
  	gassert.StrLenGreater(s, 10)
  	gassert.SliceLenEquals(arr, 0)
  	gassert.Zeros(p)
}
```

- #### Multiple

If you want to check multiple condition on one line, you can write below:

```go
func MyFunc(n int, s string, arr []string, p *MyStruct) {
  	// raise panic if any of conditions is true
  	gassert.New().NumLess(n, 5).StrLenGreater(s, 10).SliceLenEquals(arr, 0).Zeros(0).Panic()
  
  	if err := gassert.New().NumLess(n, 5).StrLenGreater(s, 10).SliceLenEquals(arr, 0).Zeros(0).Err(); err != nil {
    	// handle error
  	}
}
```

When you write like this, SHOULD start with `New()` and end with `Panic()` , `PanicDetails()`, `Err()`, `ErrDetails`.

- `Panic()` : raise panic if any of condition is true
- `PanicDetail()`: raise panic if any of condition is true with information
- `Err()`: return `error` if any of condition is true
- `ErrDetails()`: return `error` if any of condition is true with information

- #### Number

There is difference from general comparison statement in number. In Go, you can compare in only same types.

But `gassert` compares its values no matter what types.  For example, you can write this:

```go
func MyFunc(n int) {
  	gassert.NumLess(n, 12.34)
}
```

It may lose type-safety, but i thougth it is better idea for convenience.

- #### Thread-Safety

`gassert` is thread-safety because it makes internal event object every calls. And the object is managed by `sync.Pool`, so it is economical in memory usage.

- #### Complex condition

If you want to check complex condition which doesn't exists in `gassert`. You can just use `Go()`.

```go
func MyFunc(n int, s string) {
  	gassert.Go((n-5) % 10 == 10 || len(s)*2 > 10)
}
```

## License

[MIT](https://github.com/zajann/gassert/blob/main/LICENSE)
