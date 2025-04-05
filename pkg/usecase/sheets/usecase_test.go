package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	mockusecase "github.com/mickey-mickser/google-sheets-project/pkg/usecase/sheets/mocks"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/sheets/v4"
)

func TestSheetUseCase_BatchUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSheetUseCase := mockusecase.NewMockSheetUseCase(ctrl)

	jsonBody := []byte(`{"requests": []}`)

	mockSheetUseCase.EXPECT().BatchUpdate(context.Background(), "sheetId", jsonBody).
		Return(&sheets.BatchUpdateSpreadsheetResponse{}, nil).
		Times(1)

	resp, err := mockSheetUseCase.BatchUpdate(context.Background(), "sheetId", jsonBody)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestSheetUseCase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSheetUseCase := mockusecase.NewMockSheetUseCase(ctrl)

	jsonBody := []byte(`{"properties": {"title": "Test Spreadsheet"}}`)

	mockSheetUseCase.EXPECT().Create(context.Background(), jsonBody).
		Return(&sheets.Spreadsheet{SpreadsheetId: "1234", SpreadsheetUrl: "https://sheets.google.com/xyz"}, nil).
		Times(1)

	resp, err := mockSheetUseCase.Create(context.Background(), jsonBody)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "1234", resp.SpreadsheetId)
}

func TestSheetUseCase_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSheetUseCase := mockusecase.NewMockSheetUseCase(ctrl)

	sheetId := "sheetId"
	param := "A1:Z1000"

	mockSheetUseCase.EXPECT().Get(context.Background(), sheetId, param).
		Return([][]interface{}{{"value1"}}, nil).
		Times(1)

	resp, err := mockSheetUseCase.Get(context.Background(), sheetId, param)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, [][]interface{}{{"value1"}}, resp)
}

func TestSheetUseCase_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSheetUseCase := mockusecase.NewMockSheetUseCase(ctrl)

	jsonBody := []byte(`{}`)

	mockSheetUseCase.EXPECT().Delete(context.Background(), "sheetId", "A1:Z1000", jsonBody).
		Return(&sheets.ClearValuesResponse{}, nil).
		Times(1)

	resp, err := mockSheetUseCase.Delete(context.Background(), "sheetId", "A1:Z1000", jsonBody)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
