package main

import "testing"

var result string

func BenchmarkAddStrings10(b *testing.B) {

	var s string
	for n := 0; n < b.N; n++ {
		s = AddStrings("yo", "hey", 1)
	}
	result = s
}

func BenchmarkAddStringswJoin10(b *testing.B) {

	var s string
	for n := 0; n < b.N; n++ {
		AddStringswJoin("yo", "hey", 1)
	}
	result = s
}

func BenchmarkAddStringswFormat10(b *testing.B) {

	var s string
	for n := 0; n < b.N; n++ {
		AddStringswFormat("yo", "hey", 1)
	}
	result = s
}

func BenchmarkAddStringswBytesBuffer10(b *testing.B) {

	var s string
	for n := 0; n < b.N; n++ {
		AddBytesBuffer("yo", "hey", 1)
	}
	result = s

}
