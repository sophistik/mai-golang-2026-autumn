package mockdemo

import "testing"

// Фейк — ручная реализация интерфейса, живёт рядом с тестом.
// Рабочая упрощённая версия: мапа вместо настоящей БД.
type fakeStore map[string]string

func (f fakeStore) Get(key string) (string, bool) {
	v, ok := f[key]
	return v, ok
}

func TestGreet(t *testing.T) {
	tests := []struct {
		name  string
		store fakeStore
		key   string
		want  string
	}{
		{"имя найдено", fakeStore{"user": "Аня"}, "user", "Привет, Аня"},
		{"нет ключа", fakeStore{}, "user", "Привет, гость"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greet(tt.store, tt.key); got != tt.want {
				t.Errorf("Greet = %q, want %q", got, tt.want)
			}
		})
	}
}
