
# Create an executable from go module

`go build`

This create an executable with the module name
```
ls
... first-app ...
```
```
./first-app
Hello world!%
```

# `main` package & function

In order to create an executable, one of the go file should be part of `main`
 package and should contain a `main` function, which denotes an entry point for the application
```
cat app.go

// `main` is a special package name which denotes main entry point of the application
package main

import "fmt"

func main() {
	fmt.Print("Hello world!")
}
```

 # Multiple files in `main` package

 When there are multiple go files in main package, only one file should have the `main` function.
 Or else `build` will fail with an error

 # Third party libraries & `main` function

If you are building third party library, then `main` function can be omitted

