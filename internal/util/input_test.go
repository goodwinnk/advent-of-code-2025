package util

import "testing"

func TestReadInput_Day01_Task1(t *testing.T) {
	got, err := ReadInput(1, "task1.txt")
	if err != nil {
		t.Fatalf("ReadInput failed: %v", err)
	}
	if len(got) == 0 {
		t.Fatalf("ReadInput returned empty content for day01/task1.txt")
	}
}
