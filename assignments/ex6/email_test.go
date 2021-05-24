package ex6

import "testing"

func TestIsValidEmailAddress(t *testing.T) {
	type input struct {
		emailAddress string
	}
	tests := []struct {
		name string
		args input
		want bool
	}{
		// TODO: Add test cases.
		{"normal", input{"lokesh@xebia.com"}, true},
		{"dot in between", input{"lokesh.nirania@xebia.com"}, true},
		{"@ absent", input{"lokesh.xebia.com"}, false},
		{"plain .com missing", input{"lokesh@xebia"}, true},
		{"dot present but not after @", input{"lokesh.nirania@xebia"}, true},
		{"post # present", input{"lokesh@nirania#xebia.com"}, false},
		{"pre # present", input{"lokesh#nirania@xebia.com"}, true},
		{"illegal characters", input{"loki@xebia$ki.com"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidEmailAddress(tt.args.emailAddress); got != tt.want {
				t.Errorf("IsValidEmailAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
