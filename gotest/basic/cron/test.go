package main

func Assign1(s []int) {
	s = []int{6, 6, 6}
}
func main() {
	/*for range time.Tick(time.Second){
		fmt.Println("周期性执行")
	}*/
	//	c := time.Tick(time.Second)

	/*for {
		<- c
		go func() {fmt.Println("--------------")}()
	}*/
	//情况一：切片传递了变量副本，变量副本地址也同样指向原地址，当给变量副本新的地址时，不会改变原来变量的地址
	/*s := []int{1, 2, 3, 4, 5, 6}
	Assign1(s)
	fmt.Println(s) // (1)*/

	//情况二：两个变量指向相同的地址，但是其中一个切片使用了append方法后，会指向另外新的地址，所以彼此之间不会相互影响
	/*s := []int{1, 2, 3} // len=3, cap=3
	a := s
	s[0] = 888
	//append操作会分配一个新的空间给s，所以不会修改a的值
	s = append(s, 4)

	fmt.Println(a, len(a), cap(a)) // 输出：[888 2 3] 3 3
	fmt.Println(s, len(s), cap(s)) // 输出：[888 2 3 4] 4 6

	fmt.Println(a, len(a), cap(a)) // 输出：[888 2 3] 3 3*/

}
