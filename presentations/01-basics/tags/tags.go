package tags

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email,omitempty"`
	Pass  string `json:"-"`
}
