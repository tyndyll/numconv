# Numconv
`numconv` is a simple package which converts between integers and the equivalent in words

## Usage
`numconv` provides a function that takes an integer value and returns the text equivalent e.g. 1 
converts to "one", and a function that takes a string value and returns the integer equivalent e.g.
"ten" converts to 10

```
import "github.com/tyndyll/numconv"

strVal, err := numconv.Itoa(10)
// strVal = ten, err = nil. 

val, err := numconv.Atoi("ten")
// val = 10, err = nil.
```

In cases where it is not possible to convert the value, the error will not be nil
