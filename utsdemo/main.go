package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main(){
	cmd:=exec.Command("sh")
	cmd.SysProcAttr=&syscall.SysProcAttr{}
	cmd.SysProcAttr.Cloneflags=syscall.CLONE_NEWNS
	cmd.Stdin=os.Stdin
	cmd.Stdout=os.Stdout
	cmd.Stderr=os.Stderr
	if err:=cmd.Run();err!=nil{
		log.Fatalln("run error=%v",err)
	}
}
