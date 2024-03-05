package main

import (
	"errors"
	"fmt"
	"strings"
)

type operate func(x, y int) int

//type caculator func(x, y int) (int, error)

func add(x, y int) int {
	return x + y
}

//1. 函数作为参数
func caculate(x, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}

	return op(x, y), nil
}

func test1() {
	//匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量或者作为立即执行函数:
	op := func(x, y int) int {
		return x + y
	}
	op(2, 3)

	func(x, y int) int {
		return x + y
	}(5, 6)

	fmt.Println(caculate(2, 3, add))
}

//2. 高级函数，把其他的函数作为结果返回
//闭包：在一个函数中存在对外来标识符的引用。所谓的外来标识符，既不代表当前函数的任何参数或结果，也不是函数内部声明的，它是直接从外边拿过来的。
//闭包函数实例，gencaculator引用了外部的op函数
func gencaculator(op operate) func(int, int) (int, error) {
	return func(x, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

func test2() {
	add1 := gencaculator(add)
	res, err := add1(3, 4)
	fmt.Printf("The result: %d (error: %v)\n", res, err)
}

//闭包函数实例，makeSuffixFunc引用了外部的suffix变量
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func test3() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt
}

func partimes(x int) func(int) {
	fmt.Printf("start: the input message is: %v\n", x)
	return func(y int) {
		fmt.Println("result is : ", x*y)
	}
}

func test4() {
	time2 := partimes(2)
	defer time2(3)
	fmt.Println("end......")
}

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	res := nums[0]
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		res = max(res, dp[i])
	}
	return res
}

func test6(nums1, nums2 []int) {
	len1, len2 := len(nums1), len(nums2)
	for i := 0; i < len2; i++ {
		nums1 = append(nums1, 0)
	}
	fmt.Println(nums1)

	length1, length2 := len(nums1), len(nums2)

	index1, index2, index3 := 0, 0, len1

	for index1 < length1 && index2 < length2 {

		if nums1[index1] <= nums2[index2] && nums1[index1] != 0 {
			index1++
		} else if nums1[index1] > nums2[index2] && nums1[index1] != 0 {
			tmp := nums1[index1]
			nums1[index1] = nums1[index3]
			nums1[index3] = tmp
			nums1[index1] = nums2[index2]

			index1++
			index2++
			index3++
		} else {
			nums1[index1] = nums2[index2]
			index1++
			index2++
		}

	}

	fmt.Println(nums1)

}

func reverselist(head *Listnode) {

	pre, cur := nil, head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	//test5()
	nums1 := []int{1, 2, 5, 7, 8, 12}
	nums2 := []int{2, 5, 6, 7, 9}
	test6(nums1, nums2)
}
