package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"sync"
	"time"
)

type List struct {
	Val  int
	Next *List
}

func rmListNum(head *List) *List {
	if head == nil || head.Next == nil {
		return nil
	}

	cur := head
	var pre *List
	for cur != nil {
		pre = cur
		for pre.Next != nil {
			if cur.Val == pre.Next.Val {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
			}
		}
		cur = cur.Next
	}

	return head
}

func print(head *List) {
	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

//找出最长连续数字串输出并返回长度，如调用maxnumstr("123abc1234a", outputstr)后返回4且outputstr中为"1234"

func findnum(numstr string) {
	n := len(numstr)

	var tmplen int = 0
	var end int = 0
	var res int = 0
	for i := 0; i < n; i++ {

		if numstr[i] >= '0' && numstr[i] <= '9' {
			tmplen++
			if res < tmplen {
				res = tmplen
				end = i
			}
		} else {
			tmplen = 0
		}

	}
	fmt.Println(res)
	fmt.Println(numstr[end-res+1 : end+1])
}

func findKNum(nums []int, k int) int {
	left, right := 0, len(nums)-1

	for {
		index := patition(nums, left, right)
		if index == k-1 {
			return nums[index]
		} else if index < k-1 {
			left = index + 1
		} else {
			right = index - 1
		}
	}
}

func patition(nums []int, left int, right int) int {
	key := nums[left]
	for left < right {

		for left < right && nums[right] >= key {
			right--
		}
		nums[left] = nums[right]

		for left < right && nums[left] <= key {
			left++
		}

		nums[right] = nums[left]
	}

	nums[left] = key
	return left

}

/*
任意子数组和的绝对值的最大值
给你一个整数数组 nums 。一个子数组 [numsl, numsl+1, ..., numsr-1, numsr] 的 和的绝对值 为 abs(numsl + numsl+1 + ... + numsr-1 + numsr) 。
请你找出 nums 中 和的绝对值 最大的任意子数组（可能为空），并返回该 最大值 。
abs(x) 定义如下：
如果 x 是负整数，那么 abs(x) = -x 。
如果 x 是非负整数，那么 abs(x) = x 。

示例 1：
输入：nums = [1,-3,2,3,-4]
输出：5
解释：子数组 [2,3] 和的绝对值最大，为 abs(2+3) = abs(5) = 5 。
示例 2：
输入：nums = [2,-5,1,-4,3,-2]
输出：8
*/

type tree struct {
	root  *tree
	left  *tree
	right *tree
	val   int
}

func findWidthofTree(root *tree) int {
	if root == nil {
		return 0
	}

	var maxwidth int
	queue := []*tree{}
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)

		if length > maxwidth {
			maxwidth = length
		}
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.left != nil {
				queue = append(queue, node.left)
			}

			if node.right != nil {
				queue = append(queue, node.right)
			}

		}
	}

	return maxwidth
}

//使用N个goroutine交替打印出1-N, 每个goroutine只能打印1个数值，比如N=3， 打印123123123123....

func getWorker(waitCh chan int, symbol int, wg *sync.WaitGroup) (next chan int) {
	notify := make(chan int)
	wg.Add(1)
	go func(waitCh chan int) {
		defer func() {
			wg.Done()
		}()
		for d := range waitCh {
			if d >= 100 {
				break
			}
			fmt.Println("goroutine:", symbol, "print", symbol)
			notify <- d + 1
		}
		close(notify)
		fmt.Println("goroutine: finish", symbol)
	}(waitCh)
	return notify
}

func demo1() {
	wg := new(sync.WaitGroup)
	start := make(chan int)
	lastCh := start
	for i := 0; i < 3; i++ {
		lastCh = getWorker(lastCh, i+1, wg)
	}
	start <- 0
	for v := range lastCh {
		start <- v
	}
	close(start)
	wg.Wait()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var pre *ListNode
	var cur *ListNode
	cur = head

	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	return pre
}

func printList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

// func DeepCopy(input interface{}) interface{} {
// 	if input == nil {
// 		return nil
// 	}

// 	switch reflect.TypeOf(input).Kind() {

// 	case reflect.Bool, reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
// 		return input

// 	case reflect.Struct:
// 		in := reflect.ValueOf(input)
// 		out := reflect.New(in.Type()).Elem()

// 		for i := 0; i < in.NumField(); i++ {
// 			out.Field(i).Set(DeepCopy(in.Field(i).Interface()))
// 		}
// 		return out.Interface()

// 	case reflect.Array, reflect.Slice:
// 		in := reflect.ValueOf(input)
// 		out := reflect.MakeSlice(in.Type(), in.Len(), in.Cap())

// 		for i := 0; i < in.Len(); i++ {
// 			out.Index(i).Set(DeepCopy(in.Index(i).Interface()))
// 		}
// 		return out.Interface()

// 	case reflect.Map:
// 		in := reflect.ValueOf(input)
// 		out := reflect.MakeMapWithSize(in.Type(), in.Len())

// 		for _, key := range in.MapKeys() {
// 			out.SetMapIndex(DeepCopy(key.Interface()).(reflect.Value), DeepCopy(in.MapIndex(key).Interface()).(reflect.Value))
// 		}
// 		return out.Interface()

// 	default:
// 		panic(fmt.Sprintf("Unable to deepcopy object of type %v", reflect.TypeOf(input)))
// 	}
// }

// Copy creates a deep copy of whatever is passed to it and returns the copy
// in an interface{}.  The returned value will need to be asserted to the
// correct type.

type Interface interface {
	DeepCopy() interface{}
}

// Iface is an alias to Copy; this exists for backwards compatibility reasons.
func Iface(iface interface{}) interface{} {
	return Copy(iface)
}
func Copy(src interface{}) interface{} {
	if src == nil {
		return nil
	}

	// Make the interface a reflect.Value
	original := reflect.ValueOf(src)

	// Make a copy of the same type as the original.
	cpy := reflect.New(original.Type()).Elem()

	// Recursively copy the original.
	copyRecursive(original, cpy)

	// Return the copy as an interface.
	return cpy.Interface()
}

// copyRecursive does the actual copying of the interface. It currently has
// limited support for what it can handle. Add as needed.
func copyRecursive(original, cpy reflect.Value) {
	// check for implement deepcopy.Interface
	if original.CanInterface() {
		if copier, ok := original.Interface().(Interface); ok {
			cpy.Set(reflect.ValueOf(copier.DeepCopy()))
			return
		}
	}

	// handle according to original's Kind
	switch original.Kind() {
	case reflect.Ptr:
		// Get the actual value being pointed to.
		originalValue := original.Elem()

		// if  it isn't valid, return.
		if !originalValue.IsValid() {
			return
		}
		cpy.Set(reflect.New(originalValue.Type()))
		copyRecursive(originalValue, cpy.Elem())

	case reflect.Interface:
		// If this is a nil, don't do anything
		if original.IsNil() {
			return
		}
		// Get the value for the interface, not the pointer.
		originalValue := original.Elem()

		// Get the value by calling Elem().
		copyValue := reflect.New(originalValue.Type()).Elem()
		copyRecursive(originalValue, copyValue)
		cpy.Set(copyValue)

	case reflect.Struct:
		t, ok := original.Interface().(time.Time)
		if ok {
			cpy.Set(reflect.ValueOf(t))
			return
		}
		// Go through each field of the struct and copy it.
		for i := 0; i < original.NumField(); i++ {
			// The Type's StructField for a given field is checked to see if StructField.PkgPath
			// is set to determine if the field is exported or not because CanSet() returns false
			// for settable fields.  I'm not sure why.  -mohae
			if original.Type().Field(i).PkgPath != "" {
				continue
			}
			copyRecursive(original.Field(i), cpy.Field(i))
		}

	case reflect.Slice:
		if original.IsNil() {
			return
		}
		// Make a new slice and copy each element.
		cpy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
		for i := 0; i < original.Len(); i++ {
			copyRecursive(original.Index(i), cpy.Index(i))
		}

	case reflect.Map:
		if original.IsNil() {
			return
		}
		cpy.Set(reflect.MakeMap(original.Type()))
		for _, key := range original.MapKeys() {
			originalValue := original.MapIndex(key)
			copyValue := reflect.New(originalValue.Type()).Elem()
			copyRecursive(originalValue, copyValue)
			copyKey := Copy(key.Interface())
			cpy.SetMapIndex(reflect.ValueOf(copyKey), copyValue)
		}

	default:
		cpy.Set(original)
	}
}

func addListNum(headA, headB *ListNode) (head *ListNode) {
	var tail *ListNode
	sum, lastSum := 0, 0

	for headA != nil || headB != nil {
		n1, n2 := 0, 0
		if headA != nil {
			n1 = headA.Val
			headA = headA.Next
		}

		if headB != nil {
			n2 = headB.Val
			headB = headB.Next
		}

		sum = n1 + n2 + lastSum
		sum, lastSum = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}

	if lastSum > 0 {
		tail.Next = &ListNode{Val: lastSum}
	}

	return head
}

func getsubstr(s string) int {
	size := len(s)

	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, size)
		dp[i][i] = 1
	}

	for i := size - 1; i >= 0; i-- {
		for j := i + 1; j < size; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
			fmt.Print(dp[i][j])
		}

	}
	return dp[0][size-1]
}

func max(a, b int) int {

	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {

	if a > b {
		return b
	} else {
		return a
	}
}

func deleteKNumofList(head *ListNode, k int) *ListNode {
	if head == nil || k < 0 {
		return head
	}
	dummyNode := &ListNode{Next: head}
	slow, fast := dummyNode, dummyNode

	for fast.Next != nil && k != 0 {
		fast = fast.Next
		k--
		fmt.Println(fast.Val)
		//fmt.Println(k)
	}

	fmt.Println(k)

	if k > 0 {
		return head
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return dummyNode.Next

}

func reverse(head *ListNode) {
	if head == nil {
		return
	}

	var pre *ListNode
	var cur *ListNode = head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	//return pre
}

func reverse2(head *ListNode, start int, end int) *ListNode {
	dummyNode := &ListNode{Next: head}
	var pre *ListNode = dummyNode

	for i := 0; i < start; i++ {
		pre = pre.Next
		fmt.Println(pre.Val)
	}

	rightNode := pre

	for j := 0; j < end-start+1; j++ {
		rightNode = rightNode.Next
		fmt.Println(rightNode.Val)
	}

	leftNode := pre.Next
	cur := rightNode.Next

	pre.Next = nil
	rightNode.Next = nil

	reverse(leftNode)

	pre.Next = rightNode
	leftNode.Next = cur

	return dummyNode.Next

}

// //Longest Substring Without Repeating Characters
// //Input: s = "abcabcbb"
// Output: 3

// Input: s = "bbbbb"
// Output: 1

func SubstrWithoutRepeat(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}

	left, right, res := 0, 0, 0
	m := map[byte]int{}
	for right < length {
		if _, ok := m[s[right]]; !ok {
			m[s[right]] = right
		} else {
			if m[s[right]]+1 >= left {
				left = m[s[right]] + 1
			}
			m[s[right]] = right
		}

		res = max(res, right-left+1)
		right++
		//fmt.Println(s[left:right])
	}
	return res
}

//Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
//Output: [[1,6],[8,10],[15,18]]

//[-3,-1, 0, 1, 2, 4] 升序数组
//平方后升序输出，要求时间复杂度O（n）
//output:[0,1,1,4,9,16]

func CalNums(nums []int) []int {
	length := len(nums)
	res := make([]int, length)

	left, right, tail := 0, length-1, length-1
	for left <= right {
		leftnum, rightnum := nums[left]*nums[left], nums[right]*nums[right]
		var tmp int
		if leftnum > rightnum {
			tmp = leftnum
			left++
		} else {
			tmp = rightnum
			right--
		}

		res[tail] = tmp
		tail--

	}
	return res
}

func getLongestComStr(s1, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	dp[0][0] = 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

func getCap(nums []int) int {
	left, right := 0, len(nums)-1
	var res int
	for left <= right {
		//minHeight := min(nums[left], nums[right])
		tmp := min(nums[left], nums[right]) * (right - left)
		res = max(res, tmp)
		if nums[left] <= nums[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

/*给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），
并返回该子数组所对应的乘积。 测试用例的答案是一个 32-位 整数。 子数组 是数组的连续子序列。*/

func getMax(nums []int) int {
	maxNum, minNum := nums[0], nums[0]
	var res int = nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mi := maxNum, minNum
		maxNum = max(nums[i], max(mx*nums[i], mi*nums[i]))
		fmt.Println("max:", maxNum)
		minNum = min(nums[i], min(mx*nums[i], mi*nums[i]))
		fmt.Println("min:", minNum)
		res = max(res, maxNum)
	}
	return res
}

var wg sync.WaitGroup

func randNum(jobChan chan<- int64) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		jobChan <- rand.Int63()
	}
	close(jobChan)
}

func sumNum(jobChan <-chan int64, resultChan chan<- int64) {
	defer wg.Done()
	for v := range jobChan {
		var s int64
		for v > 0 {
			s += v % 10
			v = v / 10
		}
		resultChan <- s
	}
}

type Pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

func NewPool(size int) *Pool {
	if size <= 0 {
		size = 1
	}

	return &Pool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

func (p *Pool) Add(delta int) {
	for i := 0; i < delta; i++ {
		p.queue <- i
	}
	for i := 0; i > delta; i-- {
		<-p.queue
	}

	p.wg.Add(delta)
}

// Done 执行完成减一
func (p *Pool) Done(i int) {
	<-p.queue
	p.wg.Done()

}

func (p *Pool) Wait() {
	p.wg.Wait()
}

func Run() {
	pool := NewPool(10)
	fmt.Println("the NumGoroutine begin is:", runtime.NumGoroutine())
	for i := 0; i < 50; i++ {
		pool.Add(1)
		go func(i int) {
			time.Sleep(5 * time.Second)
			fmt.Println("the NumGoroutine continue is:", i, runtime.NumGoroutine())
			pool.Done(i)
		}(i)
	}

	pool.Wait()
	fmt.Println("the NumGoroutine done is:", runtime.NumGoroutine())
}

func swap01strs(s string) (res int, num []byte) {
	num = []byte(s)
	left, right := 0, len(num)-1
	for left < right {
		for left < right && num[left] == '0' {
			left++
		}
		for left < right && num[right] == '1' {
			right--
		}
		if left < right {
			tmp := num[left]
			num[left] = num[right]
			num[right] = tmp
			res++
			left++
			right--
		}

	}
	fmt.Println(string(num), res)
	return res, num
}

func change(s *[]string) {
	//修改string切片中的某一个元素，再通过指针赋值回原切片，会影响到外部的切片
	tmp := []byte((*s)[0])
	fmt.Println(tmp)
	tmp[0] = byte('g')
	tmp[5] = byte('9')
	(*s)[0] = string(tmp)
	(*s)[1] = "Go"
	*s = append(*s, "yes")
}

func bubbleSort(arr []int) {
	fmt.Println("排序前：", (arr))
	// 总结规律：先内层(每一轮)再外层,内层n-1-i次,外层n-1
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			temp := 0
			if (arr)[j] > (arr)[j+1] {
				temp = (arr)[j]
				(arr)[j] = (arr)[j+1]
				(arr)[j+1] = temp
			}
		}
	}
	arr[0] = 99
	arr = append(arr, 100)
	fmt.Println("排序后：", arr)
}

func main() {

}
