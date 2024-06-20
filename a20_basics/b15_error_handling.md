# Error Handling

### Example

```go
data, err = os.ReadFile(...)
if err != nil {
    ...
}
```

### Returning error in a function

```go
func fn() (float64, err) {
    data, err := os.ReadFile(...)
    if(err != nil){
        return 1000, errors.New("Failed to read file")
    }
    // When no error
    return data, nil
}
```