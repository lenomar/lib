package main

import "fmt"

type Flags uint

const (
	FlagUp           Flags = 1 << iota // is up [位移操作] <<: 向左移位,可以看做是 乘以2的几次方
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopback                       // is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point link
	FlagMulticast                      // supports multicast access capability
)

// [and操作] &: 都为1 才是1
func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}

// [取反] ^ :有一个为1  则 就是1  两个1 或者两个0  全是 0
// [标志位操作] &^ :  a &^ b  其实就是清除标记位
func TurnDown(v *Flags) {
	*v &^= FlagUp
}
func TurnDown1(v *Flags) {
	*v = *v &^ FlagUp
}
func TurnDown2(v *Flags) {
	*v = (*v ^ FlagUp) & FlagUp
}

// [or操作] |: 只要有一个是1 那么就是1
func SetBroadcast(v *Flags) {
	*v |= FlagBroadcast
}
func IsCast(v Flags) bool {
	return v&(FlagBroadcast|FlagMulticast) != 0
}

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("FlagUp: %v %b \n", FlagUp, FlagUp)                               // 1 1
	fmt.Printf("FlagBroadcast: %v %b \n", FlagBroadcast, FlagBroadcast)          // 2 10
	fmt.Printf("FlagLoopback: %v %b \n", FlagLoopback, FlagLoopback)             // 4 100
	fmt.Printf("FlagPointToPoint: %v %b \n", FlagPointToPoint, FlagPointToPoint) // 8 1000
	fmt.Printf("FlagMulticast: %v %b \n", FlagMulticast, FlagMulticast)          // 16 10000

	fmt.Printf("Flags: %v %b \n", v, v)           // 17 10001
	fmt.Printf("IsUp: %v %b %t\n", v, v, IsUp(v)) // 17 10001 true
	fmt.Println("TurnDown:")
	TurnDown(&v)
	fmt.Printf("IsUp: %v %b %t\n", v, v, IsUp(v)) // 16 10000 false
	fmt.Println("SetBroadcast:")
	SetBroadcast(&v)
	fmt.Printf("IsUp: %v %b %t\n", v, v, IsUp(v))     // 18 10010 false
	fmt.Printf("IsCast: %v %b %t\n", v, v, IsCast(v)) // 18 10010 true
	cast := v & (FlagBroadcast | FlagMulticast)
	fmt.Printf("Cast: %v %b %t\n", cast, cast, IsCast(cast)) // 18 10010 true
	// clear := (v ^ (FlagBroadcast | FlagLoopback)) & (FlagBroadcast | FlagLoopback)
	// fmt.Printf("Clear: %v %b %t\n", clear, clear, IsCast(clear)) // 16 10000 true 20 10100
}
