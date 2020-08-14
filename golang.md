# Cheatsheet

## Language constructs

Variables
- Declaration `var x int = 5` or `x := 5`
- Types `uintXX`, `intXX`, `floatXX`, `complexXX`, `bool`, `string`
- Derived Types (pointers, array, struct, union, function, slice, interface, map, channel)

Operators
- `+ - * / % ++ --`
- `== != > < >= <=`
- `&& || !`
- `& | ^ << >>`
- `& *`


Conditions `if expression {} [else [if expression] {} ]`

Loop 
- `for [initialisation]; condition; [increment] {}`
- `for index, value := range array {}`
- `for key, value := range map {}`

Type definition `type typename struct {}`

Pointer (`&variable` memory location, `*variable` pointer)


##  Compile and run

-  Run `go run hello.go`
- Create excutable and run
  - Build
    - `go build`
    - `./{executable}`
  - Install
    - `go install`
    - `~/go/bin/{executable}`

