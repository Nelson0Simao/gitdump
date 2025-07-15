// main.go
package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Uso: gitdump <url>")
        return
    }

    url := os.Args[1]
    fmt.Printf("Baixando .git de: %s\n", url)
}

