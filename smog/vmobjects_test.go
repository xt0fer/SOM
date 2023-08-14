package smog

import (
	"reflect"
	"testing"
)

func TestNewSObject(t *testing.T) {
	type args struct {
		n    int32
		with *SObject
	}
	tests := []struct {
		name string
		args args
		want *SObject
	}{
		// TODO: Add test cases.
		{"first object", args{1, nil}, NewSObject(1, nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSObject(tt.args.n, tt.args.with); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSClass(t *testing.T) {
	u := NewUniverse()
	type args struct {
		numberOfFields int32
		u              *Universe
	}
	tests := []struct {
		name string
		args args
		want *SClass
	}{
		// TODO: Add test cases.
		{"first class", args{1, u}, NewSClass(1, u)},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSClass(tt.args.numberOfFields, tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSClass() = %v, want %v", got, tt.want)
			}
		})
	}
}
