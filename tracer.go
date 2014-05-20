package main

import (
  "fmt"
  "reflect"
)
var print = fmt.Println

func trace(s string) string {
    fmt.Println("entering:", s)
    return s
}

func un(s string) {
    fmt.Println("leaving:", s)
}

func a() {
    defer un(trace("a"))
    fmt.Println("in a")
}

func b() {
    defer un(trace("b"))
    fmt.Println("in b")
    a()
}

/*
func traced(x func(args ...interface{})) func(...interface{}) {
    wrap := func(args...interface{}) {
       defer un(trace("some func"))
       x(args...)
    }
    return wrap
}
*/

/*
func traced(x func(...interface{}) interface{}) func(...interface{}) {
    wrap := func(args...interface{})    {
       defer un(trace("some func"))
       x(args...)
    }
    return wrap
}
*/
func traced(wrapped interface{}) func(...interface{}) interface{} {
    wrapper := func(args...interface{}) interface{} {
        defer un(trace("some func"))
        print(args)
        v := reflect.ValueOf(wrapped)
        vargs := make([]reflect.Value, len(args))
        for i, arg := range args {
            vargs[i] = reflect.ValueOf(arg)
        }
        print(vargs)
        print(v)
        result := v.Call(vargs)
        print("result", result)
        if len(result) == 1 {
            var ret interface{}
            ret = result[0].Interface()
            return ret
        } else {
            ret := make([]int, len(result))
            print("ret", ret)
            var k int = 0
            print("ret0", ret[k])
            return 20,40
            //for j, res := range result {
                //ret[i] = res.Interface()
                //print("ha", i, res)
                //print(ret[j], "wtf?")
            //}
        }
        //return ret
        return 0
        
    }
    return wrapper
}


func _c(x int, y int) (int,int) {
    ret := x * y
    print("hello ", ret)
    return ret, 1
}
var c = traced(_c)
//var c = _c
    
func main() {
    //c("kevin")
    //x:=_c(2,3)
    //print(x)
    //v:=reflect.ValueOf(_c)
    //print(v.Type())
    //print(v.Kind())
    x := c(3, 4)
    print("x=", x)
}