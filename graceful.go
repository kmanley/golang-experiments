package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/kmanley/midtown"
	"os/exec"
	"time"
)

func main() {
	flag.Parse()
	c := midtown.GracefulCommand(exec.Command("gedit"))
	_ = c.Start()
	fmt.Println("waiting for gedit to stop")
	err := c.Wait(10 * time.Second)
	fmt.Printf("Wait returned, err=%v\n", err)
	if err == midtown.ErrTimeout {
		err = c.Kill(10 * time.Second)
		fmt.Println("error", err)
	}
	//fmt.Println("gedit stopped; success=", c.ProcessState.Success(), "str=", c.ProcessState.String())
	glog.Flush()

}
