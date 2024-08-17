package spt

import "testing"

func Test_FetchEnv(t *testing.T) {
	type TestType = []struct {
		envVar       string
		defaultValue string
		outputValue  string
	}

	data := TestType{
		{"MY_SECRET", "1234", "1234"},
		{"MY_TOKEN",  "qwerty", "666"},
	}

	t.Setenv("MY_TOKEN", "666")

	for _, item := range data {
		got := FetchEnv(item.envVar, item.defaultValue)
		want := item.outputValue

		if got != want {
			t.Errorf("Got %s, wanted %s", got, want)
		}
	}
}
