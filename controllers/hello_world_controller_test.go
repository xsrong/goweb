package controllers

import "testing"

func TestHome(t *testing.T) {
	controller := HelloWorldController{}
	str, err, statusCode := controller.Home()
	if str != "Hello World!" {
		t.Errorf("expected string equal to: \"Hello World!\" but got \"%s\"\n", str)
	}
	if err != nil {
		t.Error(err)
	}
	if statusCode != 200 {
		t.Errorf("expected statusCode equal to: 200 but got %d\n", statusCode)
	}
}
