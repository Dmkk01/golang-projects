package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"simple-api/user"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func TestBodyToUser(t *testing.T) {
	valid := &user.User{
		ID:   bson.NewObjectId(),
		Name: "Test User",
		Role: "Tester",
	}

	js, err := json.Marshal(valid)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	ts := []struct {
		txt string
		r   *http.Request
		u   *user.User
		err bool
		exp *user.User
	}{
		{
			txt: "nul request",
			err: true,
		},
		{
			txt: "empty request body",
			r:   &http.Request{},
			err: true,
		},
		{
			txt: "empty user",
			r: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBufferString(`{}`)),
			},
			err: true,
		},
		{
			txt: "malformed data",
			r: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBufferString(`{"id": 12}`)),
			},
			u:   &user.User{},
			err: true,
		},
		{
			txt: "valid data",
			r: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBuffer(js)),
			},
			u:   &user.User{},
			exp: valid,
		},
	}

	for _, tc := range ts {
		t.Log(tc.txt)
		err := bodyToUser(tc.r, tc.u)

		if tc.err {
			if err == nil {
				t.Error("Expected an error and got none")
			}
			continue
		}

		if err != nil {
			t.Error(err)
			continue
		}

		if !reflect.DeepEqual(tc.exp, tc.u) {
			t.Errorf("Expected %v, got %v", tc.exp, tc.u)
		}
	}
}
