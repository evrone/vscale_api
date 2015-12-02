package vscale

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestScaletList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/scalets", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		response := `[
    {
      "hostname":"cs10087.vscale.ru",
      "locked":false,
      "location":"CiDK",
      "rplan":"monster",
      "name":"mytestserver",
      "active":true,
      "keys": [
        {
            "name": "somekeyname",
            "id": 16
        }
      ],
      "public_address":{
         "netmask":"255.255.255.0",
         "gateway":"95.213.191.1",
         "address":"95.213.191.70"
      },
      "status":"started",
      "made_from":"ubuntu_14.04_64_002_master",
      "ctid":10087,
      "private_address":{}
    },
    {
      "hostname":"cs10139.vscale.ru",
      "locked":false,
      "location":"spb0",
      "rplan":"medium",
      "name":"Maroon-Waffle",
      "active":true,
      "keys": [
        {
            "name": "somekeyname",
            "id": 16
        }
      ],
      "public_address":{
         "netmask":"255.255.255.0",
         "gateway":"95.213.191.1",
         "address":"95.213.191.121"
      },
      "status":"started",
      "made_from":"ubuntu_14.04_64_002_master",
      "ctid":10139,
      "private_address":{}
    },
    {
      "hostname":"cs10299.vscale.ru",
      "locked":false,
      "location":"spb0",
      "rplan":"large",
      "name":"Hollow-Star",
      "active":true,
      "keys": [
        {
            "name": "somekeyname",
            "id": 16
        }
      ],
      "public_address":{
         "netmask":"255.255.255.0",
         "gateway":"95.213.191.1",
         "address":"95.213.191.120"
      },
      "status":"started",
      "made_from":"ubuntu_14.04_64_002_master",
      "ctid":10299,
      "private_address":{}
    }]`

		fmt.Fprint(w, response)
	})

	sclt, _, err := client.Scalet.List()
	if err != nil {
		t.Errorf("Scalet.List returned error: %v", err)
	}

	expected := &[]Scalet{
		Scalet{
			Name:           "mytestserver",
			Hostname:       "cs10087.vscale.ru",
			Locked:         false,
			Location:       "CiDK",
			Rplan:          "monster",
			Active:         true,
			Keys:           []SSHKey{SSHKey{ID: 16, Name: "somekeyname"}},
			PublicAddress:  &ScaletAddress{Netmask: "255.255.255.0", Gateway: "95.213.191.1", Address: "95.213.191.70"},
			Status:         "started",
			MadeFrom:       "ubuntu_14.04_64_002_master",
			CTID:           10087,
			PrivateAddress: &ScaletAddress{Netmask: "", Gateway: "", Address: ""},
		},
		Scalet{
			Name:           "Maroon-Waffle",
			Hostname:       "cs10139.vscale.ru",
			Locked:         false,
			Location:       "spb0",
			Rplan:          "medium",
			Active:         true,
			Keys:           []SSHKey{SSHKey{ID: 16, Name: "somekeyname"}},
			PublicAddress:  &ScaletAddress{Netmask: "255.255.255.0", Gateway: "95.213.191.1", Address: "95.213.191.121"},
			Status:         "started",
			MadeFrom:       "ubuntu_14.04_64_002_master",
			CTID:           10139,
			PrivateAddress: &ScaletAddress{Netmask: "", Gateway: "", Address: ""},
		},
		Scalet{
			Name:           "Hollow-Star",
			Hostname:       "cs10299.vscale.ru",
			Locked:         false,
			Location:       "spb0",
			Rplan:          "large",
			Active:         true,
			Keys:           []SSHKey{SSHKey{ID: 16, Name: "somekeyname"}},
			PublicAddress:  &ScaletAddress{Netmask: "255.255.255.0", Gateway: "95.213.191.1", Address: "95.213.191.120"},
			Status:         "started",
			MadeFrom:       "ubuntu_14.04_64_002_master",
			CTID:           10299,
			PrivateAddress: &ScaletAddress{Netmask: "", Gateway: "", Address: ""},
		},
	}

	if !reflect.DeepEqual(sclt, expected) {
		t.Errorf("Scalet.List returned %+v, expected %+v", sclt, expected)
	}
}

func TestGetScaletByID(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/scalets/12345", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		response := `
    {
      "private_address": {},
      "active": true,
      "rplan": "small",
      "keys": [
        {
          "name": "MacBook",
          "id": 1234
        }
      ],
      "locked": false,
      "status": "started",
      "name": "Minimum-Temple",
      "created": "01.12.2015 08:42:42",
      "ctid": 12345,
      "hostname": "cs12345.vscale.io",
      "deleted": null,
      "made_from": "ubuntu_14.04_64_002_master",
      "public_address": {
        "address": "95.213.199.29",
        "netmask": "255.255.255.0",
        "gateway": "95.213.199.1"
      },
      "location": "spb0"
    }`

		fmt.Fprint(w, response)
	})

	sclt, _, err := client.Scalet.GetByID(12345)
	if err != nil {
		t.Errorf("Scalet.GetByID returned error: %v", err)
	}

	expected := &Scalet{
		Name:           "Minimum-Temple",
		Hostname:       "cs12345.vscale.io",
		Locked:         false,
		Location:       "spb0",
		Rplan:          "small",
		Active:         true,
		Keys:           []SSHKey{SSHKey{ID: 1234, Name: "MacBook"}},
		PublicAddress:  &ScaletAddress{Netmask: "255.255.255.0", Gateway: "95.213.199.1", Address: "95.213.199.29"},
		Status:         "started",
		MadeFrom:       "ubuntu_14.04_64_002_master",
		CTID:           12345,
		PrivateAddress: &ScaletAddress{Netmask: "", Gateway: "", Address: ""},
	}

	if !reflect.DeepEqual(sclt, expected) {
		t.Errorf("Scalet.GetByID returned %+v, expected %+v", sclt, expected)
	}
}

func TestScaletCreate(t *testing.T) {
	setup()
	defer teardown()

	createRequest := &ScaletCreateRequest{
		MakeFrom: "ubuntu_14.04_64_002_master",
		Rplan:    "medium",
		DoStart:  true,
		Name:     "New-Test",
		Keys:     []int{16},
		Location: "spb0",
	}

	mux.HandleFunc("/v1/scalets", func(w http.ResponseWriter, r *http.Request) {
		v := new(ScaletCreateRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, "POST")

		if !reflect.DeepEqual(v, createRequest) {
			t.Errorf("Request body: %+v, expected: %+v", v, createRequest)
		}

		response := `
{
    "status": "defined",
    "deleted": null,
    "public_address": {},
    "active": false,
    "location": "spb0",
    "locked": true,
    "hostname": "cs11533.vscale.io",
    "created": "20.08.2015 14:57:04",
    "keys": [
      {
          "name": "somekeyname",
          "id": 16
      }
    ],
    "private_address": {},
    "made_from": "ubuntu_14.04_64_002_master",
    "name": "New-Test",
    "ctid": 11,
    "rplan": "medium"
}`

		fmt.Fprint(w, response)
	})

	scalet, _, err := client.Scalet.Create(createRequest)
	if err != nil {
		t.Errorf("Scalet.Create returned error: %v", err)
	}

	expected := &Scalet{
		Name:     "New-Test",
		Hostname: "cs11533.vscale.io",
		Locked:   true,
		Location: "spb0",
		Rplan:    "medium",
		Active:   false,
		Keys: []SSHKey{
			SSHKey{
				Name: "somekeyname",
				ID:   16,
			},
		},
		PublicAddress:  &ScaletAddress{},
		Status:         "defined",
		MadeFrom:       "ubuntu_14.04_64_002_master",
		CTID:           11,
		PrivateAddress: &ScaletAddress{},
	}

	if !reflect.DeepEqual(scalet, expected) {
		t.Errorf("Scalet.Create returned %+v, expected %+v", scalet, expected)
	}
}

func TestScaletRestart(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/scalets/10299/restart", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")

		response := `
    {
      "active":true,
      "location":"spb0",
      "ctid":10299,
      "name":"Eternal-Hungry",
      "created":"20.07.2015 16:54:31",
      "public_address":{
         "gateway":"95.213.191.1",
         "netmask":"255.255.255.0",
         "address":"95.213.191.24"
      },
      "status":"started",
      "private_address":{},
      "rplan":"medium",
      "deleted":null,
      "hostname":"cs10299.vscale.io",
      "keys": [
        {
            "name": "somekeyname",
            "id": 16
        }
      ],
      "locked":false,
      "made_from":"ubuntu_14.04_64_002_master"
    }`

		fmt.Fprint(w, response)
	})

	sclt, _, err := client.Scalet.Restart(10299)
	if err != nil {
		t.Errorf("Scalet.Restart returned error: %v", err)
	}

	expected := &Scalet{
		Name:           "Eternal-Hungry",
		Hostname:       "cs10299.vscale.io",
		Locked:         false,
		Location:       "spb0",
		Rplan:          "medium",
		Active:         true,
		Keys:           []SSHKey{SSHKey{ID: 16, Name: "somekeyname"}},
		PublicAddress:  &ScaletAddress{Netmask: "255.255.255.0", Gateway: "95.213.191.1", Address: "95.213.191.24"},
		Status:         "started",
		MadeFrom:       "ubuntu_14.04_64_002_master",
		CTID:           10299,
		PrivateAddress: &ScaletAddress{Netmask: "", Gateway: "", Address: ""},
	}

	if !reflect.DeepEqual(sclt, expected) {
		t.Errorf("Scalet.GetByID returned %+v, expected %+v", sclt, expected)
	}
}

func TestScaletRebuild(t *testing.T) {
	setup()
	defer teardown()

	rebuildRequest := &ScaletRebuildRequest{
		ID: 15508,
	}

	mux.HandleFunc("/v1/scalets/15508/rebuild", func(w http.ResponseWriter, r *http.Request) {
		v := new(ScaletRebuildRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, "PATCH")

		response := `
    {
      "public_address":{
        "netmask":"255.255.255.0",
        "gateway":"95.213.194.1",
        "address":"95.213.194.37"
      },
      "made_from":"ubuntu_14.04_64_002_master",
      "rplan":"medium",
      "location":"spb0",
      "name":"Minimum-Windshield",
      "created":"25.09.2015 12:19:05",
      "active":true,
      "locked":false,
      "status":"started",
      "ctid":15508,
      "private_address":{},
      "deleted":null,
      "keys":[
        {
          "id":72,
          "name":"key"
        }
      ],
      "hostname":"cs15508.vscale.io"
    }`

		fmt.Fprint(w, response)
	})

	sclt, _, err := client.Scalet.Rebuild(rebuildRequest)
	if err != nil {
		t.Errorf("Scalet.Restart returned error: %v", err)
	}

	expected := &Scalet{
		Name:           "Minimum-Windshield",
		Hostname:       "cs15508.vscale.io",
		Locked:         false,
		Location:       "spb0",
		Rplan:          "medium",
		Active:         true,
		Keys:           []SSHKey{SSHKey{ID: 72, Name: "key"}},
		PublicAddress:  &ScaletAddress{Netmask: "255.255.255.0", Gateway: "95.213.194.1", Address: "95.213.194.37"},
		Status:         "started",
		MadeFrom:       "ubuntu_14.04_64_002_master",
		CTID:           15508,
		PrivateAddress: &ScaletAddress{Netmask: "", Gateway: "", Address: ""},
	}

	if !reflect.DeepEqual(sclt, expected) {
		t.Errorf("Scalet.GetByID returned %+v, expected %+v", sclt, expected)
	}
}

func TestScaletHalt(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/scalets/10299/stop", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")

		response := `
    {
      "status": "started",
      "deleted": null,
      "made_from": "ubuntu_14.04_64_002_master",
      "created": "27.08.2015 14:22:39",
      "private_address": {},
      "ctid": 10299,
      "keys": [
          {
              "name": "key",
              "id": 72
          }
      ],
      "location": "spb0",
      "hostname": "cs12669.vscale.io",
      "locked": false,
      "public_address": {
          "netmask": "255.255.255.0",
          "gateway": "95.213.199.1",
          "address": "95.213.199.48"
      },
      "rplan": "medium",
      "name": "Icy-Compass",
      "active": true
    }`

		fmt.Fprint(w, response)
	})

	sclt, _, err := client.Scalet.Halt(10299)
	if err != nil {
		t.Errorf("Scalet.Halt returned error: %v", err)
	}

	expected := &Scalet{
		Name:           "Icy-Compass",
		Hostname:       "cs12669.vscale.io",
		Locked:         false,
		Location:       "spb0",
		Rplan:          "medium",
		Active:         true,
		Keys:           []SSHKey{SSHKey{ID: 72, Name: "key"}},
		PublicAddress:  &ScaletAddress{Netmask: "255.255.255.0", Gateway: "95.213.199.1", Address: "95.213.199.48"},
		Status:         "started",
		MadeFrom:       "ubuntu_14.04_64_002_master",
		CTID:           10299,
		PrivateAddress: &ScaletAddress{Netmask: "", Gateway: "", Address: ""},
	}

	if !reflect.DeepEqual(sclt, expected) {
		t.Errorf("Scalet.Halt returned %+v, expected %+v", sclt, expected)
	}
}

func TestScaletStart(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/scalets/10299/start", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")

		response := `
    {
      "location": "spb0",
      "keys": [
          {
              "name": "key",
              "id": 72
          }
      ],
      "created": "27.08.2015 14:22:39",
      "hostname": "cs12669.vscale.io",
      "ctid": 10299,
      "status": "stopped",
      "deleted": null,
      "rplan": "medium",
      "name": "Icy-Compass",
      "made_from": "ubuntu_14.04_64_002_master",
      "active": true,
      "locked": false,
      "public_address": {
          "gateway": "95.213.199.1",
          "address": "95.213.199.48",
          "netmask": "255.255.255.0"
      },
      "private_address": {}
    }`

		fmt.Fprint(w, response)
	})

	sclt, _, err := client.Scalet.Start(10299)
	if err != nil {
		t.Errorf("Scalet.Start returned error: %v", err)
	}

	expected := &Scalet{
		Name:           "Icy-Compass",
		Hostname:       "cs12669.vscale.io",
		Locked:         false,
		Location:       "spb0",
		Rplan:          "medium",
		Active:         true,
		Keys:           []SSHKey{SSHKey{ID: 72, Name: "key"}},
		PublicAddress:  &ScaletAddress{Netmask: "255.255.255.0", Gateway: "95.213.199.1", Address: "95.213.199.48"},
		Status:         "stopped",
		MadeFrom:       "ubuntu_14.04_64_002_master",
		CTID:           10299,
		PrivateAddress: &ScaletAddress{Netmask: "", Gateway: "", Address: ""},
	}

	if !reflect.DeepEqual(sclt, expected) {
		t.Errorf("Scalet.Start returned %+v, expected %+v", sclt, expected)
	}
}

func TestScaletUpdatePlan(t *testing.T) {
	setup()
	defer teardown()

	updateRplanRequest := &ScaletUpdatePlanRequest{
		ID:    10299,
		Rplan: "monster",
	}

	mux.HandleFunc("/v1/scalets/10299/upgrade", func(w http.ResponseWriter, r *http.Request) {
		v := new(ScaletUpdatePlanRequest)
		if err := json.NewDecoder(r.Body).Decode(v); err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, "POST")

		respons := `
    {
      "hostname":"cs10299.vscale.ru",
      "locked":false,
      "location":"spb0",
      "rplan":"huge",
      "name":"MyServer",
      "active":true,
      "keys": [
        {
            "name": "somekeyname",
            "id": 16
        }
      ],
      "public_address":{
         "netmask":"255.255.255.0",
         "gateway":"95.213.191.1",
         "address":"95.213.191.120"
      },
      "status":"started",
      "made_from":"ubuntu_14.04_64_002_master",
      "ctid":10299,
      "private_address":{}
    }`
		fmt.Fprint(w, respons)
	})

	sclt, _, err := client.Scalet.UpdatePlan(updateRplanRequest)
	if err != nil {
		t.Errorf("Scalet.UpdatePlan returned error: %v", err)
	}

	expected := &Scalet{
		Name:           "MyServer",
		Hostname:       "cs10299.vscale.ru",
		Locked:         false,
		Location:       "spb0",
		Rplan:          "huge",
		Active:         true,
		Keys:           []SSHKey{SSHKey{ID: 16, Name: "somekeyname"}},
		PublicAddress:  &ScaletAddress{Netmask: "255.255.255.0", Gateway: "95.213.191.1", Address: "95.213.191.120"},
		Status:         "started",
		MadeFrom:       "ubuntu_14.04_64_002_master",
		CTID:           10299,
		PrivateAddress: &ScaletAddress{Netmask: "", Gateway: "", Address: ""},
	}
	if !reflect.DeepEqual(sclt, expected) {
		t.Errorf("Scalet.UpdatePlan returned %+v, expected %+v", sclt, expected)
	}
}
