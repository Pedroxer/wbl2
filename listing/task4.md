```golang
package main
 
func main() {
    ch := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
        }
    }()
 
    for n := range ch {
        println(n)
    }
}

```

```
Программа выведет значения от 0 до 9 и уйдет в дедлок. 
```