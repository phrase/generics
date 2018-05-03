package generics

import (
	"testing"
)

func TestJoin(t *testing.T) {
	payments := []*Payment{
		{ID: 1, AccountID: 1},
		{ID: 2, AccountID: 1},
		{ID: 3, AccountID: 2},
	}
	accounts := []*Account{
		{1, "Account 1"},
		{2, "Account 2"},
	}

	am := Index(accounts, "ID")
	New(payments).Join(am)

	tests := []struct{ Has, Want interface{} }{
		{payments[0] != nil, true},
		{payments[0].AccountID, 1},
		{payments[1].AccountID, 1},
		{payments[2].AccountID, 2},
	}
	for i, tc := range tests {
		if tc.Want != tc.Has {
			t.Errorf("%d: want %#v (%T), was %#v (%T)", i+1, tc.Want, tc.Want, tc.Has, tc.Has)
		}
	}
}

func TestJoinWithSlice(t *testing.T) {
	payments := []*Payment{
		{ID: 1, AccountID: 1},
		{ID: 2, AccountID: 1},
		{ID: 3, AccountID: 2},
	}
	accounts := []*Account{
		{1, "Account 1"},
		{2, "Account 2"},
	}
	New(payments).Join(accounts)

	tests := []struct{ Has, Want interface{} }{
		{payments[0] != nil, true},
		{payments[0].AccountID, 1},
		{payments[1].AccountID, 1},
		{payments[2].AccountID, 2},
	}
	for i, tc := range tests {
		if tc.Want != tc.Has {
			t.Errorf("%d: want %#v (%T), was %#v (%T)", i+1, tc.Want, tc.Want, tc.Has, tc.Has)
		}
	}
}
