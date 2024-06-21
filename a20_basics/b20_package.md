
# Package

### How to import your own package?

```go
import modulePath/packageName
```

Here `modulePath` is module path available in `go.mod` file. Example,

```go
cat go.mod
--
module example.com/first-app
go 1.22.1
--

import example.com/first-app/yourPackageNameToImport

```

### How to make a function available for other package to import?

Only functions whose name starts with **Capital** letter are available for other package to import

### How to install a third party package

1. Run `go get <package name>`
    - Example: `go get github.com/Pallinder/go-randomdata`
    - This installs the package in your project and also update the `go.mod` file & `go.sum` with the package information

2. Sample run
```go
go get github.com/Pallinder/go-randomdata                                        
go: downloading github.com/Pallinder/go-randomdata v1.2.0
go: added github.com/Pallinder/go-randomdata v1.2.0

cat go.mod 
module example.com/first-app
go 1.22.1
require github.com/Pallinder/go-randomdata v1.2.0 // indirect

cat go.sum 
github.com/Pallinder/go-randomdata v1.2.0 h1:DZ41wBchNRb/0GfsePLiSwb0PHZmT67XY00lCDlaYPg=
github.com/Pallinder/go-randomdata v1.2.0/go.mod h1:yHmJgulpD2Nfrm0cR9tI/+oAgRqCQQixsA8HyRZfV9Y=
```    

### How to use the methods in third party package in the code

Example:
```
import github.com/Pallinder/go-randomdata
```

### Go Package discovery

https://pkg.go.dev/google.golang.org/api/discovery/v1

