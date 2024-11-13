package main

import (
	"fmt"
	"sync"
	"time"
)
type intLock struct {
	val int
	sync.Mutex
}

func (n *intLock) isEven() bool {
	return n.val%2 == 0
}
func main() {
	n := &intLock{val: 0}

	go func() {
		n.Lock()
		defer n.Unlock()
		nIsEven := n.isEven()
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n.val, " is even")
			return
		}
		fmt.Println(n.val, "is odd")
	}()

	go func() {
		n.Lock()
		n.val++
		n.Unlock()
	}()

	time.Sleep(time.Second)
}