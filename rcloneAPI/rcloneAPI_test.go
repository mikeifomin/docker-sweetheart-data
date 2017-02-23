package rcloneAPI

import (
	// "fmt"
	"io/ioutil"
	"path"
	"testing"
)

func TestExtractProviderConf(t *testing.T) {
	cases := []struct {
		input                string
		provider, key, value string
	}{
		{"B2_KEY=value", "b2", "key", "value"},
		{"B2_KEY_NAME=value", "b2", "key_name", "value"},
		{"DRIVE_KEY=value", "drive", "key", "value"},
		{"B2_KEY=Value", "b2", "key", "Value"},
	}

	for _, c := range cases {
		provider, key, value := extractProviderConf(c.input)
		if provider != c.provider {
			t.Fatalf("expect %s, got %s", c.provider, provider)
		}
		if key != c.key {
			t.Fatalf("expect %s, got %s", c.key, key)
		}
		if value != c.value {
			t.Fatalf("expect %s, got %s", c.value, value)
		}
	}
}
func TestWriteConfig(t *testing.T) {
	cases := []struct {
		input    map[string][]string
		expected string
	}{
		{map[string][]string{"b2": []string{"key=value"}}, "[b2]\nkey=value"},
	}
	for _, c := range cases {
		dir, _ := ioutil.TempDir("", "test")
		filename := path.Join(dir, ".rclone.conf")
		err := writeConfig(c.input, filename)
		if err != nil {
			t.Fatalf("cant write config to %s: %s", filename, err)
		}
		contain, _ := ioutil.ReadFile(filename)
		actual := string(contain)
		if actual != c.expected {
			t.Fatalf("wrong writes: expect %s, actual %s", c.expected, actual)
		}
	}

}
