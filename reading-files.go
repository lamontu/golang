package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    dat, err := ioutil.ReadFile("./data_for_reading.txt")
    check(err)
    fmt.Print(string(dat))

    f, err := os.Open("./data_for_reading.txt")
    check(err)

    b1 := make([]byte, 6)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1))

    o2, err := f.Seek(6, 0)
    check(err)
    b2 := make([]byte, 7)
    n2, err := f.Read(b2)
    fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 8)
    n3, err := io.ReadAtLeast(f, b3, 5)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

    _, err = f.Seek(0, 0)
    check(err)

    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(9)
    check(err)
    fmt.Printf("9 bytes: %s\n", string(b4))

    f.Close()
}
