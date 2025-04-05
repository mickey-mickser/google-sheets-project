package usecase

import (
	"reflect"

	"google.golang.org/api/sheets/v4"
)

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

func (s *sheetUseCase) logRequestTypes(requests []*sheets.Request) {
	for _, req := range requests {
		v := reflect.ValueOf(*req)
		t := v.Type()

		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			if !field.IsNil() {
				s.log.Printf("batchUpdate use: %s", t.Field(i).Name)
			}
		}
	}
}
