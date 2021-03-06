package array_sort
//康拓展开
func getPermutation(n int, k int) string {
	var result1,nums []int
	for i:=0;i<n;i++{
		nums=append(nums,i+1)
	}
	findP(nums,k,n,&result1)
	var result []byte
	for i:=0;i<len(result1);i++{
		result=append(result, byte(result1[i]+'0'))
	}
	return string(result)
}

func findP(nums []int, k int,n int, result1 *[]int){
	if len(nums)==1{
		*result1=append(*result1,nums[0])
		return
	}
	c:=1
	manyk:=0
	for i:=1;i<n;i++{ //求出n-1的阶乘也就是c的值来确定取值位置，举个例子，k=31, n=5, 那么n-1的阶乘是4X3X2X1, 那么在n这层的取值位置就是manyk=(31-1)/24=1
		c=c*i
	}
	manyk=(k-1)/c     //后面同理，再举例来说n-1这轮的取值位置是(31-24-1)/6=1，注意通过(k-1)/(n-1)!来确定取值位置时k要减1，一层层下去就直接定位了目标组合
	*result1=append(*result1,nums[manyk])
	var zz []int
	for j:=0;j<len(nums);j++{
		if j!=manyk{
			zz=append(zz,nums[j])
		}
	}
	findP(zz,k-manyk*c,len(zz),result1)
}
