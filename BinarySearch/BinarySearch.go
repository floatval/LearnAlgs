// 二分查找
// 从命令行接受一组空格分割的白名单整型数字,然后接受一个要查询是否在白名单中的数字

package BinarySearch

import (
	"fmt"
	"os"
	"sort"
	strconv2 "strconv"
)

func rank(key int, arr []int) int {
	// 初始化标志位
	lo := 0
	hi := len(arr)
	mid := lo + (hi-lo)/2

	// 进行二分查找
	for lo <= hi {
		hi = mid - 1
		if key > arr[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}
	// 所查找对象不存在
	return -1
}

func Output() {
	// TODO 将从命令行获取白名单改为从文本文件中获取白名单.
	// 通过命令行获取白名单并排序
	whiteListStr := os.Args[1:]
	sort.Strings(whiteListStr)

	// 将 []string 转化为 []int
	whiteList, err := sliceAtoi(whiteListStr)
	if err != nil {
		fmt.Println(whiteList, err)
	}

	// 判断命令行是否传入参数
	if len(whiteListStr) > 0 {
		var key int
		fmt.Println("请输入要查询的键值")
		fmt.Scanln(&key)
		if rank(key, whiteList) < 0 {
			fmt.Printf("输入的key值: %d,不在白名单中", key)
		} else {
			fmt.Printf("输入的 key 值: %d,在白名单中", key)
		}
	}
}

func sliceAtoi(sa []string) ([]int, error) {
	// 将 []string 转化为 []int
	// TODO 将从命令行接收白名单改为从一个文本中接收白名单

	// 初始化一个存放所有白名单的切片
	si := make([]int, 0, len(sa))
	for _, nu := range sa {
		num, err := strconv2.Atoi(nu)
		if err != nil {
			return si, err
		}
		si = append(si, num)
	}
	return si, nil
}
