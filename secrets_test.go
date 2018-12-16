package gosecrets

import (
	"reflect"
	"testing"
)

func Test_randBelow(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "randBelow",
			args: args{n: 100},
		},
		{
			name: "randBelow",
			args: args{n: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randBelow(tt.args.n); got >= tt.args.n {
				t.Errorf("randBelow() = %v, want %v", got, tt.args.n)
			}
		})
	}
}

func Test_tokenHex(t *testing.T) {
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
		tt.want, _ = tokenHex(tt.args.n)
		t.Run(tt.name, func(t *testing.T) {
			got, err := tokenHex(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("tokenHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("tokenHex() = %v and want %v", got, tt.want)
			}
		})
	}
}

func Test_tokeBytes(t *testing.T) {
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
		tt.want, _ = tokeBytes(tt.args.n)
		t.Run(tt.name, func(t *testing.T) {
			got, err := tokeBytes(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("tokeBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), len(tt.want)) {
				t.Errorf("tokeBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tokenUrlSafe(t *testing.T) {
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
		tt.want, _ = tokenUrlSafe(tt.args.n)
		t.Run(tt.name, func(t *testing.T) {
			got, err := tokenUrlSafe(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("tokenUrlSafe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("tokenUrlSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}
