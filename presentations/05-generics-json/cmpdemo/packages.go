package cmpdemo

import (
	"cmp"
	"maps"
	"slices"
)

func demo() {
	cmp.Compare(1, 2) // -1, 0 или 1
	xs := []int{3, 1, 2}
	slices.Sort(xs) // Ordered — функция не нужна
	_ = slices.Contains(xs, 2)
	m := map[string]int{"go": 1, "c++": 2}
	// мапа не даёт порядка — сортируем ключи:
	keys := slices.Sorted(maps.Keys(m))
	_ = keys
}
