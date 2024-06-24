
# Type alias

### Type alias can be used to add additional methods to built-in types

```go
package main

type str string

func (text str) log() {

}

func main(){
    var name str
}
```