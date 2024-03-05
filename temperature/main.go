package main

import "fmt"

func temprature(nums []int) []int {
	var length int = len(nums)
	var res []int = make([]int, length)
	for i := 0; i < len(nums); i++ {
		k := 0
		for j := i + 1; j < length; j++ {
			k++
			if nums[j] > nums[i] {
				res[i] = k
				break
			}
			if j == length {
				res[i] = 0
			}
		}
	}
	return res
}

func tempraturetwo(nums []int) []int {
	var length int = len(nums)
	var res []int = make([]int, length)
	k, left, right := 0, 0, 1
	for right < length {
		//right := left + 1
		k++
		if nums[right] > nums[left] {
			res[left] = k
			left++
			right = left + 1
			k = 0
			continue
		}
		right++
		//k++
	}

	return res
}

/*
如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为 摆动序列 。第一个差（如果存在的话）可能是正数或负数。仅有一个元素或者含两个不等元素的序列也视作摆动序列。

例如， [1, 7, 4, 9, 2, 5] 是一个 摆动序列 ，因为差值 (6, -3, 5, -7, 3) 是正负交替出现的。

相反，[1, 4, 7, 2, 5] 和 [1, 7, 4, 5, 5] 不是摆动序列，第一个序列是因为它的前两个差值都是正数，第二个序列是因为它的最后一个差值为零。
子序列 可以通过从原始序列中删除一些（也可以不删除）元素来获得，剩下的元素保持其原始顺序。

给你一个整数数组 nums ，返回 nums 中作为 摆动序列 的 最长子序列的长度 。


输入：nums = [1,7,4,9,2,5]
输出：6
解释：整个序列均为摆动序列，各元素之间的差值为 (6, -3, 5, -7, 3) 。


示例 2：

输入：nums = [1,17,5,10,13,15,10,5,16,8]
输出：7
解释：这个序列包含几个长度为 7 摆动序列。
其中一个是 [1, 17, 10, 13, 10, 16, 8] ，各元素之间的差值为 (16, -7, 3, -3, 6, -8) 。

示例 3：

输入：nums = [1,2,3,4,5,6,7,8,9]
输出：2


提示：

1 <= nums.length <= 1000
0 <= nums[i] <= 1000*/

func calnums(nums []int) int {
	length := len(nums)
	if length == 1 {
		return 1
	}
	if length == 2 {
		return 2
	}
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	var res = 0
	stack := make([]int, 0)
	for i := 0; i < length; i++ {
		tmp := 0
		for j := i + 1; j < length; j++ {
			mul := nums[j] - nums[j-1]
			stack = append(stack, mul)
			if len(stack) > 1 && mul*stack[len(stack)-1] > 0 {
				fmt.Println(stack[len(stack)-1], mul)
				tmp = j - i + 1
				break
			}
		}

		res = max(res, tmp)
	}
	return res
}

/*func calnums2(nums []int) int {
	length := len(nums)
	if length == 1 {
		return 1
	}
	if length == 2 {
		return 2
	}
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	var res = 0
	left, right := 0, 0
	stack := make([]int, 0)
	for right < length {
		tmp := nums[right] - nums[left]
		stack = append(stack, tmp)

		if len(stack) > 0 && tmp*stack[len(stack)-1] > 0 {
			sub := j - i + 1
			res = max(res, sub)
		}
		right++
		left++

	}
	return res
}*/

func mergeLists(left, right int, lists *[][]int) *[][]int {
	return nil
}

func mergeTwoList(l1, l2 []int) []int {
	length1, length2 := len(l1), len(l2)
	tmpList := make([]int, length1+length2)

	for len(l1) != 0 && len(l2) != 0 {
		num1 := l1[len(l1)-1]
		num2 := l2[len(l2)-1]
		fmt.Println(num1, num2)
		if num1 < num2 {
			tmpList = append(tmpList, num2)
			l2 = l2[:len(l2)-1]
		} else {
			tmpList = append(tmpList, num1)
			l1 = l1[:len(l1)-1]
		}
	}

	if len(l1) == 0 {
		tmpList = append(tmpList, l2...)
	}
	if len(l2) == 0 {
		tmpList = append(tmpList, l1...)
	}

	return tmpList
}

type List struct {
	Val  int
	Next *List
}

func rmListNum(head *List) *List {
	if head == nil || head.Next == nil {
		return nil
	}

	cur := head.Next
	var pre *List = nil
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

func main() {

	node5 := &List{Val: 5, Next: nil}
	node4 := &List{Val: 2, Next: node5}
	node3 := &List{Val: 2, Next: node4}
	node2 := &List{Val: 2, Next: node3}
	node1 := &List{Val: 1, Next: node2}
	print(rmListNum(node1))
	//t1 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	//fmt.Println(mergeTwoList(t1[0], t1[1]))
}
