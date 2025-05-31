package sheetUsecase

import (
	"context"
	"fmt"
	"github.com/mickey-mickser/google-sheets-project/pkg/models"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var rangeRegexp = regexp.MustCompile(`^([^!]+)!([A-Z]+)(\d+)(?::([A-Z]+)(\d+))?$`)

// columnToIndex converts column letters (e.g. "A", "AA") to a 1-based index
func columnToIndex(col string) int {
	col = strings.ToUpper(col)
	idx := 0
	for i := 0; i < len(col); i++ {
		idx = idx*26 + int(col[i]-'A') + 1
	}
	return idx
}

// atoi wraps strconv.Atoi and logs errors without exiting
func (s *sheetUseCase) atoi(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		s.log.Printf("atoi func failed: %s", err)
	}
	return n
}

// buildValueRanges constructs ValueRange slices from UpdateRequests
func (s *sheetUseCase) buildValueRanges(updates []models.UpdateRequest) ([]*sheets.ValueRange, error) {
	valueRanges := make([]*sheets.ValueRange, 0, len(updates))

	for _, u := range updates {
		vr, err := s.buildSingleValueRange(u, rangeRegexp)
		if err != nil {
			return nil, err
		}
		valueRanges = append(valueRanges, vr)
	}

	return valueRanges, nil
}

// buildSingleValueRange parses a single UpdateRequest into a ValueRange
func (s *sheetUseCase) buildSingleValueRange(u models.UpdateRequest, re *regexp.Regexp) (*sheets.ValueRange, error) {
	rangeStr := u.Range
	if !strings.Contains(rangeStr, "!") {
		rangeStr = "Sheet1!" + rangeStr
	}

	m := re.FindStringSubmatch(rangeStr)
	if m == nil {
		return nil, fmt.Errorf("invalid range format: %s", u.Range)
	}

	sheetName := m[1]
	startColStr, startRowStr := m[2], m[3]
	endColStr, endRowStr := m[4], m[5]

	if endColStr == "" {
		endColStr, endRowStr = startColStr, startRowStr
	}

	startCol := columnToIndex(startColStr)
	endCol := columnToIndex(endColStr)
	startRow := s.atoi(startRowStr)
	endRow := s.atoi(endRowStr)

	rowCount := endRow - startRow + 1
	colCount := endCol - startCol + 1

	if rowCount <= 0 || colCount <= 0 {
		return nil, fmt.Errorf("invalid range: start > end")
	}

	values := generateValues(rowCount, colCount, u.Value)

	return &sheets.ValueRange{
		Range:          fmt.Sprintf("%s!%s%d:%s%d", sheetName, startColStr, startRow, endColStr, endRow),
		MajorDimension: "ROWS",
		Values:         values,
	}, nil
}

// generateValues fills a rows×cols matrix with the same val
func generateValues(rows, cols int, val interface{}) [][]interface{} {
	values := make([][]interface{}, rows)
	for i := 0; i < rows; i++ {
		row := make([]interface{}, cols)
		for j := 0; j < cols; j++ {
			row[j] = val
		}
		values[i] = row
	}
	return values
}

// grantSheetAccess grants Drive permissions using an absolute, cleaned credential path
func (s *sheetUseCase) grantSheetAccess(ctx context.Context, sheetID string) error {
	return s.permSetter.SetPermission(ctx, sheetID)
}

func (r *realPermissionSetter) SetPermission(ctx context.Context, sheetID string) error {
	credPath := r.CredPath
	if _, err := os.Stat(credPath); err != nil {
		credPath = "." + credPath
	}
	srv, err := drive.NewService(ctx, option.WithCredentialsFile(credPath))
	if err != nil {
		return fmt.Errorf("failed to create Google Drive client: %w", err)
	}

	perm := &drive.Permission{
		Type:         r.Type,
		Role:         r.Role,
		EmailAddress: r.EmailAddress,
	}

	_, err = srv.Permissions.Create(sheetID, perm).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("failed to set permission: %w", err)
	}

	return nil
}
func (s *sheetUseCase) SetPermissionSetter(ps PermissionSetter) {
	s.permSetter = ps
}
