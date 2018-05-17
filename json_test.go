package main

import (
	"encoding/json"
	"testing"
)

type Object = map[string]interface{}
type Array = []interface{}

const jsonString = `
{
	"header": "main",
	"elements": [
		{
			"name": "aaa",
			"size": 1,
			"active": true
		},
		{
			"name": "bbb",
			"size": 2,
			"active": false
		},
		{
			"name": "ccc",
			"size": 3,
			"active": true
		}
	]
}
`

type A struct {
	Header   string
	Elements []Element
}

type Element struct {
	Name   string
	Size   int
	Active bool
}

func UnmarshalUnstructuredJSON() (result interface{}) {
	if err := json.Unmarshal([]byte(jsonString), &result); err != nil {
		panic(err)
	}
	return
}

func UnmarshalStructuredJSON() *A {
	var a A
	if err := json.Unmarshal([]byte(jsonString), &a); err != nil {
		panic(err)
	}
	return &a
}

func TestUnmarshalUnstructuredJSON(t *testing.T) {
	r := UnmarshalUnstructuredJSON()
	if r.(Object)["elements"].(Array)[2].(Object)["name"].(string) != "ccc" {
		t.Fail()
	}
}

func TestUnmarshalStructuredJSON(t *testing.T) {
	a := UnmarshalStructuredJSON()
	if a.Elements[2].Name != "ccc" {
		t.Fail()
	}
}

func BenchmarkUnmarshalUnstructuredJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UnmarshalUnstructuredJSON()
	}
}

func BenchmarkUnmarshalStructuredJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UnmarshalStructuredJSON()
	}
}

func BenchmarkMarshalUnstructuredJSON(b *testing.B) {
	r := UnmarshalUnstructuredJSON()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(r)
	}
}

func BenchmarkMarshalStructuredJSON(b *testing.B) {
	r := UnmarshalStructuredJSON()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(r)
	}
}
