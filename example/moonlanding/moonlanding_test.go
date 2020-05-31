package moonlanding

import (
	"errors"
	"testing"

	"database.example/models/users"
)

func TestGoodMoonLanding(t *testing.T) {

	mock := users.SetupMock()

	mock.RegisterResultCallback("FindByEmail", func() (*users.User, error) {
		return &users.User{ID: "7", Email: "test@test.com", Password: "somehash"}, nil
	})

	moonlandingLog := MakeMoonGreatAgain()

	expect := "test@test.com is landing on the moon\nAfter 4 hours landing finished\n"

	if moonlandingLog != expect {
		t.Errorf("Was expecting to see moon landing but got '%s'", moonlandingLog)
	}

}

func TestBadMoonLanding(t *testing.T) {

	mock := users.SetupMock()

	mock.RegisterResultCallback("FindByEmail", func() (*users.User, error) {
		return nil, errors.New("Can't land because of reasons")
	})

	moonlandingLog := MakeMoonGreatAgain()

	expect := "Houstong we have a problem Can't land because of reasons\nDamage!\n"

	if moonlandingLog != expect {
		t.Errorf("Was expecting to see crash but got '%s'", moonlandingLog)
	}

}
