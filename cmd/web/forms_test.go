package main

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Has(t *testing.T) {
	// #1 failed case
	form := NewForm(nil)

	val := form.Has("whatever")
	if val {
		t.Error("Form shows has filed when it should not")
	}

	// #2 success case
	postedData := url.Values{}
	postedData.Add("username", "meet")

	form = NewForm(postedData)
	val = form.Has("username")
	if !val {
		t.Error("Shows form doesn't have field, but it should")
	}
}

func TestForm_Required(t *testing.T) {
	// #1 Failed case
	req := httptest.NewRequest("POST", "/endpoint", nil)
	req.PostForm = nil

	form := NewForm(req.PostForm)
	form.Required("username", "password", "email")
	if form.Valid() {
		t.Error("Form show valid when required fields are missing")
	}

	postedData := url.Values{}
	postedData.Add("username", "meet")
	postedData.Add("password", "meet")
	postedData.Add("email", "meet.soni@meetsoni.me")

	req = httptest.NewRequest("POST", "/endpoint", nil)
	req.PostForm = postedData
	form = NewForm(req.PostForm)
	form.Required("username", "password", "email")
	if !form.Valid() {
		t.Error("Form show invalid when required fields are already there")
	}
}

func TestForm_Check(t *testing.T) {
	form := NewForm(nil)
	form.Check(false, "password", "password cannot be blank")
	if form.Valid() {
		t.Error("Valid() return false it should be return true")
	}

	form = NewForm(nil)
	form.Check(true, "password", "password cannot be blank")
	if !form.Valid() {
		t.Error("Valid() return true it should be return false")
	}
}

func Test_Errors_Get(t *testing.T) {
	form := NewForm(nil)
	form.Check(false, "password", "password cannot be blank")
	s := form.Errors.Get("password")
	if len(s) == 0 {
		t.Error("Should have an error returned from get, but don't")
	}

	s = form.Errors.Get("username")
	if len(s) > 0 {
		t.Error("Should not have an error, but got one")
	}
}
