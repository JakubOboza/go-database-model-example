package users

import "testing"

func TestFindByEmail(t *testing.T) {

	mock := SetupMock()

	mock.RegisterResultCallback("FindByEmail", func() (*User, error) {
		return &User{ID: "7", Email: "test@test.com", Password: "somehash"}, nil
	})

	user, err := Store.FindByEmail("test@test.com")

	if mock.CallsCount("FindByEmail") != 1 {
		t.Errorf("Was expecting to see 1 call but got %d", mock.CallsCount("FindByEmail"))
	}

	if err != nil {
		t.Errorf("Error should be nil for mock")
	}

	if user.ID != "7" {
		t.Errorf("Expected ID to be 7 but got %s", user.ID)
	}

	if user.Email != "test@test.com" {
		t.Errorf("Expected Email to be test2test.com but got %s", user.Email)
	}

}

func TestCreateUser(t *testing.T) {

	fooName := "CreateUser"

	mock := SetupMock()

	mock.RegisterResultCallback(fooName, func() (*User, error) {
		return &User{ID: "69", Email: "test@test.com", Password: "somehash"}, nil
	})

	user, err := Store.CreateUser("test@test.com", "somePassword")

	if mock.CallsCount(fooName) != 1 {
		t.Errorf("Was expecting to see 1 call but got %d", mock.CallsCount(fooName))
	}

	if err != nil {
		t.Errorf("Error should be nil for mock")
	}

	if user.ID != "69" {
		t.Errorf("Expected ID to be 69 but got %s", user.ID)
	}

	if user.Email != "test@test.com" {
		t.Errorf("Expected Email to be test2test.com but got %s", user.Email)
	}
}
