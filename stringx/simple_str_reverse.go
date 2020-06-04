package stringx

func rvsString(in string) (out string) {
	runes := []rune(in)
	ln := len(runes)

	for i := 0; i < ln/2; i++ {
		runes[i], runes[ln-1-i] = runes[ln-1-i], runes[i]
	}

	return string(runes)
}
