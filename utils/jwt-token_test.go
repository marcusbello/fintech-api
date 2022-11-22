package utils

import "testing"

func TestValidateToken(t *testing.T) {
	type args struct {
		tokenString string
		userName    string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Compare username with user from Token",
			args: args{
				tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Imp1d29uIiwiaXNzIjoidGVzdCIsImV4cCI6MTY2ODgzOTQ0OX0.iA7Yb5saSc6nFQgduFpUIo3d4vPXethKQbE1MQvwmrE",
				userName:    "juwonna",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Test Invalid Token",
			args: args{
				tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Imp1d29uIiwiaXNzIjoidGVzdCIsImV4cCI6MTUxNjIzOTAyMn0.xCJCX4_U79rxW2fGyc-DfnyOFEBEnB6eUKDJu0jED\nbk",
				userName:    "",
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, got := ValidateToken(tt.args.tokenString, tt.args.userName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
