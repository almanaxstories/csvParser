package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBlocksRendering(t *testing.T) {
	testTable := []struct {
		element  string
		capacity int
		expected []string
	}{
		{
			element:  "balaboliha",
			capacity: 3,
			expected: []string{"bal", "abo", "lih", "a"},
		},
		{
			element:  "conjuration",
			capacity: 4,
			expected: []string{"conj", "urat", "ion"},
		},
		{
			element:  "voldemar",
			capacity: 2,
			expected: []string{"vo", "ld", "em", "ar"},
		},
	}

	for _, item := range testTable {
		result := renderBlock(item.element, item.capacity)
		comparationResult := reflect.DeepEqual(result, item.expected)
		if comparationResult == true {
			fmt.Println("Test passed successfully!")
		} else {
			fmt.Println("Test failed!")
			fmt.Println(result)
		}
	}
}

func TestBlockPrettyfycation(t *testing.T) {
	testTable := []struct {
		elements []string
		capacity int
		expected []string
	}{
		{
			elements: []string{"bal", "abo", "lih", "a"},
			capacity: 3,
			expected: []string{"| bal |", "| abo |", "| lih |", "| a   |"},
		},
		{
			elements: []string{"conj", "urat", "ion"},
			capacity: 4,
			expected: []string{"| conj |", "| urat |", "| ion  |"},
		},
		{
			elements: []string{"vo", "ld", "em", "ar"},
			capacity: 2,
			expected: []string{"| vo |", "| ld |", "| em |", "| ar |"},
		},
	}

	for _, item := range testTable {
		result := fillStrings(item.elements, item.capacity)
		slicesComparationResult := reflect.DeepEqual(result, item.expected)

		if slicesComparationResult == true {
			fmt.Println("Test passed successfully!")
		} else {
			fmt.Println("Test failed!")
		}
	}
}

func TestGlobalArrInitialization(t *testing.T) {
	testTable := []struct {
		size     int
		expected []string
	}{
		{
			size:     5,
			expected: []string{"", "", "", "", ""},
		},
		{
			size:     10,
			expected: []string{"", "", "", "", "", "", "", "", "", ""},
		},
		{
			size:     3,
			expected: []string{"", "", ""},
		},
	}

	for _, item := range testTable {
		result := initGlobalArr(item.size)
		slicesComparationResult := reflect.DeepEqual(result, item.expected)

		if slicesComparationResult == true {
			fmt.Println("Test passed successfully!")
		} else {
			fmt.Println("Test failed!")
		}
	}
}

func TestMaxSubstringsPerRecordCalculation(t *testing.T) {
	testTable := []struct {
		record   []string
		capcaity int
		expected int
	}{
		{
			record:   []string{"hgpt", "ngysljpt", "jsh", "mbyghpdtbh"},
			capcaity: 4,
			expected: 3,
		},
		{
			record:   []string{"hgpt", "ngysljpt", "jsh", "mbyghpdtbh"},
			capcaity: 2,
			expected: 5,
		},
		{
			record:   []string{"hgpt", "ngysljpttoplbmnh", "jsh", "mbyghpdtbh"},
			capcaity: 3,
			expected: 6,
		},
	}

	for _, item := range testTable {
		result := maxStringsPerRecord(item.record, item.capcaity)

		if result != item.expected {
			fmt.Println("Test 4 failed!")
			fmt.Println(result)
		} else {
			fmt.Println("Test 4 passed!")
		}
	}
}
