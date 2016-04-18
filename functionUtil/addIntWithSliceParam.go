package functionUtil

func Add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

