package functions

import (
	"fmt"
	"math"
	"sort"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums3 := append(nums1, nums2...)
	sort.Ints(nums3)
	fmt.Println(nums3)
	length := len(nums1) + len(nums2)
	if length%2 == 0 {
		fmt.Println((nums3[length/2-1] + nums3[length/2]))
		return float64(float64(nums3[length/2-1]+nums3[length/2]) / 2)
	} else {
		return float64(nums3[int(math.Floor(float64(length/2)))])
	}

}
