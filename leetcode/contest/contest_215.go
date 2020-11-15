package contest

import "sort"

//5601. 设计有序流
type OrderedStream struct {
	Data map[int]string //保存数据
	Max  int            //最大指针值
	Ptr  int            //当前指针
}

func Constructor(n int) OrderedStream {
	data := make(map[int]string)
	return OrderedStream{
		Data: data,
		Max:  n,
		Ptr:  1,
	}
}

func (this *OrderedStream) Insert(id int, value string) []string {
	this.Data[id] = value
	ret := []string{}
	if this.Ptr == id {
		for i := id; i <= this.Max; i++ {
			val, has := this.Data[i]
			if has {
				ret = append(ret, val)
			} else {
				this.Ptr = i
				break
			}
		}
	}
	return ret
}

//5603. 确定两个字符串是否接近
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	chars1 := []byte(word1)        //字母数组1
	chars2 := []byte(word2)        //字母数组2
	charCnt1 := make(map[byte]int) //字符计数map1
	charCnt2 := make(map[byte]int) //字符计数map2
	for i := 0; i < len(chars1); i++ {
		charCnt1[chars1[i]]++
		charCnt2[chars2[i]]++
	}
	cntList1 := []int{} //数量数组1
	cntList2 := []int{} //数量数组2
	for k, v1 := range charCnt1 {
		v2, has := charCnt2[k]
		if !has { //出现不存在字母直接返回
			return false
		}
		cntList1 = append(cntList1, v1)
		cntList2 = append(cntList2, v2)
	}
	sort.Ints(cntList1)
	sort.Ints(cntList2)
	for k := range cntList1 {
		if cntList1[k] != cntList2[k] {
			return false
		}
	}
	return true
}
