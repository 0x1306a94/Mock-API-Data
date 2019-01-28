package config

import "testing"

func TestConfig_Load(t *testing.T) {

	path := "../cmd/conf.yaml"
	conf, err := Load(path)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("\nMockAddr: %v \nMockPort: %v \nDashboardAddr: %v \nDashboardPort: %v \nDBPath: %v",
			conf.MockAddr,
			conf.MockPort,
			conf.DashboardAddr,
			conf.DashboardPort,
			conf.DBPath)
	}
}
