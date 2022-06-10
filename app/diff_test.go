package main

import (
	"testing"
)

func TestCompareSame(t *testing.T) {
	testData := "a"
	testDataB := "a"
	str, err := compare(testData, testDataB)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
	if str != "[{Equal a}]" {
		t.Log("error should be [{Equal a}], but got", err)
		t.Fail()
	}
}
