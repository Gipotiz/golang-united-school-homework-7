package coverage

import (
	"github.com/stretchr/testify/assert"
	"os"
	"sort"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestPeople(t *testing.T) {

	persons := People{
		{
			firstName: "Петров",
			lastName:  "Аркадий",
			birthDay:  time.Date(1950, 6, 4, 0, 0, 0, 0, time.UTC),
		},
		{
			firstName: "Сидоров",
			lastName:  "Генадий",
			birthDay:  time.Date(1967, 1, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			firstName: "Васильев",
			lastName:  "Иван",
			birthDay:  time.Date(1952, 7, 11, 0, 0, 0, 0, time.UTC),
		},
	}

	tests := []struct {
		name     string
		body     People
		expected People
	}{
		// 1 ...
		{
			name: "sortDateOrderByDESC",
			body: persons,
			expected: People{
				persons[1],
				persons[2],
				persons[0],
			},
		},
		// 2 ...
		{
			name: "sortFirstNameOrderByASC",
			body: People{
				persons[0],
				persons[1],
				{
					firstName: "Васильев",
					lastName:  "Иван",
					birthDay:  time.Date(1967, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			expected: People{
				{
					firstName: "Васильев",
					lastName:  "Иван",
					birthDay:  time.Date(1967, 1, 2, 0, 0, 0, 0, time.UTC),
				},
				persons[1],
				persons[0],
			},
		},
		// 3 ...
		{
			name: "sortLastNameOrderByASC",
			body: People{
				persons[0],
				persons[1],
				{
					firstName: "Сидоров",
					lastName:  "Николай",
					birthDay:  time.Date(1967, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			expected: People{
				persons[1],
				{
					firstName: "Сидоров",
					lastName:  "Николай",
					birthDay:  time.Date(1967, 1, 2, 0, 0, 0, 0, time.UTC),
				},
				persons[0],
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			sort.Sort(testCase.body)

			assert.Equal(t, testCase.body, testCase.expected)
		})
	}

}

func TestNew(t *testing.T) {

	tests := []struct {
		name     string
		body     string
		expected *Matrix
		wantErr  bool
	}{
		{
			name:     "Ok",
			body:     "11 11\n22 22 \n 33 33   ",
			expected: &Matrix{3, 2, []int{11, 11, 22, 22, 33, 33}},
		},
		{
			name:    "rows different lengths",
			body:    "11 11\n22 22 \n 33",
			wantErr: true,
		},
		{
			name:    "not a symbol",
			body:    "11 11\n22 a2 \n 33 33",
			wantErr: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			matrix, err := New(testCase.body)

			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, matrix, testCase.expected)
			}
		})
	}
}

func TestMatrix_Rows(t *testing.T) {
	test := struct {
		body     string
		expected [][]int
	}{
		body:     "11 11\n22 22 \n 33 33   ",
		expected: [][]int{{11, 11}, {22, 22}, {33, 33}},
	}

	matrix, _ := New(test.body)
	resRow := matrix.Rows()

	assert.Equal(t, resRow, test.expected)
}

func TestMatrix_Cols(t *testing.T) {
	test := struct {
		body     string
		expected [][]int
	}{
		body:     "11 11\n22 22 \n 33 33  ",
		expected: [][]int{{11, 22, 33}, {11, 22, 33}},
	}

	matrix, _ := New(test.body)
	resRow := matrix.Cols()

	assert.Equal(t, resRow, test.expected)
}

func TestMatrix_Set(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		expected bool
	}{
		{
			name:     "Ok",
			body:     "11 11\n22 22 \n 33 33",
			expected: true,
		},
		{
			name:     "Wrong Input",
			body:     "11 11\n22 22 \n 33 33",
			expected: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			matrix, _ := New(testCase.body)

			if testCase.expected {
				res := matrix.Set(0, 1, 356)
				assert.Equal(t, res, testCase.expected)
			} else {
				res := matrix.Set(5, 1, 356)
				assert.Equal(t, res, testCase.expected)
			}
		})
	}

}
