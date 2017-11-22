package main

import "crypto/sha1"
import "fmt"

func main() {
    s := "sha1 this string"
    fmt.Println(s)
    fmt.Println()

    h := sha1.New()
    h.Write([]byte(s))
    fmt.Println(h)
    fmt.Println()

    bs := h.Sum(nil)
    fmt.Println(bs)
    fmt.Printf("%x\n", bs)
}
