package main

func checkInclusion(s1 string, s2 string) bool {
	var c1 [26]byte
	var c2 [26]byte
	len1 := len(s1)
	len2 := len(s2)
	for i := 0; i < len1; i++ {
		c1[s1[i] - 'a']++
	}
	for i := 0; i < len2; i++ {
		if i >= len1 {
			c2[s2[i - len1] - 'a']--
		}
		c2[s2[i] - 'a']++
		c11 := c1[:]
		c22 := c2[:]
		if check(c11, c22){
			return true
		}
	}
	return false
}
func check(c1, c2 []byte) bool  {
	for i, v := range c1 {
		if v != c2[i] {
			return false
		}
	}
	return true
}
