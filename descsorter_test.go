package descsort

import (
	"reflect"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDescInts(t *testing.T) {
	received := []int{7, 3, 5, 0, 1, 6, 8, 4, 2, 9}
	expected := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	t.Log("Given the need to test if this function sorts the given int slice in descending order")
	{
		DescInts(received)
		if !reflect.DeepEqual(received, expected) {
			t.Logf("Expected %+v\n, Got %+v", expected, received)
			t.Fatal("Should be sorted in descending order", ballotX)
		}

		t.Log("Should be sorted in descending order", checkMark)
	}
}

func TestDescFloat64s(t *testing.T) {
	received := []float64{7.2, -3.1, -5.9, 0.9, 1.1, 6.3, 8.1, 4.2, 2.9, 9.2}
	expected := []float64{9.2, 8.1, 7.2, 6.3, 4.2, 2.9, 1.1, 0.9, -3.1, -5.9}

	t.Log("Given the need to test if this function sorts the given float64 slice in descending order")
	{
		DescFloat64s(received)
		if !reflect.DeepEqual(received, expected) {
			t.Logf("Expected %+v\n, Got %+v", expected, received)
			t.Fatal("Should be sorted in descending order", ballotX)
		}

		t.Log("Should be sorted in descending order", checkMark)
	}
}

func TestDescStrings(t *testing.T) {
	received := []string{"a", "v", "s", "r", "e"}
	expected := []string{"v", "s", "r", "e", "a"}

	t.Log("Given the need to test if this function sorts the given string slice in descending order")
	{
		DescStrings(received)
		if !reflect.DeepEqual(received, expected) {
			t.Logf("Expected %+v\n, Got %+v", expected, received)
			t.Fatal("Should be sorted in descending order", ballotX)
		}

		t.Log("Should be sorted in descending order", checkMark)
	}
}
