package users

import "testing"

import "time"

import "encoding/json"

func TestUserJsonSerialization(t *testing.T) {

	testTime, _ := time.Parse(
		time.RFC3339,
		"1985-01-31T22:20:13+00:00")

	user := &User{ID: "7", Email: "test@test.com", Password: "somehash", CreatedAt: testTime, UpdatedAt: time.Now()}

	jsonBytes, err := json.Marshal(user)

	if err != nil {
		t.Errorf("Expected to be able to serialize User to json bu got error %s", err)
	}

	jsonStr := string(jsonBytes)

	if jsonStr != "{\"id\":\"7\",\"email\":\"test@test.com\",\"created_at\":\"1985-01-31T22:20:13Z\"}" {
		t.Errorf("Expected serialized but got %s", jsonStr)
	}

}
