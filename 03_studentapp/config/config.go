package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Config struct {
	AppName        string `json:"app_name"`
	DefaultFaculty string `json:"default_faculty"`
	DefaultCourse  int    `json:"default_course"`
	SeedStudents   bool   `json:"seed_students"`
}

// Load JSON файлын оқып, баптауларды тексереді.
func Load(path string) (Config, error) {
	var cfg Config
	file, err := os.Open(path)
	if err != nil {
		return cfg, fmt.Errorf("конфигурацияны ашу: %w", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&cfg); err != nil {
		return cfg, fmt.Errorf("JSON оқу: %w", err)
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return cfg, fmt.Errorf("файлда бір ғана JSON объектісі болуы керек")
	}
	cfg.AppName = strings.TrimSpace(cfg.AppName)
	cfg.DefaultFaculty = strings.TrimSpace(cfg.DefaultFaculty)
	if cfg.AppName == "" || cfg.DefaultFaculty == "" {
		return cfg, fmt.Errorf("app_name және default_faculty бос болмауы керек")
	}
	if cfg.DefaultCourse < 1 || cfg.DefaultCourse > 6 {
		return cfg, fmt.Errorf("default_course 1 мен 6 аралығында болуы керек")
	}
	return cfg, nil
}
