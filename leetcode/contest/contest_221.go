package contest

import "container/heap"

//5637. 判断字符串的两半是否相似
func halvesAreAlike(s string) bool {
	tagretMap := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	left := 0
	right := len(s) - 1
	leftCount := 0
	rightCount := 0
	for left < right {
		if tagretMap[s[left]] {
			leftCount++
		}
		if tagretMap[s[right]] {
			rightCount++
		}
		left++
		right--
	}
	return leftCount == rightCount
}

//优先队列
type Data struct {
	Date int
	Num  int
}
type queue []Data

func (q queue) Len() int {
	return len(q)
}
func (q queue) Less(i, j int) bool {
	return q[i].Date < q[j].Date
}
func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
func (q *queue) Push(x interface{}) {
	*q = append(*q, x.(Data))
}
func (q *queue) Pop() interface{} {
	ret := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return ret
}

//5638. 吃苹果的最大数目
func eatenApples(apples []int, days []int) (ans int) {
	q := queue{}
	for i := 0; i < len(apples) || len(q) != 0; i++ {
		//添加
		if i < len(apples) && apples[i] != 0 {
			heap.Push(&q, Data{
				Date: i + days[i],
				Num:  apples[i],
			})
		}
		//清除过期
		for len(q) != 0 && q[0].Date == i {
			heap.Pop(&q)
		}
		//吃一个苹果
		if len(q) != 0 && q[0].Num > 0 {
			ans++
			q[0].Num--
			//清除数量为0
			for len(q) != 0 && q[0].Num == 0 {
				heap.Pop(&q)
			}
		}
	}
	return
}
