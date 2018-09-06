// 欧几里得算法求最大公约数
package gcd

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	// 第一种情况,两者相等或,两者中的一个可被另一个整除
	r := a % b
	if r == 0 {
		return b
	}

	// 第二种,二者无法整除
	return gcd(a, r)
}
