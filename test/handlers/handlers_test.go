package handlersTest

import (
	"testing"

	"../../main/handlers"
)

func TestGetBlogs(t *testing.T) {

	actualResults := handlers.GetMessage()
	if actualResults != "abcd" {
		t.Errorf("TestGetBlogs failed, actual: %s, expected: %s.", actualResults, "abcd")
	}
}
