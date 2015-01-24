package main

import "github.com/golang/glog"
import "flag"

func main() {
	flag.Parse()
	glog.Info("this is an info message")
	glog.V(1).Info("this is level 1")
	glog.V(2).Info("this is level 2")
	glog.V(3).Info("this is level 3")
	glog.V(4).Info("this is level 4")
	glog.V(5).Info("this is level 5")
	glog.Warning("this is a warning")
	glog.Error("this is an error")
	//glog.Fatal("this is fatal")
	glog.Flush()
}
