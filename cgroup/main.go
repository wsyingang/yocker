package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
)

const cgroupMemoryHierarchyMount="/sys/fs/cgroup/memory"
func main(){
	if os.Args[0]=="/proc/self/exe"{
		fmt.Printf("current pid=%d\n",syscall.Getpid())
		cmd:=exec.Command("sh","-c",`stress --vm-bytes 200m --vm-keep -m 1`)
		cmd.SysProcAttr=&syscall.SysProcAttr{}
		cmd.Stdin=os.Stdin
		cmd.Stdout=os.Stdout
		cmd.Stderr=os.Stderr
		if err:=cmd.Run();err!=nil{
			log.Fatalln(err)
			os.Exit(-1)
		}

	}
	cmd:=exec.Command("/proc/self/exe")
	cmd.SysProcAttr=&syscall.SysProcAttr{
		Cloneflags:                 syscall.CLONE_NEWUTS|syscall.CLONE_NEWPID|syscall.CLONE_NEWNS,
	}
	cmd.Stdin=os.Stdin
	cmd.Stdout=os.Stdout
	cmd.Stderr=os.Stderr
	if err:=cmd.Start();err!=nil{
		fmt.Println("ERROR",err)
		os.Exit(1)
	}else{
		fmt.Printf("current Pid=%v",cmd.Process.Pid)
		cgroupPath := path.Join(cgroupMemoryHierarchyMount, "testmemoryLimit")
		if err := os.Mkdir(cgroupPath, 0755);err!=nil{
			log.Fatalln("Mkdir fail",err)
			os.Exit(-1)
		}
		taskPath:=path.Join(cgroupPath,"tasks")
		pid:=strconv.Itoa(cmd.Process.Pid)
		ioutil.WriteFile(taskPath,[]byte(pid),0644)
	}
	cmd.Process.Wait()
}
