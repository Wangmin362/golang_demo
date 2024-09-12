package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/coin-change-ii/description/

// 问题转换：背包的容量为amount，物品为coins，其中每个物品可以取无限次，因此是一个完全背包问题
// 由于本题目是求组合为amount的组合数，因此是一个组合问题，并不是排列问题，因此1,2,1和2,1,1看成是一个组合
func change(amount int, coins []int) int {
	// dp[j]为容量为j的背包，从0..i个硬币当中可以装满容量为j背包的组合次数
	// dp[j] += dp[j - coins[i]]  例如当背包的容量为5时，那么dp[5]就等于
	//       物品         组合次数
	//        1           dp[4] // 如果我们已经知道了背包容量为4的组合次数，那么这个背包再装1，此时背包容量就是5，整好就是我们要求的dp[5]
	//        2           dp[3] // 如果我们已经知道了背包容量为3的组合次数，那么这个背包再装2，此时背包容量就是5，整好就是我们要求的dp[5]
	//        3           dp[2] // 如果我们已经知道了背包容量为2的组合次数，那么这个背包再装3，此时背包容量就是5，整好就是我们要求的dp[5]
	//        4           dp[1] // 如果我们已经知道了背包容量为1的组合次数，那么这个背包再装4，此时背包容量就是5，整好就是我们要求的dp[5]
	//        5           dp[0] // 如果我们已经知道了背包容量为0的组合次数，那么这个背包再装5，此时背包容量就是5，整好就是我们要求的dp[5]
	// 最终结果为dp[amount]
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
		fmt.Println(dp)
	}

	return dp[amount]
}

func change02(amount int, coins []int) int {
	// dp[j] += dp[j-coins[i]]
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}

	return dp[amount]
}

// 题目分析：背包的容量为amount, 物品可以取无数次，因此是完全背包问题
// 明确定义：dp[j]表示容量为j的背包，使用coins可以凑成j的组合数
// 转移方程：dp[j] += dp[j-nums[i]]  也就是总数为j的背包，若当前由1元硬币，那么由dp[j-1]个数量，若由2元硬币，那么有dp[j-2]个数量
// 初始化：dp[j] = 1
// 遍历顺序：先物品，后容量，容量从小到大，每个物品可以取多次
// dp数组大小：amount+1
// 返回值：dp[amount]
func change0912(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}

	return dp[amount]
}

func TestChange(t *testing.T) {
	var testdata = []struct {
		amount int
		coins  []int
		want   int
	}{
		{amount: 5, coins: []int{1, 2, 5}, want: 4},
	}
	for _, tt := range testdata {
		get := change0912(tt.amount, tt.coins)
		if get != tt.want {
			t.Fatalf("coins:%v, amount:%v, want:%v, get:%v", tt.coins, tt.amount, tt.want, get)
		}
	}
}
