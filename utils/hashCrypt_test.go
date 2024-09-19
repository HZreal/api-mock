package utils

/**
 * @Author elastic·H
 * @Date 2024-08-09
 * @File: hashCrypt_test.go
 * @Description:
 */

import (
	"testing"
)

func Test_encryptAndValidate(t *testing.T) {
	type args struct {
		pwd string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{pwd: "123456"}, want: true},
		{name: "test2", args: args{pwd: "123456"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encryptAndValidate(tt.args.pwd); got != tt.want {
				t.Errorf("encryptAndValidate() = %v, want %v", got, tt.want)
			}
		})
	}
}
