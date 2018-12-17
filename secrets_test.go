package gosecrets

import (
	"reflect"
	"testing"
)

func Test_TokenHex(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "token bytes",
			args:    args{n: 10},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.want, _ = TokenHex(tt.args.n)
		t.Run(tt.name, func(t *testing.T) {
			got, err := TokenHex(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("TokenHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("TokenHex() = %v and want %v", got, tt.want)
			}
		})
	}
}

func Test_TokeBytes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "token bytes",
			args:    args{n: 100},
			want:    []byte{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.want, _ = TokeBytes(tt.args.n)
		t.Run(tt.name, func(t *testing.T) {
			got, err := TokeBytes(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("TokeBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), len(tt.want)) {
				t.Errorf("TokeBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_TokenURLSafe(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "token url safae",
			args:    args{n: 8},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.want, _ = TokenURLSafe(tt.args.n)
		t.Run(tt.name, func(t *testing.T) {
			got, err := TokenURLSafe(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("TokenURLSafe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("TokenURLSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandBelow(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "RandBelow_1",
			args:    args{n: 1},
			wantErr: false,
		},
		{
			name:    "RandBelow_2",
			args:    args{n: 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RandBelow(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("RandBelow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got <= 0 {
				t.Errorf("RandBelow() = %v, must be more than 0", got)
			}
		})
	}
}
