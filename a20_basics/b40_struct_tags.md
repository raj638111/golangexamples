
# Struct Tags

1. Metadata can be added to struct fields which can be picked up by application code (or) libraries. An example would creating a json file with different field names then the one in the struct
2. Use backtick (``) to add metadata

### Example 1

```go
type Note struct {
    firstName string    `json:"first_name"`
    Content string
}
```