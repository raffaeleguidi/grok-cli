package main

import (
    "fmt"
    "github.com/gemsi/grok"
    "bufio"
    "strings"
    "log"
    "os"
    "io/ioutil"
    "regexp"
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

func array2string(array []string) (string) {
    return strings.Join(array, "\n")
}

func readLinesWithRegex(path string, newlineRegexp string, cb LineCallBack) (error) {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()

    r, _ := regexp.Compile(newlineRegexp)

    scanner := bufio.NewScanner(file)

    var lines []string //keep it for multiline
    for scanner.Scan() {
        thisLine := scanner.Text()
        if !r.Match([]byte(thisLine)) {
            fmt.Println("not newline", thisLine)
            if (len(lines) > 0) {
                cb(array2string(lines))
                lines = []string{}
            }
        } else {
            if (len(lines) > 0) {
                cb(array2string(lines))
                lines = []string{}
            }
            fmt.Println("newline", thisLine)
            lines = append(lines, scanner.Text())
        }
    }
    if (len(lines) > 0) {
        fmt.Println("lastline", array2string(lines))
        cb(array2string(lines))
    }
    return scanner.Err()
}

func loadPatternsFile(filename string, g grok.Grok) (error){
    err := readLines(filename, func(line string) {
        n := strings.Index(line, " ")
        name := line[:n]
        body := line[n+1:]
        g.AddPattern(name, body)
    })
    return err
}

func loadPatternsDir(patternsDir string, g grok.Grok) (error){
    files, _ := ioutil.ReadDir(patternsDir)
    for _, f := range files {
        fmt.Println("loading file", patternsDir + f.Name())
        err := loadPatternsFile(patternsDir + f.Name(), g)
        if (err != nil) {
            return err
        }
    }
    return nil
}

func main() {
    if (len(os.Args[1:]) < 2) {
        fmt.Println("--------------------------------")
        fmt.Println("grok")
        fmt.Println("--------------------------------")
        fmt.Println("usage: ")
        fmt.Println("\tgrok <filename> \"<pattern>\" [patternsDir] [newLinePattern]\n\r")
        fmt.Println("*error* filename and pattern are required arguments")
        return
    }

    g := grok.New()

    file := os.Args[1]
    pattern := os.Args[2]

    fmt.Println("...scanning", file, "for pattern", pattern)

    if len(os.Args[1:]) >= 3 {
        patternsDir := os.Args[3]
        err := loadPatternsDir(patternsDir, *g) //g.AddPatternsFromPath(patternsDir) // not working!
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
//    err := readLinesWithRegex(file, "^\\d", func(line string) {
        log.Println("--- newline ---", line)
        values, _ := g.Parse(pattern, line)
        for k, v := range values {
            log.Println(fmt.Sprintf("%+15s: %s", k, v))
        }
    })

    if (err != nil) {
        log.Fatal(err)
    }
}

