package helper_test

import (
	"testing"
	"time"

	"github.com/pikomonde/i-view-nityo/helper"
	"github.com/pikomonde/i-view-nityo/model"
	"github.com/stretchr/testify/assert"
)

func TestHashUserPassword(t *testing.T) {
	now := time.Now().UnixNano()
	user := model.User{
		Username:  "username",
		Password:  "password",
		CreatedAt: now,
	}
	helper.HashUserPassword(user)
}

func TestRandomString(t *testing.T) {
	str := helper.RandomString(10)
	assert.Equal(t, 10, len(str))
}
