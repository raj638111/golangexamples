
# `return`
---

```go
return a
return a, b
```

# Implicit return

**In function definition, tag a variable name (this creates the variable) along with the type**
```go
func f1() (v1 string) {
    ...
    v1 = "some str"
    // Note: No explicit return statement here    
}

 (or)

func f1() (v1 string) {
    ...
    v1 = "some str"
    return // This returns v1    
} 
```

# Storing the return values

```go
value := fn()
v1, v2 := fn()
```

# Empty return 

```go
func main() {
    return
}
```