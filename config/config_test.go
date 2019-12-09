package config

import "testing"

func TestReadConfig(t *testing.T) {
	_, err := ReadConfig("")
	if err == nil {
		t.Fatal("not implemented")
	}
	conf, err := ReadConfig("test_assets/test.toml")
	if err != nil {
		t.Fatal("not implemented")
	}
}
