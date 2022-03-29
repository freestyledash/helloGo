package algorithm

import (
	"fmt"
	"strconv"
)

type Student struct {
}

func TestMap() {
	m := make(map[string]Student)
	fmt.Println(m["123"])
}

func 判断数组中数据是否重复(nums []int) bool {

	var tempMap = make(map[string]int)
	for _, item := range nums {
		itemString := strconv.Itoa(item)
		if tempMap[itemString] == 0 {
			tempMap[itemString] = 1
		} else {
			return true
		}
	}
	return false
}

func 两数的交集(nums1 []int, nums2 []int) []int {
	tempMap1 := make(map[int]int)
	for _, item := range nums1 {
		if tempMap1[item] == 0 {
			tempMap1[item] = 1
		} else {
			tempMap1[item] = tempMap1[item] + 1
		}
	}

	tempMap2 := make(map[int]int)
	for _, item := range nums2 {
		if tempMap2[item] == 0 {
			tempMap2[item] = 1
		} else {
			tempMap2[item] = tempMap2[item] + 1
		}
	}

	int64s := []int{}

	for k, time1 := range tempMap1 {
		time2 := tempMap2[k]
		var min int
		if time1 > time2 {
			min = time2
		} else {
			min = time1
		}
		for i := 0; i < min; i++ {
			int64s = append(int64s, k)
		}

	}
	return int64s
}

func HaveAMapTest() {

	ints := []int{1, 5, -2, -4, 0}
	b := 判断数组中数据是否重复(ints)
	fmt.Println(b)

}
