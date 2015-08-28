package monitor_test

import (
	"testing"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/influxdb/influxdb/monitor"
)

func TestConfig_Parse(t *testing.T) {
	// Parse configuration.
	var c monitor.Config
	if _, err := toml.Decode(`
store-enabled=true
store-database="the_db"
store-interval="10m"
store-address="server1"
`, &c); err != nil {
		t.Fatal(err)
	}

	// Validate configuration.
	if !c.StoreEnabled {
		t.Fatalf("unexpected store-enabled: %s", c.StoreEnabled)
	} else if c.StoreDatabase != "the_db" {
		t.Fatalf("unexpected store-database: %s", c.StoreDatabase)
	} else if time.Duration(c.StoreInterval) != 10*time.Minute {
		t.Fatalf("unexpected store-interval:  %s", c.StoreInterval)
	} else if c.StoreAddress != "server1" {
		t.Fatalf("unexpected store-address: %s", c.StoreAddress)
	}
}
