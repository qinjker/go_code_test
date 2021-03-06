package main

import (
	"fmt"
	"github.com/qinjker/newmath"
	"math"
	"math/rand"
)

type Vertex struct {
	x1 int
	y1 int
}

var quit chan int = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	quit <- 0
}

func main() {
	v1 := Vertex{2, 3}
	p1 := &v1
	p1.x1 = 1e9
	fmt.Println(v1)
	i, j := 42, 1000
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	p = &j
	*p = *p / 20
	fmt.Println(j)

	// 开两个goroutine跑函数loop, loop函数负责打印10个数
	go loop()
	go loop()

	for i := 0; i < 2; i++ {
		<-quit
	}
	fmt.Printf("sdfdsf12321sdfdsf313,Sqrt(2)=%v\n", newmath.Sqrt(2))
	fmt.Println("我的随机数", rand.Intn(10))

	fmt.Printf(" %g problems.\n", math.Nextafter(2, 3))

	fmt.Println(math.Pi)

	fmt.Printf("2+3的和为=%v\n", newmath.Add(2, 3))
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("sum = %v\n", sum)

	newmath.SelfPrint()

	fmt.Printf("递归算法 5的阶乘%v\n", newmath.Recur(5))
}
