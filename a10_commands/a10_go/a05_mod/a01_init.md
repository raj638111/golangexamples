

# Create a module out of a go project

`go mod init <module name>`

- In order to build a project, we need to create a module of the project
- This can be github url where the module is hosted (or) any name
```
go mod init example.com/first-app

go: creating new go.mod: module example.com/first-app
go: to add module requirements and sums:
	go mod tidy
```

- This creates a module file `go.mod` which indicates module has been created for the project
```
cat go.mod                                
module example.com/first-app

go 1.22.1
```


# Examples

```
go mod init

go: cannot determine module path for source directory /Users/rguna/ws/golangexamples (outside GOPATH, module path must be specified)

Example usage:
	'go mod init example.com/m' to initialize a v0 or v1 module
	'go mod init example.com/m/v2' to initialize a v2 module

Run 'go help mod init' for more information.
```


