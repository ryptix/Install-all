package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
    fmt.Println("Hey Friend :3");
    fmt.Println("Let's start by updating your system");
    
    cmd := exec.Command("pacman","-Syyyu","--noconfirm");
    stdout, err := cmd.StdoutPipe();
    if err != nil {log.Fatalln(err);}

    if err := cmd.Start(); err != nil {
		log.Fatal(err);
	}

    readBuffer(stdout);
}

func readBuffer(buff io.ReadCloser){
    for {
        // Make a buffer of byte
        data := make([]byte, 1024);
        n, err := buff.Read(data);
        if err == io.EOF {break;}
        if err != nil {log.Fatalln(err);}
        // byte to string
        str := string(data[:n]);
        fmt.Print(str);

    }
}
