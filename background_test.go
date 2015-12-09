package vscale

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestLocations(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/locations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		response := `
    [
      {
        "rplans": [
          "small",
          "medium",
          "large",
          "huge",
          "monster"
        ],
        "id": "spb0",
        "active": true,
        "description": "Цветочная 21",
        "private_networking": false,
        "templates": [
          "debian_8.1_64_001_master",
          "centos_7.1_64_001_master",
          "ubuntu_14.04_64_002_master"
        ]
      }
    ]`

		fmt.Fprint(w, response)
	})

	locations, _, err := client.Background.Locations()
	if err != nil {
		t.Errorf("Background.Locations returned error: %v", err)
	}

	expected := &[]Location{
		{
			ID:                "spb0",
			Active:            true,
			Description:       "\u0426\u0432\u0435\u0442\u043e\u0447\u043d\u0430\u044f 21",
			Rplans:            []string{"small", "medium", "large", "huge", "monster"},
			Templates:         []string{"debian_8.1_64_001_master", "centos_7.1_64_001_master", "ubuntu_14.04_64_002_master"},
			PrivateNetworking: false,
		},
	}

	if !reflect.DeepEqual(locations, expected) {
		t.Errorf("Background.Locations returned %+v, expected: %+v", locations, expected)
	}
}

func TestImages(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/images", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		response := `[{"rplans": ["small", "medium", "large", "huge", "monster"], "active": true, "size": 2048, "locations": ["spb0"], "id": "debian_8.1_64_001_master", "description": "Debian_8.1_64_001_master"}, {"rplans": ["small", "medium", "large", "huge", "monster"], "active": true, "size": 2048, "locations": ["spb0"], "id": "centos_7.1_64_001_master", "description": "Centos_7.1_64_001_master"}, {"rplans": ["small", "medium", "large", "huge", "monster"], "active": true, "size": 2048, "locations": ["spb0"], "id": "ubuntu_14.04_64_002_master", "description": "Ubuntu_14.04_64_002_master"}]`
		fmt.Fprint(w, response)
	})

	images, _, err := client.Background.Images()
	if err != nil {
		t.Errorf("Background.Images returned error: %v", err)
	}
	expected := &[]Image{
		{
			ID:          "debian_8.1_64_001_master",
			Active:      true,
			Description: "Debian_8.1_64_001_master",
			Rplans:      []string{"small", "medium", "large", "huge", "monster"},
			Locations:   []string{"spb0"},
			Size:        2048,
		},
		{
			ID:          "centos_7.1_64_001_master",
			Active:      true,
			Description: "Centos_7.1_64_001_master",
			Rplans:      []string{"small", "medium", "large", "huge", "monster"},
			Locations:   []string{"spb0"},
			Size:        2048,
		},
		{
			ID:          "ubuntu_14.04_64_002_master",
			Active:      true,
			Description: "Ubuntu_14.04_64_002_master",
			Rplans:      []string{"small", "medium", "large", "huge", "monster"},
			Locations:   []string{"spb0"},
			Size:        2048,
		},
	}

	if !reflect.DeepEqual(images, expected) {
		t.Errorf("Background.Images returned %+v, expected: %+v", images, expected)
	}
}
