package main

import (
	"log"
	"time"
)

func main() {
	bigSlowOperation()
	double(4)
	log.Println(triple(4))
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")()

	time.Sleep(10 * time.Second)

}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)

	return func() { log.Printf("exit %s (%s)", msg,time.Since(start) ) }
}

func double( x int) (result int) {
	defer func() { log.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

func triple(x int) ( result int) {
	defer func() {result += x} ()
	return double(x)
}