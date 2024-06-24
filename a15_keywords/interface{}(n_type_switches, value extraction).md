
# `interface{}` / `any` type


### Interface{} in function parameter

This allows for any types to be passed as an argument to the method

```go
func print(value interface{}){
    fmt.Println(value)
}

(or)

func print(value any){
    fmt.Println(value)
}
```

### Type switches to extract type information

```go
func print(value interface{}){
    switch value.(type) {
        case int: 
            fmt.Println("Ingeger", value)
        default: // Default can also be omitted completely
            fmt.Println("rest", value)            
    }
}
```

## Extract value from interface{} type

```go
func print(value interface{}){
    typedVal, ok := value.(int)
    // ok: true if type cast succeeds
    // typedVal: Actual value as int
}
```
