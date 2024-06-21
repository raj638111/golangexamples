# `var`


### Type inference & Explicit type annotation 

```
var str = "some value"
var str string = "some vaue"
```

### Omitting `var` keyword using `:=`


```
str := "some value" // Convention
```

Note: Type annotation does not work with `:=`

### Multi variable declaration in single line

```
var str1, str2 string = "value1", "value2"
var str, no = "value1", 5
no1, no2 := 5, 6
```

# Using Pointers

### How to declare

```go
// Example 1
no := 2
myPointer := &no

// Example 2
var no int = 2
var myPointer *int = &no
```

### Pointers as function parameter

```go
// Example 1
func f1(age* int) int {
    return *age - 2
} 

// Example 2
func f1(age* int) int {
    *age = *age - 2
} 

// Example 3
var choice int
fmt.Scan(&choice)

```
