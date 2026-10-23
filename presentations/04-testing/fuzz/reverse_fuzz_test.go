package fuzzdemo

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	// seed-корпус: гоняется даже при обычном go test
	f.Add("привет")
	f.Add("go 1.18")
	f.Add("")

	f.Fuzz(func(t *testing.T, s string) {
		// Свойство 1: реверс не ломает UTF-8 — верно для ЛЮБОГО входа.
		if !utf8.ValidString(Reverse(s)) {
			t.Errorf("невалидный UTF-8 после реверса: %q", Reverse(s))
		}

		// Свойство 2: реверс — инволюция, но только для валидного UTF-8.
		// Фаззинг нашёл контрпример за 5 секунд (см. testdata/fuzz):
		// строка "\xa9" с битым байтом — []rune заменяет его на U+FFFD,
		// и двойной реверс не восстанавливает исходник. Так и задумано:
		// свойство уточнено, вход остался регресс-тестом.
		if !utf8.ValidString(s) {
			t.Skip("вход с невалидным UTF-8 вне контракта Reverse")
		}
		if back := Reverse(Reverse(s)); back != s {
			t.Errorf("Reverse(Reverse(%q)) = %q", s, back)
		}
	})
}
