package jsondemo

import (
	"encoding/json"
	"fmt"
)

type Duration struct {
	Seconds int
}

// MarshalJSON: в JSON уедет "90s", а не 90.
func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%ds", d.Seconds))
}

// UnmarshalJSON — зеркально: json.Unmarshal в string,
// strconv.Atoi; ошибки — через fmt.Errorf("...: %w", err).
