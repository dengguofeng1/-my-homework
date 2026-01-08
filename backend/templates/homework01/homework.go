package main

import (
	"fmt"
	"sync"
)

// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	// 用map来记录每个数字出现的次数
	countMap := make(map[int]int)

	// 遍历数组，统计每个数字出现的次数
	for i := 0; i < len(nums); i++ {
		// 如果这个数字已经在map中，次数加1
		countMap[nums[i]] = countMap[nums[i]] + 1
		fmt.Println(countMap)
	}

	// 再遍历map，找出只出现一次的数字
	for num, count := range countMap {
		if count == 1 {
			return num
		}
	}

	return 0
}

// 2. 回文数
// 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	// 负数不是回文数
	if x < 0 {
		return false
	}

	// 保存原始数字
	original := x
	// 用来存储反转后的数字
	reversed := 0

	// 反转数字
	for x > 0 {
		// 取出最后一位数字
		lastDigit := x % 10
		// 把这一位加到反转数字中
		reversed = reversed*10 + lastDigit
		// 去掉原数字的最后一位
		x = x / 10
	}

	// 比较原始数字和反转后的数字是否相等
	return original == reversed
}

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func IsValid(s string) bool {
	// 用切片来模拟栈
	stack := []rune{}

	// 遍历字符串中的每个字符
	for _, char := range s {
		// 如果是左括号，就放入栈中
		if char == '(' || char == '{' || char == '[' {

			stack = append(stack, char)
			fmt.Printf("append '%c' -> stack: ", char)
			for _, c := range stack {
				fmt.Printf("%c ", c)
			}
			fmt.Println()
		} else {
			// 如果是右括号，检查栈是否为空
			if len(stack) == 0 {
				return false
			}

			// 取出栈顶元素
			top := stack[len(stack)-1]
			fmt.Printf("pop '%c' -> stack: ", top)
			stack = stack[:len(stack)-1]
			for _, c := range stack {
				fmt.Printf("%c ", c)
			}
			fmt.Println()

			// 检查括号是否匹配
			if char == ')' && top != '(' {
				return false
			}
			if char == '}' && top != '{' {
				return false
			}
			if char == ']' && top != '[' {
				return false
			}
		}
	}

	// 最后栈应该是空的
	return len(stack) == 0
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	// 如果数组为空，返回空字符串
	if len(strs) == 0 {
		return ""
	}

	// 用第一个字符串作为基准
	prefix := strs[0]

	// 和后面的每个字符串比较
	for i := 1; i < len(strs); i++ {
		// 不断缩短前缀，直到当前字符串以这个前缀开头
		for len(prefix) > 0 {
			// 检查当前字符串是否以prefix开头
			if len(strs[i]) >= len(prefix) {
				isMatch := true
				for j := 0; j < len(prefix); j++ {
					if strs[i][j] != prefix[j] {
						isMatch = false
						break
					}
				}
				if isMatch {
					break
				}
			}
			// 如果不匹配，就把前缀缩短一个字符
			prefix = prefix[:len(prefix)-1]
		}

		// 如果前缀已经是空字符串了，直接返回
		if prefix == "" {
			return ""
		}
	}

	return prefix
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	// 从最后一位开始处理
	for i := len(digits) - 1; i >= 0; i-- {
		// 如果当前位小于9，直接加1后返回
		if digits[i] < 9 {
			digits[i] = digits[i] + 1
			return digits
		}
		// 如果当前位是9，变成0，继续处理前一位（进位）
		digits[i] = 0
	}

	// 如果所有位都是9，需要在最前面加一个1
	result := make([]int, len(digits)+1)
	result[0] = 1
	// 其他位都是0（默认值）
	return result
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	// 如果数组为空，返回0
	if len(nums) == 0 {
		return 0
	}

	// 慢指针，指向不重复元素应该放的位置
	i := 0

	// 快指针，用来遍历数组
	for j := 1; j < len(nums); j++ {
		// 如果当前元素和前一个不重复的元素不同
		if nums[j] != nums[i] {
			// 慢指针前进一位
			i = i + 1
			// 把当前元素放到慢指针的位置
			nums[i] = nums[j]
		}
	}

	// 返回不重复元素的个数（慢指针位置+1）
	return i + 1
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	// 如果没有区间，直接返回
	if len(intervals) == 0 {
		return intervals
	}

	// 先按照区间的起始位置排序（用冒泡排序，简单易懂）
	for i := 0; i < len(intervals); i++ {
		for j := 0; j < len(intervals)-1-i; j++ {
			if intervals[j][0] > intervals[j+1][0] {
				// 交换位置
				temp := intervals[j]
				intervals[j] = intervals[j+1]
				intervals[j+1] = temp
			}
		}
	}

	// 用来存储合并后的区间
	result := [][]int{}
	// 先把第一个区间放进去
	result = append(result, intervals[0])

	// 遍历剩余的区间
	for i := 1; i < len(intervals); i++ {
		// 获取结果中的最后一个区间
		last := result[len(result)-1]
		// 当前区间
		current := intervals[i]

		// 如果当前区间的起始位置小于等于最后一个区间的结束位置，说明有重叠
		if current[0] <= last[1] {
			// 合并区间：结束位置取两者中的较大值
			if current[1] > last[1] {
				result[len(result)-1][1] = current[1]
			}
		} else {
			// 没有重叠，直接添加当前区间
			result = append(result, current)
		}
	}

	return result
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	// 用map来记录每个数字和它的位置
	numMap := make(map[int]int)
	println(numMap)

	// 遍历数组
	for i := 0; i < len(nums); i++ {
		// 计算需要找的另一个数
		needed := target - nums[i]

		// 看看这个数是否在map中
		if index, found := numMap[needed]; found {
			// 找到了，返回两个数的位置
			return []int{index, i}
		}

		// 把当前数字和位置记录到map中
		numMap[nums[i]] = i
	}

	// 没找到，返回空
	return nil
}

func main() {
	// fmt.Println("=== 1. 只出现一次的数字 ===")
	// nums1 := []int{4, 1, 2, 1, 2}
	// fmt.Printf("输入: %v\n", nums1)
	// fmt.Printf("只出现一次的数字: %d\n\n", SingleNumber(nums1))

	fmt.Println("=== 2. 回文数 (使用协程) ===")
	palindromeChan := make(chan int)
	var wg sync.WaitGroup

	// 启动协程检查 1 到 100000 的回文数
	wg.Add(1)
	go func() {
		defer wg.Done()
		count := 0
		for i := 1; i <= 100000; i++ {
			if IsPalindrome(i) {
				count++
				if count <= 1000 { // 只打印前10个
					palindromeChan <- i
				}
			}
		}
		close(palindromeChan)
	}()

	//从 channel 接收数据

	for num := range palindromeChan {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
	wg.Wait()

	// 再启动协程统计总数
	wg.Add(1)
	go func() {
		defer wg.Done()
		count := 0
		for i := 1; i <= 100000; i++ {
			if IsPalindrome(i) {
				count++
			}
		}
		fmt.Printf("1 到 100000 中共有 %d 个回文数\n\n", count)
	}()
	wg.Wait()

	// fmt.Println("=== 3. 有效的括号 ===")
	// str := "()[]{}{]"
	// fmt.Printf("%s 是有效的吗? %v\n\n", str, IsValid(str))

	// fmt.Println("=== 4. 最长公共前缀 ===")
	// strs := []string{"flower", "flow", "flight"}
	// fmt.Printf("输入: %v\n", strs)
	// fmt.Printf("最长公共前缀: %s\n\n", LongestCommonPrefix(strs))

	// 	fmt.Println("=== 5. 加一 ===")
	// 	digits := []int{1, 2, 9}
	// 	fmt.Printf("输入: %v\n", digits)
	// 	fmt.Printf("加一后: %v\n\n", PlusOne(digits))

	// 	fmt.Println("=== 6. 删除有序数组中的重复项 ===")
	// 	nums2 := []int{1, 1, 2, 2, 3}
	// 	fmt.Printf("输入: %v\n", nums2)
	// 	length := RemoveDuplicates(nums2)
	// 	fmt.Printf("删除重复后的长度: %d\n", length)
	// 	fmt.Printf("结果数组: %v\n\n", nums2[:length])

	// 	fmt.Println("=== 7. 合并区间 ===")
	// 	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	// 	fmt.Printf("输入: %v\n", intervals)
	// 	fmt.Printf("合并后: %v\n\n", Merge(intervals))

	// fmt.Println("=== 8. 两数之和 ===")
	// nums3 := []int{2, 7, 11, 15}
	// target := 9
	// fmt.Printf("输入: %v, 目标: %d\n", nums3, target)
	// fmt.Printf("结果: %v\n", TwoSum(nums3, target))
}
