package main

import "syscall"
import "os"
import "os/exec"

//os/exec 和 syscall的区别
func main() {
    binary,lookErr := exec.LookPath("ls")
    if lookErr != nil {
        panic(lookErr)
    }

    args := []string{"ls","-a","-l","-h"}

    env := os.Environ()
    execErr := syscall.Exec(binary,args,env)
    if execErr != nil {
        panic(execErr)
    }
}
