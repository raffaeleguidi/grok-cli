package main

import (
    "net"
    "io"
    "log"
//    "fmt"
)

func proxy(l net.Conn, remote net.Conn) {
    go func() {
        _, err := io.Copy(l, remote)
        if err != nil {
            log.Fatalf("io.Copy(l, remote) failed: %v", err)
        }
        log.Println("ekkime")
    }()
    go func() {
        _, err := io.Copy(remote, l)
        if err != nil {
            log.Fatalf("io.Copy(remote, l) failed: %v", err)
        }
        log.Println("remote connection closed")
    }()
}

func getRemote() (net.Conn, error) {
    remote, err := net.Dial("tcp", "localhost:27017")
    if err != nil {
		log.Fatalf("Unable to connect %s", err)
	}
    log.Println("remote connection opened")
    return remote, err
}

func main() {
    local, err := net.Listen("tcp", "localhost:9999")
    if err != nil {
		log.Fatalf("Unable to connect %s", err)
	}
    defer local.Close()

    log.Println("proxy started")
    for {
        conn, err := local.Accept()
        if err != nil {
            log.Fatalf("listen Accept failed %s", err)
        }
        log.Println("incoming connection accepted")

        remote, err := getRemote()
        if (err == nil) {
            proxy(conn, remote)
        }
        defer remote.Close()
    }
}
