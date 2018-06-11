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
		{1, nil, "Invalid Action type int"},
		{"test", nil, "Action type test does not have any mapping"},
		{
			"server-start",
			&ServerStartAction{
				BaseAction: BaseAction{
					typeName: "server-start",
					expect: &Expectation{
						output: []string{
							"{{date}} NOTICE: {{app}} is running, pid {{pid}}",
							"{{date}} NOTICE: ready to handle connections",
						},
					},
				},
			},
			"",
		},
		{
			map[string]interface{}{"Type": "server-start"},
			&ServerStartAction{
				BaseAction: BaseAction{
					typeName: "server-start",
					expect: &Expectation{
						output: []string{
							"{{date}} NOTICE: {{app}} is running, pid {{pid}}",
							"{{date}} NOTICE: ready to handle connections",
						},
					},
				},
			},
			"",
		},
		{map[string]interface{}{"NoType": "test"}, nil, "Action does not have Type field"},
		{map[string]interface{}{"Type": 1}, nil, "Action Type has to be a string and not int"},
		{
			map[string]interface{}{
				"Type": "server-start",
				"Expect": map[string]interface{}{
					"Response": map[string]interface{}{
						"Body": "done",
					},
				},
			},
			&ServerStartAction{
				BaseAction: BaseAction{
					typeName: "server-start",
					expect: &Expectation{
						response: &ResponseExpectation{
							body: "done",
						},
					},
				},
			},
			"",
		},
		{
			map[string]interface{}{
				"Type": "server-start",
				"Expect": map[string]interface{}{
					"Response": "done",
				},
			},
			&ServerStartAction{
				BaseAction: BaseAction{
					typeName: "server-start",
					expect: &Expectation{
						response: &ResponseExpectation{
							body: "done",
						},
					},
				},
			},
			"",
		},
		{
			map[string]interface{}{
				"Type": "server-start",
				"Expect": map[string]interface{}{
					"Response": "done",
				},
			},
			&ServerStartAction{
				BaseAction: BaseAction{
					typeName: "server-start",
					expect: &Expectation{
						response: &ResponseExpectation{
							body: "done",
						},
					},
				},
			},
			"",
		},
		{
			map[string]interface{}{
				"Type": "server-start",
				"Expect": map[string]interface{}{
					"Output": []string{
						"NOTICE: starting",
					},
				},
			},
			&ServerStartAction{
				BaseAction: BaseAction{
					typeName: "server-start",
					expect: &Expectation{
						output: []string{
							"NOTICE: starting",
						},
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
		// check action
		if !reflect.DeepEqual(a, tt) {
			t.Errorf("The action %s does not match expected %s", a, tt)
		}
	}

}
