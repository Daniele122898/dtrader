package bc

func CountRune(b []byte, d rune) int  {
	count := 0
	c := uint8(d)
	for _, r := range b {
		if r == c { // count numbers of '{' for example
			count++
		}
	}
	return count
}
