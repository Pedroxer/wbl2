```golang
package main
 
import (
  "fmt"
)
 
func main() {
  var s = []string{"1", "2", "3"}
  modifySlice(s)
  fmt.Println(s)
}
 
func modifySlice(i []string) {
  i[0] = "3"
  i = append(i, "4")
  i[1] = "5"
  i = append(i, "6")
}

```

```
[3 2 3]. Потому что после первого append создаётся копия и дальнейшие изменения со слайсом не будут отражены.
```