package main

import "github.com/golang/glog"
import "flag"

func main() {
	flag.Parse()
	glog.Info("this is an info message")
	glog.V(1).Info("this is level 1")
	glog.V(5).Info("this is level 5")
	glog.V(10).Info("this is level 10")
	glog.Warning("this is a warning")
	glog.Error("this is an error")
	//glog.Fatal("this is fatal")
	glog.Flush()
}
