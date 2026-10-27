package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestConfigLoading(t *testing.T) {
	cases := []struct {
		name, json string
		valid      bool
	}{
		{"valid", `{"app_name":"Test","default_faculty":"IT","default_course":2,"seed_students":true}`, true},
		{"invalid_json", `{`, false},
		{"missing_fields", `{}`, false},
		{"invalid_course", `{"app_name":"T","default_faculty":"IT","default_course":7}`, false},
		{"unknown_field", `{"app_name":"T","default_faculty":"IT","default_course":2,"typo":1}`, false},
		{"trailing_json", `{"app_name":"T","default_faculty":"IT","default_course":2} {}`, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "config.json")
			if err := os.WriteFile(path, []byte(tc.json), 0600); err != nil {
				t.Fatal(err)
			}
			cfg, err := Load(path)
			if (err == nil) != tc.valid {
				t.Fatalf("valid=%v err=%v", tc.valid, err)
			}
			if tc.valid && (cfg.AppName != "Test" || cfg.DefaultCourse != 2 || !cfg.SeedStudents) {
				t.Fatal("wrong config")
			}
		})
	}
	if _, err := Load(filepath.Join(t.TempDir(), "missing.json")); err == nil {
		t.Fatal("missing file accepted")
	}
}
