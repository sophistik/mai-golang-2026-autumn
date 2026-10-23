package tabledemo

func Sign(n int) string {
	switch {
	case n > 0:
		return "плюс"
	case n < 0:
		return "минус"
	default:
		return "ноль"
	}
}
