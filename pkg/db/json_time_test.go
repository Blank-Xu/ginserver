package db

import (
	"encoding/json"
	"testing"
	"time"
)

type Person struct {
	Birthday JSONTime `json:"birthday"`
}

const dataJson = `{"birthday":"2019-01-27 20:00:00"}`

func TestJSONTime_MarshalJSON(t *testing.T) {
	p := new(Person)
	p.Birthday = JSONTime(time.Now())
	js, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(js))
}

func TestJSONTime_UnmarshalJSON(t *testing.T) {
	p := new(Person)
	if err := json.Unmarshal([]byte(dataJson), p); err != nil {
		t.Fatal(err)
	}
	t.Log(p.Birthday)
}
