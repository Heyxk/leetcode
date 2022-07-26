package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MinStack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinStack_Push(t *testing.T) {
	type fields struct {
		stack []int
		min   []int
	}
	type args struct {
		x int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &MinStack{
				stack: tt.fields.stack,
				min:   tt.fields.min,
			}
			ms.Push(tt.args.x)
		})
	}
}

func TestMinStack_Pop(t *testing.T) {
	type fields struct {
		stack []int
		min   []int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &MinStack{
				stack: tt.fields.stack,
				min:   tt.fields.min,
			}
			ms.Pop()
		})
	}
}

func TestMinStack_Top(t *testing.T) {
	type fields struct {
		stack []int
		min   []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &MinStack{
				stack: tt.fields.stack,
				min:   tt.fields.min,
			}
			if got := ms.Top(); got != tt.want {
				t.Errorf("MinStack.Top() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinStack_Min(t *testing.T) {
	type fields struct {
		stack []int
		min   []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &MinStack{
				stack: tt.fields.stack,
				min:   tt.fields.min,
			}
			if got := ms.Min(); got != tt.want {
				t.Errorf("MinStack.Min() = %v, want %v", got, tt.want)
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
		{"case1", []string{"MinStack", "push", "push", "push", "min", "pop", "top", "min"}, []interface{}{nil, -2, 0, -3, nil, nil, nil, nil}, []interface{}{nil,nil,nil,nil,-3,nil,0,-2}},
		{"case2", []string{"MinStack", "push", "push", "push", "min", "top", "pop", "min"}, []interface{}{nil, -2, 0, -1, nil, nil, nil, nil}, []interface{}{nil, nil, nil, nil, -2, -1, nil, -2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ms MinStack
			got := []interface{}{}
			for k, v := range tt.operation {
				switch v {
				case "MinStack":
					ms = Constructor()
					got = append(got, nil)
				case "push":
					ms.Push(tt.input[k].(int))
					got = append(got, nil)
				case "pop":
					ms.Pop()
					got = append(got, nil)
				case "min":
					got = append(got, ms.Min())
				case "top":
					got = append(got, ms.Top())
				}
				// t.Logf("%s ms=======%v", v, ms)
				// t.Logf("%s got=======%v", v, got)
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
		{"case1", []string{"MinStack", "push", "push", "push", "min", "pop", "top", "min"}, []interface{}{nil, -2, 0, -3, nil, nil, nil, nil}, []interface{}{nil, nil, nil, nil, -3, nil, 0, -2}},
		{"case2", []string{"MinStack", "push", "push", "push", "min", "top", "pop", "min"}, []interface{}{nil, -2, 0, -1, nil, nil, nil, nil}, []interface{}{nil, nil, nil, nil, -2, -1, nil, -2}},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			var ms MinStack
			for k, v := range tt.operation {
				switch v {
				case "MinStack":
					ms = Constructor()
				case "push":
					ms.Push(tt.input[k].(int))
				case "pop":
					ms.Pop()
				case "min":
					ms.Min()
				case "top":
					ms.Top()
				}
			}
		}
	}
}
