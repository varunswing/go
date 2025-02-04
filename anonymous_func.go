package main
import "fmt"

// Passing anonymous function as an argument
func gfg(i func(p, q string) string) {
    fmt.Println(i("Geeks", "for"))
}
func anonymous_func() {
    value := func(p, q string) string {
        return p + q + "Geeks"
    }
    gfg(value)
}