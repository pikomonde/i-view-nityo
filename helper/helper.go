package helper

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"

	"github.com/pikomonde/i-view-nityo/model"
)

func HashUserPassword(user model.User) string {
	input := fmt.Sprintf("user::%s::created_at::%d::%s", user.Username, user.CreatedAt, user.Password)
	h := sha256.New()
	h.Write([]byte(input))
	return fmt.Sprintf("%x", h.Sum(nil))
}

const alphanum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var rnd = rand.NewSource(time.Now().UnixNano())

func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphanum[rnd.Int63()%int64(len(alphanum))]
	}
	return string(b)
}
