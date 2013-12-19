package main

import (
	"encoding/json"
	"testing"
)

func TestValidateInput(t *testing.T) {

	jsonString := "{\"string_value\":\"hello world\", \"int_value\":23.23423412312, \"bool_value\":true, \"obj_value\": {\"string_value\":\"foo bar\", \"int_value\":3.14, \"bool_value\":false } }"

	var m map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &m)
	if err != nil {
		t.Error(err.Error())
	}

	t.Log(m)

	s := Session{}
	err = ValidateInput(m, &s)
	if err != nil {
		t.Error(err.Error())
	}

	t.Logf("%v", s)
}
