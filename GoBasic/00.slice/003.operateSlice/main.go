package main

import "fmt"

/*
1.s1 := slice[2:5]->截取slice左闭右开,索引2-5的区间元素,容量默认到slice最后,此时s1->[2,3,4]
并且s1[2]=20,所以是s1最终为->[2,3,20]。

2.s2 := s1[2:6:7]->截取s1左闭右开,索引2-6的区间元素,cap到开区间索引7(到index6),因为s1的底层
数组为slice,所以s2同样指向slice。此时s2->[4,5,6,7]。
s2两个append,最终s2->[4,5,6,7,100,200]

3.因为s1,s2指向的都是slice,根据上下文,当append 100发生时,s2还有容量,所以不会扩容,但是会改变底层数组的值,此时slice[8]=100,
当append 200发生时,s2进行了自动扩容,所以不会影响底层数组,最后s1[2]=20,同样也影响了底层数组slice相应的元素,
但是s2发生扩容后,远走高飞了,并不会影响s2。
*/

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:7]

	s2 = append(s2, 100)
	s2 = append(s2, 200)

	s1[2] = 20

	fmt.Println(s1)    //[2,3,20]
	fmt.Println(s2)    //[4,5,6,7,100,200]
	fmt.Println(slice) //[0 1 2 3 20 5 6 7 100 9]
}
