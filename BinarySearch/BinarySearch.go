// 二分查找
// 从命令行接受一组空格分割的白名单整型数字,然后接受一个要查询是否在白名单中的数字

package BinarySearch

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	strconv2 "strconv"
	"strings"
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
	files := os.Args[1:]
	whiteList, blackList := readfile(files)
	for _, black := range blackList {
		if rank(black, whiteList) < 0 {
			//fmt.Println("该选项不在白名单中: %d", black)

		}
	}
}

// 将 []string 转化为 []int
func sliceatoi(sa []string) ([]int, error) {
	// 初始化一个存放所有白名单的切片
	si := make([]int, 0, len(sa))
	for _, nu := range sa {
		num, err := strconv2.Atoi(nu)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			return si, err
		}
		si = append(si, num)
	}
	return si, nil
}

// 从文件中读取数据
func readfile(files []string) ([]int, []int) {
	//　未从命令行中传递文件名
	var whitelist, blacklist []int
	var errw, errb error
	if len(files) == 0 {
		fmt.Println("请输入需要读取的文件名")
		os.Exit(1)
	} else {
		whitelistfile := files[0]
		blacklistfile := files[1]
		fmt.Println(whitelistfile, blacklistfile)
		// 读取文件
		f, err := ioutil.ReadFile(whitelistfile)
		f1, err1 := ioutil.ReadFile(blacklistfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read whitefile: %v\n", err)
		} else if err1 != nil {
			fmt.Fprintf(os.Stderr, "read blackfile %v\n", err1)
		} else {
			// if can read blacklistfile and whitelistfile then convert them to []int
			// convert byte to string
			whiteFile := string(f)
			// convert byte to []string And remove whitespace and \n
			strArry := strings.Fields(whiteFile)
			// sorted []string
			sort.Strings(strArry)
			//  convert []string to []int
			whitelist, errw = sliceatoi(strArry)
			if errw != nil {
				fmt.Fprintf(os.Stderr, "convert file whitelist %v\n", err)
			}
			// read blacklist file
			blackfile := string(f1)
			strarry1 := strings.Fields(blackfile)
			// sorted strarry1
			sort.Strings(strarry1)
			blacklist, errb = sliceatoi(strarry1)
			if errb != nil {
				fmt.Fprintf(os.Stderr, "convert  blacklist:%v\n", err1)
			}
		}
	}
	return whitelist, blacklist
}
