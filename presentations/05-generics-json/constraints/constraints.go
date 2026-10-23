package constraintsdemo

// Constraint — это интерфейс: перечисляем разрешённые типы.
type Number interface {
	~int | ~int64 | ~float64
}

func Sum[T Number](xs []T) T {
	var total T
	for _, v := range xs {
		total += v
	}
	return total
}

// type Celsius float64 подошёл бы с ~float64; без ~ — нет.
