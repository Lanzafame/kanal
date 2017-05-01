package kanal_test

import (
	"reflect"
	"testing"

	"github.com/lanzafame/kanal"
)

func TestLoadConfig(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *kanal.Config
	}{
		{"basic config",
			args{
				"./fixtures/kanal.toml",
			},
			&kanal.Config{
				Sources: map[string]kanal.SourceConfig{
					"github": kanal.SourceConfig{
						Name: "github",
						URL:  "http://github.com/",
					},
					"trello": kanal.SourceConfig{
						Name: "trello",
						URL:  "http://trello.com/",
					},
				},
				Stores: map[string]kanal.StoreConfig{
					"postgresql": kanal.StoreConfig{
						Name: "postgresql",
						URL:  "psql://kanal",
					},
					"boltdb": kanal.StoreConfig{
						Name: "boltdb",
						URL:  "bltdb://kanal",
					},
				},
				Mappings: map[string]kanal.MappingConfig{
					"github": kanal.MappingConfig{
						Stores:         []string{"postgresql"},
						Transformation: "nil",
					},
					"trello": kanal.MappingConfig{
						Stores:         []string{"boltdb"},
						Transformation: "nil",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kanal.LoadConfig(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
