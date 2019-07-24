package easy

func judgeCircle(moves string) bool {
	row, col := 0, 0
	for _, v := range moves {
		if v == 'L' {
			row--
		} else if v == 'R' {
			row++
		} else if v == 'U' {
			col++
		} else if v == 'D' {
			col--
		}
	}
	return row == 0 && col == 0
}
