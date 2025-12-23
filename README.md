# math-helpers

Библиотека с базовыми математическими функциями на Go.

## Пример использования

```go
package main

import (
    "fmt"
    "github.com/mkrtchyan0007/math-helpers/mathhelpers"
)

func main() {
    fmt.Println(mathhelpers.Factorial(5)) // 120
    fmt.Println(mathhelpers.IsPrime(7))   // true
    fmt.Println(mathhelpers.Sum([]int{1,2,3,4})) // 10
}
