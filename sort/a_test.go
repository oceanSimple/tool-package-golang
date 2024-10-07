package mySort

import "testing"

func TestInsertionSort(t *testing.T) {
	TestSortFunc("InsertionSort", InsertionSort)
}

func TestShellSort(t *testing.T) {
	TestSortFunc("ShellSort", ShellSort)
}
