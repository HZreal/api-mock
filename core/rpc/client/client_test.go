package main

import "testing"

/**
 * @Author elastic·H
 * @Date 2024-08-09
 * @File: client_test.go
 * @Description:
 */

func Test_invoke(t *testing.T) {
	type args struct {
		x int32
		y int32
	}
	tests := []struct {
		name         string
		args         args
		wantSum      int32
		wantSubtract int32
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{x: 10, y: 20}, wantSum: 30, wantSubtract: -10},
		{name: "case2", args: args{x: 10, y: 5}, wantSum: 15, wantSubtract: 5},
		{name: "case3", args: args{x: 20, y: 20}, wantSum: 40, wantSubtract: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSum, gotSubtract := invoke(tt.args.x, tt.args.y)
			if gotSum != tt.wantSum {
				t.Errorf("invoke() gotSum = %v, want %v", gotSum, tt.wantSum)
			}
			if gotSubtract != tt.wantSubtract {
				t.Errorf("invoke() gotSubtract = %v, want %v", gotSubtract, tt.wantSubtract)
			}
		})
	}
}

func Test_invoke2(t *testing.T) {
	type args struct {
		num int32
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{num: 120}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			invoke2(tt.args.num)
		})
	}
}

func Test_invoke3(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{arr: []int32{1, 2, 3, 4, 5}}, want: 3.00},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invoke3(tt.args.arr); got != tt.want {
				t.Errorf("invoke3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_invoke4(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			invoke4()
		})
	}
}
