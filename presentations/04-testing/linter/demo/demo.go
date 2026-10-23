package demo

// Пакет с нарушениями для демонстрации линтера noprint.
func Hello() {
	println("привет") // ⚠️ noprint: встроенный println — только для отладки
	print("мир")      // ⚠️ и print тоже
}
