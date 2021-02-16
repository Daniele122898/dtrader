package bc

// Counts the number of times a rune appears in a byte slice. O(n)
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

// Counts the number of times a string appears in a byte slice.  O(n)
func CountString(b []byte, s string) int {
	count := 0
	n := len(s)
	rb := make([]uint8, 0, n)
	// translate the string runes to uints
	for i, r := range s {
		rb[i] = uint8(r)
	}
	// Start the counting
	l := len(b)
	for i := 0; i < l; i++ {
		r := b[i]
		// Now we can properly search the residual values.
		if r == rb[0] {
			if l - i < n {
				break
			}
			// We have enough headroom to check the next ones
			for j := 1; j<n; j++ {
				r2 := b[i+j]
				if r2 != rb[j] {
					i += j - 1
					break
				}
				// We have reached the last rune and it is equal
				// thus we have matched the entire string and can increase
				// the counter
				if j == n-1 {
					count++
					i += j-1
					break
				}
			}

		}
	}
	return count
}
