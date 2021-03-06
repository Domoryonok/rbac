package rbac

import (
	"encoding/json"
	"testing"
)

func TestPermission(t *testing.T) {
	R := New(nil) //NewConsoleLogger()
	crudActions := []Action{Create, Read, Update, Delete}
	usersPerm, err := R.RegisterPermission("users", "User resource", crudActions...)
	if err != nil {
		t.Fatalf("can not register users permission, err: %v", err)
	}
	if len(usersPerm.Actions()) != len(crudActions) {
		t.Fatalf("user permission actions are not valid, expected %d items, got %d items", len(crudActions), len(usersPerm.Actions()))
	}

	_, err = json.Marshal(usersPerm)
	if err != nil {
		t.Fatalf("users json marshall failed with %v", err)
	}
}
