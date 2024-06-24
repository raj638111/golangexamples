# Interface

- Is used to avoid code duplication

### Example: Create an interface

**Note**: 
If an interface has one method, it is a common Convention to have the same name for interface with an additional `r` (`Save()`, `Saver`)
```go
type saver interface{
    Save(...) error // We do not add function definition in interface
}
```

### Example: Use the interface
**Note**: Now any `struct` that has the `Save` signature implemented can be passed to this method `saveData()`
```go
func saveData(data saver) {
    data.Save(...)
}
```

### Example: Interface that embeds other interface

```go
type saver interface{
    Save() error
}
type displayer interface{
    Display()
}

type outputtable interface{
    saver
    displayer
} 
// (or)
type outputtable interface{
    saver
    Display()
} 
```