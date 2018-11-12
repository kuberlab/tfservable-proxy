package tf

import (
	"encoding/json"
	"fmt"
	"testing"
)

func sum(arr []int32) int64 {
	summ := int64(0)
	for _, i := range arr {
		summ += int64(i)
	}
	return summ
}

var scalar = `5`

func TestParseTensorScalar(t *testing.T) {
	var raw = fmt.Sprintf(`{"inputs": {"input": {"dtype": 4, "data": %v}}}`, scalar)
	modelData := ModelData{}
	err := json.Unmarshal([]byte(raw), &modelData)
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range modelData.Inputs {
		if k != "input" {
			t.Fatal("must be input")
		}
		tproto, err := v.Tensor()
		if err != nil {
			t.Fatal(err)
		}
		if len(tproto.TensorShape.Dim) != 0 {
			t.Fatal("Rank must be 0, got", len(tproto.TensorShape.Dim))
		}
		if sum(tproto.IntVal) != int64(5) {
			t.Fatal("sum must be 10, got", sum(tproto.IntVal))
		}
	}
}

var t1D = `[3, 1, 4, 2, 0]`

func TestParseTensor1D(t *testing.T) {
	var raw = fmt.Sprintf(`{"inputs": {"input": {"dtype": 4, "data": %v}}}`, t1D)
	modelData := ModelData{}
	err := json.Unmarshal([]byte(raw), &modelData)
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range modelData.Inputs {
		if k != "input" {
			t.Fatal("must be input")
		}
		tproto, err := v.Tensor()
		if err != nil {
			t.Fatal(err)
		}
		if len(tproto.TensorShape.Dim) != 1 {
			t.Fatal("Rank must be 1, got", len(tproto.TensorShape.Dim))
		}
		if sum(tproto.IntVal) != int64(10) {
			t.Fatal("sum must be 10")
		}
	}
}

var t2D = `[[0, 0, 4, 0, 0],
       [0, 0, 3, 0, 0],
       [0, 0, 2, 0, 0],
       [0, 0, 1, 0, 0]]`

func TestParseTensor2D(t *testing.T) {
	var raw = fmt.Sprintf(`{"inputs": {"input": {"dtype": 4, "data": %v}}}`, t2D)
	modelData := ModelData{}
	err := json.Unmarshal([]byte(raw), &modelData)
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range modelData.Inputs {
		if k != "input" {
			t.Fatal("must be input")
		}
		tproto, err := v.Tensor()
		if err != nil {
			t.Fatal(err)
		}
		if len(tproto.TensorShape.Dim) != 2 {
			t.Fatal("Rank must be 2, got", len(tproto.TensorShape.Dim))
		}
		if sum(tproto.IntVal) != int64(10) {
			t.Fatal("sum must be 10")
		}
	}
}

// 3d tensor: 5,4,3
var t3D = `[[[0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0],
        [0, 0, 3, 0, 0],
        [0, 0, 0, 0, 0]],

       [[0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0],
        [0, 0, 4, 0, 0],
        [0, 0, 0, 0, 0]],

       [[0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0],
        [0, 0, 5, 0, 0],
        [0, 0, 0, 0, 0]]]`

func TestParseTensor3D(t *testing.T) {
	var raw = fmt.Sprintf(`{"inputs": {"input": {"dtype": 4, "data": %v}}}`, t3D)
	modelData := ModelData{}
	err := json.Unmarshal([]byte(raw), &modelData)
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range modelData.Inputs {
		if k != "input" {
			t.Fatal("must be input")
		}
		tproto, err := v.Tensor()
		if err != nil {
			t.Fatal(err)
		}
		if len(tproto.TensorShape.Dim) != 3 {
			t.Fatal("Rank must be 3, got", len(tproto.TensorShape.Dim))
		}
		if sum(tproto.IntVal) != int64(12) {
			t.Fatal("sum must be 12")
		}
	}
}

var t4D = `[[[[0, 0, 0, 0, 0],
         [0, 0, 0, 0, 0],
         [0, 0, 1, 0, 0],
         [0, 0, 0, 0, 0]],

        [[0, 0, 0, 0, 0],
         [0, 0, 0, 0, 0],
         [0, 0, 2, 0, 0],
         [0, 0, 0, 0, 0]],

        [[0, 0, 0, 0, 0],
         [0, 0, 3, 0, 0],
         [0, 0, 0, 0, 0],
         [0, 0, 0, 0, 0]]],


       [[[0, 0, 0, 0, 0],
         [0, 0, 0, 0, 0],
         [0, 0, 4, 0, 0],
         [0, 0, 0, 0, 0]],

        [[0, 0, 0, 0, 0],
         [0, 0, 0, 0, 0],
         [0, 0, 5, 0, 0],
         [0, 0, 0, 0, 0]],

        [[0, 0, 0, 0, 0],
         [0, 0, 0, 0, 0],
         [0, 0, 6, 0, 0],
         [0, 0, 0, 0, 0]]]]`

func TestParseTensor4D(t *testing.T) {
	var raw = fmt.Sprintf(`{"inputs": {"input": {"dtype": 4, "data": %v}}}`, t4D)
	modelData := ModelData{}
	err := json.Unmarshal([]byte(raw), &modelData)
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range modelData.Inputs {
		if k != "input" {
			t.Fatal("must be input")
		}
		tproto, err := v.Tensor()
		if err != nil {
			t.Fatal(err)
		}
		if len(tproto.TensorShape.Dim) != 4 {
			t.Fatal("Rank must be 4, got", len(tproto.TensorShape.Dim))
		}
		if sum(tproto.IntVal) != int64(21) {
			t.Fatal("sum must be 21, got", sum(tproto.IntVal))
		}
	}
}

var t5D = `[[[[[0, 0, 0, 0, 0],
          [0, 0, 0, 0, 0],
          [0, 0, 12, 0, 0],
          [0, 0, 0, 0, 0]],

         [[0, 0, 11, 0, 0],
          [0, 0, 0, 0, 0],
          [0, 0, 0, 0, 0],
          [0, 0, 0, 0, 0]],

         [[0, 0, 10, 0, 0],
          [0, 0, 0, 0, 0],
          [0, 0, 0, 0, 0],
          [0, 0, 0, 0, 0]]],


        [[[0, 0, 0, 0, 0],
          [0, 0, 0, 0, 0],
          [0, 0, 9, 0, 0],
          [0, 0, 0, 0, 0]],

         [[0, 0, 0, 0, 0],
          [0, 0, 0, 0, 0],
          [0, 0, 8, 0, 0],
          [0, 0, 0, 0, 0]],

         [[0, 0, 0, 0, 0],
          [0, 0, 7, 0, 0],
          [0, 0, 0, 0, 0],
          [0, 0, 0, 0, 0]]]],



       [[[[0, 0, 0, 0, 0],
          [0, 0, 6, 0, 0],
          [0, 0, 0, 0, 0],
          [0, 0, 0, 0, 0]],

         [[0, 0, 0, 0, 0],
          [0, 0, 5, 0, 0],
          [0, 0, 0, 0, 0],
          [0, 0, 0, 0, 0]],

         [[0, 0, 0, 0, 0],
          [0, 0, 4, 0, 0],
          [0, 0, 0, 0, 0],
          [0, 0, 0, 0, 0]]],


        [[[0, 0, 0, 0, 0],
          [0, 0, 0, 0, 0],
          [0, 0, 3, 0, 0],
          [0, 0, 0, 0, 0]],

         [[0, 0, 0, 0, 0],
          [0, 0, 0, 0, 0],
          [0, 0, 2, 0, 0],
          [0, 0, 0, 0, 0]],

         [[0, 0, 0, 0, 0],
          [0, 0, 0, 0, 0],
          [0, 0, 1, 0, 0],
          [0, 0, 0, 0, 0]]]]]`

func TestParseTensor5D(t *testing.T) {
	var raw = fmt.Sprintf(`{"inputs": {"input": {"dtype": 4, "data": %v}}}`, t5D)
	modelData := ModelData{}
	err := json.Unmarshal([]byte(raw), &modelData)
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range modelData.Inputs {
		if k != "input" {
			t.Fatal("must be input")
		}
		tproto, err := v.Tensor()
		if err != nil {
			t.Fatal(err)
		}
		if len(tproto.TensorShape.Dim) != 5 {
			t.Fatal("Rank must be 5, got", len(tproto.TensorShape.Dim))
		}
		if sum(tproto.IntVal) != int64(78) {
			t.Fatal("sum must be 78, got", sum(tproto.IntVal))
		}
	}
}
