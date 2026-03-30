package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{
			"size 100",
			100,
			100,
		},
		{
			"zero size",
			0,
			0,
		},
		{
			"one size",
			1,
			1,
		},
		{
			"negative size",
			-1,
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Len(t, generateRandomElements(tt.size), tt.want)
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			"many elements",
			[]int{1, 5, 2, 3, 9, 4, 6, 7, 8},
			9,
		},
		{
			"one element",
			[]int{10},
			10,
		},
		{
			"zero elements",
			[]int{},
			0,
		},
		{
			"same elements",
			[]int{15, 15, 15},
			15,
		},
		{
			"negative elements",
			[]int{-1, -2, -3},
			-1,
		},
		{
			"mixed elements",
			[]int{1, -2, 6},
			6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maximum(tt.input))
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			"zero",
			[]int{},
			0,
		},
		{
			"less CHUNKS",
			[]int{3, 1, 2},
			3,
		},
		{
			"equal CHUNKS",
			[]int{3, 1, 4, 1, 5, 8, 2, 3},
			8,
		},
		{
			"greater CHUNKS",
			[]int{3, 1, 4, 1, 5, 8, 2, 3, 6, 7, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxChunks(tt.input))
		})
	}
}
