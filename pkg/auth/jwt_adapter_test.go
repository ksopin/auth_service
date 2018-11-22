package auth

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags

	fmt.Println("Before run all the tests in this file")
	res := m.Run()
	fmt.Println("After run all the tests in this file")
	os.Exit(res)
}

func TestJwtAdapter(t *testing.T) {
	adapter := NewJwtAdapter()

	s, err := adapter.Generate(&Identity{1, "test"})
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	err = adapter.Validate(s)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	identity, err := adapter.Read(s)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if identity.Id != 1 {
		t.Error("unexpected identity Id")
	}

	if identity.Name != "test" {
		t.Error("unexpected identity name")
	}

}
