package main

import (
	"reflect"
	"testing"
)

// Chunk tests on table driven
func TestChunk(t *testing.T) {
	tests := []struct {
		name string
		args []int
		size int
		want [][]int
	}{
		{
			name: "chunk 1",
			args: []int{1, 2, 3, 4, 5},
			size: 2,
			want: [][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			name: "chunk 2",
			args: []int{1, 2, 3, 4, 5},
			size: 3,
			want: [][]int{{1, 2, 3}, {4, 5}},
		},
		{
			name: "chunk 3",
			args: []int{1, 2, 3, 4, 5},
			size: 1,
			want: [][]int{{1}, {2}, {3}, {4}, {5}},
		},
		{
			name: "chunk 4",
			args: []int{1, 2, 3, 4, 5},
			size: 0,
			want: [][]int{{1, 2, 3, 4, 5}},
		},
		{
			name: "chunk 5",
			args: []int{1, 2, 3, 4, 5},
			size: -1,
			want: [][]int{{1, 2, 3, 4, 5}},
		},
		{
			name: "chunk 6",
			args: []int{},
			size: 2,
			want: [][]int{{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunk(tt.args, tt.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}
