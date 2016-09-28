package reqid

import "math/rand"

func randomId(len int) string {
	var id = ""

	for i := len; i >= 0; i-- {
		id = id + string(CHARS[rand.Intn(64)])
	}

	return id
}
