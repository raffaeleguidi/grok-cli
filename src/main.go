package main

import (
    "fmt"
    "github.com/gemsi/grok"
    "bufio"
    "log"
    "os"
)

type LineCallBack func(line string)

func info(s string){
    log.Println("INFO:", s)
}
func err(s string){
    log.Println("ERROR:", s)
}

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

    if (len(os.Args[1:]) < 1) {
        err("filename is a required argument")
        return
    }

    g := grok.New()

    err := readLines(os.Args[1], func(line string) {
        info("--- got line --------------------------------------")
        values, _ := g.Parse("%{COMMONAPACHELOG}", line)
        for k, v := range values {
            info(fmt.Sprintf("%+15s: %s", k, v))
        }
    })

    if (err != nil) {
        fmt.Println(err)
    }
}

