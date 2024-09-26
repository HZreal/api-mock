package service

/**
 * @Author elastic·H
 * @Date 2024-09-25
 * @File: requestTaskService_test.go
 * @Description:
 */

import (
	"fmt"
	"testing"
)

func t() {
	fmt.Println("t")
}

func TestSendReq(t *testing.T) {
	type args struct {
		uri    string
		method string
		body   string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendReq(tt.args.uri, tt.args.method, tt.args.body)
		})
	}
}
