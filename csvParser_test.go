package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAStringAndASizeAreGivenTestReturnsTheStringSlicedByAVolumeOfTheSize(t *testing.T) {
	testTable := []struct {
		element  string
		size     int
		expected []string
	}{
		{
			element:  "balaboliha",
			size:     3,
			expected: []string{"bal", "abo", "lih", "a"},
		},
		{
			element:  "conjuration",
			size:     4,
			expected: []string{"conj", "urat", "ion"},
		},
		{
			element:  "voldemar",
			size:     2,
			expected: []string{"vo", "ld", "em", "ar"},
		},
	}

	for index, item := range testTable {
		result := sliceAString(item.element, item.size)
		comparationResult := reflect.DeepEqual(result, item.expected)
		if comparationResult == true {
			fmt.Printf("Test №%v of sliceAString func has passed successfully!\n", index)
		} else {
			fmt.Printf("Test №%v of sliceAString func has failed!\n", index)
			//ffmt.Println(result)
		}
	}
}

func TestASliceWithACuttedStringAndCapacityAreGivenReturnsSameSliceWithEachElContainedWithSymbolAndWhitespaceIfElementLenIsLessThanCapacityAddsWhitespacesToEqualThem(t *testing.T) {
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

	for index, item := range testTable {
		result := makeABlock(item.elements, item.capacity)
		slicesComparationResult := reflect.DeepEqual(result, item.expected)

		if slicesComparationResult == true {
			fmt.Printf("Test №%v of makeABlock func has passed successfully!\n", index)
		} else {
			fmt.Printf("Test №%v of makeABlock func has failed!\n", index)
		}
	}
}

func TestReturnsCreatedEmptySliceWithGivenSize(t *testing.T) {
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

	for index, item := range testTable {
		result := makeASlice(item.size)
		slicesComparationResult := reflect.DeepEqual(result, item.expected)

		if slicesComparationResult == true {
			fmt.Printf("Test №%v of makeASlice func has passed successfully!\n", index)
		} else {
			fmt.Printf("Test №%v of makeASlice func has failed!\n", index)
		}
	}
}

func TestSliceOfStringsAndCapacityOfElementAreGivenFuncFindsLengthOfTheBiggestStringAndReturnsItDividedByCapacityAndRoundedToBiggestNum(t *testing.T) {
	testTable := []struct {
		record   []string
		capacity int
		expected int
	}{
		{
			record:   []string{"hgpt", "ngysljpt", "jsh", "mbyghpdtbh"},
			capacity: 4,
			expected: 3,
		},
		{
			record:   []string{"hgpt", "ngysljpt", "jsh", "mbyghpdtbh"},
			capacity: 2,
			expected: 5,
		},
		{
			record:   []string{"hgpt", "ngysljpttoplbmnh", "jsh", "mbyghpdtbh"},
			capacity: 3,
			expected: 6,
		},
		{
			record:   []string{"hgpt23ghyphbnvt", "ngysljpttoplbmnh932847823hfsdfjklhds", "jsh", "mbyghpdtbh"},
			capacity: 5,
			expected: 8,
		},
		{
			record:   []string{"egkkelrg3894g43kgr 9304t34jto43otjo  9483t9h43tn u t4398gegndfngdj jngkjerkjgkejbrg lgdkfgldfngljdaslkflksd", "kjsdfksdbf ksjdfbksdjbfjk 38793278egujdfg782 309h934rtwefkjb 2309trjehfg", "df489ut3gob7954t9jgg9 h 498gh9erogh h9f2398 fn u3hefhef e3 foiuhe3odfkgjn489gjndfgjkdfngj489g he3hf8h 98e3f98h 76679e3h 98h3e 983eh89 h3e", "dngjdhng45 in 54g in7 videb78db398i deb8732ef8b83f e3"},
			capacity: 11,
			expected: 13,
		},
		{
			record:   []string{"0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789", "012345678901234567890123456789", "0123456789"},
			capacity: 10,
			expected: 10,
		},
		{
			record:   []string{"00123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789", "012345678901234567890123456789", "0123456789"},
			capacity: 10,
			expected: 11,
		},
	}

	for index, item := range testTable {
		result := maxStringsPerRecord(item.record, item.capacity)

		if result != item.expected {
			fmt.Printf("Test №%v of maxStringsPerRecord func has failed!\n", index)
			//fmt.Println(result)
		} else {
			fmt.Printf("Test №%v of maxStringsPerRecord func has passed successfully!\n", index)
		}
	}
}

func TestTwoArraysAreGivenReturnsThemMergedIntoOne(t *testing.T) {
	testTable := []struct {
		globalArr []string
		currenArr []string
		emptyCell string
		expected  []string
	}{
		{
			globalArr: []string{"| Can |", "|  yo |", "| u p |", "| lea |", "| se  |", "| sto |", "| p?  |"},
			currenArr: []string{"| I'm |", "|  ti |", "| red |", "| .   |"},
			emptyCell: "|     |",
			expected:  []string{"| Can || I'm |", "|  yo ||  ti |", "| u p || red |", "| lea || .   |", "| se  ||     |", "| sto ||     |", "| p?  ||     |"},
		},
	}

	for index, item := range testTable {
		result := mergeArrays(item.globalArr, item.currenArr, item.emptyCell)

		if reflect.DeepEqual(result, item.expected) == true {
			fmt.Printf("Test №%v of mergeArrays func has passed successfully!\n", index)
		} else {
			fmt.Printf("Test №%v of mergeArrays func has failed!\n", index)
		}
	}

	/*result := mergeArrays(testTable[0].globalArr, testTable[0].currenArr, testTable[0].emptyCell)

	for _, item := range result {
		fmt.Println(item)
	}*/
}

/*func TestElementsFormatting(t *testing.T){
	testTable := []struct{

	}
}*/
