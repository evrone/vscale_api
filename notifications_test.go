package vscale

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNotificationsSettingsGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/billing/notify", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		response := `{"notify_balance": 0, "status": "ok"}`
		fmt.Fprint(w, response)
	})

	notify, _, err := client.Notifications.Get()
	if err != nil {
		t.Errorf("Notifications.Get returns error: %v", err)
	}

	expected := &NotificationsSettings{
		NotifyBalance: 0,
		Status:        "ok",
	}
	if !reflect.DeepEqual(notify, expected) {
		t.Errorf("Notifications.Get returned %+v, expected %+v", notify, expected)
	}
}

func TestSetNotifyBalance(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/billing/notify", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")

		response := `{"notify_balance": 10000, "status": "ok"}`
		fmt.Fprint(w, response)
	})

	notify, _, err := client.Notifications.SetNotifyBalance(10000)
	if err != nil {
		t.Errorf("Notifications.SetNotifyBalance returns error: %v", err)
	}

	expected := &NotificationsSettings{
		NotifyBalance: 10000,
		Status:        "ok",
	}
	if !reflect.DeepEqual(notify, expected) {
		t.Errorf("Notifications.SetNotifyBalance returned %+v, expected %+v", notify, expected)
	}
}
