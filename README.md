# MockinGO
Mock helpers for Go unit tests by https://sneat.team/

![Go](https://github.com/sneat-team/mockingo/workflows/Go/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/sneat-team/mockingo)](https://goreportcard.com/report/github.com/sneat-team/mockingo) [![GoDoc](https://godoc.org/github.com/sneat-team/mockingo?status.svg)](https://godoc.org/github.com/sneat-team/mockingo)


## Example of usage

**my_func.go** - code to test
``` 
import (
  "mydependency"
  "strconv"
  "strings"
)

// Pointer to dependency function so we can replace it with a mock in unit tests
var someFunc = mydependency.SomeFunc 

// This func gets 2 strings
// and repeatedely calls `mydependency.SomeFunc` as many times as requested
// and passing the 2 arguments with an iteration number added as a suffix
// and concatenate all output in the result string.
//
// **Example**: if `SomeFunc` simply concatenate 2 arguments then
// 
//   MyFunc("a", "b", 3) => "a1b1a2b2a3b3"
//
func MyFunc(a, b string, repeat int) string {
  s := make([]string, repeat)
  for i := 0; i < repeat; i++ {
    n := strconv.Itoa(i + 1)
    
    // Use pointer to the dependency instead of direct call so it can be replaced with a mock in unit tests
    s[i] = someFunc(a + n, b + n)
  }
  return strings.Join(s, "")
}
```

**my_func_test.go**
``` 
import (
  "testing"
  "strconv"
)

func TestMyFunc(t *testing.T) {
  
  // Create mock helper to store history of calls with passed arguments.
  // Parameter `t *testing.T` is needed to call `t.Helper()` inside to improve call stack reporting.
  // This reduces boilerplate code in your unit tests.
  mockedSomeFunc = mockingo.NewMockedFunc(t, "someFunc") // Name argument is used in error output if asserts are failing
  
  // replace pointer to our dependency with a mock that uses helper to store information about call
  someFunc = func(a, b string) string {
    mockedSomeFunc.Called(mockingo.NewArgument("a", a), mockingo.NewArgument("b", b))
    return a + b
  }
  
  const (
    expectedResult = "A1B1A2B2A3B3"
    numberOfCalls = 3 // The 3d argument to `MyFunc()`
  )

  // Call to function we test and assert it result 
  if result := MyFunc("A", "B", numberOfCalls); result != expectedResult {
    t.Errorf("Expected result is '%v', actual result is '%v'", expectedResult, result)
  }
  
  // Because we passed `t *testing.T` to the `NewMockedFunc()` it will report problem line propery
  mockedSomeFunc.AssertCalledAtLeastOnce(true) // True means is fatal if condition is not met
  mockedSomeFunc.AssertCalledExactly(numberOfCalls) // Verify our dependency was called expected number of times
  
  // Let's verify arguments for each call
  for i, call := range mockedSomeFunc.Calls() {
    n := strconv.Itoa(i+1)
    args := call.Args() // Get arguments passed to the mocked call
    verifyArg := func(name, value string) { // Just a small helper func to verify arguments
      t.Helper()
      if expected := value + n; args[name].(string) != expected {
        t.Errorf("Expected value  '%v' for '%v' argument for the call #%v, got: %v", expected, name, i, args[name])
      }
    }
    verifyArg("a", "A")
    verifyArg("b", "B")
  }
}
```
