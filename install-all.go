package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
    "os"
    "strconv"
)

func main() {
    fmt.Println("Hey Friend :3");
    if getProcessOwner() != "root\n" {
        log.Fatalln("Sorry you need to be root");
    }

    fmt.Println("Let's start by updating your system");
    
    cmd := exec.Command("pacman","-Syyyu","--noconfirm");
    stdout, err := cmd.StdoutPipe();
    if err != nil {log.Fatalln(err);}

    if err := cmd.Start(); err != nil {log.Fatal(err);}

    readBuffer(stdout);
    
    if len(os.Args) >= 1{
    
    }else {
        installBasePackage();
    }
}

func installBasePackage(){
    cmd := exec.Command("pacman", "-S","--noconfirm","--needed","$(comm -12 <(pacman -Slq|sort) <(sort pkglist.txt))");
    
    stdout, err := cmd.StdoutPipe();
    if err != nil {log.Fatalln(err);}
    if err := cmd.Start(); err != nil {log.Fatal(err);}

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

func getProcessOwner() string {
    stdout, err := exec.Command("ps", "-o", "user=", "-p", strconv.Itoa(os.Getpid())).Output()
    if err != nil {log.Fatalln(err);}
    return string(stdout)
}
