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
		{"test", BaseAction{typeName: "test"}, ""},
		{1, nil, "Invalid Action type"},
		{map[string]interface{}{"Type": "test"}, BaseAction{typeName: "test"}, ""},
		{map[string]interface{}{"NoType": "test"}, nil, "Action does not have Type field"},
		{map[string]interface{}{"Type": 1}, nil, "Action Type has to be a string"},
	}

	for _, table := range tables {
		a, e := CreateAction(table.v)
		at, tt := reflect.TypeOf(a), reflect.TypeOf(table.t)
		if at != tt {
			t.Errorf("The expected struct type of action is %s, instead %s returned", tt, at)
		}
		if e == nil {
			if table.e != "" {
				t.Errorf("Expected error '%s' but no error returned", table.e)
			}
			if a.Type() != table.t.Type() {
				t.Errorf("The expected type of action is %s, instead %s returned",
					table.t.Type(), a.Type())
			}
		} else {
			if e.Error() != table.e {
				t.Errorf("Expected error '%s' but instead error '%s' returned",
					e.Error(), table.e)
			}
		}
	}

}
