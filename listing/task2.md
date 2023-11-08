```golang
package main
 
import (
    "fmt"
)
 
func test() (x int) {
    defer func() {
        x++
    }()
    x = 1
    return
}
 
 
func anotherTest() int {
    var x int
    defer func() {
        x++
    }()
    x = 1
    return x
}
 
 
func main() {
    fmt.Println(test())
    fmt.Println(anotherTest())
}
```

```
2, 1.  
При выходе из test сработает defer, далее, после выполнения anotherTest, выполнится defer из anotherTest
Defer в test изменит значение х на +1. 
Defer в anotherTest не изменит x, так как в defer'e будет лежать копия локальной переменной,
в отличие от ситуации в Test, где есть именованное возвращаемое значение, которое живет до конца функции.  
```