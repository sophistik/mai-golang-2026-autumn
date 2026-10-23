package fuzzdemo

// Reverse разворачивает строку по рунам: UTF-8 остаётся валидным.
// Наивный побайтовый реверс сломал бы многобайтовые символы —
// именно такие вещи и ловит фаззинг.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
