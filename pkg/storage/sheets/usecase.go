package usecase

import (
	"context"
	"fmt"
	"google.golang.org/api/sheets/v4"
)

type SheetUseCase interface {
	BatchUpdate(ctx context.Context, sheetId string, operations []map[string]interface{}) (*sheets.BatchUpdateSpreadsheetResponse, error)
}
type sheetUseCase struct {
	cli *sheets.Service
}

func NewSheetUse(cli *sheets.Service) SheetUseCase {
	return &sheetUseCase{cli: cli}
}
func (s *sheetUseCase) BatchUpdate(ctx context.Context, sheetId string, operations []map[string]interface{}) (*sheets.BatchUpdateSpreadsheetResponse, error) {
	batchUpdateRequest, err := createBatchUpdateRequest(operations)

	if err != nil {
		return nil, fmt.Errorf("create batch update spreadsheet request: %w", err)
	}
	resp, err := s.cli.Spreadsheets.BatchUpdate(sheetId, batchUpdateRequest).Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("error while updating data in Google Sheets: %w", err)
	}

	return resp, nil
}

func createBatchUpdateRequest(operations []map[string]interface{}) (*sheets.BatchUpdateSpreadsheetRequest, error) {
	var batchUpdateRequest sheets.BatchUpdateSpreadsheetRequest
	for _, op := range operations {
		for key, value := range op {
			switch key {
			case "updateCells":
				updateCells, ok := value.(map[string]interface{})
				if !ok {
					fmt.Println("Error: Failed to convert to map")
					continue
				}

				rangeValue, ok := updateCells["range"].(map[string]interface{})
				if !ok {
					fmt.Println("Error: Could not convert 'range' to map")
					continue
				}
				fmt.Println(updateCells["rows"])
				rowsValue, ok := updateCells["rows"].([]interface{})
				if !ok {
					fmt.Println("Error: Could not convert 'rows' to interface slice")
					continue
				}

				var rowData []*sheets.RowData

				for _, rows := range rowsValue {
					rowMap, ok := rows.(map[string]interface{})
					if !ok {
						fmt.Println("Error: Failed to convert row to map")
						continue
					}

					values, ok := rowMap["values"].([]interface{})
					if !ok {
						fmt.Println("Error: Could not convert 'values' to interface slice")
						continue
					}

					var cellValues []*sheets.CellData
					for _, value := range values {
						cellData, ok := value.(map[string]interface{})
						if !ok {
							fmt.Println("Error: Failed to convert value to map")
							continue
						}

						userEnteredValue, ok := cellData["userEnteredValue"].(map[string]interface{})
						if !ok {
							fmt.Println("Error: Failed to retrieve 'userEnteredValue'")
							continue
						}

						stringValue, ok := userEnteredValue["stringValue"].(string)
						if !ok {
							fmt.Println("Error: Failed to extract 'stringValue'")
							continue
						}

						cellValues = append(cellValues, &sheets.CellData{
							UserEnteredValue: &sheets.ExtendedValue{StringValue: &stringValue},
						})
					}

					rowData = append(rowData, &sheets.RowData{
						Values: cellValues,
					})
				}

				updateRequest := &sheets.Request{
					UpdateCells: &sheets.UpdateCellsRequest{
						Range: &sheets.GridRange{
							SheetId:          toInt64(rangeValue["sheetId"]),
							StartRowIndex:    toInt64(rangeValue["startRowIndex"]),
							EndRowIndex:      toInt64(rangeValue["endRowIndex"]),
							StartColumnIndex: toInt64(rangeValue["startColumnIndex"]),
							EndColumnIndex:   toInt64(rangeValue["endColumnIndex"]),
						},
						Rows:   rowData,
						Fields: "*",
					},
				}

				batchUpdateRequest.Requests = append(batchUpdateRequest.Requests, updateRequest)
			case "appendCells":
				appendCells, ok := value.(map[string]interface{})
				if !ok {
					fmt.Println("Error: Cold note envelope range to map")
					continue
				}
				rangeValue, ok := appendCells["range"].(map[string]interface{})
				if !ok {
					fmt.Println("Error: Could not convert 'range' to map")
					continue
				}
				values, ok := rangeValue["values"].([]interface{})
				if !ok {
					fmt.Println("Error: 'values' is not an array")
					continue
				}

				var rowData []*sheets.RowData
				for _, v := range values {
					cellValues, ok := v.([]interface{})
					if !ok {
						fmt.Println("Error: Element 'values' is not an array")
						continue
					}

					var cells []*sheets.CellData
					for _, cell := range cellValues {
						strVal, ok := cell.(string)
						if !ok {
							fmt.Println("Error: Cell value is not a string")
							continue
						}

						cells = append(cells, &sheets.CellData{
							UserEnteredValue: &sheets.ExtendedValue{StringValue: &strVal},
						})
					}

					rowData = append(rowData, &sheets.RowData{Values: cells})
				}

				appendRequest := &sheets.Request{
					AppendCells: &sheets.AppendCellsRequest{
						SheetId: rangeValue["sheetId"].(int64),
						Rows:    rowData,
						Fields:  "userEnteredValue",
					},
				}
				batchUpdateRequest.Requests = append(batchUpdateRequest.Requests, appendRequest)
			case "deleteDimension":
				deleteDimension, ok := value.(map[string]interface{})
				if !ok {
					fmt.Println("Error: Failed to convert 'deleteDimension' in map")
					continue
				}
				rangeValue, ok := deleteDimension["range"].(map[string]interface{})
				if !ok {
					fmt.Println("Error: Failed to convert 'range' in map")

				}

				deleteRequest := &sheets.Request{
					DeleteDimension: &sheets.DeleteDimensionRequest{
						Range: &sheets.DimensionRange{

							SheetId:    toInt64(rangeValue["sheetId"]),
							Dimension:  toString(rangeValue["dimension"]),
							StartIndex: toInt64(rangeValue["startIndex"]),
							EndIndex:   toInt64(rangeValue["endIndex"]),
						},
					},
				}
				batchUpdateRequest.Requests = append(batchUpdateRequest.Requests, deleteRequest)
			default:
				return nil, fmt.Errorf("unsupported operation: %v\n", key)
			}
		}
	}

	return &batchUpdateRequest, nil
}
func toInt64(value interface{}) int64 {
	switch v := value.(type) {
	case float64:
		return int64(v)
	case int:
		return int64(v)
	case int64:
		return v
	default:
		fmt.Println("Cast to int64 error:", value)
		return 0
	}
}

func toString(value interface{}) string {
	str, ok := value.(string)
	if !ok {
		fmt.Println("Cast to string error:", value)
		return ""
	}
	return str
}
