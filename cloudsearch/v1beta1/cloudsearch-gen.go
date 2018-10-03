// Package cloudsearch provides access to the Cloud Search API.
//
// See https://gsuite.google.com/products/cloud-search/
//
// Usage example:
//
//   import "github.com/nilleb/google-api-go-client/cloudsearch/v1beta1"
//   ...
//   cloudsearchService, err := cloudsearch.New(oauthHttpClient)
package cloudsearch // import "github.com/nilleb/google-api-go-client/cloudsearch/v1beta1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "cloudsearch:v1beta1"
const apiName = "cloudsearch"
const apiVersion = "v1beta1"
const basePath = "https://cloudsearch.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// Index and serve your organization's data with Cloud Search
	CloudSearchScope = "https://www.googleapis.com/auth/cloud_search"

	// New Service: https://www.googleapis.com/auth/cloud_search.indexing
	CloudSearchIndexingScope = "https://www.googleapis.com/auth/cloud_search.indexing"

	// Search your organization's data in the Cloud Search index
	CloudSearchQueryScope = "https://www.googleapis.com/auth/cloud_search.query"

	// New Service: https://www.googleapis.com/auth/cloud_search.settings
	CloudSearchSettingsScope = "https://www.googleapis.com/auth/cloud_search.settings"

	// New Service:
	// https://www.googleapis.com/auth/cloud_search.settings.indexing
	CloudSearchSettingsIndexingScope = "https://www.googleapis.com/auth/cloud_search.settings.indexing"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Indexing = NewIndexingService(s)
	s.Query = NewQueryService(s)
	s.Settings = NewSettingsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Indexing *IndexingService

	Query *QueryService

	Settings *SettingsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewIndexingService(s *Service) *IndexingService {
	rs := &IndexingService{s: s}
	rs.Datasources = NewIndexingDatasourcesService(s)
	return rs
}

type IndexingService struct {
	s *Service

	Datasources *IndexingDatasourcesService
}

func NewIndexingDatasourcesService(s *Service) *IndexingDatasourcesService {
	rs := &IndexingDatasourcesService{s: s}
	rs.Items = NewIndexingDatasourcesItemsService(s)
	return rs
}

type IndexingDatasourcesService struct {
	s *Service

	Items *IndexingDatasourcesItemsService
}

func NewIndexingDatasourcesItemsService(s *Service) *IndexingDatasourcesItemsService {
	rs := &IndexingDatasourcesItemsService{s: s}
	return rs
}

type IndexingDatasourcesItemsService struct {
	s *Service
}

func NewQueryService(s *Service) *QueryService {
	rs := &QueryService{s: s}
	return rs
}

type QueryService struct {
	s *Service
}

func NewSettingsService(s *Service) *SettingsService {
	rs := &SettingsService{s: s}
	rs.Datasources = NewSettingsDatasourcesService(s)
	return rs
}

type SettingsService struct {
	s *Service

	Datasources *SettingsDatasourcesService
}

func NewSettingsDatasourcesService(s *Service) *SettingsDatasourcesService {
	rs := &SettingsDatasourcesService{s: s}
	return rs
}

type SettingsDatasourcesService struct {
	s *Service
}

// BooleanOperatorOptions: Used to provide a search operator for boolean
// properties. This is
// optional. Search operators let users restrict the query to specific
// fields
// relevant to the type of item being searched.
type BooleanOperatorOptions struct {
	// OperatorName: Indicates the operator name required in the query in
	// order to isolate the
	// boolean property. For example, if operatorName is *closed* and
	// the
	// property's name is *isClosed*, then queries
	// like
	// *closed:&lt;value&gt;* will show results only where the value of
	// the
	// property named *isClosed* matches *&lt;value&gt;*. By contrast,
	// a
	// search that uses the same *&lt;value&gt;* without an operator will
	// return
	// all items where *&lt;value&gt;* matches the value of any String
	// properties
	// or text within the content field for the item.
	// The operator name can only contain lowercase letters (a-z).
	// The maximum length is 32 characters.
	OperatorName string `json:"operatorName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "OperatorName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "OperatorName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BooleanOperatorOptions) MarshalJSON() ([]byte, error) {
	type NoMethod BooleanOperatorOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BooleanPropertyOptions: Options for boolean properties.
type BooleanPropertyOptions struct {
	// OperatorOptions: If set, describes how the boolean should be used as
	// a search operator.
	OperatorOptions *BooleanOperatorOptions `json:"operatorOptions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "OperatorOptions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "OperatorOptions") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *BooleanPropertyOptions) MarshalJSON() ([]byte, error) {
	type NoMethod BooleanPropertyOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BooleanValue: Boolean value.
type BooleanValue struct {
	Value bool `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Value") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Value") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BooleanValue) MarshalJSON() ([]byte, error) {
	type NoMethod BooleanValue
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type CompositeFilter struct {
	// LogicOperator: The logic operator of the sub filter.
	//
	// Possible values:
	//   "AND" - Logical operators, which can only be applied to sub
	// filters.
	//   "OR"
	//   "NOT" - NOT can only be applied on a single sub filter.
	LogicOperator string `json:"logicOperator,omitempty"`

	// SubFilters: Sub filters.
	SubFilters []*Filter `json:"subFilters,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LogicOperator") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LogicOperator") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CompositeFilter) MarshalJSON() ([]byte, error) {
	type NoMethod CompositeFilter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DataSource: Data source is a logical namespace for items to be
// indexed.
// All items must belong to a data source.  This is the prerequisite
// before
// items can be indexed into Cloud Search.
type DataSource struct {
	// DisableModifications: If true, Indexing API rejects any modification
	// calls to this data source
	// such as create, update, and delete.
	// Disabling this does not imply halting process of previously
	// accepted data.
	DisableModifications bool `json:"disableModifications,omitempty"`

	// DisableServing: Disable serving any search or assist
	// results.
	// Disabling serving does not imply disabling API call.
	DisableServing bool `json:"disableServing,omitempty"`

	// DisplayName: Required. Display name of the data source
	// The maximum length is 300 characters.
	DisplayName string `json:"displayName,omitempty"`

	// IndexingServiceAccounts: List of service accounts that has indexing
	// access.
	IndexingServiceAccounts []string `json:"indexingServiceAccounts,omitempty"`

	// ItemsVisibility: This restricts visibility to items at a data source
	// level to the
	// disjunction of users/groups mentioned with the field. Note that,
	// this
	// does not ensure access to a specific item, as the users needs to have
	// ACL
	// permissions on the item too. This ensures a high level access on the
	// entire
	// data source, and that the individual items are not shared outside
	// this
	// visibility.
	ItemsVisibility []*GSuitePrincipal `json:"itemsVisibility,omitempty"`

	// Name: Name of the data source resource.
	// Format: datasources/{source_id}.
	// <br />The name is ignored when creating a data source.
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "DisableModifications") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DisableModifications") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *DataSource) MarshalJSON() ([]byte, error) {
	type NoMethod DataSource
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Date: Represents a whole calendar date, for example a date of birth.
// The time of day and time zone are either specified elsewhere or are
// not significant. The date is relative to the [Proleptic Gregorian
// Calendar](https://en.wikipedia.org/wiki/Proleptic_Gregorian_calendar).
//  The date must be a valid calendar date between the year 1 and 9999.
type Date struct {
	// Day: Day of month. Must be from 1 to 31 and valid for the year and
	// month.
	Day int64 `json:"day,omitempty"`

	// Month: Month of year. Must be from 1 to 12, or 0 if specifying a date
	// without a
	// month.
	Month int64 `json:"month,omitempty"`

	// Year: Year of date. Must be from 1 to 9999.
	Year int64 `json:"year,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Day") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Day") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Date) MarshalJSON() ([]byte, error) {
	type NoMethod Date
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DateOperatorOptions: Optional. Provides a search operator for date
// properties.
// Search operators let users restrict the query to specific fields
// relevant
// to the type of item being searched.
type DateOperatorOptions struct {
	// GreaterThanOperatorName: Indicates the operator name required in the
	// query in order to isolate the
	// date property using the greater-than operator. For example,
	// if
	// greaterThanOperatorName is *closedafter* and the property's name
	// is
	// *closeDate*, then queries like *closedafter:&lt;value&gt;* will
	// show results only where the value of the property named *closeDate*
	// is
	// later than *&lt;value&gt;*.
	// The operator name can only contain lowercase letters (a-z).
	// The maximum length is 32 characters.
	GreaterThanOperatorName string `json:"greaterThanOperatorName,omitempty"`

	// LessThanOperatorName: Indicates the operator name required in the
	// query in order to isolate the
	// date property using the less-than operator. For example,
	// if
	// lessThanOperatorName is *closedbefore* and the property's name
	// is
	// *closeDate*, then queries like *closedbefore:&lt;value&gt;* will
	// show results only where the value of the property named *closeDate*
	// is
	// earlier than *&lt;value&gt;*.
	// The operator name can only contain lowercase letters (a-z).
	// The maximum length is 32 characters.
	LessThanOperatorName string `json:"lessThanOperatorName,omitempty"`

	// OperatorName: Indicates the actual string required in the query in
	// order to isolate the
	// date property. For example, suppose an issue tracking schema
	// object
	// has a property named *closeDate* that specifies an operator with
	// an
	// operatorName of *closedon*. For searches on that data, queries
	// like
	// *closedon:&lt;value&gt;* will show results only where the value of
	// the
	// *closeDate* property matches *&lt;value&gt;*. By contrast, a
	// search that uses the same *&lt;value&gt;* without an operator will
	// return
	// all items where *&lt;value&gt;* matches the value of any
	// String
	// properties or text within the content field for the indexed
	// datasource.
	// The operator name can only contain lowercase letters (a-z).
	// The maximum length is 32 characters.
	OperatorName string `json:"operatorName,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "GreaterThanOperatorName") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "GreaterThanOperatorName")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *DateOperatorOptions) MarshalJSON() ([]byte, error) {
	type NoMethod DateOperatorOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DatePropertyOptions: Options for date properties.
type DatePropertyOptions struct {
	// OperatorOptions: If set, describes how the date should be used as a
	// search operator.
	OperatorOptions *DateOperatorOptions `json:"operatorOptions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "OperatorOptions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "OperatorOptions") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *DatePropertyOptions) MarshalJSON() ([]byte, error) {
	type NoMethod DatePropertyOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DateValues: List of date values.
type DateValues struct {
	// Values: The maximum number of elements is 100.
	Values []*Date `json:"values,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Values") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Values") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DateValues) MarshalJSON() ([]byte, error) {
	type NoMethod DateValues
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DocumentThumbnailOptions: Options to control thumbnail size in
// document thumbnails.
type DocumentThumbnailOptions struct {
	// ThumbnailUrlSizeInPx: The size of the document thumbnails.
	ThumbnailUrlSizeInPx int64 `json:"thumbnailUrlSizeInPx,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "ThumbnailUrlSizeInPx") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ThumbnailUrlSizeInPx") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *DocumentThumbnailOptions) MarshalJSON() ([]byte, error) {
	type NoMethod DocumentThumbnailOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DoublePropertyOptions: Options for double properties.
type DoublePropertyOptions struct {
}

// DoubleValues: List of double values.
type DoubleValues struct {
	// Values: The maximum number of elements is 100.
	Values []float64 `json:"values,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Values") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Values") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DoubleValues) MarshalJSON() ([]byte, error) {
	type NoMethod DoubleValues
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DriveFollowUpRestrict: Drive follow-up search restricts (e.g.
// "followup:suggestions").
type DriveFollowUpRestrict struct {
	// Possible values:
	//   "UNSPECIFIED"
	//   "FOLLOWUP_SUGGESTIONS"
	//   "FOLLOWUP_ACTION_ITEMS"
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Type") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Type") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DriveFollowUpRestrict) MarshalJSON() ([]byte, error) {
	type NoMethod DriveFollowUpRestrict
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DriveLocationRestrict: Drive location search restricts (e.g.
// "is:starred").
type DriveLocationRestrict struct {
	// Possible values:
	//   "UNSPECIFIED"
	//   "TRASHED"
	//   "STARRED"
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Type") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Type") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DriveLocationRestrict) MarshalJSON() ([]byte, error) {
	type NoMethod DriveLocationRestrict
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DriveMimeTypeRestrict: Drive mime-type search restricts (e.g.
// "type:pdf").
type DriveMimeTypeRestrict struct {
	// Possible values:
	//   "UNSPECIFIED"
	//   "PDF"
	//   "DOCUMENT"
	//   "PRESENTATION"
	//   "SPREADSHEET"
	//   "FORM"
	//   "DRAWING"
	//   "SCRIPT"
	//   "MAP"
	//   "IMAGE"
	//   "AUDIO"
	//   "VIDEO"
	//   "FOLDER"
	//   "ARCHIVE"
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Type") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Type") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DriveMimeTypeRestrict) MarshalJSON() ([]byte, error) {
	type NoMethod DriveMimeTypeRestrict
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DriveTimeSpanRestrict: The time span search restrict (e.g.
// "after:2017-09-11 before:2017-09-12").
type DriveTimeSpanRestrict struct {
	// Possible values:
	//   "UNSPECIFIED"
	//   "TODAY"
	//   "YESTERDAY"
	//   "LAST_7_DAYS"
	//   "LAST_30_DAYS" - Not Enabled
	//   "LAST_90_DAYS" - Not Enabled
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Type") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Type") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DriveTimeSpanRestrict) MarshalJSON() ([]byte, error) {
	type NoMethod DriveTimeSpanRestrict
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// EmailAddress: A person's email address.
type EmailAddress struct {
	// EmailAddress: The email address.
	EmailAddress string `json:"emailAddress,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EmailAddress") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EmailAddress") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *EmailAddress) MarshalJSON() ([]byte, error) {
	type NoMethod EmailAddress
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// EnumOperatorOptions: Used to provide a search operator for enum
// properties. This is
// optional. Search operators let users restrict the query to specific
// fields
// relevant to the type of item being searched. For example, if you
// provide no
// operator for a *priority* enum property with possible values *p0* and
// *p1*,
// a query that contains the term *p0* will return items that have *p0*
// as the
// value of the *priority* property, as well as any items that contain
// the
// string *p0* in other fields. If you provide an operator name for the
// enum,
// such as *priority*, then search users can use that operator to
// refine
// results to only items that have *p0* as this property's value, with
// the
// query *priority:p0*.
type EnumOperatorOptions struct {
	// OperatorName: Indicates the operator name required in the query in
	// order to isolate the
	// enum property. For example, if operatorName is *priority* and
	// the
	// property's name is *priorityVal*, then queries
	// like
	// *priority:&lt;value&gt;* will show results only where the value of
	// the
	// property named *priorityVal* matches *&lt;value&gt;*. By contrast,
	// a
	// search that uses the same *&lt;value&gt;* without an operator will
	// return
	// all items where *&lt;value&gt;* matches the value of any
	// String
	// properties or text within the content field for the item.
	// The operator name can only contain lowercase letters (a-z).
	// The maximum length is 32 characters.
	OperatorName string `json:"operatorName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "OperatorName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "OperatorName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *EnumOperatorOptions) MarshalJSON() ([]byte, error) {
	type NoMethod EnumOperatorOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// EnumPropertyOptions: Options for enum properties, which allow you to
// define a restricted set of
// strings to match user queries, set rankings for those string values,
// and
// define an operator name to be paired with those strings so that users
// can
// narrow results to only items with a specific value. For example, for
// items in
// a request tracking system with priority information, you could define
// *p0* as
// an allowable enum value and tie this enum to the operator name
// *priority* so
// that search users could add *priority:p0* to their query to restrict
// the set
// of results to only those items indexed with the value *p0*.
type EnumPropertyOptions struct {
	// OperatorOptions: If set, describes how the enum should be used as a
	// search operator.
	OperatorOptions *EnumOperatorOptions `json:"operatorOptions,omitempty"`

	// OrderedRanking: Used to specify the ordered ranking for the
	// enumeration that determines how
	// the integer values provided in the possible EnumValuePairs are used
	// to rank
	// results. If specified, integer values must be provided for all
	// possible
	// EnumValuePair values given for this property. Can only be used
	// if
	// isRepeatable
	// is false.
	//
	// Possible values:
	//   "NO_ORDER" - There is no ranking order for the property. Results
	// will not be adjusted
	// by this property's value.
	//   "ASCENDING" - This property is ranked in ascending order. Lower
	// values indicate lower
	// ranking.
	//   "DESCENDING" - This property is ranked in descending order. Lower
	// values indicate
	// higher ranking.
	OrderedRanking string `json:"orderedRanking,omitempty"`

	// PossibleValues: The list of possible values for the enumeration
	// property. All
	// EnumValuePairs must provide a string value. If you specify an integer
	// value
	// for one EnumValuePair, then all possible EnumValuePairs must provide
	// an
	// integer value. Both the string value and integer value must be unique
	// over
	// all possible values. Once set, possible values cannot be removed
	// or
	// modified. If you supply an ordered ranking and think you might
	// insert
	// additional enum values in the future, leave gaps in the initial
	// integer
	// values to allow adding a value in between previously registered
	// values.
	// The maximum number of elements is 100.
	PossibleValues []*EnumValuePair `json:"possibleValues,omitempty"`

	// ForceSendFields is a list of field names (e.g. "OperatorOptions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "OperatorOptions") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *EnumPropertyOptions) MarshalJSON() ([]byte, error) {
	type NoMethod EnumPropertyOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// EnumValuePair: The enumeration value pair defines two things: a
// required string value and
// an optional integer value. The string value defines the necessary
// query
// term required to retrieve that item, such as *p0* for a priority
// item.
// The integer value determines the ranking of that string value
// relative
// to other enumerated values for the same property. For example, you
// might
// associate *p0* with *0* and define another enum pair such as *p1* and
// *1*.
// You must use the integer value in combination with
// ordered ranking
// to set the ranking of a given value relative to
// other enumerated values for the same property name. Here, a ranking
// order
// of DESCENDING for *priority* properties results in a ranking
// boost
// for items indexed with a value of *p0* compared to items indexed with
// a
// value of *p1*. Without a specified ranking order, the integer value
// has no
// effect on item ranking.
type EnumValuePair struct {
	// IntegerValue: The integer value of the EnumValuePair which must be
	// non-negative.
	// Optional.
	IntegerValue int64 `json:"integerValue,omitempty"`

	// StringValue: The string value of the EnumValuePair.
	// The maximum length is 32 characters.
	StringValue string `json:"stringValue,omitempty"`

	// ForceSendFields is a list of field names (e.g. "IntegerValue") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "IntegerValue") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *EnumValuePair) MarshalJSON() ([]byte, error) {
	type NoMethod EnumValuePair
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// EnumValues: List of enum values.
type EnumValues struct {
	// Values: The maximum allowable length for string values is 32
	// characters.
	// The maximum number of elements is 100.
	Values []string `json:"values,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Values") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Values") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *EnumValues) MarshalJSON() ([]byte, error) {
	type NoMethod EnumValues
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FacetBucket: A bucket in a facet is the basic unit of operation. A
// bucket can comprise
// either a single value OR a contiguous range of values, depending on
// the
// type of the field bucketed.
// FacetBucket is currently used only for returning the response object.
type FacetBucket struct {
	Range *Range `json:"range,omitempty"`

	Value *Value `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Range") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Range") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *FacetBucket) MarshalJSON() ([]byte, error) {
	type NoMethod FacetBucket
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FacetOptions: Current implementation of FacetedSearch will support
// ranges only for
// operators on numeric date/time types. There will be one FacetResult
// for every
// source_name/object_type/operator_name.
type FacetOptions struct {
	// ObjectType: If object_type is set, only those objects of that type
	// will be used to
	// compute facets. If empty, then all objects will be used to compute
	// facets.
	ObjectType string `json:"objectType,omitempty"`

	// OperatorName: Name of the operator chosen for faceting.
	// @see
	// cloudsearch.SchemaPropertyOptions
	OperatorName string `json:"operatorName,omitempty"`

	// SourceName: Source name to facet on.
	SourceName string `json:"sourceName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ObjectType") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ObjectType") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *FacetOptions) MarshalJSON() ([]byte, error) {
	type NoMethod FacetOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FacetResult: Source specific facet response
type FacetResult struct {
	// Bucket: FacetBuckets for values in response containing atleast a
	// single result.
	Bucket []*FacetBucket `json:"bucket,omitempty"`

	// ObjectType: Object type for which facet results are returned. Can be
	// empty.
	ObjectType string `json:"objectType,omitempty"`

	// OperatorName: Name of the operator chosen for faceting.
	// @see
	// cloudsearch.SchemaPropertyOptions
	OperatorName string `json:"operatorName,omitempty"`

	// SourceName: Source name for which facet results are returned. Will
	// not be empty.
	SourceName string `json:"sourceName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Bucket") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Bucket") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *FacetResult) MarshalJSON() ([]byte, error) {
	type NoMethod FacetResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type FieldViolation struct {
	// Description: Description of the error.
	Description string `json:"description,omitempty"`

	// Field: Path of field with violation.
	Field string `json:"field,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *FieldViolation) MarshalJSON() ([]byte, error) {
	type NoMethod FieldViolation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Filter: A generic way of expressing filters in a query, which
// supports two
// approaches: <br/><br/>
// **1. Setting a ValueFilter.** The name must match an operator_name
// defined in
// the schema for your data source.
// <br/>
// **2. Setting a CompositeFilter.** The filters are evaluated
// using the logical operator. The top-level operators can only be
// either an AND
// or a NOT. AND can appear only at the top-most level. OR can appear
// only under
// a top-level AND.
type Filter struct {
	CompositeFilter *CompositeFilter `json:"compositeFilter,omitempty"`

	ValueFilter *ValueFilter `json:"valueFilter,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CompositeFilter") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CompositeFilter") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Filter) MarshalJSON() ([]byte, error) {
	type NoMethod Filter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FilterOptions: Filter options to be applied on query.
type FilterOptions struct {
	// Filter: Generic filter to restrict the search, such as `lang:en`,
	// `site:xyz`.
	Filter *Filter `json:"filter,omitempty"`

	// ObjectType: If object_type is set, only objects of that type are
	// returned. This should
	// correspond to the name of the object that was registered within
	// the
	// definition of schema. The maximum length is 256 characters.
	ObjectType string `json:"objectType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Filter") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Filter") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *FilterOptions) MarshalJSON() ([]byte, error) {
	type NoMethod FilterOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FilterPropertyOptions: Options for filter properties, which are
// specifically designed to allow
// search users to filter query results by a pre-defined filter key. The
// filter
// type is simliar to the string type, but filters are only used to
// match
// queries where the filter key is used as a search operator and will
// not match
// queries where the filter key is not specified.
type FilterPropertyOptions struct {
	// OperatorName: Indicates the operator name required in the query in
	// order to isolate the
	// filter property. For example, if operatorName is *subject* and
	// the
	// property's name is *subjectLine*, then queries
	// like
	// *subject:&lt;value&gt;* will show results only where the value of
	// the
	// property named *subjectLine* matches *&lt;value&gt;*. In contrast
	// to
	// the String properties, a search that uses the same *&lt;value&gt;*
	// without an
	// operator will not return items where the *subjectLine* property
	// has
	// value equal to *&lt;value&gt;*. In this way, filter properties
	// must
	// specify the operator in order to be active for the search query.
	// The operator name can only contain lowercase letters (a-z).
	// The maximum length is 32 characters.
	OperatorName string `json:"operatorName,omitempty"`

	// Tokenization: Indicates how to tokenize the filter's value.
	//
	// Possible values:
	//   "UNSPECIFIED" - Invalid value.
	//   "TEXT" - The filter's value is tokenized as text.
	//   "ATOM" - The filter's value is not tokenized at all and treated as
	// a single token.
	Tokenization string `json:"tokenization,omitempty"`

	// ForceSendFields is a list of field names (e.g. "OperatorName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "OperatorName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *FilterPropertyOptions) MarshalJSON() ([]byte, error) {
	type NoMethod FilterPropertyOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FilterValues: List of filter values.
type FilterValues struct {
	// Values: The maximum allowable length for string values is 2048
	// characters.
	// The maximum number of elements is 100.
	Values []string `json:"values,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Values") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Values") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *FilterValues) MarshalJSON() ([]byte, error) {
	type NoMethod FilterValues
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FreshnessOptions: Indicates which freshness property to use when
// adjusting search ranking for
// an item. Fresher, more recent dates indicate higher quality. Use
// the
// freshness option property that best works with your data. For
// fileshare
// documents, last modified time is most relevant. For calendar event
// data,
// the time when the event occurs is a more relevant freshness
// indicator. In
// this way, calendar events that occur closer to the time of the search
// query
// are considered higher quality and ranked accordingly.
type FreshnessOptions struct {
	// FreshnessProperty: This property indicates the freshness level of the
	// object in the index.
	// If set, this property must be a
	// timestamp type.
	// Otherwise, the Indexing API uses
	// lastModifiedTime
	// as the freshness indicator.
	// The maximum length is 256 characters.
	FreshnessProperty string `json:"freshnessProperty,omitempty"`

	// SecsWhenStale: The duration (in seconds) after which an object should
	// be considered
	// stale.
	SecsWhenStale string `json:"secsWhenStale,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FreshnessProperty")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FreshnessProperty") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *FreshnessOptions) MarshalJSON() ([]byte, error) {
	type NoMethod FreshnessOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type GSuitePrincipal struct {
	// GsuiteDomain: This principal represents all users of the G Suite
	// domain of the
	// customer.
	GsuiteDomain bool `json:"gsuiteDomain,omitempty"`

	// GsuiteGroupEmail: This principal references a G Suite group account
	GsuiteGroupEmail string `json:"gsuiteGroupEmail,omitempty"`

	// GsuiteUserEmail: This principal references a G Suite user account
	GsuiteUserEmail string `json:"gsuiteUserEmail,omitempty"`

	// ForceSendFields is a list of field names (e.g. "GsuiteDomain") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "GsuiteDomain") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GSuitePrincipal) MarshalJSON() ([]byte, error) {
	type NoMethod GSuitePrincipal
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type IndexItemRequest struct {
	// Item: Name of the item.
	// Format:
	// datasources/{source_id}/items/{item_id}
	Item *Item `json:"item,omitempty"`

	// Mode: This update request should go into incremental priority or
	// backfill
	// priority.
	//
	// Possible values:
	//   "UNSPECIFIED" - Priority is not specified in the update
	// request.
	// Leaving priority unspecified results in an update failure.
	//   "SYNCHRONOUS" - For real-time updates.
	//   "ASYNCHRONOUS" - For bulk backfill of entire or parts of the
	// repository.
	Mode string `json:"mode,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Item") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Item") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IndexItemRequest) MarshalJSON() ([]byte, error) {
	type NoMethod IndexItemRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntegerOperatorOptions: Used to provide a search operator for integer
// properties. This is
// optional. Search operators let users restrict the query to specific
// fields
// relevant to the type of item being searched.
type IntegerOperatorOptions struct {
	// GreaterThanOperatorName: Indicates the operator name required in the
	// query in order to isolate the
	// integer property using the greater-than operator. For example,
	// if
	// greaterThanOperatorName is *priorityabove* and the property's name
	// is
	// *priorityVal*, then queries like *priorityabove:&lt;value&gt;*
	// will
	// show results only where the value of the property named *priorityVal*
	// is
	// greater than *&lt;value&gt;*.
	// The operator name can only contain lowercase letters (a-z).
	// The maximum length is 32 characters.
	GreaterThanOperatorName string `json:"greaterThanOperatorName,omitempty"`

	// LessThanOperatorName: Indicates the operator name required in the
	// query in order to isolate the
	// integer property using the less-than operator. For example,
	// if
	// lessThanOperatorName is *prioritybelow* and the property's name
	// is
	// *priorityVal*, then queries like *prioritybelow:&lt;value&gt;*
	// will
	// show results only where the value of the property named *priorityVal*
	// is
	// less than *&lt;value&gt;*.
	// The operator name can only contain lowercase letters (a-z).
	// The maximum length is 32 characters.
	LessThanOperatorName string `json:"lessThanOperatorName,omitempty"`

	// OperatorName: Indicates the operator name required in the query in
	// order to isolate the
	// integer property. For example, if operatorName is *priority* and
	// the
	// property's name is *priorityVal*, then queries
	// like
	// *priority:&lt;value&gt;* will show results only where the value of
	// the
	// property named *priorityVal* matches *&lt;value&gt;*. By contrast,
	// a
	// search that uses the same *&lt;value&gt;* without an operator will
	// return
	// all items where *&lt;value&gt;* matches the value of any String
	// properties
	// or text within the content field for the item.
	// The operator name can only contain lowercase letters (a-z).
	// The maximum length is 32 characters.
	OperatorName string `json:"operatorName,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "GreaterThanOperatorName") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "GreaterThanOperatorName")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *IntegerOperatorOptions) MarshalJSON() ([]byte, error) {
	type NoMethod IntegerOperatorOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntegerPropertyOptions: Options for integer properties.
type IntegerPropertyOptions struct {
	// MaximumValue: The maximum value of the property. The minimum and
	// maximum values for the
	// property are used to rank results according to the
	// ordered ranking.
	// Indexing requests with values greater than the maximum are accepted
	// and
	// ranked with the same weight as items indexed with the maximum value.
	MaximumValue int64 `json:"maximumValue,omitempty,string"`

	// MinimumValue: The minimum value of the property. The minimum and
	// maximum values for the
	// property are used to rank results according to the
	// ordered ranking.
	// Indexing requests with values less than the minimum are accepted
	// and
	// ranked with the same weight as items indexed with the minimum value.
	MinimumValue int64 `json:"minimumValue,omitempty,string"`

	// OperatorOptions: If set, describes how the integer should be used as
	// a search operator.
	OperatorOptions *IntegerOperatorOptions `json:"operatorOptions,omitempty"`

	// OrderedRanking: Used to specify the ordered ranking for the integer.
	// Can only be used if
	// isRepeatable
	// is false.
	//
	// Possible values:
	//   "NO_ORDER" - There is no ranking order for the property. Results
	// will not be adjusted
	// by this property's value.
	//   "ASCENDING" - This property is ranked in ascending order. Lower
	// values indicate lower
	// ranking.
	//   "DESCENDING" - This property is ranked in descending order. Lower
	// values indicate
	// higher ranking.
	OrderedRanking string `json:"orderedRanking,omitempty"`

	// ForceSendFields is a list of field names (e.g. "MaximumValue") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "MaximumValue") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntegerPropertyOptions) MarshalJSON() ([]byte, error) {
	type NoMethod IntegerPropertyOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntegerValues: List of integer values.
type IntegerValues struct {
	// Values: The maximum number of elements is 100.
	Values googleapi.Int64s `json:"values,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Values") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Values") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntegerValues) MarshalJSON() ([]byte, error) {
	type NoMethod IntegerValues
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Interaction: Represents an interaction between a user and an item.
type Interaction struct {
	// InteractionTime: The time when the user acted on the item.  If
	// multiple actions of the same
	// type exist for a single user, only the most recent action is
	// recorded.
	InteractionTime string `json:"interactionTime,omitempty"`

	// Principal: The user that acted on the item.
	Principal *Principal `json:"principal,omitempty"`

	// Possible values:
	//   "UNSPECIFIED" - Invalid value.
	//   "VIEW" - This interaction indicates the user viewed the item.
	//   "EDIT" - This interaction indicates the user edited the item.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "InteractionTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "InteractionTime") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Interaction) MarshalJSON() ([]byte, error) {
	type NoMethod Interaction
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Item: Represents a single object that is an item in the search index,
// such as a
// file, folder, or a database record.
type Item struct {
	// Acl: Access control list for this item.
	Acl *ItemAcl `json:"acl,omitempty"`

	// Content: Item content to be indexed and made text searchable.
	Content *ItemContent `json:"content,omitempty"`

	// ItemType: Type for this item.
	//
	// Possible values:
	//   "UNSPECIFIED"
	//   "CONTENT_ITEM" - An item that is indexed for the only purpose of
	// serving information.
	// These items cannot be referred in
	// containerName
	// or inheritAclFrom
	// fields.
	//   "CONTAINER_ITEM" - An item that gets indexed and whose purpose is
	// to supply other items
	// with ACLs and/or contain other items.
	//   "VIRTUAL_CONTAINER_ITEM" - An item that does not get indexed, but
	// otherwise has the same purpose
	// as CONTAINER_ITEM.
	ItemType string `json:"itemType,omitempty"`

	// Metadata: Metadata information.
	Metadata *ItemMetadata `json:"metadata,omitempty"`

	// Name: Name of the Item.
	// Format:
	// datasources/{source_id}/items/{item_id}
	// <br />This is a required field.
	// The maximum length is 2048 characters.
	Name string `json:"name,omitempty"`

	// Payload: Additional state connector can store for this item.
	// The maximum length is 10000 bytes.
	Payload string `json:"payload,omitempty"`

	// Queue: Queue this item belongs to.
	// The maximum length is 100 characters.
	Queue string `json:"queue,omitempty"`

	// Status: Status of the item.
	// Output only field.
	Status *ItemStatus `json:"status,omitempty"`

	// StructuredData: The structured data for the item that should conform
	// to a registered
	// object definition in the schema for the data source.
	StructuredData *ItemStructuredData `json:"structuredData,omitempty"`

	// Version: Required. The indexing system stores the version from the
	// datasource as a
	// byte string and compares the Item version in the index
	// to the version of the queued Item using lexical ordering.
	// <br /><br />
	// Cloud Search Indexing won't index or delete any queued item with
	// a version value that is less than or equal to the version of
	// the
	// currently indexed item.
	// The maximum length for this field is 1024 bytes.
	Version string `json:"version,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Acl") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Acl") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Item) MarshalJSON() ([]byte, error) {
	type NoMethod Item
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ItemAcl: Access control list information for the item.
type ItemAcl struct {
	// AclInheritanceType: Sets the type of access rules to apply when an
	// item inherits its ACL from a
	// parent. This should always be set in tandem with
	// the
	// inheritAclFrom
	// field. Also, when the
	// inheritAclFrom field
	// is set, this field should be set to a valid AclInheritanceType.
	//
	// Possible values:
	//   "NOT_APPLICABLE" - The default value when this item does not
	// inherit an ACL.
	// Use NOT_APPLICABLE when
	// inheritAclFrom
	// is empty.  An item without ACL inheritance can still have ACLs
	// supplied
	// by its own readers and deniedReaders fields.
	//   "CHILD_OVERRIDE" - During an authorization conflict, the ACL of the
	// child item determines
	// its read access.
	//   "PARENT_OVERRIDE" - During an authorization conflict, the ACL of
	// the parent item
	// specified in the
	// inheritAclFrom
	// field determines read access.
	//   "BOTH_PERMIT" - Access is granted only if this item and the parent
	// item specified in
	// the inheritAclFrom
	// field both permit read access.
	AclInheritanceType string `json:"aclInheritanceType,omitempty"`

	// DeniedReaders: List of readers denied access.  This is a list
	// of
	// principals who shouldn't be allowed read/write access to the
	// document
	// and should therefore not see that document in search results.
	// This list overrides access for any corresponding principals in the
	// readers
	// field.
	// The maximum number of elements is 100.
	DeniedReaders []*Principal `json:"deniedReaders,omitempty"`

	// InheritAclFrom: Name of the item to inherit the Access Permission
	// List (ACL) from.
	// Note: ACL inheritance *only* provides access permissions
	// to child items and does not define structural relationships, nor does
	// it
	// provide convenient ways to delete large groups of items.
	// Deleting an ACL parent from the index only alters the access
	// permissions of
	// child items that reference the parent in the
	// inheritAclFrom
	// field. The item is still in the index, but may not
	// visible in search results. By contrast, deletion of a container
	// item
	// also deletes all items that reference the container via
	// the
	// containerName
	// field.
	// The maximum length for this field is 2048 characters.
	InheritAclFrom string `json:"inheritAclFrom,omitempty"`

	// Owners: Optional. List of owners for the item. This field has no
	// bearing on
	// document access permissions currently. It does, however, supply
	// the capabilities for the `from:` and `owner:` query operators for
	// this
	// document. For example, all of the following queries can be used
	// to
	// match documents with supplied owners fields: "owner:me",
	// "from:username",
	// "owner:username@example.com." Slight ranking boosts also occur
	// for
	// documents for which the search user is an owner.
	// The maximum number of elements is 5.
	Owners []*Principal `json:"owners,omitempty"`

	// Readers: List of readers with access. Required if the item is not a
	// container
	// and no
	// inheritAclFrom field
	// is set. This is a list of
	// principals who should be allowed read/write access to the
	// document
	// and should therefore see that document in search results.
	// The maximum number of elements is 1000.
	Readers []*Principal `json:"readers,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AclInheritanceType")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AclInheritanceType") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ItemAcl) MarshalJSON() ([]byte, error) {
	type NoMethod ItemAcl
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ItemContent: Content of an item to be indexed and surfaced by Cloud
// Search.
type ItemContent struct {
	// ContentDataRef: Upload reference ID of a previously uploaded content
	// via write method.
	ContentDataRef *UploadItemRef `json:"contentDataRef,omitempty"`

	// Possible values:
	//   "UNSPECIFIED" - Invalid value.
	//   "HTML" - contentFormat is HTML.
	//   "TEXT" - contentFormat is free text.
	//   "RAW" - contentFormat is raw bytes.
	ContentFormat string `json:"contentFormat,omitempty"`

	// Hash: Hashing info calculated and provided by the API client for
	// content.
	// Can be used with the items.push method to calculate modified
	// state.
	// The maximum length is 2048 characters.
	Hash string `json:"hash,omitempty"`

	// InlineContent: Content that is supplied inlined within the update
	// method.
	// The maximum length is 102400 bytes (100 KiB).
	InlineContent string `json:"inlineContent,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ContentDataRef") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ContentDataRef") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ItemContent) MarshalJSON() ([]byte, error) {
	type NoMethod ItemContent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ItemMetadata: Available metadata fields for the item.
type ItemMetadata struct {
	// ContainerName: The name of the container for this item.
	// Deletion of the container item leads to automatic deletion of
	// this
	// item.  Note: ACLs are not inherited from a container item.
	// To provide ACL inheritance for an item, use the
	// inheritAclFrom
	// field. The maximum length is 2048 characters.
	ContainerName string `json:"containerName,omitempty"`

	// ContentLanguage: The BCP-47 language code for the item, such as
	// "en-US" or "sr-Latn". For
	// more information,
	// see
	// http://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	// Th
	// e maximum length is 32 characters.
	ContentLanguage string `json:"contentLanguage,omitempty"`

	// CreationTime: The time when the item was created in the source
	// repository.
	CreationTime string `json:"creationTime,omitempty"`

	// Hash: Hashing value provided by the API caller.
	// This can be used with the
	// items.push
	// method to calculate modified state.
	// The maximum length is 2048 characters.
	Hash string `json:"hash,omitempty"`

	// Interactions: A list of interactions for the item.  Interactions are
	// used to improve
	// Search quality, but are not exposed to end users.
	// The maximum number of elements is 1000.
	Interactions []*Interaction `json:"interactions,omitempty"`

	// LastModifiedTime: The time when the item was last modified in the
	// source repository.
	LastModifiedTime string `json:"lastModifiedTime,omitempty"`

	// MimeType: The original mime-type of
	// ItemContent.content
	// in the source repository.
	// The maximum length is 256 characters.
	MimeType string `json:"mimeType,omitempty"`

	// ObjectType: The type of the item.  This should correspond to the name
	// of an object
	// definition in the schema registered for the data source.  For
	// example, if
	// the schema for the data source contains an object definition with
	// name
	// 'document', then item indexing requests for objects of that type
	// should set
	// objectType to 'document'.
	// The maximum length is 256 characters.
	ObjectType string `json:"objectType,omitempty"`

	// SearchQualityMetadata: Search Quality metadata of the item.
	SearchQualityMetadata *SearchQualityMetadata `json:"searchQualityMetadata,omitempty"`

	// SourceRepositoryUrl: Link to the source repository serving the data.
	// &#83;earch results apply
	// this link to the title.
	// The maximum length is 2048 characters.
	SourceRepositoryUrl string `json:"sourceRepositoryUrl,omitempty"`

	// Title: The title of the item.  If given, this will be the displayed
	// title of the
	// Search result.
	// The maximum length is 2048 characters.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ContainerName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ContainerName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ItemMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod ItemMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ItemStatus: This contains item's status and any errors.
type ItemStatus struct {
	// Code: Status code.
	//
	// Possible values:
	//   "CODE_UNSPECIFIED" - Input-only value.  Used with
	// Items.list
	// to list all items in the queue, regardless of status.
	//   "ERROR" - Error encountered by Cloud Search while processing this
	// item.
	// Details of the error are in
	// repositoryError.
	//   "MODIFIED" - Item has been modified in the repository, and is out
	// of date with
	// the version previously accepted into Cloud Search.
	//   "NEW_ITEM" - Item is known to exist in the repository, but is not
	// yet accepted by
	// Cloud Search.
	// An item can be in this state when
	// Items.push
	// has been called for
	// an item of this name that did not exist previously.
	//   "ACCEPTED" - API has accepted the up-to-date data of this item.
	Code string `json:"code,omitempty"`

	// ProcessingError: Error details in case the item is in ERROR state.
	ProcessingError []*ProcessingError `json:"processingError,omitempty"`

	// RepositoryError: Repository error reported by connector.
	RepositoryError []*RepositoryError `json:"repositoryError,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ItemStatus) MarshalJSON() ([]byte, error) {
	type NoMethod ItemStatus
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ItemStructuredData: Available structured data fields for the item.
type ItemStructuredData struct {
	// Hash: Hashing value provided by the API caller.
	// This can be used with the
	// items.push
	// method to calculate modified state.
	// The maximum length is 2048 characters.
	Hash string `json:"hash,omitempty"`

	// Object: The structured data object that should conform to a
	// registered object
	// definition in the schema for the data source.
	Object *StructuredDataObject `json:"object,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Hash") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Hash") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ItemStructuredData) MarshalJSON() ([]byte, error) {
	type NoMethod ItemStructuredData
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ListDataSourceResponse struct {
	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	Sources []*DataSource `json:"sources,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListDataSourceResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListDataSourceResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ListItemsResponse struct {
	Items []*Item `json:"items,omitempty"`

	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListItemsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListItemsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// MatchRange: Matched range of a snippet [start, end).
type MatchRange struct {
	// End: End of the match in the snippet.
	End int64 `json:"end,omitempty"`

	// Start: Starting position of the match in the snippet.
	Start int64 `json:"start,omitempty"`

	// ForceSendFields is a list of field names (e.g. "End") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "End") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *MatchRange) MarshalJSON() ([]byte, error) {
	type NoMethod MatchRange
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Metadata: Metadata of a matched search result.
type Metadata struct {
	// CreateTime: The creation time for this document or object in the
	// search result.
	CreateTime string `json:"createTime,omitempty"`

	// Fields: Indexed fields in structured data, returned as a generic
	// named property.
	Fields []*NamedProperty `json:"fields,omitempty"`

	// MimeType: Mime type of the search result.
	MimeType string `json:"mimeType,omitempty"`

	// Owner: Owner (usually creator) of the document or object of the
	// search result.
	Owner *Person `json:"owner,omitempty"`

	// Source: The named source for the result, such as Gmail.
	Source *Source `json:"source,omitempty"`

	// ThumbnailUrl: The thumbnail URL of the result.
	ThumbnailUrl string `json:"thumbnailUrl,omitempty"`

	// UpdateTime: The last modified date for the object in the search
	// result.
	UpdateTime string `json:"updateTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreateTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Metadata) MarshalJSON() ([]byte, error) {
	type NoMethod Metadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Name: A person's name.
type Name struct {
	// DisplayName: The read-only display name formatted according to the
	// locale specified by
	// the viewer's account or the <code>Accept-Language</code> HTTP header.
	DisplayName string `json:"displayName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DisplayName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DisplayName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Name) MarshalJSON() ([]byte, error) {
	type NoMethod Name
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// NamedProperty: A typed name-value pair for structured data.  The type
// of the value should
// be the same as the registered type for the `name` property in the
// object
// definition of `objectType`.
type NamedProperty struct {
	BooleanValue *BooleanValue `json:"booleanValue,omitempty"`

	DateValues *DateValues `json:"dateValues,omitempty"`

	DoubleValues *DoubleValues `json:"doubleValues,omitempty"`

	EnumValues *EnumValues `json:"enumValues,omitempty"`

	FilterValues *FilterValues `json:"filterValues,omitempty"`

	IntegerValues *IntegerValues `json:"integerValues,omitempty"`

	// Name: The name of the property.  This name should correspond to the
	// name of the
	// property that was registered for object definition in the schema.
	// The maximum allowable length for this property is 256 characters.
	Name string `json:"name,omitempty"`

	ObjectValues *ObjectValues `json:"objectValues,omitempty"`

	StringValues *StringValues `json:"stringValues,omitempty"`

	TimestampValues *TimestampValues `json:"timestampValues,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BooleanValue") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BooleanValue") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *NamedProperty) MarshalJSON() ([]byte, error) {
	type NoMethod NamedProperty
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ObjectDefinition: The definition for an object within a data source.
type ObjectDefinition struct {
	// Name: Name for the object, which then defines its type. Item indexing
	// requests
	// should set the
	// objectType field
	// equal to this value. For example, if *name* is *Document*, then
	// indexing
	// requests for items of type Document should set
	// objectType equal to
	// *Document*. Each object definition must be uniquely named within a
	// schema.
	// The name must start with a letter and can only contain letters (A-Z,
	// a-z)
	// or numbers (0-9).
	// The maximum length is 256 characters.
	Name string `json:"name,omitempty"`

	// Options: The optional object-specific options.
	Options *ObjectOptions `json:"options,omitempty"`

	// PropertyDefinitions: The property definitions for the object.
	// The maximum number of elements is 1000.
	PropertyDefinitions []*PropertyDefinition `json:"propertyDefinitions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Name") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ObjectDefinition) MarshalJSON() ([]byte, error) {
	type NoMethod ObjectDefinition
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ObjectOptions: The options for an object.
type ObjectOptions struct {
	// FreshnessOptions: The freshness options for an object.
	FreshnessOptions *FreshnessOptions `json:"freshnessOptions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FreshnessOptions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FreshnessOptions") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ObjectOptions) MarshalJSON() ([]byte, error) {
	type NoMethod ObjectOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ObjectPropertyOptions: Options for object properties.
type ObjectPropertyOptions struct {
	// SubobjectProperties: The properties of the sub-object. These
	// properties represent a nested
	// object. For example, if this property represents a postal address,
	// the
	// subobjectProperties might be named *street*, *city*, and *state*.
	// The maximum number of elements is 1000.
	SubobjectProperties []*PropertyDefinition `json:"subobjectProperties,omitempty"`

	// ForceSendFields is a list of field names (e.g. "SubobjectProperties")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "SubobjectProperties") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ObjectPropertyOptions) MarshalJSON() ([]byte, error) {
	type NoMethod ObjectPropertyOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ObjectValues: List of object values.
type ObjectValues struct {
	// Values: The maximum number of elements is 100.
	Values []*StructuredDataObject `json:"values,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Values") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Values") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ObjectValues) MarshalJSON() ([]byte, error) {
	type NoMethod ObjectValues
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PeoplePhotoOptions: Options to control photos object in person.
type PeoplePhotoOptions struct {
	// PeoplePhotoUrlSizeInPx: The image size in pixels of photo_url for
	// Person.
	PeoplePhotoUrlSizeInPx int64 `json:"peoplePhotoUrlSizeInPx,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "PeoplePhotoUrlSizeInPx") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "PeoplePhotoUrlSizeInPx")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *PeoplePhotoOptions) MarshalJSON() ([]byte, error) {
	type NoMethod PeoplePhotoOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PeopleSuggestion: A people suggestion.
type PeopleSuggestion struct {
	// Person: Suggested person.
	Person *Person `json:"person,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Person") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Person") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PeopleSuggestion) MarshalJSON() ([]byte, error) {
	type NoMethod PeopleSuggestion
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Person: Object to represent a person.
type Person struct {
	// EmailAddresses: The person's email addresses
	EmailAddresses []*EmailAddress `json:"emailAddresses,omitempty"`

	// Name: The resource name of the person to provide information
	// about.
	// See <a
	// href="https://developers.google.com/people/api/rest/v1/people/get">
	// Pe
	// ople.get</a> from Google People API.
	Name string `json:"name,omitempty"`

	// PersonNames: The person's name
	PersonNames []*Name `json:"personNames,omitempty"`

	// Photos: A person's read-only photo. A picture shown next to the
	// person's name to
	// help others recognize the person in search results.
	Photos []*Photo `json:"photos,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EmailAddresses") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EmailAddresses") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Person) MarshalJSON() ([]byte, error) {
	type NoMethod Person
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Photo: A person's photo.
type Photo struct {
	// Url: The URL of the photo.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Url") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Url") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Photo) MarshalJSON() ([]byte, error) {
	type NoMethod Photo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PollItemsRequest struct {
	// Limit: Maximum number of items to return.
	// <br />The maximum and the default value is 1000
	Limit int64 `json:"limit,omitempty"`

	// Queue: Queue name to fetch items from.  If unspecified, PollItems
	// will
	// fetch from 'default' queue.
	// The maximum length is 100 characters.
	Queue string `json:"queue,omitempty"`

	// StatusCodes: Limit the items polled to the ones with these statuses.
	//
	// Possible values:
	//   "CODE_UNSPECIFIED" - Input-only value.  Used with
	// Items.list
	// to list all items in the queue, regardless of status.
	//   "ERROR" - Error encountered by Cloud Search while processing this
	// item.
	// Details of the error are in
	// repositoryError.
	//   "MODIFIED" - Item has been modified in the repository, and is out
	// of date with
	// the version previously accepted into Cloud Search.
	//   "NEW_ITEM" - Item is known to exist in the repository, but is not
	// yet accepted by
	// Cloud Search.
	// An item can be in this state when
	// Items.push
	// has been called for
	// an item of this name that did not exist previously.
	//   "ACCEPTED" - API has accepted the up-to-date data of this item.
	StatusCodes []string `json:"statusCodes,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Limit") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Limit") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PollItemsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod PollItemsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PollItemsResponse struct {
	// Items: Set of items from the queue available for connector to
	// process.
	// <br />These items have the following subset of fields populated: <br
	// />
	// <br />version
	// <br />metadata.hash
	// <br />structured_data.hash
	// <br />content.hash
	// <br />payload
	// <br />status
	// <br />queue
	Items []*Item `json:"items,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PollItemsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod PollItemsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Principal: Reference to a user, group, or domain.
type Principal struct {
	// GroupResourceName: This principal is a group identified using an
	// external identity.
	// The name field must specify the group resource name with this
	// format:
	// identitysources/{source_id}/groups/{ID}
	GroupResourceName string `json:"groupResourceName,omitempty"`

	// GsuitePrincipal: This principal is a GSuite user, group or domain.
	GsuitePrincipal *GSuitePrincipal `json:"gsuitePrincipal,omitempty"`

	// UserResourceName: This principal is a user identified using an
	// external identity.
	// The name field must specify the user resource name with this
	// format:
	// identitysources/{source_id}/users/{ID}
	UserResourceName string `json:"userResourceName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "GroupResourceName")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "GroupResourceName") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Principal) MarshalJSON() ([]byte, error) {
	type NoMethod Principal
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ProcessingError struct {
	// Code: Error code indicating the nature of the error.
	//
	// Possible values:
	//   "PROCESSING_ERROR_CODE_UNSPECIFIED" - Input only value.  Use this
	// value in Items.
	//   "MALFORMED_REQUEST" - Item's ACL, metadata, or content is malformed
	// or in invalid state.
	// FieldViolations contains more details on where the problem is.
	//   "UNSUPPORTED_CONTENT_FORMAT" - Countent format is unsupported.
	//   "INDIRECT_BROKEN_ACL" - Items with incomplete ACL information due
	// to inheriting other
	// items with broken ACL or having groups with unmapped descendants.
	//   "ACL_CYCLE" - ACL inheritance graph formed a cycle.
	Code string `json:"code,omitempty"`

	// ErrorMessage: Description of the error.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// FieldViolations: In case the item fields are invalid, this field
	// contains the details
	// about the validation errors.
	FieldViolations []*FieldViolation `json:"fieldViolations,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ProcessingError) MarshalJSON() ([]byte, error) {
	type NoMethod ProcessingError
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PropertyDefinition: The definition of a property within an object.
type PropertyDefinition struct {
	BooleanPropertyOptions *BooleanPropertyOptions `json:"booleanPropertyOptions,omitempty"`

	DatePropertyOptions *DatePropertyOptions `json:"datePropertyOptions,omitempty"`

	DoublePropertyOptions *DoublePropertyOptions `json:"doublePropertyOptions,omitempty"`

	EnumPropertyOptions *EnumPropertyOptions `json:"enumPropertyOptions,omitempty"`

	FilterPropertyOptions *FilterPropertyOptions `json:"filterPropertyOptions,omitempty"`

	IntegerPropertyOptions *IntegerPropertyOptions `json:"integerPropertyOptions,omitempty"`

	// IsRepeatable: Indicates if multiple values are allowed for the
	// property. For example, a
	// document only has one description but can have multiple comments.
	// Cannot be
	// true for properties whose type is a boolean.
	// If set to false, properties that contain more than one value will
	// cause the
	// indexing request for that item to be rejected.
	IsRepeatable bool `json:"isRepeatable,omitempty"`

	// IsReturnable: Indicates if the property identifies data that should
	// be returned in search
	// results via the Query API. If set to *true*, indicates that Query
	// API
	// users can use matching property fields in results. However, storing
	// fields
	// requires more space allocation and uses more bandwidth for search
	// queries,
	// which impacts performance over large datasets. Set to *true* here
	// only if
	// the field is needed for search results. Cannot be true for
	// properties
	// whose type is an object.
	IsReturnable bool `json:"isReturnable,omitempty"`

	// Name: The name of the property. Item indexing requests sent to the
	// Indexing API
	// should set the property name
	// equal to this value. For example, if name is *subject_line*, then
	// indexing
	// requests for document items with subject fields should set the
	// name for that field equal to
	// *subject_line*. Use the name as the identifier for the object
	// property.
	// Once registered as a property for an object, you cannot re-use this
	// name
	// for another property within that object.
	// The name must start with a letter and can only contain letters (A-Z,
	// a-z)
	// or numbers (0-9).
	// The maximum length is 256 characters.
	Name string `json:"name,omitempty"`

	ObjectPropertyOptions *ObjectPropertyOptions `json:"objectPropertyOptions,omitempty"`

	StringPropertyOptions *StringPropertyOptions `json:"stringPropertyOptions,omitempty"`

	TimestampPropertyOptions *TimestampPropertyOptions `json:"timestampPropertyOptions,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "BooleanPropertyOptions") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BooleanPropertyOptions")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *PropertyDefinition) MarshalJSON() ([]byte, error) {
	type NoMethod PropertyDefinition
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PushItem: Represents an item to be pushed to the indexing queue.
type PushItem struct {
	// ContentHash: Content hash of the item according to the repository. If
	// specified, this is
	// used to determine how to modify this
	// item's status. Setting this field and the
	// type field results in argument
	// error.
	// The maximum length is 2048 characters.
	ContentHash string `json:"contentHash,omitempty"`

	// MetadataHash: Metadata hash of the item according to the repository.
	// If specified, this
	// is used to determine how to modify this
	// item's status. Setting this field and the
	// type field results in argument
	// error.
	// The maximum length is 2048 characters.
	MetadataHash string `json:"metadataHash,omitempty"`

	// Payload: Provides additional document state information for the
	// connector,
	// such as an alternate repository ID and other metadata.
	// The maximum length is 8192 bytes.
	Payload string `json:"payload,omitempty"`

	// Queue: Queue to which this item belongs to.  The <code>default</code>
	// queue is
	// chosen if this field is not specified. The maximum length is
	// 512 characters.
	Queue string `json:"queue,omitempty"`

	// RepositoryError: Populate this field to store Connector or repository
	// error details.
	// This information is displayed in the Admin Console.
	// This field may only be populated when the
	// Type is
	// REPOSITORY_ERROR.
	RepositoryError *RepositoryError `json:"repositoryError,omitempty"`

	// StructuredDataHash: Structured data hash of the item according to the
	// repository. If specified,
	// this is used to determine how to modify this item's status. Setting
	// this
	// field and the type field
	// results in argument error.
	// The maximum length is 2048 characters.
	StructuredDataHash string `json:"structuredDataHash,omitempty"`

	// Type: The type of the push operation that defines the push behavior.
	//
	// Possible values:
	//   "UNSPECIFIED" - Default UNSPECIFIED.  Specifies that the push
	// operation should not modify
	// ItemStatus
	//   "MODIFIED" - Indicates that the repository document has been
	// modified or updated since
	// the previous
	// index
	// call. This changes status to
	// MODIFIED state for
	// an existing item. If this is called on a non existing item, the
	// status is
	// changed to
	// NEW_ITEM.
	//   "NOT_MODIFIED" - Item in the repository has not been modified since
	// the last update
	// call.  This push operation will set status to
	// ACCEPTED state.
	//   "REPOSITORY_ERROR" - Connector is facing a repository error
	// regarding this item.  Change
	// status to
	// ERROR
	// state. Item is unreserved and rescheduled at a future time determined
	// by
	// exponential backoff.
	//   "REQUEUE" - Call push with REQUEUE only for items that have been
	// reserved.
	// This action unreserves the item and resets its available time to
	// the
	// wall clock time.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ContentHash") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ContentHash") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PushItem) MarshalJSON() ([]byte, error) {
	type NoMethod PushItem
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PushItemRequest struct {
	// Item: Item to push onto the queue.
	Item *PushItem `json:"item,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Item") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Item") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PushItemRequest) MarshalJSON() ([]byte, error) {
	type NoMethod PushItemRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type QueryInterpretation struct {
	// Possible values:
	//   "NONE" - No natural language interpretation or the natural language
	// interpretation
	// is not used to fetch the search results.
	//   "BLEND" - The natural language results is mixed with results from
	// original query.
	//   "REPLACE" - The results only contain natural language results.
	InterpretationType string `json:"interpretationType,omitempty"`

	// InterpretedQuery: The interpretation of the query used in search. For
	// example, query "email
	// from john" will be interpreted as "from:john source:mail"
	InterpretedQuery string `json:"interpretedQuery,omitempty"`

	// ForceSendFields is a list of field names (e.g. "InterpretationType")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "InterpretationType") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *QueryInterpretation) MarshalJSON() ([]byte, error) {
	type NoMethod QueryInterpretation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// QueryInterpretationOptions: Options to interpret user query.
type QueryInterpretationOptions struct {
	// DisableNlInterpretation: Flag to disable query parsing. Default is
	// false, Set to true to disable
	// query parsing.
	DisableNlInterpretation bool `json:"disableNlInterpretation,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "DisableNlInterpretation") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DisableNlInterpretation")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *QueryInterpretationOptions) MarshalJSON() ([]byte, error) {
	type NoMethod QueryInterpretationOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// QueryItem: Information relevant only to a query entry.
type QueryItem struct {
	// IsSynthetic: True if the text was generated by means other than a
	// previous user search.
	IsSynthetic bool `json:"isSynthetic,omitempty"`

	// ForceSendFields is a list of field names (e.g. "IsSynthetic") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "IsSynthetic") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *QueryItem) MarshalJSON() ([]byte, error) {
	type NoMethod QueryItem
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// QuerySuggestion: A completed query suggestion.
type QuerySuggestion struct {
	// Suggestion: Suggested query.
	Suggestion string `json:"suggestion,omitempty"`

	// Type: Type of the suggestion.
	//
	// Possible values:
	//   "GENERIC" - Suggestion from documents in corpus.
	//   "HISTORY" - Personal query history.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Suggestion") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Suggestion") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *QuerySuggestion) MarshalJSON() ([]byte, error) {
	type NoMethod QuerySuggestion
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Range: The range of values [start, end) for which faceting is applied
type Range struct {
	End *Value `json:"end,omitempty"`

	Start *Value `json:"start,omitempty"`

	// ForceSendFields is a list of field names (e.g. "End") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "End") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Range) MarshalJSON() ([]byte, error) {
	type NoMethod Range
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RepositoryError: Errors when the connector is communicating to the
// source repository.
type RepositoryError struct {
	// ErrorMessage: Message that describes the error. The maximum allowable
	// length
	// of the message is 8192 characters.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// HttpStatusCode: Error codes.  Matches the definition of HTTP status
	// codes.
	HttpStatusCode int64 `json:"httpStatusCode,omitempty"`

	// Type: Type of error.
	//
	// Possible values:
	//   "UNKNOWN" - Unknown error.
	//   "NETWORK_ERROR" - Unknown or unreachable host.
	//   "DNS_ERROR" - DNS problem, such as the DNS server is not
	// responding.
	//   "CONNECTION_ERROR" - Cannot connect to the repository server.
	//   "AUTHENTICATION_ERROR" - Failed authentication due to incorrect
	// credentials.
	//   "AUTHORIZATION_ERROR" - Service account is not authorized for the
	// repository.
	//   "SERVER_ERROR" - Repository server error.
	//   "QUOTA_EXCEEDED" - Quota exceeded.
	//   "SERVICE_UNAVAILABLE" - Server temporarily unavailable.
	//   "CLIENT_ERROR" - Client-related error, such as an invalid request
	// from the connector to
	// the repository server.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ErrorMessage") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ErrorMessage") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RepositoryError) MarshalJSON() ([]byte, error) {
	type NoMethod RepositoryError
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RequestOptions: Shared request options for all RPC methods.
type RequestOptions struct {
	// DocumentThumbnailOptions: The options to configure document thumbnail
	// url.
	DocumentThumbnailOptions *DocumentThumbnailOptions `json:"documentThumbnailOptions,omitempty"`

	// LanguageCode: The BCP-47 language code, such as "en-US" or
	// "sr-Latn".
	// For more information,
	// see
	// http://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	// Fo
	// r translations.
	LanguageCode string `json:"languageCode,omitempty"`

	// PeoplePhotoOptions: The options to configure photo profile url.
	PeoplePhotoOptions *PeoplePhotoOptions `json:"peoplePhotoOptions,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "DocumentThumbnailOptions") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DocumentThumbnailOptions")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *RequestOptions) MarshalJSON() ([]byte, error) {
	type NoMethod RequestOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RestrictItem: Information relevant only to a restrict entry.
// NextId: 7
type RestrictItem struct {
	DriveFollowUpRestrict *DriveFollowUpRestrict `json:"driveFollowUpRestrict,omitempty"`

	DriveLocationRestrict *DriveLocationRestrict `json:"driveLocationRestrict,omitempty"`

	DriveMimeTypeRestrict *DriveMimeTypeRestrict `json:"driveMimeTypeRestrict,omitempty"`

	DriveTimeSpanRestrict *DriveTimeSpanRestrict `json:"driveTimeSpanRestrict,omitempty"`

	// SearchOperator: The search restrict (e.g. "after:2017-09-11
	// before:2017-09-12").
	SearchOperator string `json:"searchOperator,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "DriveFollowUpRestrict") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DriveFollowUpRestrict") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *RestrictItem) MarshalJSON() ([]byte, error) {
	type NoMethod RestrictItem
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Schema: The schema definition for a data source.
type Schema struct {
	// ObjectDefinitions: The list of top-level objects for the data
	// source.
	// The maximum number of elements is 10.
	ObjectDefinitions []*ObjectDefinition `json:"objectDefinitions,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ObjectDefinitions")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ObjectDefinitions") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Schema) MarshalJSON() ([]byte, error) {
	type NoMethod Schema
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SearchQualityMetadata: Search Quality metadata of the item.
type SearchQualityMetadata struct {
	// Quality: An indication of the quality of the item, used to influence
	// Search quality.
	// Value should be between 0.0 (lowest quality) and 1.0 (highest
	// quality).
	Quality float64 `json:"quality,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Quality") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Quality") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SearchQualityMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod SearchQualityMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *SearchQualityMetadata) UnmarshalJSON(data []byte) error {
	type NoMethod SearchQualityMetadata
	var s1 struct {
		Quality gensupport.JSONFloat64 `json:"quality"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Quality = float64(s1.Quality)
	return nil
}

// SearchRequest: The search API request.
type SearchRequest struct {
	// FacetOptions: The options to enable facet.
	FacetOptions []*FacetOptions `json:"facetOptions,omitempty"`

	// PageSize: Maximum number of search results to return in one
	// page.
	// Valid values are between 1 and 100, inclusive.
	// Default value is 10.
	PageSize int64 `json:"pageSize,omitempty"`

	// Query: The raw query string.
	// See supported search operators in the [Cloud
	// search
	// Cheat
	// Sheet](https://gsuite.google.com/learning-center/products
	// /cloudsearch/cheat-sheet/)
	Query string `json:"query,omitempty"`

	// QueryInterpretationOptions: Options to interpret the user query.
	QueryInterpretationOptions *QueryInterpretationOptions `json:"queryInterpretationOptions,omitempty"`

	// RequestOptions: Request options.
	RequestOptions *RequestOptions `json:"requestOptions,omitempty"`

	// SortOptions: The options for sorting the search results
	SortOptions *SortOptions `json:"sortOptions,omitempty"`

	// SourceOptions: The sources to use for querying. You must specifiy at
	// least one source
	// to search.
	SourceOptions []*SearchSourceOptions `json:"sourceOptions,omitempty"`

	// Start: Starting index of the results.
	Start int64 `json:"start,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FacetOptions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FacetOptions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SearchRequest) MarshalJSON() ([]byte, error) {
	type NoMethod SearchRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SearchResponse: The search API response.
type SearchResponse struct {
	// FacetResults: Repeated facet results.
	FacetResults []*FacetResult `json:"facetResults,omitempty"`

	// HasMoreResults: Whether there are more search results matching the
	// query.
	HasMoreResults bool `json:"hasMoreResults,omitempty"`

	// QueryInterpretation: Query interpretation result for user query.
	// Empty if query parsing is
	// disabled.
	QueryInterpretation *QueryInterpretation `json:"queryInterpretation,omitempty"`

	// Results: Results from a search query.
	Results []*SearchResult `json:"results,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "FacetResults") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FacetResults") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SearchResponse) MarshalJSON() ([]byte, error) {
	type NoMethod SearchResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SearchResult: Results containing indexed information for a document.
type SearchResult struct {
	// Metadata: Metadata of the search result.
	Metadata *Metadata `json:"metadata,omitempty"`

	// Snippet: The concatenation of all snippets (summaries) available for
	// this result.
	Snippet *Snippet `json:"snippet,omitempty"`

	// Title: Title of the search result.
	Title string `json:"title,omitempty"`

	// Url: The URL of the result.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Metadata") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Metadata") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SearchResult) MarshalJSON() ([]byte, error) {
	type NoMethod SearchResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SearchSourceOptions: Configure a source to query the search API.
type SearchSourceOptions struct {
	// FilterOptions: Retrictions applied to this source.
	// The maximum number of elements is 20.
	FilterOptions []*FilterOptions `json:"filterOptions,omitempty"`

	// Source: The source of the search.
	Source *Source `json:"source,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FilterOptions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FilterOptions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SearchSourceOptions) MarshalJSON() ([]byte, error) {
	type NoMethod SearchSourceOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Snippet: Snippet of the search result, which summarizes the content
// of the resulting
// page.
type Snippet struct {
	// MatchRanges: The matched ranges in the snippet.
	MatchRanges []*MatchRange `json:"matchRanges,omitempty"`

	// Snippet: The snippet.
	Snippet string `json:"snippet,omitempty"`

	// ForceSendFields is a list of field names (e.g. "MatchRanges") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "MatchRanges") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Snippet) MarshalJSON() ([]byte, error) {
	type NoMethod Snippet
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type SortOptions struct {
	// OperatorName: Name of the operator corresponding to the field(s) to
	// sort on
	OperatorName string `json:"operatorName,omitempty"`

	// SortOrder: Ascending is the default sort order
	//
	// Possible values:
	//   "ASCENDING"
	//   "DESCENDING"
	SortOrder string `json:"sortOrder,omitempty"`

	// ForceSendFields is a list of field names (e.g. "OperatorName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "OperatorName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SortOptions) MarshalJSON() ([]byte, error) {
	type NoMethod SortOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Source: Defines sources for the suggest/search APIs.
type Source struct {
	// Name: Source name for content indexed by the
	// Indexing API
	Name string `json:"name,omitempty"`

	// PredefinedSource: Predefined content source for Google Apps.
	//
	// Possible values:
	//   "GENERIC" - Default source. Used for suggest API only. Ignored in
	// search API.
	//   "PEOPLE" - People source. Used for suggest API only. Ignored in
	// search API.
	//   "GOOGLE_DRIVE"
	//   "GOOGLE_GMAIL"
	//   "GOOGLE_SITES"
	//   "GOOGLE_GROUPS"
	//   "GOOGLE_CALENDAR"
	//   "GOOGLE_KEEP"
	PredefinedSource string `json:"predefinedSource,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Name") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Source) MarshalJSON() ([]byte, error) {
	type NoMethod Source
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// StartUploadItemRequest: Start upload file request.
type StartUploadItemRequest struct {
}

// StringOperatorOptions: Used to provide a search operator for string
// properties. This is optional.
// Search operators let users restrict the query to specific fields
// relevant
// to the type of item being searched.
type StringOperatorOptions struct {
	// OperatorName: Indicates the operator name required in the query in
	// order to isolate the
	// string property. For example, if operatorName is *subject* and
	// the
	// property's name is *subjectLine*, then queries
	// like
	// *subject:&lt;value&gt;* will show results only where the value of
	// the
	// property named *subjectLine* matches *&lt;value&gt;*. By contrast,
	// a
	// search that uses the same *&lt;value&gt;* without an operator will
	// return
	// all items where *&lt;value&gt;* matches the value of any String
	// properties
	// or text within the content field for the item.
	// The operator name can only contain lowercase letters (a-z).
	// The maximum length is 32 characters.
	OperatorName string `json:"operatorName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "OperatorName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "OperatorName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *StringOperatorOptions) MarshalJSON() ([]byte, error) {
	type NoMethod StringOperatorOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// StringPropertyOptions: Options for string properties. String property
// values will be indexed and
// used for full text retrieval.
type StringPropertyOptions struct {
	// OperatorOptions: If set, describes how the string should be used as a
	// search operator.
	OperatorOptions *StringOperatorOptions `json:"operatorOptions,omitempty"`

	// TokenImportance: Indicates the search quality importance of the
	// tokens within the
	// field. Can only be used if Tokenization is TEXT. For HTML
	// tokenization,
	// HTML tags around the tokens indicate the importance.
	//
	// Possible values:
	//   "NORMAL" - Treat the text like body text.
	//   "HIGHEST" - Treat the text like the title of the item.
	//   "HIGH" - Treat the text with higher importance than body text.
	//   "LOW" - Treat the text with lower importance than body text.
	TokenImportance string `json:"tokenImportance,omitempty"`

	// Tokenization: Indicates how to tokenize the string's value.
	//
	// Possible values:
	//   "UNSPECIFIED" - Invalid value.
	//   "TEXT" - The string is tokenized as text.
	//   "HTML" - The string is tokenized as HTML.
	Tokenization string `json:"tokenization,omitempty"`

	// ForceSendFields is a list of field names (e.g. "OperatorOptions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "OperatorOptions") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *StringPropertyOptions) MarshalJSON() ([]byte, error) {
	type NoMethod StringPropertyOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// StringValues: List of string values.
type StringValues struct {
	// Values: The maximum allowable length for string values is 2048
	// characters.
	// The maximum number of string elements is 100.
	Values []string `json:"values,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Values") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Values") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *StringValues) MarshalJSON() ([]byte, error) {
	type NoMethod StringValues
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// StructuredDataObject: A structured data object consisting of named
// properties.
type StructuredDataObject struct {
	// Properties: The properties for the object.
	// The maximum number of elements is 1000.
	Properties []*NamedProperty `json:"properties,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Properties") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Properties") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *StructuredDataObject) MarshalJSON() ([]byte, error) {
	type NoMethod StructuredDataObject
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SuggestRequest: Request of suggest API.
type SuggestRequest struct {
	// Query: Partial query for the completion suggestion.
	Query string `json:"query,omitempty"`

	// RequestOptions: Request options, such as image options for people
	// info and thumbnails.
	RequestOptions *RequestOptions `json:"requestOptions,omitempty"`

	// SourceOptions: suggest options.
	SourceOptions []*SuggestSourceOptions `json:"sourceOptions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Query") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Query") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SuggestRequest) MarshalJSON() ([]byte, error) {
	type NoMethod SuggestRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SuggestResponse: Response of the suggest API.
type SuggestResponse struct {
	// PeopleSuggestions: List of people suggestions.
	PeopleSuggestions []*PeopleSuggestion `json:"peopleSuggestions,omitempty"`

	// QuerySuggestions: List of query text suggestions.
	QuerySuggestions []*QuerySuggestion `json:"querySuggestions,omitempty"`

	// SuggestResults: List of suggestion results.
	SuggestResults []*SuggestResult `json:"suggestResults,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "PeopleSuggestions")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "PeopleSuggestions") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *SuggestResponse) MarshalJSON() ([]byte, error) {
	type NoMethod SuggestResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SuggestResult: One suggestion result.
type SuggestResult struct {
	// Name: The name of the suggestion.
	Name string `json:"name,omitempty"`

	PeopleSuggestion *PeopleSuggestion `json:"peopleSuggestion,omitempty"`

	QuerySuggestion *QuerySuggestion `json:"querySuggestion,omitempty"`

	// SuggestedQuery: The suggested query that will be used for search,
	// when the user
	// clicks on the suggestion
	SuggestedQuery string `json:"suggestedQuery,omitempty"`

	// Possible values:
	//   "DEFAULT" - Suggestion with default type.
	//   "HISTORY" - Personal query history.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Name") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SuggestResult) MarshalJSON() ([]byte, error) {
	type NoMethod SuggestResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SuggestSourceOptions: Configure a source to query the suggest API.
type SuggestSourceOptions struct {
	// MaxNumResults: Maximum number of results to return from this source.
	MaxNumResults int64 `json:"maxNumResults,omitempty"`

	// Source: The source of the suggestion.
	Source *Source `json:"source,omitempty"`

	// ForceSendFields is a list of field names (e.g. "MaxNumResults") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "MaxNumResults") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SuggestSourceOptions) MarshalJSON() ([]byte, error) {
	type NoMethod SuggestSourceOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TimestampOperatorOptions: Used to provide a search operator for
// timestamp properties. This is
// optional. Search operators let users restrict the query to specific
// fields
// relevant to the type of item being searched.
type TimestampOperatorOptions struct {
	// GreaterThanOperatorName: Indicates the operator name required in the
	// query in order to isolate the
	// timestamp property using the greater-than operator. For example,
	// if
	// greaterThanOperatorName is *closedafter* and the property's name
	// is
	// *closeDate*, then queries like *closedafter:&lt;value&gt;* will
	// show results only where the value of the property named *closeDate*
	// is
	// later than *&lt;value&gt;*.
	// The operator name can only contain lowercase letters (a-z).
	// The maximum length is 32 characters.
	GreaterThanOperatorName string `json:"greaterThanOperatorName,omitempty"`

	// LessThanOperatorName: Indicates the operator name required in the
	// query in order to isolate the
	// timestamp property using the less-than operator. For example,
	// if
	// lessThanOperatorName is *closedbefore* and the property's name
	// is
	// *closeDate*, then queries like *closedbefore:&lt;value&gt;* will
	// show results only where the value of the property named *closeDate*
	// is
	// earlier than *&lt;value&gt;*.
	// The operator name can only contain lowercase letters (a-z).
	// The maximum length is 32 characters.
	LessThanOperatorName string `json:"lessThanOperatorName,omitempty"`

	// OperatorName: Indicates the operator name required in the query in
	// order to isolate the
	// timestamp property. For example, if operatorName is *closedon* and
	// the
	// property's name is *closeDate*, then queries
	// like
	// *closedon:&lt;value&gt;* will show results only where the value of
	// the
	// property named *closeDate* matches *&lt;value&gt;*. By contrast,
	// a
	// search that uses the same *&lt;value&gt;* without an operator will
	// return
	// all items where *&lt;value&gt;* matches the value of any String
	// properties
	// or text within the content field for the item.
	// The operator name can only contain lowercase letters (a-z).
	// The maximum length is 32 characters.
	OperatorName string `json:"operatorName,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "GreaterThanOperatorName") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "GreaterThanOperatorName")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *TimestampOperatorOptions) MarshalJSON() ([]byte, error) {
	type NoMethod TimestampOperatorOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TimestampPropertyOptions: Options for timestamp properties.
type TimestampPropertyOptions struct {
	// OperatorOptions: If set, describes how the timestamp should be used
	// as a search operator.
	OperatorOptions *TimestampOperatorOptions `json:"operatorOptions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "OperatorOptions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "OperatorOptions") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *TimestampPropertyOptions) MarshalJSON() ([]byte, error) {
	type NoMethod TimestampPropertyOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TimestampValues: List of timestamp values.
type TimestampValues struct {
	// Values: The maximum number of elements is 100.
	Values []string `json:"values,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Values") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Values") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TimestampValues) MarshalJSON() ([]byte, error) {
	type NoMethod TimestampValues
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type UnreserveItemsRequest struct {
	// Queue: Name of a queue to unreserve items from.
	Queue string `json:"queue,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Queue") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Queue") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UnreserveItemsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod UnreserveItemsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type UpdateDataSourceRequest struct {
	Source *DataSource `json:"source,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Source") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Source") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UpdateDataSourceRequest) MarshalJSON() ([]byte, error) {
	type NoMethod UpdateDataSourceRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type UpdateSchemaRequest struct {
	// Schema: The new schema for the source.
	Schema *Schema `json:"schema,omitempty"`

	// ValidateOnly: If true, the request will be validated without side
	// effects.
	ValidateOnly bool `json:"validateOnly,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Schema") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Schema") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UpdateSchemaRequest) MarshalJSON() ([]byte, error) {
	type NoMethod UpdateSchemaRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UploadItemRef: Represents an upload session reference.
// This reference is created via upload
// method.
// Updating of item content may refer to this uploaded content
// via
// contentDataRef.
type UploadItemRef struct {
	// ResourceName: Name of the content reference.
	// The maximum length is 2048 characters.
	ResourceName string `json:"resourceName,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ResourceName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ResourceName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UploadItemRef) MarshalJSON() ([]byte, error) {
	type NoMethod UploadItemRef
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Value: Definition of a single value with generic type.
type Value struct {
	BooleanValue bool `json:"booleanValue,omitempty"`

	DateValue *Date `json:"dateValue,omitempty"`

	DoubleValue float64 `json:"doubleValue,omitempty"`

	IntegerValue int64 `json:"integerValue,omitempty,string"`

	StringValue string `json:"stringValue,omitempty"`

	TimestampValue string `json:"timestampValue,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BooleanValue") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BooleanValue") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Value) MarshalJSON() ([]byte, error) {
	type NoMethod Value
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *Value) UnmarshalJSON(data []byte) error {
	type NoMethod Value
	var s1 struct {
		DoubleValue gensupport.JSONFloat64 `json:"doubleValue"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.DoubleValue = float64(s1.DoubleValue)
	return nil
}

type ValueFilter struct {
	// OperatorName: The `operator_name` applied to the query, such as
	// *price_greater_than*.
	// The filter can work against both types of filters defined in the
	// schema
	// for your data source:
	// <br/><br/>
	// 1. `operator_name`, where the query filters results by the
	// property
	// that matches the value.
	// <br/>
	// 2. `greater_than_operator_name` or `less_than_operator_name` in
	// your
	// schema. The query filters the results for the property values that
	// are
	// greater than or less than  the supplied value in the query.
	OperatorName string `json:"operatorName,omitempty"`

	// Value: The value to be compared with.
	Value *Value `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "OperatorName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "OperatorName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ValueFilter) MarshalJSON() ([]byte, error) {
	type NoMethod ValueFilter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "cloudsearch.indexing.datasources.deleteSchema":

type IndexingDatasourcesDeleteSchemaCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// DeleteSchema: Deletes the schema of a data source.
func (r *IndexingDatasourcesService) DeleteSchema(name string) *IndexingDatasourcesDeleteSchemaCall {
	c := &IndexingDatasourcesDeleteSchemaCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *IndexingDatasourcesDeleteSchemaCall) Fields(s ...googleapi.Field) *IndexingDatasourcesDeleteSchemaCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *IndexingDatasourcesDeleteSchemaCall) Context(ctx context.Context) *IndexingDatasourcesDeleteSchemaCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *IndexingDatasourcesDeleteSchemaCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *IndexingDatasourcesDeleteSchemaCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/indexing/{+name}/schema")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.indexing.datasources.deleteSchema" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *IndexingDatasourcesDeleteSchemaCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the schema of a data source.",
	//   "flatPath": "v1beta1/indexing/datasources/{datasourcesId}/schema",
	//   "httpMethod": "DELETE",
	//   "id": "cloudsearch.indexing.datasources.deleteSchema",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name of the data source to delete Schema.  Format:\ndatasources/{source_id}",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/indexing/{+name}/schema",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.settings",
	//     "https://www.googleapis.com/auth/cloud_search.settings.indexing"
	//   ]
	// }

}

// method id "cloudsearch.indexing.datasources.getSchema":

type IndexingDatasourcesGetSchemaCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetSchema: Gets the schema of a data source.
func (r *IndexingDatasourcesService) GetSchema(name string) *IndexingDatasourcesGetSchemaCall {
	c := &IndexingDatasourcesGetSchemaCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *IndexingDatasourcesGetSchemaCall) Fields(s ...googleapi.Field) *IndexingDatasourcesGetSchemaCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *IndexingDatasourcesGetSchemaCall) IfNoneMatch(entityTag string) *IndexingDatasourcesGetSchemaCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *IndexingDatasourcesGetSchemaCall) Context(ctx context.Context) *IndexingDatasourcesGetSchemaCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *IndexingDatasourcesGetSchemaCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *IndexingDatasourcesGetSchemaCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/indexing/{+name}/schema")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.indexing.datasources.getSchema" call.
// Exactly one of *Schema or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Schema.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *IndexingDatasourcesGetSchemaCall) Do(opts ...googleapi.CallOption) (*Schema, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Schema{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the schema of a data source.",
	//   "flatPath": "v1beta1/indexing/datasources/{datasourcesId}/schema",
	//   "httpMethod": "GET",
	//   "id": "cloudsearch.indexing.datasources.getSchema",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name of the data source to get Schema.  Format:\ndatasources/{source_id}",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/indexing/{+name}/schema",
	//   "response": {
	//     "$ref": "Schema"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.settings",
	//     "https://www.googleapis.com/auth/cloud_search.settings.indexing"
	//   ]
	// }

}

// method id "cloudsearch.indexing.datasources.updateSchema":

type IndexingDatasourcesUpdateSchemaCall struct {
	s                   *Service
	name                string
	updateschemarequest *UpdateSchemaRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// UpdateSchema: Updates the schema of a data source.
func (r *IndexingDatasourcesService) UpdateSchema(name string, updateschemarequest *UpdateSchemaRequest) *IndexingDatasourcesUpdateSchemaCall {
	c := &IndexingDatasourcesUpdateSchemaCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.updateschemarequest = updateschemarequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *IndexingDatasourcesUpdateSchemaCall) Fields(s ...googleapi.Field) *IndexingDatasourcesUpdateSchemaCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *IndexingDatasourcesUpdateSchemaCall) Context(ctx context.Context) *IndexingDatasourcesUpdateSchemaCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *IndexingDatasourcesUpdateSchemaCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *IndexingDatasourcesUpdateSchemaCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.updateschemarequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/indexing/{+name}/schema")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.indexing.datasources.updateSchema" call.
// Exactly one of *Schema or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Schema.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *IndexingDatasourcesUpdateSchemaCall) Do(opts ...googleapi.CallOption) (*Schema, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Schema{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the schema of a data source.",
	//   "flatPath": "v1beta1/indexing/datasources/{datasourcesId}/schema",
	//   "httpMethod": "PUT",
	//   "id": "cloudsearch.indexing.datasources.updateSchema",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name of the data source to update Schema.  Format:\ndatasources/{source_id}",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/indexing/{+name}/schema",
	//   "request": {
	//     "$ref": "UpdateSchemaRequest"
	//   },
	//   "response": {
	//     "$ref": "Schema"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.settings",
	//     "https://www.googleapis.com/auth/cloud_search.settings.indexing"
	//   ]
	// }

}

// method id "cloudsearch.indexing.datasources.items.delete":

type IndexingDatasourcesItemsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes Item resource for the
// specified resource name.
func (r *IndexingDatasourcesItemsService) Delete(name string) *IndexingDatasourcesItemsDeleteCall {
	c := &IndexingDatasourcesItemsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Version sets the optional parameter "version": Required. The
// incremented version of the item to delete from the index.
// The indexing system stores the version from the datasource as a
// byte string and compares the Item version in the index
// to the version of the queued Item using lexical ordering.
// <br /><br />
// Cloud Search Indexing won't delete any queued item with
// a version value that is less than or equal to
// the version of the currently indexed item.
// The maximum length for this field is 1024 bytes.
func (c *IndexingDatasourcesItemsDeleteCall) Version(version string) *IndexingDatasourcesItemsDeleteCall {
	c.urlParams_.Set("version", version)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *IndexingDatasourcesItemsDeleteCall) Fields(s ...googleapi.Field) *IndexingDatasourcesItemsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *IndexingDatasourcesItemsDeleteCall) Context(ctx context.Context) *IndexingDatasourcesItemsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *IndexingDatasourcesItemsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *IndexingDatasourcesItemsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/indexing/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.indexing.datasources.items.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *IndexingDatasourcesItemsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes Item resource for the\nspecified resource name.",
	//   "flatPath": "v1beta1/indexing/datasources/{datasourcesId}/items/{itemsId}",
	//   "httpMethod": "DELETE",
	//   "id": "cloudsearch.indexing.datasources.items.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. Name of the item to delete.\nFormat: datasources/{source_id}/items/{item_id}",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+/items/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "version": {
	//       "description": "Required. The incremented version of the item to delete from the index.\nThe indexing system stores the version from the datasource as a\nbyte string and compares the Item version in the index\nto the version of the queued Item using lexical ordering.\n\u003cbr /\u003e\u003cbr /\u003e\nCloud Search Indexing won't delete any queued item with\na version value that is less than or equal to\nthe version of the currently indexed item.\nThe maximum length for this field is 1024 bytes.",
	//       "format": "byte",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/indexing/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.indexing"
	//   ]
	// }

}

// method id "cloudsearch.indexing.datasources.items.get":

type IndexingDatasourcesItemsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets Item resource by item name.
func (r *IndexingDatasourcesItemsService) Get(name string) *IndexingDatasourcesItemsGetCall {
	c := &IndexingDatasourcesItemsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *IndexingDatasourcesItemsGetCall) Fields(s ...googleapi.Field) *IndexingDatasourcesItemsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *IndexingDatasourcesItemsGetCall) IfNoneMatch(entityTag string) *IndexingDatasourcesItemsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *IndexingDatasourcesItemsGetCall) Context(ctx context.Context) *IndexingDatasourcesItemsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *IndexingDatasourcesItemsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *IndexingDatasourcesItemsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/indexing/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.indexing.datasources.items.get" call.
// Exactly one of *Item or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Item.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *IndexingDatasourcesItemsGetCall) Do(opts ...googleapi.CallOption) (*Item, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Item{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets Item resource by item name.",
	//   "flatPath": "v1beta1/indexing/datasources/{datasourcesId}/items/{itemsId}",
	//   "httpMethod": "GET",
	//   "id": "cloudsearch.indexing.datasources.items.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name of the item to get info.\nFormat: datasources/{source_id}/items/{item_id}",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+/items/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/indexing/{+name}",
	//   "response": {
	//     "$ref": "Item"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.indexing"
	//   ]
	// }

}

// method id "cloudsearch.indexing.datasources.items.index":

type IndexingDatasourcesItemsIndexCall struct {
	s                *Service
	name             string
	indexitemrequest *IndexItemRequest
	urlParams_       gensupport.URLParams
	ctx_             context.Context
	header_          http.Header
}

// Index: Updates Item ACL, metadata, and
// content. It will insert the Item if it
// does not exist.
// This method does not support partial updates.  Fields with no
// provided
// values are cleared out in the Cloud Search index.
func (r *IndexingDatasourcesItemsService) Index(name string, indexitemrequest *IndexItemRequest) *IndexingDatasourcesItemsIndexCall {
	c := &IndexingDatasourcesItemsIndexCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.indexitemrequest = indexitemrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *IndexingDatasourcesItemsIndexCall) Fields(s ...googleapi.Field) *IndexingDatasourcesItemsIndexCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *IndexingDatasourcesItemsIndexCall) Context(ctx context.Context) *IndexingDatasourcesItemsIndexCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *IndexingDatasourcesItemsIndexCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *IndexingDatasourcesItemsIndexCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.indexitemrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/indexing/{+name}:index")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.indexing.datasources.items.index" call.
// Exactly one of *Item or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Item.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *IndexingDatasourcesItemsIndexCall) Do(opts ...googleapi.CallOption) (*Item, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Item{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates Item ACL, metadata, and\ncontent. It will insert the Item if it\ndoes not exist.\nThis method does not support partial updates.  Fields with no provided\nvalues are cleared out in the Cloud Search index.",
	//   "flatPath": "v1beta1/indexing/datasources/{datasourcesId}/items/{itemsId}:index",
	//   "httpMethod": "POST",
	//   "id": "cloudsearch.indexing.datasources.items.index",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name of the Item. Format:\ndatasources/{source_id}/items/{item_id}\n\u003cbr /\u003eThis is a required field.\nThe maximum length is 2048 characters.",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+/items/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/indexing/{+name}:index",
	//   "request": {
	//     "$ref": "IndexItemRequest"
	//   },
	//   "response": {
	//     "$ref": "Item"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.indexing"
	//   ]
	// }

}

// method id "cloudsearch.indexing.datasources.items.list":

type IndexingDatasourcesItemsListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists all or a subset of Item resources.
func (r *IndexingDatasourcesItemsService) List(name string) *IndexingDatasourcesItemsListCall {
	c := &IndexingDatasourcesItemsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Brief sets the optional parameter "brief": When set to true, the
// indexing system only populates the following
// fields:
// name,
// version,
// metadata.hash,
// structured_data.hash,
// content.ha
// sh.
// <br />If this value is false, then all the fields are populated in
// Item.
func (c *IndexingDatasourcesItemsListCall) Brief(brief bool) *IndexingDatasourcesItemsListCall {
	c.urlParams_.Set("brief", fmt.Sprint(brief))
	return c
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// items to fetch in a request.
// The max value is 10000 when brief is true.  The max value is 10 if
// brief
// is false.
func (c *IndexingDatasourcesItemsListCall) PageSize(pageSize int64) *IndexingDatasourcesItemsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous List request, if any.
func (c *IndexingDatasourcesItemsListCall) PageToken(pageToken string) *IndexingDatasourcesItemsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *IndexingDatasourcesItemsListCall) Fields(s ...googleapi.Field) *IndexingDatasourcesItemsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *IndexingDatasourcesItemsListCall) IfNoneMatch(entityTag string) *IndexingDatasourcesItemsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *IndexingDatasourcesItemsListCall) Context(ctx context.Context) *IndexingDatasourcesItemsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *IndexingDatasourcesItemsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *IndexingDatasourcesItemsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/indexing/{+name}/items")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.indexing.datasources.items.list" call.
// Exactly one of *ListItemsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListItemsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *IndexingDatasourcesItemsListCall) Do(opts ...googleapi.CallOption) (*ListItemsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListItemsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all or a subset of Item resources.",
	//   "flatPath": "v1beta1/indexing/datasources/{datasourcesId}/items",
	//   "httpMethod": "GET",
	//   "id": "cloudsearch.indexing.datasources.items.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "brief": {
	//       "description": "When set to true, the indexing system only populates the following fields:\nname,\nversion,\nmetadata.hash,\nstructured_data.hash,\ncontent.hash.\n\u003cbr /\u003eIf this value is false, then all the fields are populated in Item.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "name": {
	//       "description": "Name of the Data Source to list Items.  Format:\ndatasources/{source_id}",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Maximum number of items to fetch in a request.\nThe max value is 10000 when brief is true.  The max value is 10 if brief\nis false.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The next_page_token value returned from a previous List request, if any.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/indexing/{+name}/items",
	//   "response": {
	//     "$ref": "ListItemsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.indexing"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *IndexingDatasourcesItemsListCall) Pages(ctx context.Context, f func(*ListItemsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "cloudsearch.indexing.datasources.items.poll":

type IndexingDatasourcesItemsPollCall struct {
	s                *Service
	name             string
	pollitemsrequest *PollItemsRequest
	urlParams_       gensupport.URLParams
	ctx_             context.Context
	header_          http.Header
}

// Poll: Polls for unreserved items from the indexing queue and marks
// a
// set as reserved, starting with items that have
// the oldest timestamp from the highest priority
// ItemStatus.
// The priority order is as follows: <br />
// ERROR
// <br />
// MODIFIED
// <br />
// NEW_ITEM
// <br />
// ACCEPTED
// <br />
// Reserving items ensures that polling from other threads
// cannot create overlapping sets.
//
// After handling the reserved items, the client should put items
// back
// into the unreserved state, either by calling
// index,
// or by calling
// push with
// the type REQUEUE.
//
// Items automatically become available (unreserved) after 4 hours even
// if no
// update or push method is called.
func (r *IndexingDatasourcesItemsService) Poll(name string, pollitemsrequest *PollItemsRequest) *IndexingDatasourcesItemsPollCall {
	c := &IndexingDatasourcesItemsPollCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.pollitemsrequest = pollitemsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *IndexingDatasourcesItemsPollCall) Fields(s ...googleapi.Field) *IndexingDatasourcesItemsPollCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *IndexingDatasourcesItemsPollCall) Context(ctx context.Context) *IndexingDatasourcesItemsPollCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *IndexingDatasourcesItemsPollCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *IndexingDatasourcesItemsPollCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.pollitemsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/indexing/{+name}/items:poll")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.indexing.datasources.items.poll" call.
// Exactly one of *PollItemsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *PollItemsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *IndexingDatasourcesItemsPollCall) Do(opts ...googleapi.CallOption) (*PollItemsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &PollItemsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Polls for unreserved items from the indexing queue and marks a\nset as reserved, starting with items that have\nthe oldest timestamp from the highest priority\nItemStatus.\nThe priority order is as follows: \u003cbr /\u003e\nERROR\n\u003cbr /\u003e\nMODIFIED\n\u003cbr /\u003e\nNEW_ITEM\n\u003cbr /\u003e\nACCEPTED\n\u003cbr /\u003e\nReserving items ensures that polling from other threads\ncannot create overlapping sets.\n\nAfter handling the reserved items, the client should put items back\ninto the unreserved state, either by calling\nindex,\nor by calling\npush with\nthe type REQUEUE.\n\nItems automatically become available (unreserved) after 4 hours even if no\nupdate or push method is called.",
	//   "flatPath": "v1beta1/indexing/datasources/{datasourcesId}/items:poll",
	//   "httpMethod": "POST",
	//   "id": "cloudsearch.indexing.datasources.items.poll",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name of the Data Source to poll items.\nFormat: datasources/{source_id}",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/indexing/{+name}/items:poll",
	//   "request": {
	//     "$ref": "PollItemsRequest"
	//   },
	//   "response": {
	//     "$ref": "PollItemsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.indexing"
	//   ]
	// }

}

// method id "cloudsearch.indexing.datasources.items.push":

type IndexingDatasourcesItemsPushCall struct {
	s               *Service
	name            string
	pushitemrequest *PushItemRequest
	urlParams_      gensupport.URLParams
	ctx_            context.Context
	header_         http.Header
}

// Push: Pushes an items onto a queue for later polling and updating.
func (r *IndexingDatasourcesItemsService) Push(name string, pushitemrequest *PushItemRequest) *IndexingDatasourcesItemsPushCall {
	c := &IndexingDatasourcesItemsPushCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.pushitemrequest = pushitemrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *IndexingDatasourcesItemsPushCall) Fields(s ...googleapi.Field) *IndexingDatasourcesItemsPushCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *IndexingDatasourcesItemsPushCall) Context(ctx context.Context) *IndexingDatasourcesItemsPushCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *IndexingDatasourcesItemsPushCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *IndexingDatasourcesItemsPushCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.pushitemrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/indexing/{+name}:push")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.indexing.datasources.items.push" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *IndexingDatasourcesItemsPushCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Pushes an items onto a queue for later polling and updating.",
	//   "flatPath": "v1beta1/indexing/datasources/{datasourcesId}/items/{itemsId}:push",
	//   "httpMethod": "POST",
	//   "id": "cloudsearch.indexing.datasources.items.push",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name of the item to\npush into the indexing queue.\u003cbr /\u003e\nFormat: datasources/{source_id}/items/{ID}\n\u003cbr /\u003eThis is a required field.\nThe maximum length is 2048 characters.",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+/items/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/indexing/{+name}:push",
	//   "request": {
	//     "$ref": "PushItemRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.indexing"
	//   ]
	// }

}

// method id "cloudsearch.indexing.datasources.items.unreserve":

type IndexingDatasourcesItemsUnreserveCall struct {
	s                     *Service
	name                  string
	unreserveitemsrequest *UnreserveItemsRequest
	urlParams_            gensupport.URLParams
	ctx_                  context.Context
	header_               http.Header
}

// Unreserve: Unreserves all items from a queue, making them all
// eligible to be
// polled.  This method is useful for resetting the indexing queue
// after a connector has been restarted.
func (r *IndexingDatasourcesItemsService) Unreserve(name string, unreserveitemsrequest *UnreserveItemsRequest) *IndexingDatasourcesItemsUnreserveCall {
	c := &IndexingDatasourcesItemsUnreserveCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.unreserveitemsrequest = unreserveitemsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *IndexingDatasourcesItemsUnreserveCall) Fields(s ...googleapi.Field) *IndexingDatasourcesItemsUnreserveCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *IndexingDatasourcesItemsUnreserveCall) Context(ctx context.Context) *IndexingDatasourcesItemsUnreserveCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *IndexingDatasourcesItemsUnreserveCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *IndexingDatasourcesItemsUnreserveCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.unreserveitemsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/indexing/{+name}/items:unreserve")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.indexing.datasources.items.unreserve" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *IndexingDatasourcesItemsUnreserveCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Unreserves all items from a queue, making them all eligible to be\npolled.  This method is useful for resetting the indexing queue\nafter a connector has been restarted.",
	//   "flatPath": "v1beta1/indexing/datasources/{datasourcesId}/items:unreserve",
	//   "httpMethod": "POST",
	//   "id": "cloudsearch.indexing.datasources.items.unreserve",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name of the Data Source to unreserve all items.\nFormat: datasources/{source_id}",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/indexing/{+name}/items:unreserve",
	//   "request": {
	//     "$ref": "UnreserveItemsRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.indexing"
	//   ]
	// }

}

// method id "cloudsearch.indexing.datasources.items.upload":

type IndexingDatasourcesItemsUploadCall struct {
	s                      *Service
	name                   string
	startuploaditemrequest *StartUploadItemRequest
	urlParams_             gensupport.URLParams
	ctx_                   context.Context
	header_                http.Header
}

// Upload: Creates an upload session for uploading item content. For
// items smaller
// than 100 KiB, it's easier to embed the content
// inline within
// index.
func (r *IndexingDatasourcesItemsService) Upload(name string, startuploaditemrequest *StartUploadItemRequest) *IndexingDatasourcesItemsUploadCall {
	c := &IndexingDatasourcesItemsUploadCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.startuploaditemrequest = startuploaditemrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *IndexingDatasourcesItemsUploadCall) Fields(s ...googleapi.Field) *IndexingDatasourcesItemsUploadCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *IndexingDatasourcesItemsUploadCall) Context(ctx context.Context) *IndexingDatasourcesItemsUploadCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *IndexingDatasourcesItemsUploadCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *IndexingDatasourcesItemsUploadCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.startuploaditemrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/indexing/{+name}:upload")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.indexing.datasources.items.upload" call.
// Exactly one of *UploadItemRef or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *UploadItemRef.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *IndexingDatasourcesItemsUploadCall) Do(opts ...googleapi.CallOption) (*UploadItemRef, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &UploadItemRef{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates an upload session for uploading item content. For items smaller\nthan 100 KiB, it's easier to embed the content\ninline within\nindex.",
	//   "flatPath": "v1beta1/indexing/datasources/{datasourcesId}/items/{itemsId}:upload",
	//   "httpMethod": "POST",
	//   "id": "cloudsearch.indexing.datasources.items.upload",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name of the Data Source to start a resumable upload.\nFormat: datasources/{source_id}",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+/items/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/indexing/{+name}:upload",
	//   "request": {
	//     "$ref": "StartUploadItemRequest"
	//   },
	//   "response": {
	//     "$ref": "UploadItemRef"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.indexing"
	//   ]
	// }

}

// method id "cloudsearch.query.search":

type QuerySearchCall struct {
	s             *Service
	searchrequest *SearchRequest
	urlParams_    gensupport.URLParams
	ctx_          context.Context
	header_       http.Header
}

// Search: The Cloud Search Query API provides the search method, which
// returns
// the most relevant results from a user query.  The results can come
// from
// G Suite Apps, such as Gmail or Google Drive, or they can come from
// data
// that you have indexed from a third party.
func (r *QueryService) Search(searchrequest *SearchRequest) *QuerySearchCall {
	c := &QuerySearchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.searchrequest = searchrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *QuerySearchCall) Fields(s ...googleapi.Field) *QuerySearchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *QuerySearchCall) Context(ctx context.Context) *QuerySearchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *QuerySearchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *QuerySearchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.searchrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/query/search")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.query.search" call.
// Exactly one of *SearchResponse or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *SearchResponse.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *QuerySearchCall) Do(opts ...googleapi.CallOption) (*SearchResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &SearchResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "The Cloud Search Query API provides the search method, which returns\nthe most relevant results from a user query.  The results can come from\nG Suite Apps, such as Gmail or Google Drive, or they can come from data\nthat you have indexed from a third party.",
	//   "flatPath": "v1beta1/query/search",
	//   "httpMethod": "POST",
	//   "id": "cloudsearch.query.search",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1beta1/query/search",
	//   "request": {
	//     "$ref": "SearchRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.query"
	//   ]
	// }

}

// method id "cloudsearch.query.suggest":

type QuerySuggestCall struct {
	s              *Service
	suggestrequest *SuggestRequest
	urlParams_     gensupport.URLParams
	ctx_           context.Context
	header_        http.Header
}

// Suggest: The Cloud Search Query API provides suggested text
// query completions and people query suggestions.
func (r *QueryService) Suggest(suggestrequest *SuggestRequest) *QuerySuggestCall {
	c := &QuerySuggestCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.suggestrequest = suggestrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *QuerySuggestCall) Fields(s ...googleapi.Field) *QuerySuggestCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *QuerySuggestCall) Context(ctx context.Context) *QuerySuggestCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *QuerySuggestCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *QuerySuggestCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.suggestrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/query/suggest")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.query.suggest" call.
// Exactly one of *SuggestResponse or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *SuggestResponse.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *QuerySuggestCall) Do(opts ...googleapi.CallOption) (*SuggestResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &SuggestResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "The Cloud Search Query API provides suggested text\nquery completions and people query suggestions.",
	//   "flatPath": "v1beta1/query/suggest",
	//   "httpMethod": "POST",
	//   "id": "cloudsearch.query.suggest",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1beta1/query/suggest",
	//   "request": {
	//     "$ref": "SuggestRequest"
	//   },
	//   "response": {
	//     "$ref": "SuggestResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.query"
	//   ]
	// }

}

// method id "cloudsearch.settings.datasources.create":

type SettingsDatasourcesCreateCall struct {
	s          *Service
	datasource *DataSource
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Data source
// Create data source
func (r *SettingsDatasourcesService) Create(datasource *DataSource) *SettingsDatasourcesCreateCall {
	c := &SettingsDatasourcesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.datasource = datasource
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SettingsDatasourcesCreateCall) Fields(s ...googleapi.Field) *SettingsDatasourcesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SettingsDatasourcesCreateCall) Context(ctx context.Context) *SettingsDatasourcesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SettingsDatasourcesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SettingsDatasourcesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.datasource)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/settings/datasources")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.settings.datasources.create" call.
// Exactly one of *DataSource or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *DataSource.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *SettingsDatasourcesCreateCall) Do(opts ...googleapi.CallOption) (*DataSource, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &DataSource{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Data source\nCreate data source",
	//   "flatPath": "v1beta1/settings/datasources",
	//   "httpMethod": "POST",
	//   "id": "cloudsearch.settings.datasources.create",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1beta1/settings/datasources",
	//   "request": {
	//     "$ref": "DataSource"
	//   },
	//   "response": {
	//     "$ref": "DataSource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.settings",
	//     "https://www.googleapis.com/auth/cloud_search.settings.indexing"
	//   ]
	// }

}

// method id "cloudsearch.settings.datasources.delete":

type SettingsDatasourcesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes a data source
func (r *SettingsDatasourcesService) Delete(name string) *SettingsDatasourcesDeleteCall {
	c := &SettingsDatasourcesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SettingsDatasourcesDeleteCall) Fields(s ...googleapi.Field) *SettingsDatasourcesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SettingsDatasourcesDeleteCall) Context(ctx context.Context) *SettingsDatasourcesDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SettingsDatasourcesDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SettingsDatasourcesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/settings/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.settings.datasources.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *SettingsDatasourcesDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a data source",
	//   "flatPath": "v1beta1/settings/datasources/{datasourcesId}",
	//   "httpMethod": "DELETE",
	//   "id": "cloudsearch.settings.datasources.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name of the data source.\nFormat: datasources/{source_id}.",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/settings/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.settings",
	//     "https://www.googleapis.com/auth/cloud_search.settings.indexing"
	//   ]
	// }

}

// method id "cloudsearch.settings.datasources.get":

type SettingsDatasourcesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets data source
func (r *SettingsDatasourcesService) Get(name string) *SettingsDatasourcesGetCall {
	c := &SettingsDatasourcesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SettingsDatasourcesGetCall) Fields(s ...googleapi.Field) *SettingsDatasourcesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *SettingsDatasourcesGetCall) IfNoneMatch(entityTag string) *SettingsDatasourcesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SettingsDatasourcesGetCall) Context(ctx context.Context) *SettingsDatasourcesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SettingsDatasourcesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SettingsDatasourcesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/settings/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.settings.datasources.get" call.
// Exactly one of *DataSource or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *DataSource.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *SettingsDatasourcesGetCall) Do(opts ...googleapi.CallOption) (*DataSource, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &DataSource{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets data source",
	//   "flatPath": "v1beta1/settings/datasources/{datasourcesId}",
	//   "httpMethod": "GET",
	//   "id": "cloudsearch.settings.datasources.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name of the data source resource.\nFormat: datasources/{source_id}.",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/settings/{+name}",
	//   "response": {
	//     "$ref": "DataSource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.settings",
	//     "https://www.googleapis.com/auth/cloud_search.settings.indexing"
	//   ]
	// }

}

// method id "cloudsearch.settings.datasources.list":

type SettingsDatasourcesListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists data source
func (r *SettingsDatasourcesService) List() *SettingsDatasourcesListCall {
	c := &SettingsDatasourcesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// PageToken sets the optional parameter "pageToken": Starting index of
// the results.
func (c *SettingsDatasourcesListCall) PageToken(pageToken string) *SettingsDatasourcesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SettingsDatasourcesListCall) Fields(s ...googleapi.Field) *SettingsDatasourcesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *SettingsDatasourcesListCall) IfNoneMatch(entityTag string) *SettingsDatasourcesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SettingsDatasourcesListCall) Context(ctx context.Context) *SettingsDatasourcesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SettingsDatasourcesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SettingsDatasourcesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/settings/datasources")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.settings.datasources.list" call.
// Exactly one of *ListDataSourceResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListDataSourceResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *SettingsDatasourcesListCall) Do(opts ...googleapi.CallOption) (*ListDataSourceResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListDataSourceResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists data source",
	//   "flatPath": "v1beta1/settings/datasources",
	//   "httpMethod": "GET",
	//   "id": "cloudsearch.settings.datasources.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "pageToken": {
	//       "description": "Starting index of the results.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/settings/datasources",
	//   "response": {
	//     "$ref": "ListDataSourceResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.settings",
	//     "https://www.googleapis.com/auth/cloud_search.settings.indexing"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *SettingsDatasourcesListCall) Pages(ctx context.Context, f func(*ListDataSourceResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "cloudsearch.settings.datasources.update":

type SettingsDatasourcesUpdateCall struct {
	s                       *Service
	name                    string
	updatedatasourcerequest *UpdateDataSourceRequest
	urlParams_              gensupport.URLParams
	ctx_                    context.Context
	header_                 http.Header
}

// Update: Updates data source
func (r *SettingsDatasourcesService) Update(name string, updatedatasourcerequest *UpdateDataSourceRequest) *SettingsDatasourcesUpdateCall {
	c := &SettingsDatasourcesUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.updatedatasourcerequest = updatedatasourcerequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SettingsDatasourcesUpdateCall) Fields(s ...googleapi.Field) *SettingsDatasourcesUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SettingsDatasourcesUpdateCall) Context(ctx context.Context) *SettingsDatasourcesUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SettingsDatasourcesUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SettingsDatasourcesUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.updatedatasourcerequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/settings/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.settings.datasources.update" call.
// Exactly one of *DataSource or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *DataSource.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *SettingsDatasourcesUpdateCall) Do(opts ...googleapi.CallOption) (*DataSource, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &DataSource{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates data source",
	//   "flatPath": "v1beta1/settings/datasources/{datasourcesId}",
	//   "httpMethod": "PUT",
	//   "id": "cloudsearch.settings.datasources.update",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name of the data source resource.\nFormat: datasources/{source_id}.\n\u003cbr /\u003eThe name is ignored when creating a data source.",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/settings/{+name}",
	//   "request": {
	//     "$ref": "UpdateDataSourceRequest"
	//   },
	//   "response": {
	//     "$ref": "DataSource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.settings",
	//     "https://www.googleapis.com/auth/cloud_search.settings.indexing"
	//   ]
	// }

}
