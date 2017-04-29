package kanal_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/lanzafame/kanal"
)

func TestNewConfig(t *testing.T) {
	conf := kanal.NewConfig()
	conftype := strings.Split(reflect.TypeOf(conf).String(), ".")
	if conftype[1] != "Config" {
		t.Errorf("NewConfig() = %v", conftype)
	}
}

//func TestConfig_LoadConfig_DefaultConfigFile(t *testing.T) {
//	expected := []string{"test", "first", "second"}
//	conf := kanal.NewConfig()
//	conf.LoadConfig()
//	for i, key := range conf.Keys() {
//		if expected[i] != key {
//			t.Errorf("LoadConfig(): expected %v, got %v", key, expected[i])
//		}
//	}
//}
//
//func TestConfig_LoadConfig_NonDefaultConfigFile(t *testing.T) {
//	expected := []string{"test", "first", "second"}
//	conf := kanal.NewConfig()
//	conf.LoadConfig("nonkanal.toml")
//	for i, key := range conf.Keys() {
//		if expected[i] != key {
//			t.Errorf("LoadConfig(): expected %v, got %v", key, expected[i])
//		}
//	}
//}
