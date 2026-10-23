package jsondemo

import "encoding/json"

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email,omitempty"` // пустое — не писать
	Pass  string `json:"-"`               // в JSON не попадает
}

func Marshal() ([]byte, error) {
	u := User{ID: 42, Pass: "secret"}
	// {"id":42} — без pass и без пустого email
	return json.Marshal(u)
}
