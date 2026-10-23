package jsondemo

import "encoding/json"

type Session struct {
	Token string // экспортируется
	note  string // приватное — json его не видит
}

func MarshalSession() ([]byte, error) {
	s := Session{Token: "abc", note: "внутренняя заметка"}
	return json.Marshal(s) // {"Token":"abc"} — note пропало молча
}
