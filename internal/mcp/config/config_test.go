package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name     string
		env      map[string]string
		want     *Config
		wantErr  bool
	}{
		{
			name: "defaults",
			env:  map[string]string{},
			want: &Config{
				APIKey:    "",
				APISecret: "",
				Testnet:   false,
				Port:      "30001",
			},
		},
		{
			name: "testnet true",
			env:  map[string]string{"BYBIT_TESTNET": "true"},
			want: &Config{Testnet: true, Port: "30001"},
		},
		{
			name: "custom port",
			env:  map[string]string{"MCP_PORT": "8080"},
			want: &Config{Port: "8080"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.env {
				os.Setenv(k, v)
			}
			defer func() {
				for k := range tt.env {
					os.Unsetenv(k)
				}
			}()

			cfg := Load()
			if cfg.APIKey != tt.want.APIKey || cfg.Testnet != tt.want.Testnet || cfg.Port != tt.want.Port {
				t.Errorf("Load() = %+v, want %+v", cfg, tt.want)
			}
		})
	}
}
