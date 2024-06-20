# Available types

1. `int`
2. `uint`
3. `float64`
4. `string`
5. `bool`
6. `err`

# Type inference & Explicit type annotation 

```
var str = "some value"
var str string = "some vaue"
```

# Omitting `var` keyword using `:=`


```
str := "some value" // Convention
```

Note: Type annotation does not work with `:=`

# Multi variable declaration in single line

```
var str1, str2 string = "value1", "value2"
var str, no = "value1", 5
no1, no2 := 5, 6
```

### string

- Double quotes or backticks are allowed
```
var str = "some value"
var str = `mutli line 
string`
```



