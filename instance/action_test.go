package instance

import (
	"reflect"
	"testing"
)

func TestCreateAction(t *testing.T) {
	tables := []struct {
		v interface{}
		t Action
		e string
	}{
		{"test", &BaseAction{typeName: "test"}, ""},
		{1, nil, "Invalid Action type int"},
		{map[string]interface{}{"Type": "test"}, &BaseAction{typeName: "test"}, ""},
		{map[string]interface{}{"NoType": "test"}, nil, "Action does not have Type field"},
		{map[string]interface{}{"Type": 1}, nil, "Action Type has to be a string and not int"},
		{
			map[string]interface{}{
				"Type": "test",
				"Expect": map[string]interface{}{
					"Response": map[string]interface{}{
						"Body": "done",
					},
				},
			},
			&BaseAction{
				typeName: "test",
				expect: &Expectation{
					response: &ResponseExpectation{
						body: "done",
					},
				},
			},
			"",
		},
		{
			map[string]interface{}{
				"Type": "test",
				"Expect": map[string]interface{}{
					"Response": "done",
				},
			},
			&BaseAction{
				typeName: "test",
				expect: &Expectation{
					response: &ResponseExpectation{
						body: "done",
					},
				},
			},
			"",
		},
		{
			map[string]interface{}{
				"Type": "test",
				"Expect": map[string]interface{}{
					"Output": []string{
						"NOTICE: starting",
					},
				},
			},
			&BaseAction{
				typeName: "test",
				expect: &Expectation{
					output: []string{
						"NOTICE: starting",
					},
				},
			},
			"",
		},
	}

	for _, table := range tables {
		a, e := CreateAction(table.v)
		tt := table.t
		aType, ttType := reflect.TypeOf(a), reflect.TypeOf(tt)
		if aType != ttType {
			t.Errorf("The expected struct type of action is %s, instead %s returned",
				ttType, aType)
		}
		if e != nil {
			if table.e == "" {
				t.Errorf("Expected no error but instead error '%s' returned",
					e.Error())
			} else if e.Error() != table.e {
				t.Errorf("Expected error '%s' but instead error '%s' returned",
					table.e, e.Error())
			}
			continue
		}
		if table.e != "" {
			t.Errorf("Expected error '%s' but no error returned", table.e)
		}
		// we can cast to base action as every tested action should be that type
		ba, ok := a.(*BaseAction)
		if !ok {
			t.Error("The action is not a subclass of BaseAction")
		}
		btt := tt.(*BaseAction)
		// check type name
		if ba.typeName != btt.typeName {
			t.Errorf("The expected type name of action is %s, instead %s returned",
				ba.typeName, btt.typeName)
		}
		// check expectation
		if !reflect.DeepEqual(ba.expect, btt.expect) {
			t.Errorf("The expectation %s does not match expected %s",
				ba.expect, btt.expect)
		}
	}

}
