package main

import (
    "fmt"
    "github.com/gemsi/grok"
    "bufio"
    "strings"
    "log"
    "os"
)

type LineCallBack func(line string)

func readLines(path string, cb LineCallBack) (error) {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    //var lines []string //keep it for multiline
    for scanner.Scan() {
        cb(scanner.Text())
      //lines = append(lines, scanner.Text())
    }
    return scanner.Err()
}

func main() {
    if (len(os.Args[1:]) < 2) {
        fmt.Println("--------------------------------")
        fmt.Println("grok")
        fmt.Println("--------------------------------")
        fmt.Println("usage: ")
        fmt.Println("\tgrok <filename> \"<pattern>\" [patternsFile] [newLinePattern]\n\r")
        fmt.Println("*error* filename and pattern are required arguments")
        return
    }

    g := grok.New()

    file := os.Args[1]
    pattern := os.Args[2]

    fmt.Println("...scanning", file, "for pattern", pattern)

    if len(os.Args[1:]) >= 3 {
//        patternsDir := os.Args[3]
//        g.AddPatternsFromPath(patternsDir) // bit wirjubg
        patternsFile := os.Args[3]
        err := readLines(patternsFile, func(line string) {
            n := strings.Index(line, " ")
            name := line[:n]
            body := line[n+1:]
            g.AddPattern(name, body)
        })

        if (err != nil) {
            log.Fatalf("oops", err)
        }


    }

    if len(os.Args[1:]) >= 4 {
        newLinePattern := os.Args[4]
        fmt.Println("newline pattern:", newLinePattern)
        fmt.Println("*warning* multiline pattern matching is not yet implemented")
    }
    // yet to be implemented

    err := readLines(file, func(line string) {
        log.Println("--- newline --------------------------------------")
        //values, _ := g.Parse("%{COMMONAPACHELOG}", line)
        values, _ := g.Parse(pattern, line)
        for k, v := range values {
            log.Println(fmt.Sprintf("%+15s: %s", k, v))
        }
    })

    if (err != nil) {
        log.Fatal(err)
    }
}

