package main

import "fmt"
var print = fmt.Println

func makeEvenGenerator() func() uint {
    i := uint(0)
    return func() (ret uint) { //wtf ret uint?
        ret = i  //wtf?
        i += 2
        return
    }
}

func main() {
    nextEven := makeEvenGenerator()
    print(nextEven()) // 0
    print(nextEven()) // 2
    print(nextEven()) // 4
}
