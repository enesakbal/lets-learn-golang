# Functions In Go Language

Go does not support named parameters in functions like Dart. In Go, function parameters are positional, meaning you must pass arguments in the order they are defined in the function signature.

***In Go, the capitalization of the first letter of a function name determines its visibility (export status) outside its package.*** This is part of Go's simple but powerful access control system.

```go
// Only visible within the same package
func helperFunction() {
    // Can only be called from within this package
}

// Visible to other packages
func PublicFunction() {
    // Can be called from any package that imports this one
}
```


#  Basic Features and Key Considerations

## 1. Function Declaration
In Go, a function is declared using the `func` keyword followed by the function name, parameters, return types, and the function body.
### Syntax:
```go
func functionName(param1 Type, param2 Type) ReturnType {
    // Function body
}

func add(a int, b int) int {
    return a + b
}
```

## 2. Multiple Return Values
Go allows functions to return multiple values like Dart Language. This feature is commonly used for returning both a result and an error.

### Syntax:
```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

## 3. Named Return Values
Functions can name their return values, which acts as documentation and allows "naked" returns.

### Syntax:
```go
func Divide(a, b float64) (result float64, err error) {
    if b == 0 {
        err = errors.New("cannot divide by zero")
        return // naked return
    }
    // 
    result = a / b // It's named return value.
    return
}
```

## 4. Variadic Functions
Functions can accept variable number of arguments of the same type.
***The variadic parameter must be the last parameter.***

```go
func Sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Usage
sum := Sum(1, 2, 3, 4) // returns 10
```

## 5. Closures

Functions can access variables from their surrounding scope.
***Use case: Useful for creating middleware and generators.***

```go
Copy
func counter() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    c := counter()
    fmt.Println(c()) // 1
    fmt.Println(c()) // 2
}
```
