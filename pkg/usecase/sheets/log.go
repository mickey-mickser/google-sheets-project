package sheetUsecase

import (
	"github.com/mickey-mickser/google-sheets-project/pkg/models"
	"github.com/sirupsen/logrus"
	"reflect"
	"strings"

	"google.golang.org/api/sheets/v4"
)

const (
	// RequestTypeAdd represents add operations
	RequestTypeAdd = "Add"
	// RequestTypeUpdate represents update operations
	RequestTypeUpdate = "Update"
	// RequestTypeDelete represents delete operations
	RequestTypeDelete = "Delete"
)

type RequestMetrics struct {
	TotalRequests  int
	AddRequests    int
	UpdateRequests int
	DeleteRequests int
	OtherRequests  int
}

// RequestBatchUpdate represents all possible Google Sheets API batch update operations.
// Each field corresponds to a specific operation type that can be performed on a spreadsheet.
type RequestBatchUpdate struct {
	AddBanding                   *sheets.AddBandingRequest                   `json:"addBanding,omitempty"`
	AddChart                     *sheets.AddChartRequest                     `json:"addChart,omitempty"`
	AddConditionalFormatRule     *sheets.AddConditionalFormatRuleRequest     `json:"addConditionalFormatRule,omitempty"`
	AddDataSource                *sheets.AddDataSourceRequest                `json:"addDataSource,omitempty"`
	AddDimensionGroup            *sheets.AddDimensionGroupRequest            `json:"addDimensionGroup,omitempty"`
	AddFilterView                *sheets.AddFilterViewRequest                `json:"addFilterView,omitempty"`
	AddNamedRange                *sheets.AddNamedRangeRequest                `json:"addNamedRange,omitempty"`
	AddProtectedRange            *sheets.AddProtectedRangeRequest            `json:"addProtectedRange,omitempty"`
	AddSheet                     *sheets.AddSheetRequest                     `json:"addSheet,omitempty"`
	AddSlicer                    *sheets.AddSlicerRequest                    `json:"addSlicer,omitempty"`
	AppendCells                  *sheets.AppendCellsRequest                  `json:"appendCells,omitempty"`
	AppendDimension              *sheets.AppendDimensionRequest              `json:"appendDimension,omitempty"`
	AutoFill                     *sheets.AutoFillRequest                     `json:"autoFill,omitempty"`
	AutoResizeDimensions         *sheets.AutoResizeDimensionsRequest         `json:"autoResizeDimensions,omitempty"`
	CancelDataSourceRefresh      *sheets.CancelDataSourceRefreshRequest      `json:"cancelDataSourceRefresh,omitempty"`
	ClearBasicFilter             *sheets.ClearBasicFilterRequest             `json:"clearBasicFilter,omitempty"`
	CopyPaste                    *sheets.CopyPasteRequest                    `json:"copyPaste,omitempty"`
	CreateDeveloperMetadata      *sheets.CreateDeveloperMetadataRequest      `json:"createDeveloperMetadata,omitempty"`
	CutPaste                     *sheets.CutPasteRequest                     `json:"cutPaste,omitempty"`
	DeleteBanding                *sheets.DeleteBandingRequest                `json:"deleteBanding,omitempty"`
	DeleteConditionalFormatRule  *sheets.DeleteConditionalFormatRuleRequest  `json:"deleteConditionalFormatRule,omitempty"`
	DeleteDataSource             *sheets.DeleteDataSourceRequest             `json:"deleteDataSource,omitempty"`
	DeleteDeveloperMetadata      *sheets.DeleteDeveloperMetadataRequest      `json:"deleteDeveloperMetadata,omitempty"`
	DeleteDimension              *sheets.DeleteDimensionRequest              `json:"deleteDimension,omitempty"`
	DeleteDimensionGroup         *sheets.DeleteDimensionGroupRequest         `json:"deleteDimensionGroup,omitempty"`
	DeleteDuplicates             *sheets.DeleteDuplicatesRequest             `json:"deleteDuplicates,omitempty"`
	DeleteEmbeddedObject         *sheets.DeleteEmbeddedObjectRequest         `json:"deleteEmbeddedObject,omitempty"`
	DeleteFilterView             *sheets.DeleteFilterViewRequest             `json:"deleteFilterView,omitempty"`
	DeleteNamedRange             *sheets.DeleteNamedRangeRequest             `json:"deleteNamedRange,omitempty"`
	DeleteProtectedRange         *sheets.DeleteProtectedRangeRequest         `json:"deleteProtectedRange,omitempty"`
	DeleteRange                  *sheets.DeleteRangeRequest                  `json:"deleteRange,omitempty"`
	DeleteSheet                  *sheets.DeleteSheetRequest                  `json:"deleteSheet,omitempty"`
	DuplicateFilterView          *sheets.DuplicateFilterViewRequest          `json:"duplicateFilterView,omitempty"`
	DuplicateSheet               *sheets.DuplicateSheetRequest               `json:"duplicateSheet,omitempty"`
	FindReplace                  *sheets.FindReplaceRequest                  `json:"findReplace,omitempty"`
	InsertDimension              *sheets.InsertDimensionRequest              `json:"insertDimension,omitempty"`
	InsertRange                  *sheets.InsertRangeRequest                  `json:"insertRange,omitempty"`
	MergeCells                   *sheets.MergeCellsRequest                   `json:"mergeCells,omitempty"`
	MoveDimension                *sheets.MoveDimensionRequest                `json:"moveDimension,omitempty"`
	PasteData                    *sheets.PasteDataRequest                    `json:"pasteData,omitempty"`
	RandomizeRange               *sheets.RandomizeRangeRequest               `json:"randomizeRange,omitempty"`
	RefreshDataSource            *sheets.RefreshDataSourceRequest            `json:"refreshDataSource,omitempty"`
	RepeatCell                   *sheets.RepeatCellRequest                   `json:"repeatCell,omitempty"`
	SetBasicFilter               *sheets.SetBasicFilterRequest               `json:"setBasicFilter,omitempty"`
	SetDataValidation            *sheets.SetDataValidationRequest            `json:"setDataValidation,omitempty"`
	SortRange                    *sheets.SortRangeRequest                    `json:"sortRange,omitempty"`
	TextToColumns                *sheets.TextToColumnsRequest                `json:"textToColumns,omitempty"`
	TrimWhitespace               *sheets.TrimWhitespaceRequest               `json:"trimWhitespace,omitempty"`
	UnmergeCells                 *sheets.UnmergeCellsRequest                 `json:"unmergeCells,omitempty"`
	UpdateBanding                *sheets.UpdateBandingRequest                `json:"updateBanding,omitempty"`
	UpdateBorders                *sheets.UpdateBordersRequest                `json:"updateBorders,omitempty"`
	UpdateCells                  *sheets.UpdateCellsRequest                  `json:"updateCells,omitempty"`
	UpdateChartSpec              *sheets.UpdateChartSpecRequest              `json:"updateChartSpec,omitempty"`
	UpdateConditionalFormatRule  *sheets.UpdateConditionalFormatRuleRequest  `json:"updateConditionalFormatRule,omitempty"`
	UpdateDataSource             *sheets.UpdateDataSourceRequest             `json:"updateDataSource,omitempty"`
	UpdateDeveloperMetadata      *sheets.UpdateDeveloperMetadataRequest      `json:"updateDeveloperMetadata,omitempty"`
	UpdateDimensionGroup         *sheets.UpdateDimensionGroupRequest         `json:"updateDimensionGroup,omitempty"`
	UpdateDimensionProperties    *sheets.UpdateDimensionPropertiesRequest    `json:"updateDimensionProperties,omitempty"`
	UpdateEmbeddedObjectBorder   *sheets.UpdateEmbeddedObjectBorderRequest   `json:"updateEmbeddedObjectBorder,omitempty"`
	UpdateEmbeddedObjectPosition *sheets.UpdateEmbeddedObjectPositionRequest `json:"updateEmbeddedObjectPosition,omitempty"`
	UpdateFilterView             *sheets.UpdateFilterViewRequest             `json:"updateFilterView,omitempty"`
	UpdateNamedRange             *sheets.UpdateNamedRangeRequest             `json:"updateNamedRange,omitempty"`
	UpdateProtectedRange         *sheets.UpdateProtectedRangeRequest         `json:"updateProtectedRange,omitempty"`
	UpdateSheetProperties        *sheets.UpdateSheetPropertiesRequest        `json:"updateSheetProperties,omitempty"`
	UpdateSlicerSpec             *sheets.UpdateSlicerSpecRequest             `json:"updateSlicerSpec,omitempty"`
	UpdateSpreadsheetProperties  *sheets.UpdateSpreadsheetPropertiesRequest  `json:"updateSpreadsheetProperties,omitempty"`
}

// logRequestTypes logs non-empty fields of each UpdateRequest in the batch for debugging.
func (s *sheetUseCase) logRequestTypes(requests []models.UpdateRequest) {
	if len(requests) == 0 {
		s.log.Debug("No batch update requests to process")
		return
	}

	for i, req := range requests {
		v := reflect.ValueOf(req)
		t := v.Type()

		var nonEmptyFields []string
		for j := 0; j < v.NumField(); j++ {
			field := v.Field(j)
			fieldType := t.Field(j)
			kind := field.Kind()

			switch kind {
			case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map, reflect.Func, reflect.Chan:
				if !field.IsNil() {
					nonEmptyFields = append(nonEmptyFields, fieldType.Name)
				}
			case reflect.String:
				if field.String() != "" {
					nonEmptyFields = append(nonEmptyFields, fieldType.Name)
				}
			default:
				zero := reflect.Zero(field.Type()).Interface()
				if !reflect.DeepEqual(field.Interface(), zero) {
					nonEmptyFields = append(nonEmptyFields, fieldType.Name)
				}
			}
		}

		if len(nonEmptyFields) == 0 {
			s.log.WithFields(logrus.Fields{
				"request_index": i,
				"request_count": len(requests),
			}).Info("Processing batch update request with no set fields")
		} else {
			s.log.WithFields(logrus.Fields{
				"request_index": i,
				"request_count": len(requests),
				"request_types": nonEmptyFields,
			}).Info("Processing batch update request")
		}
	}
}

// GetActiveRequests returns a list of names of those RequestBatchUpdate fields that are not nil (i.e. actually present in the request).
func (r *RequestBatchUpdate) GetActiveRequests() []string {
	var active []string
	v := reflect.ValueOf(r).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if !v.Field(i).IsNil() {
			active = append(active, t.Field(i).Name)
		}
	}
	return active
}

// GetRequestCategory assigns the query name to one of the categories.
func GetRequestCategory(requestType string) string {
	switch {
	case strings.HasPrefix(requestType, RequestTypeAdd):
		return RequestTypeAdd
	case strings.HasPrefix(requestType, RequestTypeUpdate):
		return RequestTypeUpdate
	case strings.HasPrefix(requestType, RequestTypeDelete):
		return RequestTypeDelete
	default:
		return "Other"
	}
}

// CollectMetrics calculates metrics for existing operations.
func (r *RequestBatchUpdate) CollectMetrics() RequestMetrics {
	metrics := RequestMetrics{}
	active := r.GetActiveRequests()
	metrics.TotalRequests = len(active)

	for _, reqName := range active {
		switch GetRequestCategory(reqName) {
		case RequestTypeAdd:
			metrics.AddRequests++
		case RequestTypeUpdate:
			metrics.UpdateRequests++
		case RequestTypeDelete:
			metrics.DeleteRequests++
		default:
			metrics.OtherRequests++
		}
	}
	return metrics
}
