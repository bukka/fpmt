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
	}

	for _, table := range tables {
		a, _ := CreateAction(table.v)
		if reflect.TypeOf(a) != reflect.TypeOf(table.t) {
			t.Errorf("The expected struct type of action is %T, instead %T returned", table.t, a)
		}
		if a.Type() != table.t.Type() {
			t.Errorf("The expected type of action is %s, instead %s returned", table.t.Type(), a.Type())
		}
	}

}
