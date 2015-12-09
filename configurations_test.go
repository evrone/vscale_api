package vscale

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestRplans(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/rplans", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		response := `[{"memory": 512, "disk": 20480, "locations": ["spb0"], "network": 1024, "id": "small", "addresses": 1, "cpus": 1, "templates": ["debian_8.1_64_001_master", "centos_7.1_64_001_master", "ubuntu_14.04_64_002_master"]}, {"memory": 1024, "disk": 30720, "locations": ["spb0"], "network": 2048, "id": "medium", "addresses": 1, "cpus": 1, "templates": ["debian_8.1_64_001_master", "centos_7.1_64_001_master", "ubuntu_14.04_64_002_master"]}, {"memory": 2048, "disk": 40960, "locations": ["spb0"], "network": 3072, "id": "large", "addresses": 1, "cpus": 2, "templates": ["debian_8.1_64_001_master", "centos_7.1_64_001_master", "ubuntu_14.04_64_002_master"]}, {"memory": 4096, "disk": 61440, "locations": ["spb0"], "network": 4096, "id": "huge", "addresses": 1, "cpus": 2, "templates": ["debian_8.1_64_001_master", "centos_7.1_64_001_master", "ubuntu_14.04_64_002_master"]}, {"memory": 8192, "disk": 81920, "locations": ["spb0"], "network": 5120, "id": "monster", "addresses": 1, "cpus": 4, "templates": ["debian_8.1_64_001_master", "centos_7.1_64_001_master", "ubuntu_14.04_64_002_master"]}]`
		fmt.Fprint(w, response)
	})

	collection, _, err := client.Configurations.Rplans()
	if err != nil {
		t.Errorf("Configurations.Rplans returnes error: %v", err)
	}

	expected := &[]Rplan{
		{
			ID:        "small",
			Memory:    512,
			Disk:      20480,
			Locations: []string{"spb0"},
			Network:   1024,
			Addresses: 1,
			Cpus:      1,
			Templates: []string{"debian_8.1_64_001_master", "centos_7.1_64_001_master", "ubuntu_14.04_64_002_master"},
		},
		{
			ID:        "medium",
			Memory:    1024,
			Disk:      30720,
			Locations: []string{"spb0"},
			Network:   2048,
			Addresses: 1,
			Cpus:      1,
			Templates: []string{"debian_8.1_64_001_master", "centos_7.1_64_001_master", "ubuntu_14.04_64_002_master"},
		},
		{
			ID:        "large",
			Memory:    2048,
			Disk:      40960,
			Locations: []string{"spb0"},
			Network:   3072,
			Addresses: 1,
			Cpus:      2,
			Templates: []string{"debian_8.1_64_001_master", "centos_7.1_64_001_master", "ubuntu_14.04_64_002_master"},
		},
		{
			ID:        "huge",
			Memory:    4096,
			Disk:      61440,
			Locations: []string{"spb0"},
			Network:   4096,
			Addresses: 1,
			Cpus:      2,
			Templates: []string{"debian_8.1_64_001_master", "centos_7.1_64_001_master", "ubuntu_14.04_64_002_master"},
		},
		{
			ID:        "monster",
			Memory:    8192,
			Disk:      81920,
			Locations: []string{"spb0"},
			Network:   5120,
			Addresses: 1,
			Cpus:      4,
			Templates: []string{"debian_8.1_64_001_master", "centos_7.1_64_001_master", "ubuntu_14.04_64_002_master"},
		},
	}

	if !reflect.DeepEqual(collection, expected) {
		t.Errorf("Configurations.Rplans returned %+v, expected: %+v", collection, expected)
	}
}
