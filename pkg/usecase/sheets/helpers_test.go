package sheetUsecase

import (
	"regexp"
	"testing"

	"github.com/mickey-mickser/google-sheets-project/pkg/models"
	"github.com/sirupsen/logrus"
)

func TestRangeRegexp(t *testing.T) {
	tests := []struct {
		input   string
		matches bool
	}{
		// Positive cases
		{"Sheet1!A1", true},
		{"MySheet!B2:D4", true},
		{"Sheet1!A1:Z100", true},

		// Negative cases
		{"Sheet1", false},
		{"!A1:B2", false},
		{"Sheet1!123", false},
		{"Sheet1!A", false},
		{"Sheet1!A1B2", false},
		{"", false},
	}

	for _, tt := range tests {
		result := rangeRegexp.MatchString(tt.input)
		if result != tt.matches {
			t.Errorf("rangeRegexp.MatchString(%q) = %v; want %v", tt.input, result, tt.matches)
		}
	}
}

func TestColumnToIndex(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"A", 1},
		{"Z", 26},
		{"AA", 27},
		{"AB", 28},
		{"zz", 702},
	}
	for _, tc := range tests {
		got := columnToIndex(tc.input)
		if got != tc.expected {
			t.Errorf("columnToIndex(%q) = %d; want %d", tc.input, got, tc.expected)
		}
	}
}

func TestGenerateValues(t *testing.T) {
	rows, cols := 2, 3
	val := "X"
	matrix := generateValues(rows, cols, val)

	if len(matrix) != rows {
		t.Fatalf("expected %d rows, got %d", rows, len(matrix))
	}
	for i := 0; i < rows; i++ {
		if len(matrix[i]) != cols {
			t.Errorf("row %d: expected %d cols, got %d", i, cols, len(matrix[i]))
		}
		for j := 0; j < cols; j++ {
			if matrix[i][j] != val {
				t.Errorf("matrix[%d][%d] = %v; want %v", i, j, matrix[i][j], val)
			}
		}
	}
}

func TestAtoi(t *testing.T) {
	// Create a dummy sheetUseCase with a logger
	logger := logrus.New()
	s := &sheetUseCase{log: logger}

	// Valid integer
	if got := s.atoi("123"); got != 123 {
		t.Errorf("atoi(\"123\") = %d; want 123", got)
	}

	// Invalid integer: should return 0 without panic
	if got := s.atoi("abc"); got != 0 {
		t.Errorf("atoi(\"abc\") = %d; want 0", got)
	}
}

func TestBuildSingleValueRange(t *testing.T) {
	logger := logrus.New()
	s := &sheetUseCase{log: logger}
	re := regexp.MustCompile(`^([^!]+)!([A-Z]+)(\d+)(?::([A-Z]+)(\d+))?$`)

	req := models.UpdateRequest{
		Range: "Sheet1!A1:B2",
		Value: "V",
	}
	vr, err := s.buildSingleValueRange(req, re)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if vr.Range != "Sheet1!A1:B2" {
		t.Errorf("Range = %q; want %q", vr.Range, "Sheet1!A1:B2")
	}
	// Expect 2 rows x 2 cols of "V"
	if len(vr.Values) != 2 || len(vr.Values[0]) != 2 {
		t.Fatalf("expected 2x2 matrix, got %dx%d", len(vr.Values), len(vr.Values[0]))
	}
	for i := range vr.Values {
		for j := range vr.Values[i] {
			if vr.Values[i][j] != "V" {
				t.Errorf("vr.Values[%d][%d] = %v; want %v", i, j, vr.Values[i][j], "V")
			}
		}
	}
}

func TestBuildValueRanges(t *testing.T) {
	logger := logrus.New()
	s := &sheetUseCase{log: logger}
	updates := []models.UpdateRequest{
		{Range: "S!A1:A1", Value: "1"},
		{Range: "S!B2:C3", Value: "2"},
	}
	rs, err := s.buildValueRanges(updates)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(rs) != len(updates) {
		t.Errorf("expected %d ranges, got %d", len(updates), len(rs))
	}
}

//	func TestBuildSingleValueRange_Invalid(t *testing.T) {
//		logger := logrus.New()
//		s := &sheetUseCase{log: logger}
//		_, err := s.buildSingleValueRange(models.UpdateRequest{Range: "invalid"}, rangeRegexp)
//		if err == nil {
//			t.Error("expected error for invalid range format, got nil")
//		}
//	}
func TestBuildSingleValueRange_Invalid(t *testing.T) {
	logger := logrus.New()
	s := &sheetUseCase{log: logger}
	_, err := s.buildSingleValueRange(models.UpdateRequest{Range: "invalid"}, rangeRegexp)
	if err == nil {
		t.Error("expected error for invalid range format, got nil")
	}
}
