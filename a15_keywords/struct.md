# Create custom struct using `type` & `struct`

[[_TOC_]]

### Define a struct & create a struct instance

```go
type user struct{ // User can be used to make the struct available for import
    firstName string
    lastName string
    birthDate string
    createdAt time.Time
}

var appUser user

// Create an instance: Method 1
appUser = user {
    firstName: "fname",
    lastName: "lname",
    birthDate: "2023",
    createAd: time.Now(), // Note: Comma is also needed for the last field
}

// Create an instance: Method 2
appUser = user {
    "fname",
    "lname",
    "2023",
    time.Now(),
}

// Create an instance with null value
appUser = user{}

```

### Add method to `struct` using Receiver argument

```go
type user struct {
    firstName: string
}


// - Here (u user) is called Receiver argument / parameter
// - ** This receiver argument is a copy of the struct instance. not the original struct instance. Refer pointer section for an example of how to refer the actual instance
func (u user) print() { 
    fmt.Println(u.firstName)
}

appUser := user{
    "somefirstname"
}

appUser.print()

```

# Struct Pointers

### Passing struct as pointers

```go
type user struct {
    firstName: string
}


func print(u *user) {
    fmt.Println(u.firstName) // Note: The way we refer the value of struct pointer is different from regular type pointer
    (or)
    fmt.Println((*u).firstName)
}

```

### Using pointer Receiver argument to change the actual struct instance


```go
type user struct {
    firstName: string
}

func (u *user) print() { 
    fmt.Println(u.firstName)
}

func (u *user) clear() {
    u.firstName = ""
}

appUser := user{
    "somefirstname"
}

appUser.print()
appUser.clear()

```

# Constructor function pattern (Convention)

### Creating constructor function

It is the convention to create a constructor function by prefixing `new` to the struct name. Example: If struct name is `user`, then the constructor function will be `newUser`

```go
type user struct {
    firstName: string
}

func (u *user) print() { 
    fmt.Println(u.firstName)
}

func (u *user) clear() {
    u.firstName = ""
}

// Construction function: Returning a value 
// func newUser(firstName string) user {
//     return user{firstName}
// }

// Constructor function: Returning a pointer
func newUser(firstName: String) *user {
    return &user{firstName}
}

var appUser *user = newUser("somename")

appUser.print()
appUser.clear()

```

### Constructor function for validation

```go
func newUser(firstName) (*user, error) {
    ...
}
```

# Packaging a struct

### Example
`cat user.go`
```go
package user
type User struct {
    FirstName: string // Exportable field
    lastName: String  // Not exportable
}

// Note: Once a struct with constructor fucntion is moved to a new package, we can just name the constructor function as `New` (Convention)
//func NewUser(firstName: String) *User {
func New(firstName: String) *User {    
    return &User{firstName}
}
```

`cat main.go`
```go
package main
import example.com/first-app/user

var u *user.User = user.New("somename")
```

# Nested struct (Embedding a struct within another struct)

### Way 1: Example: Anonymous embedding (Recommended and convenient)
```go
type Admin struct {
    email string
    User // Anonymous embedding
}

func newAdmin(email string) Admin{
    return Admin{
        email: email,
        User: User{...}
    }
}

var admin = new Admin(...)
// Note: A method / field available in User can be called directly when using anonymous embedding
//admin.User.Clear()
admin.Clear()

```

### Way 2: Example

```go
type Admin struct {
    email string
    password string
    u User
}
func newAdmin(email string) Admin{
    return Admin{
        email: email,
        u: User{...}
    }
}
```