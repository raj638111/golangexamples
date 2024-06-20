# `func`
---

### Return void
```go
func f1() {

}
```

### Explicitly return some values
```go
func f1() float64 {
    return v1    
}

func f1() (float64, int) {
    return v1, v2    
}
```

### Implicitly return some values

```
func f1() (v1 string) { // Note: A variable name is added to type annotation
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

