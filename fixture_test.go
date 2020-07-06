package fixture

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	TestDir = "test"
)

type TestStruct struct {
	Name string `json:"name"`
}

func TestSave(t *testing.T) {
	ts := TestStruct{Name: "HERE"}

	err := Save(ts, TestDir, "save.json")
	require.NoError(t, err)
}

func TestRead(t *testing.T) {
	ts := TestStruct{Name: "HERE"}

	var comp TestStruct
	err := Read(&comp, TestDir, "read.json")
	require.NoError(t, err)

	if comp.Name != ts.Name {
		t.Fail()
	}
}
