package services

import "testing"

func TestValidateRegisterInput(t *testing.T) {
	tests := []struct {
		name    string
		input   RegisterInput
		wantErr bool
	}{
		{name: "valid", input: RegisterInput{Email: "a@b.com", Password: "secret123", DisplayName: "User", Role: "reader"}},
		{name: "short password", input: RegisterInput{Email: "a@b.com", Password: "123", DisplayName: "User", Role: "reader"}, wantErr: true},
		{name: "invalid role", input: RegisterInput{Email: "a@b.com", Password: "secret123", DisplayName: "User", Role: "hacker"}, wantErr: true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateRegisterInput(tc.input)
			if tc.wantErr && err == nil {
				t.Fatal("expected error")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}
