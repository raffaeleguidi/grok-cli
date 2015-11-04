package main

import (
    "fmt"
    "github.com/gemsi/grok"
    "bufio"
//    "fmt"
//    "log"
    "os"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) (error) {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()

    g := grok.New()

    //var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        values, _ := g.Parse("%{COMMONAPACHELOG}", scanner.Text())
      //lines = append(lines, scanner.Text())
        for k, v := range values {
            fmt.Printf("%+15s: %s\n", k, v)
        }
        fmt.Println("--------------------------------------")
    }
    return scanner.Err()
}

func main() {

//    fmt.Println(os.Args[1:])
//    fmt.Println(len(os.Args[1:]))

    if (len(os.Args[1:]) >= 1) {
        err := readLines(os.Args[1])
        if (err != nil) {
            fmt.Println(err)
        }
    }
    return

    g := grok.New()
    values, _ := g.Parse("%{COMMONAPACHELOG}", `127.0.0.1 - - [23/Apr/2014:22:58:32 +0200] "GET /index.php HTTP/1.1" 404 207`)

    for k, v := range values {
        fmt.Printf("%+15s: %s\n", k, v)
    }
}

