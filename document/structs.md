# Structs
A struct (short for structure) in Go is a user-defined composite data type that groups together variables (called fields) under a single name. Structs are used to represent real-world entities (like a Person, Car, or Book) by bundling related data fields.

### How Are Structs Different from Classes?

***Key Similarities:***

- Both structs in Go and classes in other OOP languages group related data (fields) together.
- Both have the ability to define methods for their behavior.

***Key Differences:***

- Go does not have inheritance, which is a hallmark of OOP languages like Java or Python. However, Go does allow composition using struct embedding.
- Go uses explicit method binding, where methods are associated with specific structs, not implicitly tied to the struct's instance.


## 1. Basic Struct Declaration

### Syntax:
```go
// Define a struct type
type Person struct {
    Name string
    Age  int
}

// Create an instance
p := Person{
    Name: "Alice",
    Age:  30,
}

fmt.Println(p.Name) // Access field
p.Age = 31         // Modify field
```



## 2. Embedded Structs (Composition)
In the Go programming language, Embedded Structures are a way to reuse a struct’s fields without having to inherit from it. It allows you to include the fields and methods of an existing struct into a new struct. The new struct can add additional fields and methods to those it inherits.

### Syntax:
```go
type Address struct {
    City  string
    State string
}

type Employee struct {
    Person  // Embedded struct
    Address // Embedded struct
    Salary  float64
}

// Usage
emp := Employee{
    Person: Person{Name: "Bob", Age: 40},
    Address: Address{City: "NY", State: "NY"},
    Salary: 75000,
}

fmt.Println(emp.Name) // Fields promoted from Person
fmt.Println(emp.City) // Fields promoted from Address
```


## 3. Struct Methods
In Go, methods are functions that are associated with a particular type. These methods can be attached to structs and provide behavior for the data they hold. Methods are defined with a ***receiver***, which is the instance of the struct the method acts upon.

### Syntax:
```go
func (p Person) Greet() string { ... }  // Value receiver
func (p *Person) Birthday() { ... }     // Pointer receiver
```
___Value receiver: Works on a copy (safe for read-only)___

___Pointer receiver: Modifies original struct (use for changes)___



## 4. Struct Tags
Golang provides packages in the standard library to write and retrieve structures to/from the JSON file. In the definition of structures, field tags are added to the field declaration which is used as the field name in JSON files.

[More information about tags](https://go.dev/wiki/Well-known-struct-tags)

### Syntax:
```go
type Person struct {
    Name    string `json:"name"`         // field tag for Name
    Aadhaar int    `json:"aadhaar"`       // field tag for Aadhaar
    Street  string `json:"street"`       // field tag for Street
    HouseNo int    `json:"house_number"` // field tag for HouseNO
}
```

### 🚨 Important Notes
```go
var p Person // Zero value: Person{Name: "", Age: 0}

p1 := Person{"Alice", 30}
p2 := Person{"Alice", 30}
fmt.Println(p1 == p2) // true (if all fields are comparable)

type account struct { // Unexported (private to package)
    balance float64   // Unexported field
    Owner   string    // Exported field
}

//  Anonymous Structs
temp := struct {
    ID   int
    Name string
}{
    ID:   1,
    Name: "Temp",
}
```