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
	cmd.SysProcAttr.Cloneflags=syscall.CLONE_NEWNET
	cmd.Stdin=os.Stdin
	cmd.Stderr=os.Stderr
	cmd.Stdout=os.Stdout
	if err:=cmd.Run();err!=nil{
		log.Fatalln(err)
	}
	os.Exit(-1)
}
