/*
 * @Author: vingurzhou
 * @Date: 2022-07-10 22:02:57
 * @Last Modified by: vingurzhou
 * @Last Modified time: 2022-07-10 22:17:39
 */
package main

import (
	"fmt"
	"golangPractice/functions"
)

func main() {
	fmt.Println(functions.TwoSum([]int{2, 7, 11, 15}, 9)) //1

	l1 := &functions.ListNode{Val: 2}
	l1.Next = &functions.ListNode{Val: 4}
	l1.Next.Next = &functions.ListNode{Val: 3}
	l2 := &functions.ListNode{Val: 5}
	l2.Next = &functions.ListNode{Val: 6}
	l2.Next.Next = &functions.ListNode{Val: 4}
	fmt.Println(functions.AddTwoNumbers(l1, l2)) //2

	fmt.Println(functions.LengthOfLongestSubstring("abcabcbb")) //3

	fmt.Println(functions.FindMedianSortedArrays([]int{1, 2}, []int{3, 4})) //4

	fmt.Println(functions.LongestPalindrome("ccc")) //5

	fmt.Println(functions.Reverse(123)) //7
}
