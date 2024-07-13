package problems

// BestTimeStock solves LeetCode problem
// 122. Best Time to Buy and Sell Stock II
func BestTimeStock(prices []int) int {
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		tod := prices[i]
		tom := prices[i+1]
		if tom > tod {
			profit += tom - tod
		}
	}
	return profit
}
