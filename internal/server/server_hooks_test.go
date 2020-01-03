package server

import (
	"reflect"
	"testing"
)

func TestNewNoOpHook(t *testing.T) {
	type args struct {
		server Server
	}
	tests := []struct {
		name string
		args args
		want SHooks
	}{
		{"ShouldGetNoOpHook", args{Server{}}, NoOpHook{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNoOpHook(tt.args.server); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNoOpHook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoOpHook_PreCmdHook(t *testing.T) {
	type args struct {
		cmd Message
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"NoErrOnAnything", args{Message{}}, false},
		{"NoErrOnUnknownCmd", args{Message{_command:"XYZ"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NoOpHook{}
			if err := n.PreCmdHook(tt.args.cmd); (err != nil) != tt.wantErr {
				t.Errorf("PreCmdHook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNoOpHook_PreStartHook(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"NoErrOnAnything", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NoOpHook{}
			if err := n.PreStartHook(); (err != nil) != tt.wantErr {
				t.Errorf("PreStartHook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}