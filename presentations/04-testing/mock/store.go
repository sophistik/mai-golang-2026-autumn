package mockdemo

// Интерфейс объявляет ПОТРЕБИТЕЛЬ — функция Greet (лекция 2).
// Благодаря этому тест может подставить фейк, не трогая продакшн-код.
type Store interface {
	Get(key string) (string, bool)
}

func Greet(s Store, key string) string {
	if v, ok := s.Get(key); ok {
		return "Привет, " + v
	}
	return "Привет, гость"
}
