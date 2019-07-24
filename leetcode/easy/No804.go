package easy

var rep = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	m := map[string]int{}
	for i := range words {
		var key string
		for _, v := range words[i] {
			key += rep[v-'a']
		}
		m[key]++
	}

	return len(m)
}
