package config

import (
	"encoding/json"
	"os"
)

// DBConfig maps to the "db" section of local.json
type DBConfig struct {
  Host     string `json:"host"`
  Port     int    `json:"port"`
  User     string `json:"user"`
  Password string `json:"password"`
  DBName   string `json:"dbname"`
  SSLMode  string `json:"sslmode"`
}

// Config holds all app config sections
type Config struct {
  DB DBConfig `json:"db"`
}

// LoadConfig reads a JSON config file into a Config struct
func LoadConfig(path string) (*Config, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var cfg Config
  decoder := json.NewDecoder(file)
  if err := decoder.Decode(&cfg); err != nil {
    return nil, err
  }
  return &cfg, nil
}