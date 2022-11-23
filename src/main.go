package main

import (
    "os"
    "fmt"
    "os/exec"
    "log"
    "bytes"
    "github.com/uia-worker/is105komp/mycmd"
)

func main() {
    fmt.Println(mycmd.Print())
    fmt.Println("test")
    fmt.Println(os.Getpid())
    cmd := exec.Command("ls")
    
    log.Printf("Running command and waiting for it to finish...")

    var out bytes.Buffer
    cmd.Stdout = &out
	
    err := cmd.Run()
    	
    log.Printf("Command finished with error: %v", err)
    fmt.Println(out.String())

}


