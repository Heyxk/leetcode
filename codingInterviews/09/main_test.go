package main

import (
	"reflect"
	"testing"
)

func TestCQueue_AppendTail(t *testing.T) {
	type fields struct {
		inStack  []int
		outStack []int
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{"case1", fields{inStack: []int{}, outStack: []int{}}, args{1}},
		{"case2", fields{inStack: []int{0}, outStack: []int{}}, args{1}},
		{"case3", fields{inStack: []int{}, outStack: []int{0}}, args{1}},
		{"case4", fields{inStack: []int{0, 1}, outStack: []int{1}}, args{1}},
		{"case5", fields{inStack: []int{1, 2, 3, 4, 5, 6, 7}, outStack: []int{1}}, args{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &CQueue{
				inStack:  tt.fields.inStack,
				outStack: tt.fields.outStack,
			}
			this.AppendTail(tt.args.value)
		})
	}
}

func TestCQueue_DeleteHead(t *testing.T) {
	type fields struct {
		inStack  []int
		outStack []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{"case1", fields{inStack: []int{}, outStack: []int{}}, -1},
		{"case2", fields{inStack: []int{0}, outStack: []int{}}, 0},
		{"case3", fields{inStack: []int{}, outStack: []int{0}}, 0},
		{"case4", fields{inStack: []int{0, 1}, outStack: []int{1}}, 1},
		{"case5", fields{inStack: []int{1, 2, 3, 4, 5, 6, 7}, outStack: []int{1, 2, 3, 4}}, 1},
		{"case6", fields{inStack: []int{5, 6, 7}, outStack: []int{3, 4}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &CQueue{
				inStack:  tt.fields.inStack,
				outStack: tt.fields.outStack,
			}
			if got := this.DeleteHead(); got != tt.want {
				t.Errorf("CQueue.DeleteHead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeetCode(t *testing.T) {
	tests := []struct {
		name      string
		operation []string
		input     []interface{}
		want      []interface{}
	}{
		// TODO: Add test cases.
		{"case1", []string{"CQueue", "appendTail", "deleteHead", "deleteHead", "deleteHead"}, []interface{}{nil, 3, nil, nil, nil}, []interface{}{nil, nil, 3, -1, -1}},
		{
			"case2",
			[]string{"CQueue", "deleteHead", "appendTail", "deleteHead", "deleteHead", "deleteHead", "deleteHead", "appendTail", "deleteHead", "appendTail", "appendTail", "deleteHead", "deleteHead", "deleteHead", "appendTail", "deleteHead", "deleteHead", "deleteHead", "deleteHead", "appendTail", "appendTail", "appendTail", "appendTail", "deleteHead", "deleteHead", "appendTail", "appendTail", "deleteHead", "deleteHead", "deleteHead", "appendTail", "appendTail", "deleteHead", "appendTail", "deleteHead", "appendTail", "appendTail", "appendTail", "deleteHead", "deleteHead", "appendTail", "appendTail", "deleteHead", "deleteHead", "deleteHead", "deleteHead", "appendTail", "appendTail", "deleteHead", "deleteHead", "appendTail", "deleteHead", "appendTail", "appendTail", "appendTail", "appendTail", "deleteHead", "appendTail", "deleteHead", "deleteHead", "appendTail", "appendTail", "appendTail", "deleteHead", "deleteHead", "appendTail", "appendTail", "appendTail", "deleteHead", "deleteHead", "deleteHead", "deleteHead", "appendTail", "appendTail", "deleteHead", "deleteHead", "appendTail", "deleteHead", "appendTail", "appendTail", "deleteHead", "deleteHead", "appendTail", "deleteHead", "appendTail", "appendTail", "appendTail", "deleteHead", "appendTail", "appendTail", "appendTail", "appendTail", "deleteHead", "deleteHead", "deleteHead", "appendTail", "deleteHead", "appendTail", "deleteHead", "appendTail", "appendTail"},
			[]interface{}{nil, nil, 97, nil, nil, nil, nil, 15, nil, 1, 43, nil, nil, nil, 18, nil, nil, nil, nil, 36, 69, 21, 91, nil, nil, 22, 40, nil, nil, nil, 81, 65, nil, 77, nil, 63, 96, 5, nil, nil, 35, 90, nil, nil, nil, nil, 77, 83, nil, nil, 52, nil, 2, 66, 87, 90, nil, 2, nil, nil, 33, 16, 72, nil, nil, 14, 78, 8, nil, nil, nil, nil, 3, 83, nil, nil, 13, nil, 79, 44, nil, nil, 33, nil, 55, 76, 12, nil, 91, 24, 49, 47, nil, nil, nil, 85, nil, 69, nil, 94, 52},
			[]interface{}{nil, -1, nil, 97, -1, -1, -1, nil, 15, nil, nil, 1, 43, -1, nil, 18, -1, -1, -1, nil, nil, nil, nil, 36, 69, nil, nil, 21, 91, 22, nil, nil, 40, nil, 81, nil, nil, nil, 65, 77, nil, nil, 63, 96, 5, 35, nil, nil, 90, 77, nil, 83, nil, nil, nil, nil, 52, nil, 2, 66, nil, nil, nil, 87, 90, nil, nil, nil, 2, 33, 16, 72, nil, nil, 14, 78, nil, 8, nil, nil, 3, 83, nil, 13, nil, nil, nil, 79, nil, nil, nil, nil, 44, 33, 55, nil, 76, nil, 12, nil, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var cq CQueue
			got := []interface{}{}
			for k, v := range tt.operation {
				switch v {
				case "CQueue":
					cq = Constructor()
					got = append(got, nil)
				case "appendTail":
					cq.AppendTail(tt.input[k].(int))
					got = append(got, nil)
				case "deleteHead":
					got = append(got, cq.DeleteHead())
				}
				t.Logf("%s cq=======%v", v, cq)
				t.Logf("got=======%v", got)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}

}

func BenchmarkLeetCode(b *testing.B) {
	tests := []struct {
		name      string
		operation []string
		input     []interface{}
		want      []interface{}
	}{
		// TODO: Add test cases.
		{"case1", []string{"CQueue", "appendTail", "deleteHead", "deleteHead", "deleteHead"}, []interface{}{nil, 3, nil, nil, nil}, []interface{}{nil, nil, 3, -1, -1}},
		{
			"case2",
			[]string{"CQueue", "deleteHead", "appendTail", "deleteHead", "deleteHead", "deleteHead", "deleteHead", "appendTail", "deleteHead", "appendTail", "appendTail", "deleteHead", "deleteHead", "deleteHead", "appendTail", "deleteHead", "deleteHead", "deleteHead", "deleteHead", "appendTail", "appendTail", "appendTail", "appendTail", "deleteHead", "deleteHead", "appendTail", "appendTail", "deleteHead", "deleteHead", "deleteHead", "appendTail", "appendTail", "deleteHead", "appendTail", "deleteHead", "appendTail", "appendTail", "appendTail", "deleteHead", "deleteHead", "appendTail", "appendTail", "deleteHead", "deleteHead", "deleteHead", "deleteHead", "appendTail", "appendTail", "deleteHead", "deleteHead", "appendTail", "deleteHead", "appendTail", "appendTail", "appendTail", "appendTail", "deleteHead", "appendTail", "deleteHead", "deleteHead", "appendTail", "appendTail", "appendTail", "deleteHead", "deleteHead", "appendTail", "appendTail", "appendTail", "deleteHead", "deleteHead", "deleteHead", "deleteHead", "appendTail", "appendTail", "deleteHead", "deleteHead", "appendTail", "deleteHead", "appendTail", "appendTail", "deleteHead", "deleteHead", "appendTail", "deleteHead", "appendTail", "appendTail", "appendTail", "deleteHead", "appendTail", "appendTail", "appendTail", "appendTail", "deleteHead", "deleteHead", "deleteHead", "appendTail", "deleteHead", "appendTail", "deleteHead", "appendTail", "appendTail"},
			[]interface{}{nil, nil, 97, nil, nil, nil, nil, 15, nil, 1, 43, nil, nil, nil, 18, nil, nil, nil, nil, 36, 69, 21, 91, nil, nil, 22, 40, nil, nil, nil, 81, 65, nil, 77, nil, 63, 96, 5, nil, nil, 35, 90, nil, nil, nil, nil, 77, 83, nil, nil, 52, nil, 2, 66, 87, 90, nil, 2, nil, nil, 33, 16, 72, nil, nil, 14, 78, 8, nil, nil, nil, nil, 3, 83, nil, nil, 13, nil, 79, 44, nil, nil, 33, nil, 55, 76, 12, nil, 91, 24, 49, 47, nil, nil, nil, 85, nil, 69, nil, 94, 52},
			[]interface{}{nil, -1, nil, 97, -1, -1, -1, nil, 15, nil, nil, 1, 43, -1, nil, 18, -1, -1, -1, nil, nil, nil, nil, 36, 69, nil, nil, 21, 91, 22, nil, nil, 40, nil, 81, nil, nil, nil, 65, 77, nil, nil, 63, 96, 5, 35, nil, nil, 90, 77, nil, 83, nil, nil, nil, nil, 52, nil, 2, 66, nil, nil, nil, 87, 90, nil, nil, nil, 2, 33, 16, 72, nil, nil, 14, 78, nil, 8, nil, nil, 3, 83, nil, 13, nil, nil, nil, 79, nil, nil, nil, nil, 44, 33, 55, nil, 76, nil, 12, nil, nil},
		},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			var cq CQueue
			for k, v := range tt.operation {
				switch v {
				case "CQueue":
					cq = Constructor()
				case "appendTail":
					cq.AppendTail(tt.input[k].(int))
				case "deleteHead":
				}
			}
		}
	}
}
