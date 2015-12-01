package vscale

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAccountGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/account", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		response := `
    {
      "info": {
          "actdate": "2015-07-07 08:16:47.107987",
          "country": "",
          "email": "username@example.com",
          "face_id": "1",
          "id": "1001",
          "locale": "ru",
          "middlename": "MiddleName",
          "mobile": "+79901234567",
          "name": "UserName",
          "state": "1",
          "surname": "SurName"
      }
    }`
		fmt.Fprint(w, response)

	})

	acct, _, err := client.Account.Get()
	if err != nil {
		t.Errorf("Account.Get returned error: %v", err)
	}

	expected := &Account{
		ActivateDate: "2015-07-07 08:16:47.107987",
		Country:      "",
		FaceID:       1,
		ID:           1001,
		State:        1,
		Email:        "username@example.com",
		Name:         "UserName",
		MiddleName:   "MiddleName",
		SurName:      "SurName",
	}

	if !reflect.DeepEqual(acct, expected) {
		t.Errorf("Account.Get returned %+v, expected %+v", acct, expected)
	}
}
