package user_center

import (
	"context"
	"testing"
	"time"
	"web/db/model"

	"github.com/segmentio/ksuid"
)

func TestGenerateID(t *testing.T) {
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second * 1)
		k := ksuid.New()
		t.Log(k.Value())
	}
}

func TestID(t *testing.T) {

}

func TestRegister(t *testing.T) {
	_, err := registerWithEmail(context.Background(), "1@qq.com", "321231", "111111")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGetUser(t *testing.T) {
	user := model.User{
		UserId: "22wwFYVrngG0P5EkOIowYuJAfQ2",
	}
	t.Log(getUser(context.Background(), user))
}
