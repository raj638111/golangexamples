
# `readFile`

```go
data, err := os.readFile("balance.txt")
// (or)
data, _ := os.readFile("balance.txt")
balanceStr := string(data)
```