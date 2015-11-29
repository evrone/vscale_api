package vscale

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAccount(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/account",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `{
  "info": {
      "actdate": "2015-07-07 08:16:47.107987",
      "country": "",
      "email": "username@domain.ru",
      "face_id": "1",
      "id": "1001",
      "locale": "ru",
      "middlename": "\u042E\u0437\u0435\u0440\u043E\u0432\u0438\u0447",
      "mobile": "+79901234567",
      "name": "\u041F\u043E\u043B\u044C\u0437\u043E\u0432\u0430\u0442\u0435\u043B\u044C",
      "state": "1",
      "surname": "\u0412\u0441\u043A\u0430\u043B\u0435\u0442\u043E\u0432"
  }
}`)
		},
	)

	account, err := client.Account()

	if err != nil {
		t.Errorf("Account returned error: %v", err)
	}

	want := `{
  "info": {
      "actdate": "2015-07-07 08:16:47.107987",
      "country": "",
      "email": "username@domain.ru",
      "face_id": "1",
      "id": "1001",
      "locale": "ru",
      "middlename": "\u042E\u0437\u0435\u0440\u043E\u0432\u0438\u0447",
      "mobile": "+79901234567",
      "name": "\u041F\u043E\u043B\u044C\u0437\u043E\u0432\u0430\u0442\u0435\u043B\u044C",
      "state": "1",
      "surname": "\u0412\u0441\u043A\u0430\u043B\u0435\u0442\u043E\u0432"
  }
}`
	if want != account {
		t.Errorf("Expected: %+v real: %+v", want, account)
	}
}
