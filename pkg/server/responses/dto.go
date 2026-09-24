package responses

import (
	"net/http"
)

type ClearValuesResponse struct {
	ClearedRange    string `json:"clearedRange,omitempty"`
	SpreadsheetId   string `json:"spreadsheetId,omitempty"`
	ServerResponse  `json:"-"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

type BatchUpdateValuesResponse struct {
	Responses           []*UpdateValuesResponse `json:"responses,omitempty"`
	SpreadsheetId       string                  `json:"spreadsheetId,omitempty"`
	TotalUpdatedCells   int64                   `json:"totalUpdatedCells,omitempty"`
	TotalUpdatedColumns int64                   `json:"totalUpdatedColumns,omitempty"`
	TotalUpdatedRows    int64                   `json:"totalUpdatedRows,omitempty"`
	TotalUpdatedSheets  int64                   `json:"totalUpdatedSheets,omitempty"`
	ServerResponse      `json:"-"`
	ForceSendFields     []string `json:"-"`
	NullFields          []string `json:"-"`
}
type UpdateValuesResponse struct {
	SpreadsheetId   string      `json:"spreadsheetId,omitempty"`
	UpdatedCells    int64       `json:"updatedCells,omitempty"`
	UpdatedColumns  int64       `json:"updatedColumns,omitempty"`
	UpdatedData     *ValueRange `json:"updatedData,omitempty"`
	UpdatedRange    string      `json:"updatedRange,omitempty"`
	UpdatedRows     int64       `json:"updatedRows,omitempty"`
	ServerResponse  `json:"-"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type ValueRange struct {
	MajorDimension  string          `json:"majorDimension,omitempty"`
	Range           string          `json:"range,omitempty"`
	Values          [][]interface{} `json:"values,omitempty"`
	ServerResponse  `json:"-"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type ServerResponse struct {
	HTTPStatusCode int
	Header         http.Header
}

type Spreadsheet struct {
	DataSourceSchedules []*DataSourceRefreshSchedule `json:"dataSourceSchedules,omitempty"`
	DataSources         []*DataSource                `json:"dataSources,omitempty"`
	DeveloperMetadata   []*DeveloperMetadata         `json:"developerMetadata,omitempty"`
	NamedRanges         []*NamedRange                `json:"namedRanges,omitempty"`
	Properties          *SpreadsheetProperties       `json:"properties,omitempty"`
	Sheets              []*Sheet                     `json:"sheets,omitempty"`
	SpreadsheetId       string                       `json:"spreadsheetId,omitempty"`
	SpreadsheetUrl      string                       `json:"spreadsheetUrl,omitempty"`
	ServerResponse      `json:"-"`
	ForceSendFields     []string `json:"-"`
	NullFields          []string `json:"-"`
}
type DataSourceRefreshSchedule struct {
	DailySchedule   *DataSourceRefreshDailySchedule   `json:"dailySchedule,omitempty"`
	Enabled         bool                              `json:"enabled,omitempty"`
	MonthlySchedule *DataSourceRefreshMonthlySchedule `json:"monthlySchedule,omitempty"`
	NextRun         *Interval                         `json:"nextRun,omitempty"`
	RefreshScope    string                            `json:"refreshScope,omitempty"`
	WeeklySchedule  *DataSourceRefreshWeeklySchedule  `json:"weeklySchedule,omitempty"`
	ForceSendFields []string                          `json:"-"`
	NullFields      []string                          `json:"-"`
}
type DataSourceRefreshDailySchedule struct {
	StartTime       *TimeOfDay `json:"startTime,omitempty"`
	ForceSendFields []string   `json:"-"`
	NullFields      []string   `json:"-"`
}
type TimeOfDay struct {
	Hours           int64    `json:"hours,omitempty"`
	Minutes         int64    `json:"minutes,omitempty"`
	Nanos           int64    `json:"nanos,omitempty"`
	Seconds         int64    `json:"seconds,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type DataSourceRefreshMonthlySchedule struct {
	DaysOfMonth     []int64    `json:"daysOfMonth,omitempty"`
	StartTime       *TimeOfDay `json:"startTime,omitempty"`
	ForceSendFields []string   `json:"-"`
	NullFields      []string   `json:"-"`
}
type Interval struct {
	EndTime         string   `json:"endTime,omitempty"`
	StartTime       string   `json:"startTime,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type DataSourceRefreshWeeklySchedule struct {
	DaysOfWeek      []string   `json:"daysOfWeek,omitempty"`
	StartTime       *TimeOfDay `json:"startTime,omitempty"`
	ForceSendFields []string   `json:"-"`
	NullFields      []string   `json:"-"`
}
type DataSourceColumn struct {
	Formula         string                     `json:"formula,omitempty"`
	Reference       *DataSourceColumnReference `json:"reference,omitempty"`
	ForceSendFields []string                   `json:"-"`
	NullFields      []string                   `json:"-"`
}
type DataSource struct {
	CalculatedColumns []*DataSourceColumn `json:"calculatedColumns,omitempty"`
	DataSourceId      string              `json:"dataSourceId,omitempty"`
	SheetId           int64               `json:"sheetId,omitempty"`
	Spec              *DataSourceSpec     `json:"spec,omitempty"`
	ForceSendFields   []string            `json:"-"`
	NullFields        []string            `json:"-"`
}

type DataSourceColumnReference struct {
	Name            string   `json:"name,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type DataSourceSpec struct {
	BigQuery        *BigQueryDataSourceSpec `json:"bigQuery,omitempty"`
	Looker          *LookerDataSourceSpec   `json:"looker,omitempty"`
	Parameters      []*DataSourceParameter  `json:"parameters,omitempty"`
	ForceSendFields []string                `json:"-"`
	NullFields      []string                `json:"-"`
}
type BigQueryDataSourceSpec struct {
	ProjectId       string             `json:"projectId,omitempty"`
	QuerySpec       *BigQueryQuerySpec `json:"querySpec,omitempty"`
	TableSpec       *BigQueryTableSpec `json:"tableSpec,omitempty"`
	ForceSendFields []string           `json:"-"`
	NullFields      []string           `json:"-"`
}
type BigQueryQuerySpec struct {
	RawQuery        string   `json:"rawQuery,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type LookerDataSourceSpec struct {
	Explore         string   `json:"explore,omitempty"`
	InstanceUri     string   `json:"instanceUri,omitempty"`
	Model           string   `json:"model,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type BigQueryTableSpec struct {
	DatasetId       string   `json:"datasetId,omitempty"`
	TableId         string   `json:"tableId,omitempty"`
	TableProjectId  string   `json:"tableProjectId,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type DataSourceParameter struct {
	Name            string     `json:"name,omitempty"`
	NamedRangeId    string     `json:"namedRangeId,omitempty"`
	Range           *GridRange `json:"range,omitempty"`
	ForceSendFields []string   `json:"-"`
	NullFields      []string   `json:"-"`
}
type GridRange struct {
	EndColumnIndex   int64    `json:"endColumnIndex,omitempty"`
	EndRowIndex      int64    `json:"endRowIndex,omitempty"`
	SheetId          int64    `json:"sheetId,omitempty"`
	StartColumnIndex int64    `json:"startColumnIndex,omitempty"`
	StartRowIndex    int64    `json:"startRowIndex,omitempty"`
	ForceSendFields  []string `json:"-"`
	NullFields       []string `json:"-"`
}
type DeveloperMetadata struct {
	Location        *DeveloperMetadataLocation `json:"location,omitempty"`
	MetadataId      int64                      `json:"metadataId,omitempty"`
	MetadataKey     string                     `json:"metadataKey,omitempty"`
	MetadataValue   string                     `json:"metadataValue,omitempty"`
	Visibility      string                     `json:"visibility,omitempty"`
	ServerResponse  `json:"-"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type DeveloperMetadataLocation struct {
	DimensionRange  *DimensionRange `json:"dimensionRange,omitempty"`
	LocationType    string          `json:"locationType,omitempty"`
	SheetId         int64           `json:"sheetId,omitempty"`
	Spreadsheet     bool            `json:"spreadsheet,omitempty"`
	ForceSendFields []string        `json:"-"`
	NullFields      []string        `json:"-"`
}
type DimensionRange struct {
	Dimension       string   `json:"dimension,omitempty"`
	EndIndex        int64    `json:"endIndex,omitempty"`
	SheetId         int64    `json:"sheetId,omitempty"`
	StartIndex      int64    `json:"startIndex,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type NamedRange struct {
	Name            string     `json:"name,omitempty"`
	NamedRangeId    string     `json:"namedRangeId,omitempty"`
	Range           *GridRange `json:"range,omitempty"`
	ForceSendFields []string   `json:"-"`
	NullFields      []string   `json:"-"`
}
type SpreadsheetProperties struct {
	AutoRecalc                              string                        `json:"autoRecalc,omitempty"`
	DefaultFormat                           *CellFormat                   `json:"defaultFormat,omitempty"`
	ImportFunctionsExternalUrlAccessAllowed bool                          `json:"importFunctionsExternalUrlAccessAllowed,omitempty"`
	IterativeCalculationSettings            *IterativeCalculationSettings `json:"iterativeCalculationSettings,omitempty"`
	Locale                                  string                        `json:"locale,omitempty"`
	SpreadsheetTheme                        *SpreadsheetTheme             `json:"spreadsheetTheme,omitempty"`
	TimeZone                                string                        `json:"timeZone,omitempty"`
	Title                                   string                        `json:"title,omitempty"`
	ForceSendFields                         []string                      `json:"-"`
	NullFields                              []string                      `json:"-"`
}
type CellFormat struct {
	BackgroundColor      *Color        `json:"backgroundColor,omitempty"`
	BackgroundColorStyle *ColorStyle   `json:"backgroundColorStyle,omitempty"`
	Borders              *Borders      `json:"borders,omitempty"`
	HorizontalAlignment  string        `json:"horizontalAlignment,omitempty"`
	HyperlinkDisplayType string        `json:"hyperlinkDisplayType,omitempty"`
	NumberFormat         *NumberFormat `json:"numberFormat,omitempty"`
	Padding              *Padding      `json:"padding,omitempty"`
	TextDirection        string        `json:"textDirection,omitempty"`
	TextFormat           *TextFormat   `json:"textFormat,omitempty"`
	TextRotation         *TextRotation `json:"textRotation,omitempty"`
	VerticalAlignment    string        `json:"verticalAlignment,omitempty"`
	WrapStrategy         string        `json:"wrapStrategy,omitempty"`
	ForceSendFields      []string      `json:"-"`
	NullFields           []string      `json:"-"`
}
type Color struct {
	Alpha           float64  `json:"alpha,omitempty"`
	Blue            float64  `json:"blue,omitempty"`
	Green           float64  `json:"green,omitempty"`
	Red             float64  `json:"red,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type ColorStyle struct {
	RgbColor        *Color   `json:"rgbColor,omitempty"`
	ThemeColor      string   `json:"themeColor,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type Borders struct {
	Bottom          *Border  `json:"bottom,omitempty"`
	Left            *Border  `json:"left,omitempty"`
	Right           *Border  `json:"right,omitempty"`
	Top             *Border  `json:"top,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type Border struct {
	Color           *Color      `json:"color,omitempty"`
	ColorStyle      *ColorStyle `json:"colorStyle,omitempty"`
	Style           string      `json:"style,omitempty"`
	Width           int64       `json:"width,omitempty"`
	ForceSendFields []string    `json:"-"`
	NullFields      []string    `json:"-"`
}
type NumberFormat struct {
	Pattern         string   `json:"pattern,omitempty"`
	Type            string   `json:"type,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type Padding struct {
	Bottom          int64    `json:"bottom,omitempty"`
	Left            int64    `json:"left,omitempty"`
	Right           int64    `json:"right,omitempty"`
	Top             int64    `json:"top,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type TextRotation struct {
	Angle           int64    `json:"angle,omitempty"`
	Vertical        bool     `json:"vertical,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type TextFormat struct {
	Bold                 bool        `json:"bold,omitempty"`
	FontFamily           string      `json:"fontFamily,omitempty"`
	FontSize             int64       `json:"fontSize,omitempty"`
	ForegroundColor      *Color      `json:"foregroundColor,omitempty"`
	ForegroundColorStyle *ColorStyle `json:"foregroundColorStyle,omitempty"`
	Italic               bool        `json:"italic,omitempty"`
	Link                 *Link       `json:"link,omitempty"`
	Strikethrough        bool        `json:"strikethrough,omitempty"`
	Underline            bool        `json:"underline,omitempty"`
	ForceSendFields      []string    `json:"-"`
	NullFields           []string    `json:"-"`
}

type Link struct {
	Uri             string   `json:"uri,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type Sheet struct {
	BandedRanges       []*BandedRange           `json:"bandedRanges,omitempty"`
	BasicFilter        *BasicFilter             `json:"basicFilter,omitempty"`
	Charts             []*EmbeddedChart         `json:"charts,omitempty"`
	ColumnGroups       []*DimensionGroup        `json:"columnGroups,omitempty"`
	ConditionalFormats []*ConditionalFormatRule `json:"conditionalFormats,omitempty"`
	Data               []*GridData              `json:"data,omitempty"`
	DeveloperMetadata  []*DeveloperMetadata     `json:"developerMetadata,omitempty"`
	FilterViews        []*FilterView            `json:"filterViews,omitempty"`
	Merges             []*GridRange             `json:"merges,omitempty"`
	Properties         *SheetProperties         `json:"properties,omitempty"`
	ProtectedRanges    []*ProtectedRange        `json:"protectedRanges,omitempty"`
	RowGroups          []*DimensionGroup        `json:"rowGroups,omitempty"`
	Slicers            []*Slicer                `json:"slicers,omitempty"`
	ForceSendFields    []string                 `json:"-"`
	NullFields         []string                 `json:"-"`
}
type BandedRange struct {
	BandedRangeId    int64              `json:"bandedRangeId,omitempty"`
	ColumnProperties *BandingProperties `json:"columnProperties,omitempty"`
	Range            *GridRange         `json:"range,omitempty"`
	RowProperties    *BandingProperties `json:"rowProperties,omitempty"`
	ForceSendFields  []string           `json:"-"`
	NullFields       []string           `json:"-"`
}
type BandingProperties struct {
	FirstBandColor       *Color      `json:"firstBandColor,omitempty"`
	FirstBandColorStyle  *ColorStyle `json:"firstBandColorStyle,omitempty"`
	FooterColor          *Color      `json:"footerColor,omitempty"`
	FooterColorStyle     *ColorStyle `json:"footerColorStyle,omitempty"`
	HeaderColor          *Color      `json:"headerColor,omitempty"`
	HeaderColorStyle     *ColorStyle `json:"headerColorStyle,omitempty"`
	SecondBandColor      *Color      `json:"secondBandColor,omitempty"`
	SecondBandColorStyle *ColorStyle `json:"secondBandColorStyle,omitempty"`
	ForceSendFields      []string    `json:"-"`
	NullFields           []string    `json:"-"`
}
type BasicFilter struct {
	Criteria        map[string]FilterCriteria `json:"criteria,omitempty"`
	FilterSpecs     []*FilterSpec             `json:"filterSpecs,omitempty"`
	Range           *GridRange                `json:"range,omitempty"`
	SortSpecs       []*SortSpec               `json:"sortSpecs,omitempty"`
	ForceSendFields []string                  `json:"-"`
	NullFields      []string                  `json:"-"`
}
type FilterCriteria struct {
	Condition                   *BooleanCondition `json:"condition,omitempty"`
	HiddenValues                []string          `json:"hiddenValues,omitempty"`
	VisibleBackgroundColor      *Color            `json:"visibleBackgroundColor,omitempty"`
	VisibleBackgroundColorStyle *ColorStyle       `json:"visibleBackgroundColorStyle,omitempty"`
	VisibleForegroundColor      *Color            `json:"visibleForegroundColor,omitempty"`
	VisibleForegroundColorStyle *ColorStyle       `json:"visibleForegroundColorStyle,omitempty"`
	ForceSendFields             []string          `json:"-"`
	NullFields                  []string          `json:"-"`
}
type FilterSpec struct {
	ColumnIndex               int64                      `json:"columnIndex,omitempty"`
	DataSourceColumnReference *DataSourceColumnReference `json:"dataSourceColumnReference,omitempty"`
	FilterCriteria            *FilterCriteria            `json:"filterCriteria,omitempty"`
	ForceSendFields           []string                   `json:"-"`
	NullFields                []string                   `json:"-"`
}

type BooleanCondition struct {
	Type            string            `json:"type,omitempty"`
	Values          []*ConditionValue `json:"values,omitempty"`
	ForceSendFields []string          `json:"-"`
	NullFields      []string          `json:"-"`
}
type ConditionValue struct {
	RelativeDate     string   `json:"relativeDate,omitempty"`
	UserEnteredValue string   `json:"userEnteredValue,omitempty"`
	ForceSendFields  []string `json:"-"`
	NullFields       []string `json:"-"`
}

type SortSpec struct {
	BackgroundColor           *Color                     `json:"backgroundColor,omitempty"`
	BackgroundColorStyle      *ColorStyle                `json:"backgroundColorStyle,omitempty"`
	DataSourceColumnReference *DataSourceColumnReference `json:"dataSourceColumnReference,omitempty"`
	DimensionIndex            int64                      `json:"dimensionIndex,omitempty"`
	ForegroundColor           *Color                     `json:"foregroundColor,omitempty"`
	ForegroundColorStyle      *ColorStyle                `json:"foregroundColorStyle,omitempty"`
	SortOrder                 string                     `json:"sortOrder,omitempty"`
	ForceSendFields           []string                   `json:"-"`
	NullFields                []string                   `json:"-"`
}
type EmbeddedChart struct {
	Border          *EmbeddedObjectBorder   `json:"border,omitempty"`
	ChartId         int64                   `json:"chartId,omitempty"`
	Position        *EmbeddedObjectPosition `json:"position,omitempty"`
	Spec            *ChartSpec              `json:"spec,omitempty"`
	ForceSendFields []string                `json:"-"`
	NullFields      []string                `json:"-"`
}
type EmbeddedObjectBorder struct {
	Color           *Color      `json:"color,omitempty"`
	ColorStyle      *ColorStyle `json:"colorStyle,omitempty"`
	ForceSendFields []string    `json:"-"`
	NullFields      []string    `json:"-"`
}
type EmbeddedObjectPosition struct {
	NewSheet        bool             `json:"newSheet,omitempty"`
	OverlayPosition *OverlayPosition `json:"overlayPosition,omitempty"`
	SheetId         int64            `json:"sheetId,omitempty"`
	ForceSendFields []string         `json:"-"`
	NullFields      []string         `json:"-"`
}
type OverlayPosition struct {
	AnchorCell      *GridCoordinate `json:"anchorCell,omitempty"`
	HeightPixels    int64           `json:"heightPixels,omitempty"`
	OffsetXPixels   int64           `json:"offsetXPixels,omitempty"`
	OffsetYPixels   int64           `json:"offsetYPixels,omitempty"`
	WidthPixels     int64           `json:"widthPixels,omitempty"`
	ForceSendFields []string        `json:"-"`
	NullFields      []string        `json:"-"`
}
type GridCoordinate struct {
	ColumnIndex     int64    `json:"columnIndex,omitempty"`
	RowIndex        int64    `json:"rowIndex,omitempty"`
	SheetId         int64    `json:"sheetId,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type ChartSpec struct {
	AltText                   string                     `json:"altText,omitempty"`
	BackgroundColor           *Color                     `json:"backgroundColor,omitempty"`
	BackgroundColorStyle      *ColorStyle                `json:"backgroundColorStyle,omitempty"`
	BasicChart                *BasicChartSpec            `json:"basicChart,omitempty"`
	BubbleChart               *BubbleChartSpec           `json:"bubbleChart,omitempty"`
	CandlestickChart          *CandlestickChartSpec      `json:"candlestickChart,omitempty"`
	DataSourceChartProperties *DataSourceChartProperties `json:"dataSourceChartProperties,omitempty"`
	FilterSpecs               []*FilterSpec              `json:"filterSpecs,omitempty"`
	FontName                  string                     `json:"fontName,omitempty"`
	HiddenDimensionStrategy   string                     `json:"hiddenDimensionStrategy,omitempty"`
	HistogramChart            *HistogramChartSpec        `json:"histogramChart,omitempty"`
	Maximized                 bool                       `json:"maximized,omitempty"`
	OrgChart                  *OrgChartSpec              `json:"orgChart,omitempty"`
	PieChart                  *PieChartSpec              `json:"pieChart,omitempty"`
	ScorecardChart            *ScorecardChartSpec        `json:"scorecardChart,omitempty"`
	SortSpecs                 []*SortSpec                `json:"sortSpecs,omitempty"`
	Subtitle                  string                     `json:"subtitle,omitempty"`
	SubtitleTextFormat        *TextFormat                `json:"subtitleTextFormat,omitempty"`
	SubtitleTextPosition      *TextPosition              `json:"subtitleTextPosition,omitempty"`
	Title                     string                     `json:"title,omitempty"`
	TitleTextFormat           *TextFormat                `json:"titleTextFormat,omitempty"`
	TitleTextPosition         *TextPosition              `json:"titleTextPosition,omitempty"`
	TreemapChart              *TreemapChartSpec          `json:"treemapChart,omitempty"`
	WaterfallChart            *WaterfallChartSpec        `json:"waterfallChart,omitempty"`
	ForceSendFields           []string                   `json:"-"`
	NullFields                []string                   `json:"-"`
}

type TextPosition struct {
	HorizontalAlignment string   `json:"horizontalAlignment,omitempty"`
	ForceSendFields     []string `json:"-"`
	NullFields          []string `json:"-"`
}
type BasicChartSpec struct {
	Axis             []*BasicChartAxis   `json:"axis,omitempty"`
	ChartType        string              `json:"chartType,omitempty"`
	CompareMode      string              `json:"compareMode,omitempty"`
	Domains          []*BasicChartDomain `json:"domains,omitempty"`
	HeaderCount      int64               `json:"headerCount,omitempty"`
	InterpolateNulls bool                `json:"interpolateNulls,omitempty"`
	LegendPosition   string              `json:"legendPosition,omitempty"`
	LineSmoothing    bool                `json:"lineSmoothing,omitempty"`
	Series           []*BasicChartSeries `json:"series,omitempty"`
	StackedType      string              `json:"stackedType,omitempty"`
	ThreeDimensional bool                `json:"threeDimensional,omitempty"`
	TotalDataLabel   *DataLabel          `json:"totalDataLabel,omitempty"`
	ForceSendFields  []string            `json:"-"`
	NullFields       []string            `json:"-"`
}

type BasicChartAxis struct {
	Format            *TextFormat                 `json:"format,omitempty"`
	Position          string                      `json:"position,omitempty"`
	Title             string                      `json:"title,omitempty"`
	TitleTextPosition *TextPosition               `json:"titleTextPosition,omitempty"`
	ViewWindowOptions *ChartAxisViewWindowOptions `json:"viewWindowOptions,omitempty"`
	ForceSendFields   []string                    `json:"-"`
	NullFields        []string                    `json:"-"`
}

type ChartAxisViewWindowOptions struct {
	ViewWindowMax   float64  `json:"viewWindowMax,omitempty"`
	ViewWindowMin   float64  `json:"viewWindowMin,omitempty"`
	ViewWindowMode  string   `json:"viewWindowMode,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type BasicChartDomain struct {
	Domain          *ChartData `json:"domain,omitempty"`
	Reversed        bool       `json:"reversed,omitempty"`
	ForceSendFields []string   `json:"-"`
	NullFields      []string   `json:"-"`
}

type ChartData struct {
	AggregateType   string                     `json:"aggregateType,omitempty"`
	ColumnReference *DataSourceColumnReference `json:"columnReference,omitempty"`
	GroupRule       *ChartGroupRule            `json:"groupRule,omitempty"`
	SourceRange     *ChartSourceRange          `json:"sourceRange,omitempty"`
	ForceSendFields []string                   `json:"-"`
	NullFields      []string                   `json:"-"`
}
type ChartGroupRule struct {
	DateTimeRule    *ChartDateTimeRule  `json:"dateTimeRule,omitempty"`
	HistogramRule   *ChartHistogramRule `json:"histogramRule,omitempty"`
	ForceSendFields []string            `json:"-"`
	NullFields      []string            `json:"-"`
}
type ChartSourceRange struct {
	Sources         []*GridRange `json:"sources,omitempty"`
	ForceSendFields []string     `json:"-"`
	NullFields      []string     `json:"-"`
}
type ChartDateTimeRule struct {
	Type            string   `json:"type,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type ChartHistogramRule struct {
	IntervalSize    float64  `json:"intervalSize,omitempty"`
	MaxValue        float64  `json:"maxValue,omitempty"`
	MinValue        float64  `json:"minValue,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type BasicChartSeries struct {
	Color           *Color                               `json:"color,omitempty"`
	ColorStyle      *ColorStyle                          `json:"colorStyle,omitempty"`
	DataLabel       *DataLabel                           `json:"dataLabel,omitempty"`
	LineStyle       *LineStyle                           `json:"lineStyle,omitempty"`
	PointStyle      *PointStyle                          `json:"pointStyle,omitempty"`
	Series          *ChartData                           `json:"series,omitempty"`
	StyleOverrides  []*BasicSeriesDataPointStyleOverride `json:"styleOverrides,omitempty"`
	TargetAxis      string                               `json:"targetAxis,omitempty"`
	Type            string                               `json:"type,omitempty"`
	ForceSendFields []string                             `json:"-"`
	NullFields      []string                             `json:"-"`
}

type DataLabel struct {
	CustomLabelData *ChartData  `json:"customLabelData,omitempty"`
	Placement       string      `json:"placement,omitempty"`
	TextFormat      *TextFormat `json:"textFormat,omitempty"`
	Type            string      `json:"type,omitempty"`
	ForceSendFields []string    `json:"-"`
	NullFields      []string    `json:"-"`
}
type LineStyle struct {
	Type            string   `json:"type,omitempty"`
	Width           int64    `json:"width,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type PointStyle struct {
	Shape           string   `json:"shape,omitempty"`
	Size            float64  `json:"size,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type BasicSeriesDataPointStyleOverride struct {
	Color           *Color      `json:"color,omitempty"`
	ColorStyle      *ColorStyle `json:"colorStyle,omitempty"`
	Index           int64       `json:"index,omitempty"`
	PointStyle      *PointStyle `json:"pointStyle,omitempty"`
	ForceSendFields []string    `json:"-"`
	NullFields      []string    `json:"-"`
}

type BubbleChartSpec struct {
	BubbleBorderColor      *Color      `json:"bubbleBorderColor,omitempty"`
	BubbleBorderColorStyle *ColorStyle `json:"bubbleBorderColorStyle,omitempty"`
	BubbleLabels           *ChartData  `json:"bubbleLabels,omitempty"`
	BubbleMaxRadiusSize    int64       `json:"bubbleMaxRadiusSize,omitempty"`
	BubbleMinRadiusSize    int64       `json:"bubbleMinRadiusSize,omitempty"`
	BubbleOpacity          float64     `json:"bubbleOpacity,omitempty"`
	BubbleSizes            *ChartData  `json:"bubbleSizes,omitempty"`
	BubbleTextStyle        *TextFormat `json:"bubbleTextStyle,omitempty"`
	Domain                 *ChartData  `json:"domain,omitempty"`
	GroupIds               *ChartData  `json:"groupIds,omitempty"`
	LegendPosition         string      `json:"legendPosition,omitempty"`
	Series                 *ChartData  `json:"series,omitempty"`
	ForceSendFields        []string    `json:"-"`
	NullFields             []string    `json:"-"`
}
type CandlestickChartSpec struct {
	Data            []*CandlestickData `json:"data,omitempty"`
	Domain          *CandlestickDomain `json:"domain,omitempty"`
	ForceSendFields []string           `json:"-"`
	NullFields      []string           `json:"-"`
}
type CandlestickData struct {
	CloseSeries     *CandlestickSeries `json:"closeSeries,omitempty"`
	HighSeries      *CandlestickSeries `json:"highSeries,omitempty"`
	LowSeries       *CandlestickSeries `json:"lowSeries,omitempty"`
	OpenSeries      *CandlestickSeries `json:"openSeries,omitempty"`
	ForceSendFields []string           `json:"-"`
	NullFields      []string           `json:"-"`
}
type CandlestickSeries struct {
	Data            *ChartData `json:"data,omitempty"`
	ForceSendFields []string   `json:"-"`
	NullFields      []string   `json:"-"`
}
type CandlestickDomain struct {
	Data            *ChartData `json:"data,omitempty"`
	Reversed        bool       `json:"reversed,omitempty"`
	ForceSendFields []string   `json:"-"`
	NullFields      []string   `json:"-"`
}

type DataSourceChartProperties struct {
	DataExecutionStatus *DataExecutionStatus `json:"dataExecutionStatus,omitempty"`
	DataSourceId        string               `json:"dataSourceId,omitempty"`
	ForceSendFields     []string             `json:"-"`
	NullFields          []string             `json:"-"`
}
type DataExecutionStatus struct {
	ErrorCode       string   `json:"errorCode,omitempty"`
	ErrorMessage    string   `json:"errorMessage,omitempty"`
	LastRefreshTime string   `json:"lastRefreshTime,omitempty"`
	State           string   `json:"state,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type HistogramChartSpec struct {
	BucketSize        float64            `json:"bucketSize,omitempty"`
	LegendPosition    string             `json:"legendPosition,omitempty"`
	OutlierPercentile float64            `json:"outlierPercentile,omitempty"`
	Series            []*HistogramSeries `json:"series,omitempty"`
	ShowItemDividers  bool               `json:"showItemDividers,omitempty"`
	ForceSendFields   []string           `json:"-"`
	NullFields        []string           `json:"-"`
}
type HistogramSeries struct {
	BarColor        *Color      `json:"barColor,omitempty"`
	BarColorStyle   *ColorStyle `json:"barColorStyle,omitempty"`
	Data            *ChartData  `json:"data,omitempty"`
	ForceSendFields []string    `json:"-"`
	NullFields      []string    `json:"-"`
}

type OrgChartSpec struct {
	Labels                 *ChartData  `json:"labels,omitempty"`
	NodeColor              *Color      `json:"nodeColor,omitempty"`
	NodeColorStyle         *ColorStyle `json:"nodeColorStyle,omitempty"`
	NodeSize               string      `json:"nodeSize,omitempty"`
	ParentLabels           *ChartData  `json:"parentLabels,omitempty"`
	SelectedNodeColor      *Color      `json:"selectedNodeColor,omitempty"`
	SelectedNodeColorStyle *ColorStyle `json:"selectedNodeColorStyle,omitempty"`
	Tooltips               *ChartData  `json:"tooltips,omitempty"`
	ForceSendFields        []string    `json:"-"`
	NullFields             []string    `json:"-"`
}
type PieChartSpec struct {
	Domain           *ChartData `json:"domain,omitempty"`
	LegendPosition   string     `json:"legendPosition,omitempty"`
	PieHole          float64    `json:"pieHole,omitempty"`
	Series           *ChartData `json:"series,omitempty"`
	ThreeDimensional bool       `json:"threeDimensional,omitempty"`
	ForceSendFields  []string   `json:"-"`
	NullFields       []string   `json:"-"`
}

type ScorecardChartSpec struct {
	AggregateType       string                          `json:"aggregateType,omitempty"`
	BaselineValueData   *ChartData                      `json:"baselineValueData,omitempty"`
	BaselineValueFormat *BaselineValueFormat            `json:"baselineValueFormat,omitempty"`
	CustomFormatOptions *ChartCustomNumberFormatOptions `json:"customFormatOptions,omitempty"`
	KeyValueData        *ChartData                      `json:"keyValueData,omitempty"`
	KeyValueFormat      *KeyValueFormat                 `json:"keyValueFormat,omitempty"`
	NumberFormatSource  string                          `json:"numberFormatSource,omitempty"`
	ScaleFactor         float64                         `json:"scaleFactor,omitempty"`
	ForceSendFields     []string                        `json:"-"`
	NullFields          []string                        `json:"-"`
}

type BaselineValueFormat struct {
	ComparisonType     string        `json:"comparisonType,omitempty"`
	Description        string        `json:"description,omitempty"`
	NegativeColor      *Color        `json:"negativeColor,omitempty"`
	NegativeColorStyle *ColorStyle   `json:"negativeColorStyle,omitempty"`
	Position           *TextPosition `json:"position,omitempty"`
	PositiveColor      *Color        `json:"positiveColor,omitempty"`
	PositiveColorStyle *ColorStyle   `json:"positiveColorStyle,omitempty"`
	TextFormat         *TextFormat   `json:"textFormat,omitempty"`
	ForceSendFields    []string      `json:"-"`
	NullFields         []string      `json:"-"`
}
type ChartCustomNumberFormatOptions struct {
	Prefix          string   `json:"prefix,omitempty"`
	Suffix          string   `json:"suffix,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

type KeyValueFormat struct {
	Position        *TextPosition `json:"position,omitempty"`
	TextFormat      *TextFormat   `json:"textFormat,omitempty"`
	ForceSendFields []string      `json:"-"`
	NullFields      []string      `json:"-"`
}

type TreemapChartSpec struct {
	ColorData        *ChartData              `json:"colorData,omitempty"`
	ColorScale       *TreemapChartColorScale `json:"colorScale,omitempty"`
	HeaderColor      *Color                  `json:"headerColor,omitempty"`
	HeaderColorStyle *ColorStyle             `json:"headerColorStyle,omitempty"`
	HideTooltips     bool                    `json:"hideTooltips,omitempty"`
	HintedLevels     int64                   `json:"hintedLevels,omitempty"`
	Labels           *ChartData              `json:"labels,omitempty"`
	Levels           int64                   `json:"levels,omitempty"`
	MaxValue         float64                 `json:"maxValue,omitempty"`
	MinValue         float64                 `json:"minValue,omitempty"`
	ParentLabels     *ChartData              `json:"parentLabels,omitempty"`
	SizeData         *ChartData              `json:"sizeData,omitempty"`
	TextFormat       *TextFormat             `json:"textFormat,omitempty"`
	ForceSendFields  []string                `json:"-"`
	NullFields       []string                `json:"-"`
}
type TreemapChartColorScale struct {
	MaxValueColor      *Color      `json:"maxValueColor,omitempty"`
	MaxValueColorStyle *ColorStyle `json:"maxValueColorStyle,omitempty"`
	MidValueColor      *Color      `json:"midValueColor,omitempty"`
	MidValueColorStyle *ColorStyle `json:"midValueColorStyle,omitempty"`
	MinValueColor      *Color      `json:"minValueColor,omitempty"`
	MinValueColorStyle *ColorStyle `json:"minValueColorStyle,omitempty"`
	NoDataColor        *Color      `json:"noDataColor,omitempty"`
	NoDataColorStyle   *ColorStyle `json:"noDataColorStyle,omitempty"`
	ForceSendFields    []string    `json:"-"`
	NullFields         []string    `json:"-"`
}
type WaterfallChartSpec struct {
	ConnectorLineStyle *LineStyle              `json:"connectorLineStyle,omitempty"`
	Domain             *WaterfallChartDomain   `json:"domain,omitempty"`
	FirstValueIsTotal  bool                    `json:"firstValueIsTotal,omitempty"`
	HideConnectorLines bool                    `json:"hideConnectorLines,omitempty"`
	Series             []*WaterfallChartSeries `json:"series,omitempty"`
	StackedType        string                  `json:"stackedType,omitempty"`
	TotalDataLabel     *DataLabel              `json:"totalDataLabel,omitempty"`
	ForceSendFields    []string                `json:"-"`
	NullFields         []string                `json:"-"`
}
type WaterfallChartSeries struct {
	CustomSubtotals      []*WaterfallChartCustomSubtotal `json:"customSubtotals,omitempty"`
	Data                 *ChartData                      `json:"data,omitempty"`
	DataLabel            *DataLabel                      `json:"dataLabel,omitempty"`
	HideTrailingSubtotal bool                            `json:"hideTrailingSubtotal,omitempty"`
	NegativeColumnsStyle *WaterfallChartColumnStyle      `json:"negativeColumnsStyle,omitempty"`
	PositiveColumnsStyle *WaterfallChartColumnStyle      `json:"positiveColumnsStyle,omitempty"`
	SubtotalColumnsStyle *WaterfallChartColumnStyle      `json:"subtotalColumnsStyle,omitempty"`
	ForceSendFields      []string                        `json:"-"`
	NullFields           []string                        `json:"-"`
}
type WaterfallChartDomain struct {
	Data            *ChartData `json:"data,omitempty"`
	Reversed        bool       `json:"reversed,omitempty"`
	ForceSendFields []string   `json:"-"`
	NullFields      []string   `json:"-"`
}
type WaterfallChartCustomSubtotal struct {
	DataIsSubtotal  bool     `json:"dataIsSubtotal,omitempty"`
	Label           string   `json:"label,omitempty"`
	SubtotalIndex   int64    `json:"subtotalIndex,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type WaterfallChartColumnStyle struct {
	Color           *Color      `json:"color,omitempty"`
	ColorStyle      *ColorStyle `json:"colorStyle,omitempty"`
	Label           string      `json:"label,omitempty"`
	ForceSendFields []string    `json:"-"`
	NullFields      []string    `json:"-"`
}

type DimensionGroup struct {
	Collapsed       bool            `json:"collapsed,omitempty"`
	Depth           int64           `json:"depth,omitempty"`
	Range           *DimensionRange `json:"range,omitempty"`
	ForceSendFields []string        `json:"-"`
	NullFields      []string        `json:"-"`
}
type ConditionalFormatRule struct {
	BooleanRule     *BooleanRule  `json:"booleanRule,omitempty"`
	GradientRule    *GradientRule `json:"gradientRule,omitempty"`
	Ranges          []*GridRange  `json:"ranges,omitempty"`
	ForceSendFields []string      `json:"-"`
	NullFields      []string      `json:"-"`
}

type BooleanRule struct {
	Condition       *BooleanCondition `json:"condition,omitempty"`
	Format          *CellFormat       `json:"format,omitempty"`
	ForceSendFields []string          `json:"-"`
	NullFields      []string          `json:"-"`
}
type GradientRule struct {
	Maxpoint        *InterpolationPoint `json:"maxpoint,omitempty"`
	Midpoint        *InterpolationPoint `json:"midpoint,omitempty"`
	Minpoint        *InterpolationPoint `json:"minpoint,omitempty"`
	ForceSendFields []string            `json:"-"`
	NullFields      []string            `json:"-"`
}
type InterpolationPoint struct {
	Color           *Color      `json:"color,omitempty"`
	ColorStyle      *ColorStyle `json:"colorStyle,omitempty"`
	Type            string      `json:"type,omitempty"`
	Value           string      `json:"value,omitempty"`
	ForceSendFields []string    `json:"-"`
	NullFields      []string    `json:"-"`
}

type GridData struct {
	ColumnMetadata  []*DimensionProperties `json:"columnMetadata,omitempty"`
	RowData         []*RowData             `json:"rowData,omitempty"`
	RowMetadata     []*DimensionProperties `json:"rowMetadata,omitempty"`
	StartColumn     int64                  `json:"startColumn,omitempty"`
	StartRow        int64                  `json:"startRow,omitempty"`
	ForceSendFields []string               `json:"-"`
	NullFields      []string               `json:"-"`
}
type DimensionProperties struct {
	DataSourceColumnReference *DataSourceColumnReference `json:"dataSourceColumnReference,omitempty"`
	DeveloperMetadata         []*DeveloperMetadata       `json:"developerMetadata,omitempty"`
	HiddenByFilter            bool                       `json:"hiddenByFilter,omitempty"`
	HiddenByUser              bool                       `json:"hiddenByUser,omitempty"`
	PixelSize                 int64                      `json:"pixelSize,omitempty"`
	ForceSendFields           []string                   `json:"-"`
	NullFields                []string                   `json:"-"`
}
type RowData struct {
	Values          []*CellData `json:"values,omitempty"`
	ForceSendFields []string    `json:"-"`
	NullFields      []string    `json:"-"`
}
type CellData struct {
	DataSourceFormula *DataSourceFormula  `json:"dataSourceFormula,omitempty"`
	DataSourceTable   *DataSourceTable    `json:"dataSourceTable,omitempty"`
	DataValidation    *DataValidationRule `json:"dataValidation,omitempty"`
	EffectiveFormat   *CellFormat         `json:"effectiveFormat,omitempty"`
	EffectiveValue    *ExtendedValue      `json:"effectiveValue,omitempty"`
	FormattedValue    string              `json:"formattedValue,omitempty"`
	Hyperlink         string              `json:"hyperlink,omitempty"`
	Note              string              `json:"note,omitempty"`
	PivotTable        *PivotTable         `json:"pivotTable,omitempty"`
	TextFormatRuns    []*TextFormatRun    `json:"textFormatRuns,omitempty"`
	UserEnteredFormat *CellFormat         `json:"userEnteredFormat,omitempty"`
	UserEnteredValue  *ExtendedValue      `json:"userEnteredValue,omitempty"`
	ForceSendFields   []string            `json:"-"`
	NullFields        []string            `json:"-"`
}
type DataSourceFormula struct {
	DataExecutionStatus *DataExecutionStatus `json:"dataExecutionStatus,omitempty"`
	DataSourceId        string               `json:"dataSourceId,omitempty"`
	ForceSendFields     []string             `json:"-"`
	NullFields          []string             `json:"-"`
}

type DataSourceTable struct {
	ColumnSelectionType string                       `json:"columnSelectionType,omitempty"`
	Columns             []*DataSourceColumnReference `json:"columns,omitempty"`
	DataExecutionStatus *DataExecutionStatus         `json:"dataExecutionStatus,omitempty"`
	DataSourceId        string                       `json:"dataSourceId,omitempty"`
	FilterSpecs         []*FilterSpec                `json:"filterSpecs,omitempty"`
	RowLimit            int64                        `json:"rowLimit,omitempty"`
	SortSpecs           []*SortSpec                  `json:"sortSpecs,omitempty"`
	ForceSendFields     []string                     `json:"-"`
	NullFields          []string                     `json:"-"`
}
type DataValidationRule struct {
	Condition       *BooleanCondition `json:"condition,omitempty"`
	InputMessage    string            `json:"inputMessage,omitempty"`
	ShowCustomUi    bool              `json:"showCustomUi,omitempty"`
	Strict          bool              `json:"strict,omitempty"`
	ForceSendFields []string          `json:"-"`
	NullFields      []string          `json:"-"`
}
type ExtendedValue struct {
	BoolValue       *bool       `json:"boolValue,omitempty"`
	ErrorValue      *ErrorValue `json:"errorValue,omitempty"`
	FormulaValue    *string     `json:"formulaValue,omitempty"`
	NumberValue     *float64    `json:"numberValue,omitempty"`
	StringValue     *string     `json:"stringValue,omitempty"`
	ForceSendFields []string    `json:"-"`
	NullFields      []string    `json:"-"`
}
type ErrorValue struct {
	Message         string   `json:"message,omitempty"`
	Type            string   `json:"type,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type PivotTable struct {
	Columns             []*PivotGroup                  `json:"columns,omitempty"`
	Criteria            map[string]PivotFilterCriteria `json:"criteria,omitempty"`
	DataExecutionStatus *DataExecutionStatus           `json:"dataExecutionStatus,omitempty"`
	DataSourceId        string                         `json:"dataSourceId,omitempty"`
	FilterSpecs         []*PivotFilterSpec             `json:"filterSpecs,omitempty"`
	Rows                []*PivotGroup                  `json:"rows,omitempty"`
	Source              *GridRange                     `json:"source,omitempty"`
	ValueLayout         string                         `json:"valueLayout,omitempty"`
	Values              []*PivotValue                  `json:"values,omitempty"`
	ForceSendFields     []string                       `json:"-"`
	NullFields          []string                       `json:"-"`
}
type PivotGroup struct {
	DataSourceColumnReference *DataSourceColumnReference `json:"dataSourceColumnReference,omitempty"`
	GroupLimit                *PivotGroupLimit           `json:"groupLimit,omitempty"`
	GroupRule                 *PivotGroupRule            `json:"groupRule,omitempty"`
	Label                     string                     `json:"label,omitempty"`
	RepeatHeadings            bool                       `json:"repeatHeadings,omitempty"`
	ShowTotals                bool                       `json:"showTotals,omitempty"`
	SortOrder                 string                     `json:"sortOrder,omitempty"`
	SourceColumnOffset        int64                      `json:"sourceColumnOffset,omitempty"`
	ValueBucket               *PivotGroupSortValueBucket `json:"valueBucket,omitempty"`
	ValueMetadata             []*PivotGroupValueMetadata `json:"valueMetadata,omitempty"`
	ForceSendFields           []string                   `json:"-"`
	NullFields                []string                   `json:"-"`
}
type PivotFilterCriteria struct {
	Condition        *BooleanCondition `json:"condition,omitempty"`
	VisibleByDefault bool              `json:"visibleByDefault,omitempty"`
	VisibleValues    []string          `json:"visibleValues,omitempty"`
	ForceSendFields  []string          `json:"-"`
	NullFields       []string          `json:"-"`
}
type PivotFilterSpec struct {
	ColumnOffsetIndex         int64                      `json:"columnOffsetIndex,omitempty"`
	DataSourceColumnReference *DataSourceColumnReference `json:"dataSourceColumnReference,omitempty"`
	FilterCriteria            *PivotFilterCriteria       `json:"filterCriteria,omitempty"`
	ForceSendFields           []string                   `json:"-"`
	NullFields                []string                   `json:"-"`
}

type PivotGroupLimit struct {
	ApplyOrder      int64    `json:"applyOrder,omitempty"`
	CountLimit      int64    `json:"countLimit,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type PivotGroupRule struct {
	DateTimeRule    *DateTimeRule  `json:"dateTimeRule,omitempty"`
	HistogramRule   *HistogramRule `json:"histogramRule,omitempty"`
	ManualRule      *ManualRule    `json:"manualRule,omitempty"`
	ForceSendFields []string       `json:"-"`
	NullFields      []string       `json:"-"`
}
type DateTimeRule struct {
	Type            string   `json:"type,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type PivotGroupSortValueBucket struct {
	Buckets         []*ExtendedValue `json:"buckets,omitempty"`
	ValuesIndex     int64            `json:"valuesIndex,omitempty"`
	ForceSendFields []string         `json:"-"`
	NullFields      []string         `json:"-"`
}
type PivotGroupValueMetadata struct {
	Collapsed       bool           `json:"collapsed,omitempty"`
	Value           *ExtendedValue `json:"value,omitempty"`
	ForceSendFields []string       `json:"-"`
	NullFields      []string       `json:"-"`
}
type PivotValue struct {
	CalculatedDisplayType     string                     `json:"calculatedDisplayType,omitempty"`
	DataSourceColumnReference *DataSourceColumnReference `json:"dataSourceColumnReference,omitempty"`
	Formula                   string                     `json:"formula,omitempty"`
	Name                      string                     `json:"name,omitempty"`
	SourceColumnOffset        int64                      `json:"sourceColumnOffset,omitempty"`
	SummarizeFunction         string                     `json:"summarizeFunction,omitempty"`
	ForceSendFields           []string                   `json:"-"`
	NullFields                []string                   `json:"-"`
}
type HistogramRule struct {
	End             float64  `json:"end,omitempty"`
	Interval        float64  `json:"interval,omitempty"`
	Start           float64  `json:"start,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}
type ManualRule struct {
	Groups          []*ManualRuleGroup `json:"groups,omitempty"`
	ForceSendFields []string           `json:"-"`
	NullFields      []string           `json:"-"`
}
type ManualRuleGroup struct {
	GroupName       *ExtendedValue   `json:"groupName,omitempty"`
	Items           []*ExtendedValue `json:"items,omitempty"`
	ForceSendFields []string         `json:"-"`
	NullFields      []string         `json:"-"`
}
type TextFormatRun struct {
	Format          *TextFormat `json:"format,omitempty"`
	StartIndex      int64       `json:"startIndex,omitempty"`
	ForceSendFields []string    `json:"-"`
	NullFields      []string    `json:"-"`
}

type FilterView struct {
	Criteria        map[string]FilterCriteria `json:"criteria,omitempty"`
	FilterSpecs     []*FilterSpec             `json:"filterSpecs,omitempty"`
	FilterViewId    int64                     `json:"filterViewId,omitempty"`
	NamedRangeId    string                    `json:"namedRangeId,omitempty"`
	Range           *GridRange                `json:"range,omitempty"`
	SortSpecs       []*SortSpec               `json:"sortSpecs,omitempty"`
	Title           string                    `json:"title,omitempty"`
	ForceSendFields []string                  `json:"-"`
	NullFields      []string                  `json:"-"`
}
type SheetProperties struct {
	DataSourceSheetProperties *DataSourceSheetProperties `json:"dataSourceSheetProperties,omitempty"`
	GridProperties            *GridProperties            `json:"gridProperties,omitempty"`
	Hidden                    bool                       `json:"hidden,omitempty"`
	Index                     int64                      `json:"index,omitempty"`
	RightToLeft               bool                       `json:"rightToLeft,omitempty"`
	SheetId                   int64                      `json:"sheetId,omitempty"`
	SheetType                 string                     `json:"sheetType,omitempty"`
	TabColor                  *Color                     `json:"tabColor,omitempty"`
	TabColorStyle             *ColorStyle                `json:"tabColorStyle,omitempty"`
	Title                     string                     `json:"title,omitempty"`
	ServerResponse            `json:"-"`
	ForceSendFields           []string `json:"-"`
	NullFields                []string `json:"-"`
}
type DataSourceSheetProperties struct {
	Columns             []*DataSourceColumn  `json:"columns,omitempty"`
	DataExecutionStatus *DataExecutionStatus `json:"dataExecutionStatus,omitempty"`
	DataSourceId        string               `json:"dataSourceId,omitempty"`
	ForceSendFields     []string             `json:"-"`
	NullFields          []string             `json:"-"`
}
type GridProperties struct {
	ColumnCount             int64    `json:"columnCount,omitempty"`
	ColumnGroupControlAfter bool     `json:"columnGroupControlAfter,omitempty"`
	FrozenColumnCount       int64    `json:"frozenColumnCount,omitempty"`
	FrozenRowCount          int64    `json:"frozenRowCount,omitempty"`
	HideGridlines           bool     `json:"hideGridlines,omitempty"`
	RowCount                int64    `json:"rowCount,omitempty"`
	RowGroupControlAfter    bool     `json:"rowGroupControlAfter,omitempty"`
	ForceSendFields         []string `json:"-"`
	NullFields              []string `json:"-"`
}

type ProtectedRange struct {
	Description           string       `json:"description,omitempty"`
	Editors               *Editors     `json:"editors,omitempty"`
	NamedRangeId          string       `json:"namedRangeId,omitempty"`
	ProtectedRangeId      int64        `json:"protectedRangeId,omitempty"`
	Range                 *GridRange   `json:"range,omitempty"`
	RequestingUserCanEdit bool         `json:"requestingUserCanEdit,omitempty"`
	UnprotectedRanges     []*GridRange `json:"unprotectedRanges,omitempty"`
	WarningOnly           bool         `json:"warningOnly,omitempty"`
	ForceSendFields       []string     `json:"-"`
	NullFields            []string     `json:"-"`
}
type Editors struct {
	DomainUsersCanEdit bool     `json:"domainUsersCanEdit,omitempty"`
	Groups             []string `json:"groups,omitempty"`
	Users              []string `json:"users,omitempty"`
	ForceSendFields    []string `json:"-"`
	NullFields         []string `json:"-"`
}
type Slicer struct {
	Position        *EmbeddedObjectPosition `json:"position,omitempty"`
	SlicerId        int64                   `json:"slicerId,omitempty"`
	Spec            *SlicerSpec             `json:"spec,omitempty"`
	ForceSendFields []string                `json:"-"`
	NullFields      []string                `json:"-"`
}
type SlicerSpec struct {
	ApplyToPivotTables   bool            `json:"applyToPivotTables,omitempty"`
	BackgroundColor      *Color          `json:"backgroundColor,omitempty"`
	BackgroundColorStyle *ColorStyle     `json:"backgroundColorStyle,omitempty"`
	ColumnIndex          int64           `json:"columnIndex,omitempty"`
	DataRange            *GridRange      `json:"dataRange,omitempty"`
	FilterCriteria       *FilterCriteria `json:"filterCriteria,omitempty"`
	HorizontalAlignment  string          `json:"horizontalAlignment,omitempty"`
	TextFormat           *TextFormat     `json:"textFormat,omitempty"`
	Title                string          `json:"title,omitempty"`
	ForceSendFields      []string        `json:"-"`
	NullFields           []string        `json:"-"`
}

type IterativeCalculationSettings struct {
	ConvergenceThreshold float64  `json:"convergenceThreshold,omitempty"`
	MaxIterations        int64    `json:"maxIterations,omitempty"`
	ForceSendFields      []string `json:"-"`
	NullFields           []string `json:"-"`
}
type SpreadsheetTheme struct {
	PrimaryFontFamily string            `json:"primaryFontFamily,omitempty"`
	ThemeColors       []*ThemeColorPair `json:"themeColors,omitempty"`
	ForceSendFields   []string          `json:"-"`
	NullFields        []string          `json:"-"`
}
type ThemeColorPair struct {
	Color           *ColorStyle `json:"color,omitempty"`
	ColorType       string      `json:"colorType,omitempty"`
	ForceSendFields []string    `json:"-"`
	NullFields      []string    `json:"-"`
}
