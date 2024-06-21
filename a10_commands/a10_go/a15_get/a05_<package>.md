
# Command to install a new package (Third party package)

```bash
go get <package name>

Example:

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

# How to use the methods in third party package in the code

Example:
```
import github.com/Pallinder/go-randomdata
```