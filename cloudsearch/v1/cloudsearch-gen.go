// Package cloudsearch provides access to the Cloud Search API.
//
// See https://gsuite.google.com/products/cloud-search/
//
// Usage example:
//
//   import "github.com/nilleb/google-api-go-client/cloudsearch/v1"
//   ...
//   cloudsearchService, err := cloudsearch.New(oauthHttpClient)
package cloudsearch // import "github.com/nilleb/google-api-go-client/cloudsearch/v1"

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

const apiId = "cloudsearch:v1"
const apiName = "cloudsearch"
const apiVersion = "v1"
const basePath = "https://cloudsearch.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// Index and serve your organization's data with Cloud Search
	CloudSearchScope = "https://www.googleapis.com/auth/cloud_search"

	// New Service: https://www.googleapis.com/auth/cloud_search.debug
	CloudSearchDebugScope = "https://www.googleapis.com/auth/cloud_search.debug"

	// New Service: https://www.googleapis.com/auth/cloud_search.indexing
	CloudSearchIndexingScope = "https://www.googleapis.com/auth/cloud_search.indexing"

	// Search your organization's data in the Cloud Search index
	CloudSearchQueryScope = "https://www.googleapis.com/auth/cloud_search.query"

	// New Service: https://www.googleapis.com/auth/cloud_search.settings
	CloudSearchSettingsScope = "https://www.googleapis.com/auth/cloud_search.settings"

	// New Service:
	// https://www.googleapis.com/auth/cloud_search.settings.indexing
	CloudSearchSettingsIndexingScope = "https://www.googleapis.com/auth/cloud_search.settings.indexing"

	// New Service:
	// https://www.googleapis.com/auth/cloud_search.settings.query
	CloudSearchSettingsQueryScope = "https://www.googleapis.com/auth/cloud_search.settings.query"

	// New Service: https://www.googleapis.com/auth/cloud_search.stats
	CloudSearchStatsScope = "https://www.googleapis.com/auth/cloud_search.stats"

	// New Service:
	// https://www.googleapis.com/auth/cloud_search.stats.indexing
	CloudSearchStatsIndexingScope = "https://www.googleapis.com/auth/cloud_search.stats.indexing"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Debug = NewDebugService(s)
	s.Indexing = NewIndexingService(s)
	s.Media = NewMediaService(s)
	s.Operations = NewOperationsService(s)
	s.Query = NewQueryService(s)
	s.Settings = NewSettingsService(s)
	s.Stats = NewStatsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Debug *DebugService

	Indexing *IndexingService

	Media *MediaService

	Operations *OperationsService

	Query *QueryService

	Settings *SettingsService

	Stats *StatsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewDebugService(s *Service) *DebugService {
	rs := &DebugService{s: s}
	rs.Datasources = NewDebugDatasourcesService(s)
	rs.Identitysources = NewDebugIdentitysourcesService(s)
	return rs
}

type DebugService struct {
	s *Service

	Datasources *DebugDatasourcesService

	Identitysources *DebugIdentitysourcesService
}

func NewDebugDatasourcesService(s *Service) *DebugDatasourcesService {
	rs := &DebugDatasourcesService{s: s}
	rs.Items = NewDebugDatasourcesItemsService(s)
	return rs
}

type DebugDatasourcesService struct {
	s *Service

	Items *DebugDatasourcesItemsService
}

func NewDebugDatasourcesItemsService(s *Service) *DebugDatasourcesItemsService {
	rs := &DebugDatasourcesItemsService{s: s}
	rs.Unmappedids = NewDebugDatasourcesItemsUnmappedidsService(s)
	return rs
}

type DebugDatasourcesItemsService struct {
	s *Service

	Unmappedids *DebugDatasourcesItemsUnmappedidsService
}

func NewDebugDatasourcesItemsUnmappedidsService(s *Service) *DebugDatasourcesItemsUnmappedidsService {
	rs := &DebugDatasourcesItemsUnmappedidsService{s: s}
	return rs
}

type DebugDatasourcesItemsUnmappedidsService struct {
	s *Service
}

func NewDebugIdentitysourcesService(s *Service) *DebugIdentitysourcesService {
	rs := &DebugIdentitysourcesService{s: s}
	rs.Items = NewDebugIdentitysourcesItemsService(s)
	rs.Unmappedids = NewDebugIdentitysourcesUnmappedidsService(s)
	return rs
}

type DebugIdentitysourcesService struct {
	s *Service

	Items *DebugIdentitysourcesItemsService

	Unmappedids *DebugIdentitysourcesUnmappedidsService
}

func NewDebugIdentitysourcesItemsService(s *Service) *DebugIdentitysourcesItemsService {
	rs := &DebugIdentitysourcesItemsService{s: s}
	return rs
}

type DebugIdentitysourcesItemsService struct {
	s *Service
}

func NewDebugIdentitysourcesUnmappedidsService(s *Service) *DebugIdentitysourcesUnmappedidsService {
	rs := &DebugIdentitysourcesUnmappedidsService{s: s}
	return rs
}

type DebugIdentitysourcesUnmappedidsService struct {
	s *Service
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

func NewMediaService(s *Service) *MediaService {
	rs := &MediaService{s: s}
	return rs
}

type MediaService struct {
	s *Service
}

func NewOperationsService(s *Service) *OperationsService {
	rs := &OperationsService{s: s}
	return rs
}

type OperationsService struct {
	s *Service
}

func NewQueryService(s *Service) *QueryService {
	rs := &QueryService{s: s}
	rs.Sources = NewQuerySourcesService(s)
	return rs
}

type QueryService struct {
	s *Service

	Sources *QuerySourcesService
}

func NewQuerySourcesService(s *Service) *QuerySourcesService {
	rs := &QuerySourcesService{s: s}
	return rs
}

type QuerySourcesService struct {
	s *Service
}

func NewSettingsService(s *Service) *SettingsService {
	rs := &SettingsService{s: s}
	rs.Datasources = NewSettingsDatasourcesService(s)
	rs.Searchapplications = NewSettingsSearchapplicationsService(s)
	return rs
}

type SettingsService struct {
	s *Service

	Datasources *SettingsDatasourcesService

	Searchapplications *SettingsSearchapplicationsService
}

func NewSettingsDatasourcesService(s *Service) *SettingsDatasourcesService {
	rs := &SettingsDatasourcesService{s: s}
	return rs
}

type SettingsDatasourcesService struct {
	s *Service
}

func NewSettingsSearchapplicationsService(s *Service) *SettingsSearchapplicationsService {
	rs := &SettingsSearchapplicationsService{s: s}
	return rs
}

type SettingsSearchapplicationsService struct {
	s *Service
}

func NewStatsService(s *Service) *StatsService {
	rs := &StatsService{s: s}
	rs.Index = NewStatsIndexService(s)
	return rs
}

type StatsService struct {
	s *Service

	Index *StatsIndexService
}

func NewStatsIndexService(s *Service) *StatsIndexService {
	rs := &StatsIndexService{s: s}
	rs.Datasources = NewStatsIndexDatasourcesService(s)
	return rs
}

type StatsIndexService struct {
	s *Service

	Datasources *StatsIndexDatasourcesService
}

func NewStatsIndexDatasourcesService(s *Service) *StatsIndexDatasourcesService {
	rs := &StatsIndexDatasourcesService{s: s}
	return rs
}

type StatsIndexDatasourcesService struct {
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
	// all items where *&lt;value&gt;* matches the value of any
	// String properties or text within the content field for the item.
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

type CheckAccessResponse struct {
	// HasAccess: Returns true if principal has access.  Returns false
	// otherwise.
	HasAccess bool `json:"hasAccess,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "HasAccess") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "HasAccess") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CheckAccessResponse) MarshalJSON() ([]byte, error) {
	type NoMethod CheckAccessResponse
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

// CustomerIndexStats: Aggregation of items by status code as of the
// specified date.
type CustomerIndexStats struct {
	// Date: Date for which statistics were calculated.
	Date *Date `json:"date,omitempty"`

	// ItemCountByStatus: Number of items aggregrated by status code.
	ItemCountByStatus []*ItemCountByStatus `json:"itemCountByStatus,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Date") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Date") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CustomerIndexStats) MarshalJSON() ([]byte, error) {
	type NoMethod CustomerIndexStats
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

	// DisableServing: Disable serving any search or assist results.
	DisableServing bool `json:"disableServing,omitempty"`

	// DisplayName: Required. Display name of the data source
	// The maximum length is 300 characters.
	DisplayName string `json:"displayName,omitempty"`

	// IndexingServiceAccounts: List of service accounts that have indexing
	// access.
	IndexingServiceAccounts []string `json:"indexingServiceAccounts,omitempty"`

	// ItemsVisibility: This restricts visibility to items at a data source
	// level to the
	// disjunction of users/groups mentioned with the field. Note that,
	// this
	// does not ensure access to a specific item, as users need to have
	// ACL
	// permissions on the contained items. This ensures a high level
	// access
	// on the entire data source, and that the individual items are not
	// shared
	// outside this visibility.
	ItemsVisibility []*GSuitePrincipal `json:"itemsVisibility,omitempty"`

	// Name: Name of the data source resource.
	// Format: datasources/{source_id}.
	// <br />The name is ignored when creating a data source.
	Name string `json:"name,omitempty"`

	// OperationIds: IDs of the Long Running Operations (LROs) currently
	// running for this schema.
	OperationIds []string `json:"operationIds,omitempty"`

	// ShortName: A short name or alias for the source.  This value will be
	// used to match the
	// 'source' operator. For example, if the short name is *&lt;value&gt;*
	// then
	// queries like *source:&lt;value&gt;* will only return results for
	// this
	// source. The value must be unique across all data sources. The value
	// must
	// only contain alphanumeric characters (a-zA-Z0-9). The value cannot
	// start
	// with 'google' and cannot be one of the following: mail, gmail, docs,
	// drive,
	// groups, sites, calendar, hangouts, gplus, keep.
	// Its maximum length is 32 characters.
	ShortName string `json:"shortName,omitempty"`

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

// DataSourceIndexStats: Aggregation of items by status code as of the
// specified date.
type DataSourceIndexStats struct {
	// Date: Date for which index stats were calculated. If the date of
	// request is not
	// the current date then stats calculated on the next day are returned.
	// Stats
	// are calculated close to mid night in this case. If date of request
	// is
	// current date, then real time stats are returned.
	Date *Date `json:"date,omitempty"`

	// ItemCountByStatus: Number of items aggregrated by status code.
	ItemCountByStatus []*ItemCountByStatus `json:"itemCountByStatus,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Date") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Date") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DataSourceIndexStats) MarshalJSON() ([]byte, error) {
	type NoMethod DataSourceIndexStats
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DataSourceRestriction: Restriction on Datasource.
type DataSourceRestriction struct {
	// FilterOptions: Filter options restricting the results. If multiple
	// filters
	// are present, they are grouped by object type before joining.
	// Filters with the same object type are joined conjunctively, then
	// the resulting expressions are joined disjunctively.
	//
	// The maximum number of elements is 20.
	FilterOptions []*FilterOptions `json:"filterOptions,omitempty"`

	// Source: The source of restriction.
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

func (s *DataSourceRestriction) MarshalJSON() ([]byte, error) {
	type NoMethod DataSourceRestriction
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

	// Month: Month of date. Must be from 1 to 12.
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

// DebugOptions: Shared request debug options for all cloudsearch RPC
// methods.
type DebugOptions struct {
	// EnableDebugging: If set, the request will enable debugging features
	// of Cloud Search.
	// Only turn on this field, if asked by Google to help with debugging.
	EnableDebugging bool `json:"enableDebugging,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EnableDebugging") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EnableDebugging") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *DebugOptions) MarshalJSON() ([]byte, error) {
	type NoMethod DebugOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type DeleteQueueItemsRequest struct {
	// ConnectorName: Name of connector making this call.
	// <br />Format: datasources/{source_id}/connectors/{ID}
	ConnectorName string `json:"connectorName,omitempty"`

	// DebugOptions: Common debug options.
	DebugOptions *DebugOptions `json:"debugOptions,omitempty"`

	// Queue: Name of a queue to delete items from.
	Queue string `json:"queue,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ConnectorName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ConnectorName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DeleteQueueItemsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod DeleteQueueItemsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DisplayedProperty: A reference to a top-level property within the
// object that should be
// displayed in search results. The values of the chosen properties will
// be
// displayed in the search results along with the
// dislpay label
// for that property if one is specified. If a display label is not
// specified,
// only the values will be shown.
type DisplayedProperty struct {
	// PropertyName: The name of the top-level property as defined in a
	// property definition
	// for the object. If the name is not a defined property in the schema,
	// an
	// error will be given when attempting to update the schema.
	PropertyName string `json:"propertyName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "PropertyName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "PropertyName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DisplayedProperty) MarshalJSON() ([]byte, error) {
	type NoMethod DisplayedProperty
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DoubleOperatorOptions: Used to provide a search operator for double
// properties. This is
// optional. Search operators let users restrict the query to specific
// fields
// relevant to the type of item being searched.
type DoubleOperatorOptions struct {
	// OperatorName: Indicates the operator name required in the query in
	// order to use the
	// double property in sorting or as a facet.
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

func (s *DoubleOperatorOptions) MarshalJSON() ([]byte, error) {
	type NoMethod DoubleOperatorOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DoublePropertyOptions: Options for double properties.
type DoublePropertyOptions struct {
	// OperatorOptions: If set, describes how the double should be used as a
	// search operator.
	OperatorOptions *DoubleOperatorOptions `json:"operatorOptions,omitempty"`

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

func (s *DoublePropertyOptions) MarshalJSON() ([]byte, error) {
	type NoMethod DoublePropertyOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
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
	//   "SITE"
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
// ordered
// ranking to
// set the ranking of a given value relative to other enumerated values
// for
// the same property name. Here, a ranking order of DESCENDING for
// *priority*
// properties results in a ranking boost for items indexed with a value
// of
// *p0* compared to items indexed with a value of *p1*. Without a
// specified
// ranking order, the integer value has no effect on item ranking.
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

// ErrorInfo: Error information about the response.
type ErrorInfo struct {
	ErrorMessages []*ErrorMessage `json:"errorMessages,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ErrorMessages") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ErrorMessages") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ErrorInfo) MarshalJSON() ([]byte, error) {
	type NoMethod ErrorInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ErrorMessage: Error message per source response.
type ErrorMessage struct {
	ErrorMessage string `json:"errorMessage,omitempty"`

	Source *Source `json:"source,omitempty"`

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

func (s *ErrorMessage) MarshalJSON() ([]byte, error) {
	type NoMethod ErrorMessage
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
	// Count: Number of results that match the bucket value.
	Count int64 `json:"count,omitempty"`

	// Percentage: Percent of results that match the bucket value. This
	// value is between
	// (0-100].
	// This may not be accurate and is a best effort estimate.
	Percentage int64 `json:"percentage,omitempty"`

	Value *Value `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Count") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Count") to include in API
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

// FacetOptions: Specifies operators to return facet results for. There
// will be one
// FacetResult for every source_name/object_type/operator_name
// combination.
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

	// SourceName: Source name to facet on. Format:
	// datasources/{source_id}
	// If empty, all data sources will be used.
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
	// Buckets: FacetBuckets for values in response containing atleast a
	// single result.
	Buckets []*FacetBucket `json:"buckets,omitempty"`

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

	// ForceSendFields is a list of field names (e.g. "Buckets") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Buckets") to include in
	// API requests with the JSON null value. By default, fields with empty
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
	// FreshnessDuration: The duration (in seconds) after which an object
	// should be considered
	// stale.
	FreshnessDuration string `json:"freshnessDuration,omitempty"`

	// FreshnessProperty: This property indicates the freshness level of the
	// object in the index.
	// If set, this property must be a top-level property within
	// the
	// property definitions
	// and it must be a
	// timestamp type
	// or
	// date type.
	// Otherwise, the Indexing API uses
	// updateTime
	// as the freshness indicator.
	// The maximum length is 256 characters.
	FreshnessProperty string `json:"freshnessProperty,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FreshnessDuration")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FreshnessDuration") to
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

type GetCustomerIndexStatsResponse struct {
	// Stats: Summary of indexed item counts, one for each day in the
	// requested range.
	Stats []*CustomerIndexStats `json:"stats,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Stats") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Stats") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GetCustomerIndexStatsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GetCustomerIndexStatsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type GetDataSourceIndexStatsResponse struct {
	// Stats: Summary of indexed item counts, one for each day in the
	// requested range.
	Stats []*DataSourceIndexStats `json:"stats,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Stats") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Stats") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GetDataSourceIndexStatsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GetDataSourceIndexStatsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// HtmlOperatorOptions: Used to provide a search operator for html
// properties. This is optional.
// Search operators let users restrict the query to specific fields
// relevant
// to the type of item being searched.
type HtmlOperatorOptions struct {
	// OperatorName: Indicates the operator name required in the query in
	// order to isolate the
	// html property. For example, if operatorName is *subject* and
	// the
	// property's name is *subjectLine*, then queries
	// like
	// *subject:&lt;value&gt;* will show results only where the value of
	// the
	// property named *subjectLine* matches *&lt;value&gt;*. By contrast,
	// a
	// search that uses the same *&lt;value&gt;* without an operator will
	// return
	// all items where *&lt;value&gt;* matches the value of any
	// html properties or text within the content field for the item.
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

func (s *HtmlOperatorOptions) MarshalJSON() ([]byte, error) {
	type NoMethod HtmlOperatorOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// HtmlPropertyOptions: Options for html properties.
type HtmlPropertyOptions struct {
	// OperatorOptions: If set, describes how the property should be used as
	// a search operator.
	OperatorOptions *HtmlOperatorOptions `json:"operatorOptions,omitempty"`

	// RetrievalImportance: Indicates the search quality importance of the
	// tokens within the
	// field when used for retrieval. Can only be set to DEFAULT or NONE.
	RetrievalImportance *RetrievalImportance `json:"retrievalImportance,omitempty"`

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

func (s *HtmlPropertyOptions) MarshalJSON() ([]byte, error) {
	type NoMethod HtmlPropertyOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// HtmlValues: List of html values.
type HtmlValues struct {
	// Values: The maximum allowable length for html values is 2048
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

func (s *HtmlValues) MarshalJSON() ([]byte, error) {
	type NoMethod HtmlValues
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type IndexItemRequest struct {
	// ConnectorName: Name of connector making this call.
	// <br />Format: datasources/{source_id}/connectors/{ID}
	ConnectorName string `json:"connectorName,omitempty"`

	// DebugOptions: Common debug options.
	DebugOptions *DebugOptions `json:"debugOptions,omitempty"`

	// Item: Name of the item.
	// Format:
	// datasources/{source_id}/items/{item_id}
	Item *Item `json:"item,omitempty"`

	// Mode: Required. The RequestMode for this request.
	//
	// Possible values:
	//   "UNSPECIFIED" - Priority is not specified in the update
	// request.
	// Leaving priority unspecified results in an update failure.
	//   "SYNCHRONOUS" - For real-time updates.
	//   "ASYNCHRONOUS" - For changes that are executed after the response
	// is sent back to the
	// caller.
	Mode string `json:"mode,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ConnectorName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ConnectorName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
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
	// all items where *&lt;value&gt;* matches the value of any
	// String
	// properties or text within the content field for the item.
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
	// The maximum length is 1536 characters.
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

// ItemAcl: Access control list information for the item. For more
// information
// see
// https://developers.google.com/cloud-search/docs/guides/index-your-
// data#acls
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
	// by its own readers and
	// deniedReaders fields.
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

	// DeniedReaders: List of principals who are explicitly denied access to
	// the item in search
	// results. While principals are denied access by default, use denied
	// readers
	// to handle exceptions and override the list allowed readers.
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
	// The maximum length for this field is 1536 characters.
	InheritAclFrom string `json:"inheritAclFrom,omitempty"`

	// Owners: Optional. List of owners for the item. This field has no
	// bearing on
	// document access permissions. It does, however, offer
	// a slight ranking boosts items where the querying user is an
	// owner.
	// The maximum number of elements is 5.
	Owners []*Principal `json:"owners,omitempty"`

	// Readers: List of principals who are allowed to see the item in search
	// results.
	// Optional if inheriting permissions from another item or if the
	// item
	// is not intended to be visible, such as
	// virtual containers.
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

type ItemCountByStatus struct {
	// Count: Number of items matching the status code.
	Count int64 `json:"count,omitempty,string"`

	// StatusCode: Status of the items.
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
	StatusCode string `json:"statusCode,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Count") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Count") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ItemCountByStatus) MarshalJSON() ([]byte, error) {
	type NoMethod ItemCountByStatus
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
	// field. The maximum length is 1536 characters.
	ContainerName string `json:"containerName,omitempty"`

	// ContentLanguage: The BCP-47 language code for the item, such as
	// "en-US" or "sr-Latn". For
	// more information,
	// see
	// http://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	// Th
	// e maximum length is 32 characters.
	ContentLanguage string `json:"contentLanguage,omitempty"`

	// CreateTime: The time when the item was created in the source
	// repository.
	CreateTime string `json:"createTime,omitempty"`

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

	// Keywords: Additional keywords or phrases that should match the
	// item.
	// Used internally for user generated content.
	// The maximum number of elements is 100.
	// The maximum length is 8192 characters.
	Keywords []string `json:"keywords,omitempty"`

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

	// SearchQualityMetadata: Additional search quality metadata of the item
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

	// UpdateTime: The time when the item was last modified in the source
	// repository.
	UpdateTime string `json:"updateTime,omitempty"`

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

	// ProcessingErrors: Error details in case the item is in ERROR state.
	ProcessingErrors []*ProcessingError `json:"processingErrors,omitempty"`

	// RepositoryErrors: Repository error reported by connector.
	RepositoryErrors []*RepositoryError `json:"repositoryErrors,omitempty"`

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

type ListItemNamesForUnmappedIdentityResponse struct {
	ItemNames []string `json:"itemNames,omitempty"`

	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ItemNames") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ItemNames") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListItemNamesForUnmappedIdentityResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListItemNamesForUnmappedIdentityResponse
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

// ListQuerySourcesResponse: List sources response.
type ListQuerySourcesResponse struct {
	NextPageToken string `json:"nextPageToken,omitempty"`

	Sources []*QuerySource `json:"sources,omitempty"`

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

func (s *ListQuerySourcesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListQuerySourcesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ListSearchApplicationsResponse struct {
	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	SearchApplications []*SearchApplication `json:"searchApplications,omitempty"`

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

func (s *ListSearchApplicationsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListSearchApplicationsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ListUnmappedIdentitiesResponse struct {
	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	UnmappedIdentities []*UnmappedIdentity `json:"unmappedIdentities,omitempty"`

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

func (s *ListUnmappedIdentitiesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListUnmappedIdentitiesResponse
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

// Media: Media resource.
type Media struct {
	// ResourceName: Name of the media resource.
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

func (s *Media) MarshalJSON() ([]byte, error) {
	type NoMethod Media
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Metadata: Metadata of a matched search result.
type Metadata struct {
	// CreateTime: The creation time for this document or object in the
	// search result.
	CreateTime string `json:"createTime,omitempty"`

	// DisplayOptions: Options that specify how to display a structured data
	// search result.
	DisplayOptions *ResultDisplayMetadata `json:"displayOptions,omitempty"`

	// Fields: Indexed fields in structured data, returned as a generic
	// named property.
	Fields []*NamedProperty `json:"fields,omitempty"`

	// MimeType: Mime type of the search result.
	MimeType string `json:"mimeType,omitempty"`

	// ObjectType: Object type of the search result.
	ObjectType string `json:"objectType,omitempty"`

	// Owner: Owner (usually creator) of the document or object of the
	// search result.
	Owner *Person `json:"owner,omitempty"`

	// Source: The named source for the result, such as Gmail.
	Source *Source `json:"source,omitempty"`

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

// Metaline: A metaline is a list of properties that are displayed along
// with the search
// result to provide context.
type Metaline struct {
	// Properties: The list of displayed properties for the metaline.
	Properties []*DisplayedProperty `json:"properties,omitempty"`

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

func (s *Metaline) MarshalJSON() ([]byte, error) {
	type NoMethod Metaline
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
	BooleanValue bool `json:"booleanValue,omitempty"`

	DateValues *DateValues `json:"dateValues,omitempty"`

	DoubleValues *DoubleValues `json:"doubleValues,omitempty"`

	EnumValues *EnumValues `json:"enumValues,omitempty"`

	HtmlValues *HtmlValues `json:"htmlValues,omitempty"`

	IntegerValues *IntegerValues `json:"integerValues,omitempty"`

	// Name: The name of the property.  This name should correspond to the
	// name of the
	// property that was registered for object definition in the schema.
	// The maximum allowable length for this property is 256 characters.
	Name string `json:"name,omitempty"`

	ObjectValues *ObjectValues `json:"objectValues,omitempty"`

	TextValues *TextValues `json:"textValues,omitempty"`

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

// ObjectDisplayOptions: The display options for an object.
type ObjectDisplayOptions struct {
	// Metalines: Defines the properties that will be displayed in the
	// metalines of the
	// search results. The property values will be displayed in the order
	// given
	// here. If a property holds multiple values, all of the values will
	// be
	// diplayed before the next properties. For this reason, it is a good
	// practice
	// to specify singular properties before repeated properties in this
	// list. All
	// of the properties must set
	// is_returnable
	// to true. The maximum number of elements is 3.
	Metalines []*Metaline `json:"metalines,omitempty"`

	// ObjectDisplayLabel: The user friendly label to display in the search
	// result to inidicate the
	// type of the item. This is OPTIONAL; if not given, an object label
	// will not
	// be displayed on the context line of the search results. The maximum
	// length
	// is 32 characters.
	ObjectDisplayLabel string `json:"objectDisplayLabel,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Metalines") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Metalines") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ObjectDisplayOptions) MarshalJSON() ([]byte, error) {
	type NoMethod ObjectDisplayOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ObjectOptions: The options for an object.
type ObjectOptions struct {
	// DisplayOptions: Options that determine how the object is displayed in
	// the Cloud Search
	// results page.
	DisplayOptions *ObjectDisplayOptions `json:"displayOptions,omitempty"`

	// FreshnessOptions: The freshness options for an object.
	FreshnessOptions *FreshnessOptions `json:"freshnessOptions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DisplayOptions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DisplayOptions") to
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

// Operation: This resource represents a long-running operation that is
// the result of a
// network API call.
type Operation struct {
	// Done: If the value is `false`, it means the operation is still in
	// progress.
	// If `true`, the operation is completed, and either `error` or
	// `response` is
	// available.
	Done bool `json:"done,omitempty"`

	// Error: The error result of the operation in case of failure or
	// cancellation.
	Error *Status `json:"error,omitempty"`

	// Metadata: Service-specific metadata associated with the operation.
	// It typically
	// contains progress information and common metadata such as create
	// time.
	// Some services might not provide such metadata.  Any method that
	// returns a
	// long-running operation should document the metadata type, if any.
	Metadata googleapi.RawMessage `json:"metadata,omitempty"`

	// Name: The server-assigned name, which is only unique within the same
	// service that
	// originally returns it. If you use the default HTTP mapping,
	// the
	// `name` should have the format of `operations/some/unique/name`.
	Name string `json:"name,omitempty"`

	// Response: The normal response of the operation in case of success.
	// If the original
	// method returns no data on success, such as `Delete`, the response
	// is
	// `google.protobuf.Empty`.  If the original method is
	// standard
	// `Get`/`Create`/`Update`, the response should be the resource.  For
	// other
	// methods, the response should have the type `XxxResponse`, where
	// `Xxx`
	// is the original method name.  For example, if the original method
	// name
	// is `TakeSnapshot()`, the inferred response type
	// is
	// `TakeSnapshotResponse`.
	Response googleapi.RawMessage `json:"response,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Done") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Done") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Operation) MarshalJSON() ([]byte, error) {
	type NoMethod Operation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PeopleSuggestion: A people suggestion.
type PeopleSuggestion struct {
	// Person: Suggested person. All fields of the person object might not
	// be populated.
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

	// ObfuscatedId: Obfuscated ID of a person.
	ObfuscatedId string `json:"obfuscatedId,omitempty"`

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
	// ConnectorName: Name of connector making this call.
	// <br />Format: datasources/{source_id}/connectors/{ID}
	ConnectorName string `json:"connectorName,omitempty"`

	// DebugOptions: Common debug options.
	DebugOptions *DebugOptions `json:"debugOptions,omitempty"`

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

	// ForceSendFields is a list of field names (e.g. "ConnectorName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ConnectorName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
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

	// DisplayOptions: Options that determine how the property is displayed
	// in the Cloud Search
	// results page if it is specified to be displayed in the
	// object's
	// display options
	// .
	DisplayOptions *PropertyDisplayOptions `json:"displayOptions,omitempty"`

	DoublePropertyOptions *DoublePropertyOptions `json:"doublePropertyOptions,omitempty"`

	EnumPropertyOptions *EnumPropertyOptions `json:"enumPropertyOptions,omitempty"`

	HtmlPropertyOptions *HtmlPropertyOptions `json:"htmlPropertyOptions,omitempty"`

	IntegerPropertyOptions *IntegerPropertyOptions `json:"integerPropertyOptions,omitempty"`

	// IsFacetable: Indicates that the property can be used for generating
	// facets. Cannot be
	// true for properties whose type is object. IsReturnable must be true
	// to set
	// this option.
	// Only supported for Boolean, Enum, and Text properties.
	IsFacetable bool `json:"isFacetable,omitempty"`

	// IsRepeatable: Indicates that multiple values are allowed for the
	// property. For example, a
	// document only has one description but can have multiple comments.
	// Cannot be
	// true for properties whose type is a boolean.
	// If set to false, properties that contain more than one value will
	// cause the
	// indexing request for that item to be rejected.
	IsRepeatable bool `json:"isRepeatable,omitempty"`

	// IsReturnable: Indicates that the property identifies data that should
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

	// IsSortable: Indicates that the property can be used for sorting.
	// Cannot be true for
	// properties that are repeatable. Cannot be true for properties whose
	// type
	// is object or user identifier. IsReturnable must be true to set this
	// option.
	// Only supported for Boolean, Date, Double, Integer, and
	// Timestamp
	// properties.
	IsSortable bool `json:"isSortable,omitempty"`

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

	TextPropertyOptions *TextPropertyOptions `json:"textPropertyOptions,omitempty"`

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

// PropertyDisplayOptions: The display options for a property.
type PropertyDisplayOptions struct {
	// DisplayLabel: The user friendly label for the property that will be
	// used if the property
	// is specified to be displayed in ObjectDisplayOptions. If given, the
	// display
	// label will be shown in front of the property values when the property
	// is
	// part of the object display options. For example, if the property
	// value is
	// '1', the value by itself may not be useful context for the user. If
	// the
	// display name given was 'priority', then the user will see 'priority :
	// 1' in
	// the search results which provides clear conext to search users. This
	// is
	// OPTIONAL; if not given, only the property values will be
	// displayed.
	// The maximum length is 32 characters.
	DisplayLabel string `json:"displayLabel,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DisplayLabel") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DisplayLabel") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PropertyDisplayOptions) MarshalJSON() ([]byte, error) {
	type NoMethod PropertyDisplayOptions
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
	// update
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
	// REPOSITORY_ERROR
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
	// ConnectorName: Name of connector making this call.
	// <br />Format: datasources/{source_id}/connectors/{ID}
	ConnectorName string `json:"connectorName,omitempty"`

	// DebugOptions: Common debug options.
	DebugOptions *DebugOptions `json:"debugOptions,omitempty"`

	// Item: Item to push onto the queue.
	Item *PushItem `json:"item,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ConnectorName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ConnectorName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
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
	// DisableNlInterpretation: Flag to disable natural language (NL)
	// interpretation of queries. Default is
	// false, Set to true to disable natural language interpretation.
	// NL
	// interpretation only applies to predefined datasources.
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

// QueryOperator: The definition of a operator that can be used in a
// Search/Suggest request.
type QueryOperator struct {
	// DisplayName: Display name of the operator
	DisplayName string `json:"displayName,omitempty"`

	// EnumValues: Potential list of values for the opeatror field. This
	// field is only filled
	// when we can safely enumerate all the possible values of this
	// operator.
	EnumValues []string `json:"enumValues,omitempty"`

	// GreaterThanOperatorName: Indicates the operator name that can be used
	// to  isolate the property using
	// the greater-than operator.
	GreaterThanOperatorName string `json:"greaterThanOperatorName,omitempty"`

	// IsFacetable: Can this operator be used to get facets.
	IsFacetable bool `json:"isFacetable,omitempty"`

	// IsRepeatable: Indicates if multiple values can be set for this
	// property.
	IsRepeatable bool `json:"isRepeatable,omitempty"`

	// IsReturnable: Will the property associated with this facet be
	// returned as part of search
	// results.
	IsReturnable bool `json:"isReturnable,omitempty"`

	// IsSortable: Can this operator be used to sort results.
	IsSortable bool `json:"isSortable,omitempty"`

	// IsSuggestable: Can get suggestions for this field.
	IsSuggestable bool `json:"isSuggestable,omitempty"`

	// LessThanOperatorName: Indicates the operator name that can be used to
	//  isolate the property using
	// the less-than operator.
	LessThanOperatorName string `json:"lessThanOperatorName,omitempty"`

	// OperatorName: The name of the operator.
	OperatorName string `json:"operatorName,omitempty"`

	// Type: Type of the operator.
	//
	// Possible values:
	//   "UNKNOWN" - Invalid value.
	//   "INTEGER"
	//   "DOUBLE"
	//   "TIMESTAMP"
	//   "BOOLEAN"
	//   "ENUM"
	//   "DATE"
	//   "TEXT"
	//   "HTML"
	Type string `json:"type,omitempty"`

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

func (s *QueryOperator) MarshalJSON() ([]byte, error) {
	type NoMethod QueryOperator
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// QuerySource: List of sources that the user can search using the query
// API.
type QuerySource struct {
	// DisplayName: Display name of the data source.
	DisplayName string `json:"displayName,omitempty"`

	// Operators: List of all operators applicable for this source.
	Operators []*QueryOperator `json:"operators,omitempty"`

	// ShortName: A short name or alias for the source.  This value can be
	// used with the
	// 'source' operator.
	ShortName string `json:"shortName,omitempty"`

	// Source: Name of the source
	Source *Source `json:"source,omitempty"`

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

func (s *QuerySource) MarshalJSON() ([]byte, error) {
	type NoMethod QuerySource
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// QuerySuggestion: A completed query suggestion.
type QuerySuggestion struct {
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
	// DebugOptions: Debug options of the request
	DebugOptions *DebugOptions `json:"debugOptions,omitempty"`

	// LanguageCode: The BCP-47 language code, such as "en-US" or
	// "sr-Latn".
	// For more information,
	// see
	// http://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	// Fo
	// r translations.
	LanguageCode string `json:"languageCode,omitempty"`

	// SearchApplicationId: Id of the application created using
	// SearchApplicationsService.
	SearchApplicationId string `json:"searchApplicationId,omitempty"`

	// TimeZone: Current user's time zone id, such as "America/Los_Angeles"
	// or
	// "Australia/Sydney". These IDs are defined by
	// [Unicode Common Locale Data Repository
	// (CLDR)](http://cldr.unicode.org/)
	// project, and currently available in the
	// file
	// [timezone.xml](http://unicode.org/repos/cldr/trunk/common/bcp47/t
	// imezone.xml)
	TimeZone string `json:"timeZone,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DebugOptions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DebugOptions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RequestOptions) MarshalJSON() ([]byte, error) {
	type NoMethod RequestOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ResetSearchApplicationRequest struct {
	// DebugOptions: Common debug options.
	DebugOptions *DebugOptions `json:"debugOptions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DebugOptions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DebugOptions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ResetSearchApplicationRequest) MarshalJSON() ([]byte, error) {
	type NoMethod ResetSearchApplicationRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ResponseDebugInfo: Debugging information about the response.
type ResponseDebugInfo struct {
	// FormattedDebugInfo: General debug info formatted for display.
	FormattedDebugInfo string `json:"formattedDebugInfo,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FormattedDebugInfo")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FormattedDebugInfo") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ResponseDebugInfo) MarshalJSON() ([]byte, error) {
	type NoMethod ResponseDebugInfo
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

// ResultCounts: Result count information
type ResultCounts struct {
	// SourceResultCounts: Result count information for each source with
	// results.
	SourceResultCounts []*SourceResultCount `json:"sourceResultCounts,omitempty"`

	// ForceSendFields is a list of field names (e.g. "SourceResultCounts")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "SourceResultCounts") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ResultCounts) MarshalJSON() ([]byte, error) {
	type NoMethod ResultCounts
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ResultDebugInfo: Debugging information about the result.
type ResultDebugInfo struct {
	// FormattedDebugInfo: General debug info formatted for display.
	FormattedDebugInfo string `json:"formattedDebugInfo,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FormattedDebugInfo")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FormattedDebugInfo") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ResultDebugInfo) MarshalJSON() ([]byte, error) {
	type NoMethod ResultDebugInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ResultDisplayField: Display Fields for Search Results
type ResultDisplayField struct {
	// Label: The display label for the property.
	Label string `json:"label,omitempty"`

	// OperatorName: The operator name of the property.
	OperatorName string `json:"operatorName,omitempty"`

	// Property: The name value pair for the property.
	Property *NamedProperty `json:"property,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Label") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Label") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ResultDisplayField) MarshalJSON() ([]byte, error) {
	type NoMethod ResultDisplayField
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ResultDisplayLine: The collection of fields that make up a displayed
// line
type ResultDisplayLine struct {
	Fields []*ResultDisplayField `json:"fields,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Fields") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Fields") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ResultDisplayLine) MarshalJSON() ([]byte, error) {
	type NoMethod ResultDisplayLine
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ResultDisplayMetadata struct {
	// Metalines: The metalines content to be displayed with the result.
	Metalines []*ResultDisplayLine `json:"metalines,omitempty"`

	// ObjectTypeLabel: The display label for the object.
	ObjectTypeLabel string `json:"objectTypeLabel,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Metalines") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Metalines") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ResultDisplayMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod ResultDisplayMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type RetrievalImportance struct {
	// Importance: Indicates the ranking importance given to property when
	// it is matched
	// during retrieval. Once set, the token importance of a property cannot
	// be
	// changed.
	//
	// Possible values:
	//   "DEFAULT" - Treat the match like a body text match.
	//   "HIGHEST" - Treat the match like a match against title of the item.
	//   "HIGH" - Treat the match with higher importance than body text.
	//   "LOW" - Treat the match with lower importance than body text.
	//   "NONE" - Do not match against this field during retrieval. The
	// property can still
	// be used for operator matching, faceting, and suggest if
	// desired.
	Importance string `json:"importance,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Importance") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Importance") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RetrievalImportance) MarshalJSON() ([]byte, error) {
	type NoMethod RetrievalImportance
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Schema: The schema definition for a data source.
type Schema struct {
	// ObjectDefinitions: The list of top-level objects for the data
	// source.
	// The maximum number of elements is 10.
	ObjectDefinitions []*ObjectDefinition `json:"objectDefinitions,omitempty"`

	// OperationIds: IDs of the Long Running Operations (LROs) currently
	// running for this
	// schema. After modifying the schema, wait for opeations to
	// complete
	// before indexing additional content.
	OperationIds []string `json:"operationIds,omitempty"`

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

// ScoringConfig: Scoring configurations for a source while processing
// a
// Search or
// Suggest request.
type ScoringConfig struct {
	// DisableFreshness: Whether to use freshness as a ranking signal. By
	// default, freshness is used
	// as a ranking signal.
	DisableFreshness bool `json:"disableFreshness,omitempty"`

	// DisablePersonalization: Whether to personalize the results. By
	// default, personal signals will
	// be used to boost results.
	DisablePersonalization bool `json:"disablePersonalization,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DisableFreshness") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DisableFreshness") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ScoringConfig) MarshalJSON() ([]byte, error) {
	type NoMethod ScoringConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SearchApplication: SearchApplication
type SearchApplication struct {
	// DataSourceRestrictions: Retrictions applied to the
	// configurations.
	// The maximum number of elements is 10.
	DataSourceRestrictions []*DataSourceRestriction `json:"dataSourceRestrictions,omitempty"`

	// DefaultFacetOptions: The default fields for returning facet
	// results.
	// The sources specified here also have been included
	// in
	// data_source_restrictions
	// above.
	DefaultFacetOptions []*FacetOptions `json:"defaultFacetOptions,omitempty"`

	// DefaultSortOptions: The default options for sorting the search
	// results
	DefaultSortOptions *SortOptions `json:"defaultSortOptions,omitempty"`

	// DisplayName: Display name of the Search Application.
	// The maximum length is 300 characters.
	DisplayName string `json:"displayName,omitempty"`

	// Name: Name of the Search Application.
	// <br />Format: searchapplications/{application_id}.
	Name string `json:"name,omitempty"`

	// OperationIds: IDs of the Long Running Operations (LROs) currently
	// running for this schema.
	// Output only field.
	OperationIds []string `json:"operationIds,omitempty"`

	// ScoringConfig: Configuration for ranking results.
	ScoringConfig *ScoringConfig `json:"scoringConfig,omitempty"`

	// SourceConfig: Configuration for a sources specified in
	// data_source_restrictions.
	SourceConfig []*SourceConfig `json:"sourceConfig,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "DataSourceRestrictions") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DataSourceRestrictions")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *SearchApplication) MarshalJSON() ([]byte, error) {
	type NoMethod SearchApplication
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type SearchItemsByViewUrlRequest struct {
	// DebugOptions: Common debug options.
	DebugOptions *DebugOptions `json:"debugOptions,omitempty"`

	// PageToken: The next_page_token value returned from a previous
	// request, if any.
	PageToken string `json:"pageToken,omitempty"`

	// ViewUrl: Specify the full view URL to find the corresponding
	// item.
	// The maximum length is 2048 characters.
	ViewUrl string `json:"viewUrl,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DebugOptions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DebugOptions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SearchItemsByViewUrlRequest) MarshalJSON() ([]byte, error) {
	type NoMethod SearchItemsByViewUrlRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type SearchItemsByViewUrlResponse struct {
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

func (s *SearchItemsByViewUrlResponse) MarshalJSON() ([]byte, error) {
	type NoMethod SearchItemsByViewUrlResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SearchQualityMetadata: Additional search quality metadata of the
// item.
type SearchQualityMetadata struct {
	// Quality: An indication of the quality of the item, used to influence
	// search quality.
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
	// DataSourceRestrictions: The sources to use for querying. If not
	// specified, all data sources
	// from the current search application are used.
	DataSourceRestrictions []*DataSourceRestriction `json:"dataSourceRestrictions,omitempty"`

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

	// RequestOptions: Request options, such as the search application and
	// user timezone.
	RequestOptions *RequestOptions `json:"requestOptions,omitempty"`

	// SortOptions: The options for sorting the search results
	SortOptions *SortOptions `json:"sortOptions,omitempty"`

	// Start: Starting index of the results.
	Start int64 `json:"start,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "DataSourceRestrictions") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DataSourceRestrictions")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *SearchRequest) MarshalJSON() ([]byte, error) {
	type NoMethod SearchRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SearchResponse: The search API response.
type SearchResponse struct {
	// DebugInfo: Debugging information about the response.
	DebugInfo *ResponseDebugInfo `json:"debugInfo,omitempty"`

	// ErrorInfo: Error information about the response.
	ErrorInfo *ErrorInfo `json:"errorInfo,omitempty"`

	// FacetResults: Repeated facet results.
	FacetResults []*FacetResult `json:"facetResults,omitempty"`

	// HasMoreResults: Whether there are more search results matching the
	// query.
	HasMoreResults bool `json:"hasMoreResults,omitempty"`

	// QueryInterpretation: Query interpretation result for user query.
	// Empty if query interpretation
	// is disabled.
	QueryInterpretation *QueryInterpretation `json:"queryInterpretation,omitempty"`

	// ResultCountEstimate: The estimated result count for this query.
	ResultCountEstimate int64 `json:"resultCountEstimate,omitempty,string"`

	// ResultCountExact: The exact result count for this query.
	ResultCountExact int64 `json:"resultCountExact,omitempty,string"`

	// ResultCounts: Expanded result count information.
	ResultCounts *ResultCounts `json:"resultCounts,omitempty"`

	// Results: Results from a search query.
	Results []*SearchResult `json:"results,omitempty"`

	// SpellResults: Suggested spelling for the query.
	SpellResults []*SpellResult `json:"spellResults,omitempty"`

	// StructuredResults: Structured results for the user query. These
	// results are not counted
	// against the page_size.
	StructuredResults []*StructuredResult `json:"structuredResults,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "DebugInfo") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DebugInfo") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
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
	// ClusteredResults: If source is clustered, provide list of clustered
	// results. There will only
	// be one level of clustered results. If current source is not enabled
	// for
	// clustering, this field will be empty.
	ClusteredResults []*SearchResult `json:"clusteredResults,omitempty"`

	// DebugInfo: Debugging information about this search result.
	DebugInfo *ResultDebugInfo `json:"debugInfo,omitempty"`

	// Metadata: Metadata of the search result.
	Metadata *Metadata `json:"metadata,omitempty"`

	// Snippet: The concatenation of all snippets (summaries) available for
	// this result.
	Snippet *Snippet `json:"snippet,omitempty"`

	// Title: Title of the search result.
	Title string `json:"title,omitempty"`

	// Url: The URL of the result.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ClusteredResults") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ClusteredResults") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *SearchResult) MarshalJSON() ([]byte, error) {
	type NoMethod SearchResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Snippet: Snippet of the search result, which summarizes the content
// of the resulting
// page.
type Snippet struct {
	// MatchRanges: The matched ranges in the snippet.
	MatchRanges []*MatchRange `json:"matchRanges,omitempty"`

	// Snippet: The snippet of the document.
	// The snippet of the document. May contain escaped HTML character
	// that
	// should be unescaped prior to rendering.
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
	// OperatorName: Name of the operator corresponding to the field to sort
	// on.
	// The corresponding property must be marked as
	// sortable.
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
	// Indexing API.
	Name string `json:"name,omitempty"`

	// PredefinedSource: Predefined content source for Google Apps.
	//
	// Possible values:
	//   "NONE"
	//   "QUERY_HISTORY" - Suggests queries issued by the user in the past.
	// Only valid when used
	// with the suggest API. Ignored when used in the query API.
	//   "PERSON" - Suggests people in the organization. Only valid when
	// used
	// with the suggest API. Results in an error when used in the query API.
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

// SourceConfig: Configurations for a source while processing a
// Search or
// Suggest request.
type SourceConfig struct {
	// CrowdingConfig: The crowding configuration for the source.
	CrowdingConfig *SourceCrowdingConfig `json:"crowdingConfig,omitempty"`

	// ScoringConfig: The scoring configuration for the source.
	ScoringConfig *SourceScoringConfig `json:"scoringConfig,omitempty"`

	// Source: The source for which this configuration is to be used.
	Source *Source `json:"source,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CrowdingConfig") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CrowdingConfig") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *SourceConfig) MarshalJSON() ([]byte, error) {
	type NoMethod SourceConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SourceCrowdingConfig: Set search results crowding limits. Crowding is
// a situation in which
// multiple results from the same source or host "crowd out" other
// results,
// diminishing the quality of search for users. To foster better search
// quality
// and source diversity in search results, you can set a condition to
// reduce
// repetitive results by source.
type SourceCrowdingConfig struct {
	// Field: Use a field to control results crowding. For example, if you
	// want to
	// control overly similar results from Gmail topics, use
	// `thread_id`.
	// For similar pages from Google Sites, you can use `webspace_id`.
	// When matching query results contain the same field value
	// in
	// `GenericMetadata`, crowding limits are set on those records.
	Field string `json:"field,omitempty"`

	// NumResults: Maximum number of results allowed from a source.
	// No limits will be set on results if this value is less than or equal
	// to 0.
	NumResults int64 `json:"numResults,omitempty"`

	// NumSuggestions: Maximum number of suggestions allowed from a
	// source.
	// No limits will be set on results if this value is less than or equal
	// to 0.
	NumSuggestions int64 `json:"numSuggestions,omitempty"`

	// Source: Control results by content source. This option limits the
	// total number
	// of results from a given source and ignores field-based crowding
	// control.
	Source bool `json:"source,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Field") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Field") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SourceCrowdingConfig) MarshalJSON() ([]byte, error) {
	type NoMethod SourceCrowdingConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SourceResultCount: Per source result count information.
type SourceResultCount struct {
	// HasMoreResults: Whether there are more search results for this
	// source.
	HasMoreResults bool `json:"hasMoreResults,omitempty"`

	// ResultCountEstimate: The estimated result count for this source.
	ResultCountEstimate int64 `json:"resultCountEstimate,omitempty,string"`

	// ResultCountExact: The exact result count for this source.
	ResultCountExact int64 `json:"resultCountExact,omitempty,string"`

	// Source: The source the result count information is associated with.
	Source *Source `json:"source,omitempty"`

	// ForceSendFields is a list of field names (e.g. "HasMoreResults") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "HasMoreResults") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *SourceResultCount) MarshalJSON() ([]byte, error) {
	type NoMethod SourceResultCount
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SourceScoringConfig: Set the scoring configuration. This allows
// modifying the ranking of results
// for a source.
type SourceScoringConfig struct {
	// SourceImportance: Importance of the source.
	//
	// Possible values:
	//   "DEFAULT"
	//   "LOW"
	//   "HIGH"
	SourceImportance string `json:"sourceImportance,omitempty"`

	// ForceSendFields is a list of field names (e.g. "SourceImportance") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "SourceImportance") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *SourceScoringConfig) MarshalJSON() ([]byte, error) {
	type NoMethod SourceScoringConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type SpellResult struct {
	// SuggestedQuery: The suggested spelling of the query.
	SuggestedQuery string `json:"suggestedQuery,omitempty"`

	// ForceSendFields is a list of field names (e.g. "SuggestedQuery") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "SuggestedQuery") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *SpellResult) MarshalJSON() ([]byte, error) {
	type NoMethod SpellResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// StartUploadItemRequest: Start upload file request.
type StartUploadItemRequest struct {
	// ConnectorName: Name of connector making this call.
	// <br />Format: datasources/{source_id}/connectors/{ID}
	ConnectorName string `json:"connectorName,omitempty"`

	// DebugOptions: Common debug options.
	DebugOptions *DebugOptions `json:"debugOptions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ConnectorName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ConnectorName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *StartUploadItemRequest) MarshalJSON() ([]byte, error) {
	type NoMethod StartUploadItemRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Status: The `Status` type defines a logical error model that is
// suitable for different
// programming environments, including REST APIs and RPC APIs. It is
// used by
// [gRPC](https://github.com/grpc). The error model is designed to
// be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error
// message,
// and error details. The error code should be an enum value
// of
// google.rpc.Code, but it may accept additional error codes if needed.
// The
// error message should be a developer-facing English message that
// helps
// developers *understand* and *resolve* the error. If a localized
// user-facing
// error message is needed, put the localized message in the error
// details or
// localize it in the client. The optional error details may contain
// arbitrary
// information about the error. There is a predefined set of error
// detail types
// in the package `google.rpc` that can be used for common error
// conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error
// model, but it
// is not necessarily the actual wire format. When the `Status` message
// is
// exposed in different client libraries and different wire protocols,
// it can be
// mapped differently. For example, it will likely be mapped to some
// exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety
// of
// environments, either with or without APIs, to provide a
// consistent developer experience across different
// environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the
// client,
//     it may embed the `Status` in the normal response to indicate the
// partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step
// may
//     have a `Status` message for error reporting.
//
// - Batch operations. If a client uses batch request and batch
// response, the
//     `Status` message should be used directly inside batch response,
// one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous
// operation
//     results in its response, the status of those operations should
// be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message
// `Status` could
//     be used directly after any stripping needed for security/privacy
// reasons.
type Status struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details.  There is a
	// common set of
	// message types for APIs to use.
	Details []googleapi.RawMessage `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any
	// user-facing error message should be localized and sent in
	// the
	// google.rpc.Status.details field, or localized by the client.
	Message string `json:"message,omitempty"`

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

func (s *Status) MarshalJSON() ([]byte, error) {
	type NoMethod Status
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

// StructuredResult: Structured results that are returned as part of
// search request.
type StructuredResult struct {
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

func (s *StructuredResult) MarshalJSON() ([]byte, error) {
	type NoMethod StructuredResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SuggestRequest: Request of suggest API.
type SuggestRequest struct {
	// DataSourceRestrictions: The sources to use for suggestions. If not
	// specified, all data sources
	// from the current search application are used.
	DataSourceRestrictions []*DataSourceRestriction `json:"dataSourceRestrictions,omitempty"`

	// Query: Partial query for the completion suggestion.
	Query string `json:"query,omitempty"`

	// RequestOptions: Request options, such as the search application and
	// user timezone.
	RequestOptions *RequestOptions `json:"requestOptions,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "DataSourceRestrictions") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DataSourceRestrictions")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *SuggestRequest) MarshalJSON() ([]byte, error) {
	type NoMethod SuggestRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SuggestResponse: Response of the suggest API.
type SuggestResponse struct {
	// SuggestResults: List of suggestion results.
	SuggestResults []*SuggestResult `json:"suggestResults,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "SuggestResults") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "SuggestResults") to
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
	PeopleSuggestion *PeopleSuggestion `json:"peopleSuggestion,omitempty"`

	QuerySuggestion *QuerySuggestion `json:"querySuggestion,omitempty"`

	// Source: The source of the suggestion.
	Source *Source `json:"source,omitempty"`

	// SuggestedQuery: The suggested query that will be used for search,
	// when the user
	// clicks on the suggestion
	SuggestedQuery string `json:"suggestedQuery,omitempty"`

	// ForceSendFields is a list of field names (e.g. "PeopleSuggestion") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "PeopleSuggestion") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *SuggestResult) MarshalJSON() ([]byte, error) {
	type NoMethod SuggestResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TextOperatorOptions: Used to provide a search operator for text
// properties. This is optional.
// Search operators let users restrict the query to specific fields
// relevant
// to the type of item being searched.
type TextOperatorOptions struct {
	// ExactMatchWithOperator: If true, the text value will be tokenized as
	// one atomic value in
	// operator searches and facet matches. For example, if the operator
	// name is
	// "genre" and the value is "science-fiction" the query
	// restrictions
	// "genre:science" and "genre:fiction" will not match the
	// item;
	// "genre:science-fiction" will. Value matching is case-sensitive
	// and does not remove special characters.
	// If false, the text will be tokenized. For example, if the value
	// is
	// "science-fiction" the queries "genre:science" and "genre:fiction"
	// will
	// match the item.
	ExactMatchWithOperator bool `json:"exactMatchWithOperator,omitempty"`

	// OperatorName: Indicates the operator name required in the query in
	// order to isolate the
	// text property. For example, if operatorName is *subject* and
	// the
	// property's name is *subjectLine*, then queries
	// like
	// *subject:&lt;value&gt;* will show results only where the value of
	// the
	// property named *subjectLine* matches *&lt;value&gt;*. By contrast,
	// a
	// search that uses the same *&lt;value&gt;* without an operator will
	// return
	// all items where *&lt;value&gt;* matches the value of any
	// text properties or text within the content field for the item.
	// The operator name can only contain lowercase letters (a-z).
	// The maximum length is 32 characters.
	OperatorName string `json:"operatorName,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "ExactMatchWithOperator") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExactMatchWithOperator")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *TextOperatorOptions) MarshalJSON() ([]byte, error) {
	type NoMethod TextOperatorOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TextPropertyOptions: Options for text properties.
type TextPropertyOptions struct {
	// OperatorOptions: If set, describes how the property should be used as
	// a search operator.
	OperatorOptions *TextOperatorOptions `json:"operatorOptions,omitempty"`

	// RetrievalImportance: Indicates the search quality importance of the
	// tokens within the
	// field when used for retrieval.
	RetrievalImportance *RetrievalImportance `json:"retrievalImportance,omitempty"`

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

func (s *TextPropertyOptions) MarshalJSON() ([]byte, error) {
	type NoMethod TextPropertyOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TextValues: List of text values.
type TextValues struct {
	// Values: The maximum allowable length for text values is 2048
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

func (s *TextValues) MarshalJSON() ([]byte, error) {
	type NoMethod TextValues
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
	// all items where *&lt;value&gt;* matches the value of any
	// String
	// properties or text within the content field for the item. The
	// operator
	// name can only contain lowercase letters (a-z). The maximum length is
	// 32
	// characters.
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

type UnmappedIdentity struct {
	// ExternalIdentity: The resource name for an external user.
	ExternalIdentity *Principal `json:"externalIdentity,omitempty"`

	// ResolutionStatusCode: The resolution status for the external
	// identity.
	//
	// Possible values:
	//   "CODE_UNSPECIFIED" - Input-only value.  Used to list all unmapped
	// identities regardless of
	// status.
	//   "NOT_FOUND" - The unmapped identity was not found in IDaaS, and
	// needs to be provided by
	// the user.
	//   "IDENTITY_SOURCE_NOT_FOUND" - The identity source associated with
	// the identity was either not found or
	// deleted.
	//   "IDENTITY_SOURCE_MISCONFIGURED" - IDaaS does not understand the
	// identity source, probably because the
	// schema was modified in a non compatible way.
	//   "TOO_MANY_MAPPINGS_FOUND" - The number of users associated with the
	// external identity is too large.
	//   "INTERNAL_ERROR" - Internal error.
	ResolutionStatusCode string `json:"resolutionStatusCode,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExternalIdentity") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExternalIdentity") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *UnmappedIdentity) MarshalJSON() ([]byte, error) {
	type NoMethod UnmappedIdentity
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type UnreserveItemsRequest struct {
	// ConnectorName: Name of connector making this call.
	// <br />Format: datasources/{source_id}/connectors/{ID}
	ConnectorName string `json:"connectorName,omitempty"`

	// DebugOptions: Common debug options.
	DebugOptions *DebugOptions `json:"debugOptions,omitempty"`

	// Queue: Name of a queue to unreserve items from.
	Queue string `json:"queue,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ConnectorName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ConnectorName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
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
	// DebugOptions: Common debug options.
	DebugOptions *DebugOptions `json:"debugOptions,omitempty"`

	Source *DataSource `json:"source,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DebugOptions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DebugOptions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
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
	// DebugOptions: Common debug options.
	DebugOptions *DebugOptions `json:"debugOptions,omitempty"`

	// Schema: The new schema for the source.
	Schema *Schema `json:"schema,omitempty"`

	// ValidateOnly: If true, the request will be validated without side
	// effects.
	ValidateOnly bool `json:"validateOnly,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DebugOptions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DebugOptions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
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
	// Name: Name of the content reference.
	// The maximum length is 2048 characters.
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

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

// method id "cloudsearch.debug.datasources.items.checkAccess":

type DebugDatasourcesItemsCheckAccessCall struct {
	s          *Service
	name       string
	principal  *Principal
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// CheckAccess: Checks whether an item is accessible by specified
// principal.
func (r *DebugDatasourcesItemsService) CheckAccess(name string, principal *Principal) *DebugDatasourcesItemsCheckAccessCall {
	c := &DebugDatasourcesItemsCheckAccessCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.principal = principal
	return c
}

// DebugOptionsEnableDebugging sets the optional parameter
// "debugOptions.enableDebugging": If set, the request will enable
// debugging features of Cloud Search.
// Only turn on this field, if asked by Google to help with debugging.
func (c *DebugDatasourcesItemsCheckAccessCall) DebugOptionsEnableDebugging(debugOptionsEnableDebugging bool) *DebugDatasourcesItemsCheckAccessCall {
	c.urlParams_.Set("debugOptions.enableDebugging", fmt.Sprint(debugOptionsEnableDebugging))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DebugDatasourcesItemsCheckAccessCall) Fields(s ...googleapi.Field) *DebugDatasourcesItemsCheckAccessCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DebugDatasourcesItemsCheckAccessCall) Context(ctx context.Context) *DebugDatasourcesItemsCheckAccessCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DebugDatasourcesItemsCheckAccessCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *DebugDatasourcesItemsCheckAccessCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.principal)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/debug/{+name}:checkAccess")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.debug.datasources.items.checkAccess" call.
// Exactly one of *CheckAccessResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *CheckAccessResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *DebugDatasourcesItemsCheckAccessCall) Do(opts ...googleapi.CallOption) (*CheckAccessResponse, error) {
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
	ret := &CheckAccessResponse{
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
	//   "description": "Checks whether an item is accessible by specified principal.",
	//   "flatPath": "v1/debug/datasources/{datasourcesId}/items/{itemsId}:checkAccess",
	//   "httpMethod": "POST",
	//   "id": "cloudsearch.debug.datasources.items.checkAccess",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "debugOptions.enableDebugging": {
	//       "description": "If set, the request will enable debugging features of Cloud Search.\nOnly turn on this field, if asked by Google to help with debugging.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "name": {
	//       "description": "Item name, format:\ndatasources/{source_id}/items/{item_id}",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+/items/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/debug/{+name}:checkAccess",
	//   "request": {
	//     "$ref": "Principal"
	//   },
	//   "response": {
	//     "$ref": "CheckAccessResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.debug"
	//   ]
	// }

}

// method id "cloudsearch.debug.datasources.items.searchByViewUrl":

type DebugDatasourcesItemsSearchByViewUrlCall struct {
	s                           *Service
	name                        string
	searchitemsbyviewurlrequest *SearchItemsByViewUrlRequest
	urlParams_                  gensupport.URLParams
	ctx_                        context.Context
	header_                     http.Header
}

// SearchByViewUrl: Fetches the item whose viewUrl exactly matches that
// of the URL provided
// in the request.
func (r *DebugDatasourcesItemsService) SearchByViewUrl(name string, searchitemsbyviewurlrequest *SearchItemsByViewUrlRequest) *DebugDatasourcesItemsSearchByViewUrlCall {
	c := &DebugDatasourcesItemsSearchByViewUrlCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.searchitemsbyviewurlrequest = searchitemsbyviewurlrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DebugDatasourcesItemsSearchByViewUrlCall) Fields(s ...googleapi.Field) *DebugDatasourcesItemsSearchByViewUrlCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DebugDatasourcesItemsSearchByViewUrlCall) Context(ctx context.Context) *DebugDatasourcesItemsSearchByViewUrlCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DebugDatasourcesItemsSearchByViewUrlCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *DebugDatasourcesItemsSearchByViewUrlCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.searchitemsbyviewurlrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/debug/{+name}/items:searchByViewUrl")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.debug.datasources.items.searchByViewUrl" call.
// Exactly one of *SearchItemsByViewUrlResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *SearchItemsByViewUrlResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *DebugDatasourcesItemsSearchByViewUrlCall) Do(opts ...googleapi.CallOption) (*SearchItemsByViewUrlResponse, error) {
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
	ret := &SearchItemsByViewUrlResponse{
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
	//   "description": "Fetches the item whose viewUrl exactly matches that of the URL provided\nin the request.",
	//   "flatPath": "v1/debug/datasources/{datasourcesId}/items:searchByViewUrl",
	//   "httpMethod": "POST",
	//   "id": "cloudsearch.debug.datasources.items.searchByViewUrl",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Source name, format:\ndatasources/{source_id}",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/debug/{+name}/items:searchByViewUrl",
	//   "request": {
	//     "$ref": "SearchItemsByViewUrlRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchItemsByViewUrlResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.debug"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *DebugDatasourcesItemsSearchByViewUrlCall) Pages(ctx context.Context, f func(*SearchItemsByViewUrlResponse) error) error {
	c.ctx_ = ctx
	defer func(pt string) { c.searchitemsbyviewurlrequest.PageToken = pt }(c.searchitemsbyviewurlrequest.PageToken) // reset paging to original point
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
		c.searchitemsbyviewurlrequest.PageToken = x.NextPageToken
	}
}

// method id "cloudsearch.debug.datasources.items.unmappedids.list":

type DebugDatasourcesItemsUnmappedidsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List all unmapped identities for a specific item.
func (r *DebugDatasourcesItemsUnmappedidsService) List(parent string) *DebugDatasourcesItemsUnmappedidsListCall {
	c := &DebugDatasourcesItemsUnmappedidsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// DebugOptionsEnableDebugging sets the optional parameter
// "debugOptions.enableDebugging": If set, the request will enable
// debugging features of Cloud Search.
// Only turn on this field, if asked by Google to help with debugging.
func (c *DebugDatasourcesItemsUnmappedidsListCall) DebugOptionsEnableDebugging(debugOptionsEnableDebugging bool) *DebugDatasourcesItemsUnmappedidsListCall {
	c.urlParams_.Set("debugOptions.enableDebugging", fmt.Sprint(debugOptionsEnableDebugging))
	return c
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// items to fetch in a request.
// Defaults to 100.
func (c *DebugDatasourcesItemsUnmappedidsListCall) PageSize(pageSize int64) *DebugDatasourcesItemsUnmappedidsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous List request, if any.
func (c *DebugDatasourcesItemsUnmappedidsListCall) PageToken(pageToken string) *DebugDatasourcesItemsUnmappedidsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DebugDatasourcesItemsUnmappedidsListCall) Fields(s ...googleapi.Field) *DebugDatasourcesItemsUnmappedidsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *DebugDatasourcesItemsUnmappedidsListCall) IfNoneMatch(entityTag string) *DebugDatasourcesItemsUnmappedidsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DebugDatasourcesItemsUnmappedidsListCall) Context(ctx context.Context) *DebugDatasourcesItemsUnmappedidsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DebugDatasourcesItemsUnmappedidsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *DebugDatasourcesItemsUnmappedidsListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/debug/{+parent}/unmappedids")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.debug.datasources.items.unmappedids.list" call.
// Exactly one of *ListUnmappedIdentitiesResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListUnmappedIdentitiesResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *DebugDatasourcesItemsUnmappedidsListCall) Do(opts ...googleapi.CallOption) (*ListUnmappedIdentitiesResponse, error) {
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
	ret := &ListUnmappedIdentitiesResponse{
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
	//   "description": "List all unmapped identities for a specific item.",
	//   "flatPath": "v1/debug/datasources/{datasourcesId}/items/{itemsId}/unmappedids",
	//   "httpMethod": "GET",
	//   "id": "cloudsearch.debug.datasources.items.unmappedids.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "debugOptions.enableDebugging": {
	//       "description": "If set, the request will enable debugging features of Cloud Search.\nOnly turn on this field, if asked by Google to help with debugging.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "pageSize": {
	//       "description": "Maximum number of items to fetch in a request.\nDefaults to 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The next_page_token value returned from a previous List request, if any.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "The name of the item, in the following format:\ndatasources/{source_id}/items/{ID}",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+/items/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/debug/{+parent}/unmappedids",
	//   "response": {
	//     "$ref": "ListUnmappedIdentitiesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.debug"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *DebugDatasourcesItemsUnmappedidsListCall) Pages(ctx context.Context, f func(*ListUnmappedIdentitiesResponse) error) error {
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

// method id "cloudsearch.debug.identitysources.items.listForunmappedidentity":

type DebugIdentitysourcesItemsListForunmappedidentityCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// ListForunmappedidentity: Lists names of items associated with an
// unmapped identity.
func (r *DebugIdentitysourcesItemsService) ListForunmappedidentity(parent string) *DebugIdentitysourcesItemsListForunmappedidentityCall {
	c := &DebugIdentitysourcesItemsListForunmappedidentityCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// DebugOptionsEnableDebugging sets the optional parameter
// "debugOptions.enableDebugging": If set, the request will enable
// debugging features of Cloud Search.
// Only turn on this field, if asked by Google to help with debugging.
func (c *DebugIdentitysourcesItemsListForunmappedidentityCall) DebugOptionsEnableDebugging(debugOptionsEnableDebugging bool) *DebugIdentitysourcesItemsListForunmappedidentityCall {
	c.urlParams_.Set("debugOptions.enableDebugging", fmt.Sprint(debugOptionsEnableDebugging))
	return c
}

// GroupResourceName sets the optional parameter "groupResourceName":
func (c *DebugIdentitysourcesItemsListForunmappedidentityCall) GroupResourceName(groupResourceName string) *DebugIdentitysourcesItemsListForunmappedidentityCall {
	c.urlParams_.Set("groupResourceName", groupResourceName)
	return c
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// items to fetch in a request.
// Defaults to 100.
func (c *DebugIdentitysourcesItemsListForunmappedidentityCall) PageSize(pageSize int64) *DebugIdentitysourcesItemsListForunmappedidentityCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous List request, if any.
func (c *DebugIdentitysourcesItemsListForunmappedidentityCall) PageToken(pageToken string) *DebugIdentitysourcesItemsListForunmappedidentityCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// UserResourceName sets the optional parameter "userResourceName":
func (c *DebugIdentitysourcesItemsListForunmappedidentityCall) UserResourceName(userResourceName string) *DebugIdentitysourcesItemsListForunmappedidentityCall {
	c.urlParams_.Set("userResourceName", userResourceName)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DebugIdentitysourcesItemsListForunmappedidentityCall) Fields(s ...googleapi.Field) *DebugIdentitysourcesItemsListForunmappedidentityCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *DebugIdentitysourcesItemsListForunmappedidentityCall) IfNoneMatch(entityTag string) *DebugIdentitysourcesItemsListForunmappedidentityCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DebugIdentitysourcesItemsListForunmappedidentityCall) Context(ctx context.Context) *DebugIdentitysourcesItemsListForunmappedidentityCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DebugIdentitysourcesItemsListForunmappedidentityCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *DebugIdentitysourcesItemsListForunmappedidentityCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/debug/{+parent}/items:forunmappedidentity")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.debug.identitysources.items.listForunmappedidentity" call.
// Exactly one of *ListItemNamesForUnmappedIdentityResponse or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *ListItemNamesForUnmappedIdentityResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *DebugIdentitysourcesItemsListForunmappedidentityCall) Do(opts ...googleapi.CallOption) (*ListItemNamesForUnmappedIdentityResponse, error) {
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
	ret := &ListItemNamesForUnmappedIdentityResponse{
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
	//   "description": "Lists names of items associated with an unmapped identity.",
	//   "flatPath": "v1/debug/identitysources/{identitysourcesId}/items:forunmappedidentity",
	//   "httpMethod": "GET",
	//   "id": "cloudsearch.debug.identitysources.items.listForunmappedidentity",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "debugOptions.enableDebugging": {
	//       "description": "If set, the request will enable debugging features of Cloud Search.\nOnly turn on this field, if asked by Google to help with debugging.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "groupResourceName": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Maximum number of items to fetch in a request.\nDefaults to 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The next_page_token value returned from a previous List request, if any.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "The name of the identity source, in the following format:\nidentitysources/{source_id}}",
	//       "location": "path",
	//       "pattern": "^identitysources/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userResourceName": {
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/debug/{+parent}/items:forunmappedidentity",
	//   "response": {
	//     "$ref": "ListItemNamesForUnmappedIdentityResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.debug"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *DebugIdentitysourcesItemsListForunmappedidentityCall) Pages(ctx context.Context, f func(*ListItemNamesForUnmappedIdentityResponse) error) error {
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

// method id "cloudsearch.debug.identitysources.unmappedids.list":

type DebugIdentitysourcesUnmappedidsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists unmapped user identities for an identity source.
func (r *DebugIdentitysourcesUnmappedidsService) List(parent string) *DebugIdentitysourcesUnmappedidsListCall {
	c := &DebugIdentitysourcesUnmappedidsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// DebugOptionsEnableDebugging sets the optional parameter
// "debugOptions.enableDebugging": If set, the request will enable
// debugging features of Cloud Search.
// Only turn on this field, if asked by Google to help with debugging.
func (c *DebugIdentitysourcesUnmappedidsListCall) DebugOptionsEnableDebugging(debugOptionsEnableDebugging bool) *DebugIdentitysourcesUnmappedidsListCall {
	c.urlParams_.Set("debugOptions.enableDebugging", fmt.Sprint(debugOptionsEnableDebugging))
	return c
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// items to fetch in a request.
// Defaults to 100.
func (c *DebugIdentitysourcesUnmappedidsListCall) PageSize(pageSize int64) *DebugIdentitysourcesUnmappedidsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous List request, if any.
func (c *DebugIdentitysourcesUnmappedidsListCall) PageToken(pageToken string) *DebugIdentitysourcesUnmappedidsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// ResolutionStatusCode sets the optional parameter
// "resolutionStatusCode": Limit users selection to this status.
//
// Possible values:
//   "CODE_UNSPECIFIED"
//   "NOT_FOUND"
//   "IDENTITY_SOURCE_NOT_FOUND"
//   "IDENTITY_SOURCE_MISCONFIGURED"
//   "TOO_MANY_MAPPINGS_FOUND"
//   "INTERNAL_ERROR"
func (c *DebugIdentitysourcesUnmappedidsListCall) ResolutionStatusCode(resolutionStatusCode string) *DebugIdentitysourcesUnmappedidsListCall {
	c.urlParams_.Set("resolutionStatusCode", resolutionStatusCode)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DebugIdentitysourcesUnmappedidsListCall) Fields(s ...googleapi.Field) *DebugIdentitysourcesUnmappedidsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *DebugIdentitysourcesUnmappedidsListCall) IfNoneMatch(entityTag string) *DebugIdentitysourcesUnmappedidsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DebugIdentitysourcesUnmappedidsListCall) Context(ctx context.Context) *DebugIdentitysourcesUnmappedidsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DebugIdentitysourcesUnmappedidsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *DebugIdentitysourcesUnmappedidsListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/debug/{+parent}/unmappedids")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.debug.identitysources.unmappedids.list" call.
// Exactly one of *ListUnmappedIdentitiesResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListUnmappedIdentitiesResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *DebugIdentitysourcesUnmappedidsListCall) Do(opts ...googleapi.CallOption) (*ListUnmappedIdentitiesResponse, error) {
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
	ret := &ListUnmappedIdentitiesResponse{
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
	//   "description": "Lists unmapped user identities for an identity source.",
	//   "flatPath": "v1/debug/identitysources/{identitysourcesId}/unmappedids",
	//   "httpMethod": "GET",
	//   "id": "cloudsearch.debug.identitysources.unmappedids.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "debugOptions.enableDebugging": {
	//       "description": "If set, the request will enable debugging features of Cloud Search.\nOnly turn on this field, if asked by Google to help with debugging.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "pageSize": {
	//       "description": "Maximum number of items to fetch in a request.\nDefaults to 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The next_page_token value returned from a previous List request, if any.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "The name of the identity source, in the following format:\nidentitysources/{source_id}",
	//       "location": "path",
	//       "pattern": "^identitysources/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resolutionStatusCode": {
	//       "description": "Limit users selection to this status.",
	//       "enum": [
	//         "CODE_UNSPECIFIED",
	//         "NOT_FOUND",
	//         "IDENTITY_SOURCE_NOT_FOUND",
	//         "IDENTITY_SOURCE_MISCONFIGURED",
	//         "TOO_MANY_MAPPINGS_FOUND",
	//         "INTERNAL_ERROR"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/debug/{+parent}/unmappedids",
	//   "response": {
	//     "$ref": "ListUnmappedIdentitiesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.debug"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *DebugIdentitysourcesUnmappedidsListCall) Pages(ctx context.Context, f func(*ListUnmappedIdentitiesResponse) error) error {
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

// DebugOptionsEnableDebugging sets the optional parameter
// "debugOptions.enableDebugging": If set, the request will enable
// debugging features of Cloud Search.
// Only turn on this field, if asked by Google to help with debugging.
func (c *IndexingDatasourcesDeleteSchemaCall) DebugOptionsEnableDebugging(debugOptionsEnableDebugging bool) *IndexingDatasourcesDeleteSchemaCall {
	c.urlParams_.Set("debugOptions.enableDebugging", fmt.Sprint(debugOptionsEnableDebugging))
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/indexing/{+name}/schema")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.indexing.datasources.deleteSchema" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *IndexingDatasourcesDeleteSchemaCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "flatPath": "v1/indexing/datasources/{datasourcesId}/schema",
	//   "httpMethod": "DELETE",
	//   "id": "cloudsearch.indexing.datasources.deleteSchema",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "debugOptions.enableDebugging": {
	//       "description": "If set, the request will enable debugging features of Cloud Search.\nOnly turn on this field, if asked by Google to help with debugging.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "name": {
	//       "description": "Name of the data source to delete Schema.  Format:\ndatasources/{source_id}",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/indexing/{+name}/schema",
	//   "response": {
	//     "$ref": "Operation"
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

// DebugOptionsEnableDebugging sets the optional parameter
// "debugOptions.enableDebugging": If set, the request will enable
// debugging features of Cloud Search.
// Only turn on this field, if asked by Google to help with debugging.
func (c *IndexingDatasourcesGetSchemaCall) DebugOptionsEnableDebugging(debugOptionsEnableDebugging bool) *IndexingDatasourcesGetSchemaCall {
	c.urlParams_.Set("debugOptions.enableDebugging", fmt.Sprint(debugOptionsEnableDebugging))
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/indexing/{+name}/schema")
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
	//   "flatPath": "v1/indexing/datasources/{datasourcesId}/schema",
	//   "httpMethod": "GET",
	//   "id": "cloudsearch.indexing.datasources.getSchema",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "debugOptions.enableDebugging": {
	//       "description": "If set, the request will enable debugging features of Cloud Search.\nOnly turn on this field, if asked by Google to help with debugging.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "name": {
	//       "description": "Name of the data source to get Schema.  Format:\ndatasources/{source_id}",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/indexing/{+name}/schema",
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/indexing/{+name}/schema")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.indexing.datasources.updateSchema" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *IndexingDatasourcesUpdateSchemaCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "flatPath": "v1/indexing/datasources/{datasourcesId}/schema",
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
	//   "path": "v1/indexing/{+name}/schema",
	//   "request": {
	//     "$ref": "UpdateSchemaRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
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

// ConnectorName sets the optional parameter "connectorName": Name of
// connector making this call.
// <br />Format: datasources/{source_id}/connectors/{ID}
func (c *IndexingDatasourcesItemsDeleteCall) ConnectorName(connectorName string) *IndexingDatasourcesItemsDeleteCall {
	c.urlParams_.Set("connectorName", connectorName)
	return c
}

// DebugOptionsEnableDebugging sets the optional parameter
// "debugOptions.enableDebugging": If set, the request will enable
// debugging features of Cloud Search.
// Only turn on this field, if asked by Google to help with debugging.
func (c *IndexingDatasourcesItemsDeleteCall) DebugOptionsEnableDebugging(debugOptionsEnableDebugging bool) *IndexingDatasourcesItemsDeleteCall {
	c.urlParams_.Set("debugOptions.enableDebugging", fmt.Sprint(debugOptionsEnableDebugging))
	return c
}

// Mode sets the optional parameter "mode": Required. The RequestMode
// for this request.
//
// Possible values:
//   "UNSPECIFIED"
//   "SYNCHRONOUS"
//   "ASYNCHRONOUS"
func (c *IndexingDatasourcesItemsDeleteCall) Mode(mode string) *IndexingDatasourcesItemsDeleteCall {
	c.urlParams_.Set("mode", mode)
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/indexing/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.indexing.datasources.items.delete" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *IndexingDatasourcesItemsDeleteCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "flatPath": "v1/indexing/datasources/{datasourcesId}/items/{itemsId}",
	//   "httpMethod": "DELETE",
	//   "id": "cloudsearch.indexing.datasources.items.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "connectorName": {
	//       "description": "Name of connector making this call.\n\u003cbr /\u003eFormat: datasources/{source_id}/connectors/{ID}",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "debugOptions.enableDebugging": {
	//       "description": "If set, the request will enable debugging features of Cloud Search.\nOnly turn on this field, if asked by Google to help with debugging.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "mode": {
	//       "description": "Required. The RequestMode for this request.",
	//       "enum": [
	//         "UNSPECIFIED",
	//         "SYNCHRONOUS",
	//         "ASYNCHRONOUS"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
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
	//   "path": "v1/indexing/{+name}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.indexing"
	//   ]
	// }

}

// method id "cloudsearch.indexing.datasources.items.deleteQueueItems":

type IndexingDatasourcesItemsDeleteQueueItemsCall struct {
	s                       *Service
	name                    string
	deletequeueitemsrequest *DeleteQueueItemsRequest
	urlParams_              gensupport.URLParams
	ctx_                    context.Context
	header_                 http.Header
}

// DeleteQueueItems: Deletes all items in a queue. This method is useful
// for deleting stale
// items.
func (r *IndexingDatasourcesItemsService) DeleteQueueItems(name string, deletequeueitemsrequest *DeleteQueueItemsRequest) *IndexingDatasourcesItemsDeleteQueueItemsCall {
	c := &IndexingDatasourcesItemsDeleteQueueItemsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.deletequeueitemsrequest = deletequeueitemsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *IndexingDatasourcesItemsDeleteQueueItemsCall) Fields(s ...googleapi.Field) *IndexingDatasourcesItemsDeleteQueueItemsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *IndexingDatasourcesItemsDeleteQueueItemsCall) Context(ctx context.Context) *IndexingDatasourcesItemsDeleteQueueItemsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *IndexingDatasourcesItemsDeleteQueueItemsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *IndexingDatasourcesItemsDeleteQueueItemsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.deletequeueitemsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/indexing/{+name}/items:deleteQueueItems")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.indexing.datasources.items.deleteQueueItems" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *IndexingDatasourcesItemsDeleteQueueItemsCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Deletes all items in a queue. This method is useful for deleting stale\nitems.",
	//   "flatPath": "v1/indexing/datasources/{datasourcesId}/items:deleteQueueItems",
	//   "httpMethod": "POST",
	//   "id": "cloudsearch.indexing.datasources.items.deleteQueueItems",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name of the Data Source to delete items in a queue.\nFormat: datasources/{source_id}",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/indexing/{+name}/items:deleteQueueItems",
	//   "request": {
	//     "$ref": "DeleteQueueItemsRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
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

// ConnectorName sets the optional parameter "connectorName": Name of
// connector making this call.
// <br />Format: datasources/{source_id}/connectors/{ID}
func (c *IndexingDatasourcesItemsGetCall) ConnectorName(connectorName string) *IndexingDatasourcesItemsGetCall {
	c.urlParams_.Set("connectorName", connectorName)
	return c
}

// DebugOptionsEnableDebugging sets the optional parameter
// "debugOptions.enableDebugging": If set, the request will enable
// debugging features of Cloud Search.
// Only turn on this field, if asked by Google to help with debugging.
func (c *IndexingDatasourcesItemsGetCall) DebugOptionsEnableDebugging(debugOptionsEnableDebugging bool) *IndexingDatasourcesItemsGetCall {
	c.urlParams_.Set("debugOptions.enableDebugging", fmt.Sprint(debugOptionsEnableDebugging))
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/indexing/{+name}")
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
	//   "flatPath": "v1/indexing/datasources/{datasourcesId}/items/{itemsId}",
	//   "httpMethod": "GET",
	//   "id": "cloudsearch.indexing.datasources.items.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "connectorName": {
	//       "description": "Name of connector making this call.\n\u003cbr /\u003eFormat: datasources/{source_id}/connectors/{ID}",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "debugOptions.enableDebugging": {
	//       "description": "If set, the request will enable debugging features of Cloud Search.\nOnly turn on this field, if asked by Google to help with debugging.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "name": {
	//       "description": "Name of the item to get info.\nFormat: datasources/{source_id}/items/{item_id}",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+/items/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/indexing/{+name}",
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/indexing/{+name}:index")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.indexing.datasources.items.index" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *IndexingDatasourcesItemsIndexCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "flatPath": "v1/indexing/datasources/{datasourcesId}/items/{itemsId}:index",
	//   "httpMethod": "POST",
	//   "id": "cloudsearch.indexing.datasources.items.index",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name of the Item. Format:\ndatasources/{source_id}/items/{item_id}\n\u003cbr /\u003eThis is a required field.\nThe maximum length is 1536 characters.",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+/items/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/indexing/{+name}:index",
	//   "request": {
	//     "$ref": "IndexItemRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
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

// ConnectorName sets the optional parameter "connectorName": Name of
// connector making this call.
// <br />Format: datasources/{source_id}/connectors/{ID}
func (c *IndexingDatasourcesItemsListCall) ConnectorName(connectorName string) *IndexingDatasourcesItemsListCall {
	c.urlParams_.Set("connectorName", connectorName)
	return c
}

// DebugOptionsEnableDebugging sets the optional parameter
// "debugOptions.enableDebugging": If set, the request will enable
// debugging features of Cloud Search.
// Only turn on this field, if asked by Google to help with debugging.
func (c *IndexingDatasourcesItemsListCall) DebugOptionsEnableDebugging(debugOptionsEnableDebugging bool) *IndexingDatasourcesItemsListCall {
	c.urlParams_.Set("debugOptions.enableDebugging", fmt.Sprint(debugOptionsEnableDebugging))
	return c
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// items to fetch in a request.
// The max value is 1000 when brief is true.  The max value is 10 if
// brief
// is false.
// <br />The default value is 10
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/indexing/{+name}/items")
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
	//   "flatPath": "v1/indexing/datasources/{datasourcesId}/items",
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
	//     "connectorName": {
	//       "description": "Name of connector making this call.\n\u003cbr /\u003eFormat: datasources/{source_id}/connectors/{ID}",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "debugOptions.enableDebugging": {
	//       "description": "If set, the request will enable debugging features of Cloud Search.\nOnly turn on this field, if asked by Google to help with debugging.",
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
	//       "description": "Maximum number of items to fetch in a request.\nThe max value is 1000 when brief is true.  The max value is 10 if brief\nis false.\n\u003cbr /\u003eThe default value is 10",
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
	//   "path": "v1/indexing/{+name}/items",
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/indexing/{+name}/items:poll")
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
	//   "flatPath": "v1/indexing/datasources/{datasourcesId}/items:poll",
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
	//   "path": "v1/indexing/{+name}/items:poll",
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

// Push: Pushes an item onto a queue for later polling and updating.
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/indexing/{+name}:push")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.indexing.datasources.items.push" call.
// Exactly one of *Item or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Item.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *IndexingDatasourcesItemsPushCall) Do(opts ...googleapi.CallOption) (*Item, error) {
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
	//   "description": "Pushes an item onto a queue for later polling and updating.",
	//   "flatPath": "v1/indexing/datasources/{datasourcesId}/items/{itemsId}:push",
	//   "httpMethod": "POST",
	//   "id": "cloudsearch.indexing.datasources.items.push",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name of the item to\npush into the indexing queue.\u003cbr /\u003e\nFormat: datasources/{source_id}/items/{ID}\n\u003cbr /\u003eThis is a required field.\nThe maximum length is 1536 characters.",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+/items/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/indexing/{+name}:push",
	//   "request": {
	//     "$ref": "PushItemRequest"
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/indexing/{+name}/items:unreserve")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.indexing.datasources.items.unreserve" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *IndexingDatasourcesItemsUnreserveCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "flatPath": "v1/indexing/datasources/{datasourcesId}/items:unreserve",
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
	//   "path": "v1/indexing/{+name}/items:unreserve",
	//   "request": {
	//     "$ref": "UnreserveItemsRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
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
// update.
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/indexing/{+name}:upload")
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
	//   "description": "Creates an upload session for uploading item content. For items smaller\nthan 100 KiB, it's easier to embed the content\ninline within\nupdate.",
	//   "flatPath": "v1/indexing/datasources/{datasourcesId}/items/{itemsId}:upload",
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
	//   "path": "v1/indexing/{+name}:upload",
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

// method id "cloudsearch.media.upload":

type MediaUploadCall struct {
	s            *Service
	resourceName string
	media        *Media
	urlParams_   gensupport.URLParams
	mediaInfo_   *gensupport.MediaInfo
	ctx_         context.Context
	header_      http.Header
}

// Upload: Uploads media for indexing.
//
// The upload endpoint supports direct and resumable upload protocols
// and
// is intended for large items that can not be inlined during index
// requests. To
// index large content:
//
// 1. Call upload to begin
//    a session and get the item reference.
// 1. Upload the content using the item reference's resource name.
// 1. Call index with the item
//    reference as the content.
//
// For additional information, see
// [Create a content connector using the REST
// API](https://developers.google.com/cloud-search/docs/guides/content-co
// nnector#rest).
func (r *MediaService) Upload(resourceName string, media *Media) *MediaUploadCall {
	c := &MediaUploadCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resourceName = resourceName
	c.media = media
	return c
}

// Media specifies the media to upload in one or more chunks. The chunk
// size may be controlled by supplying a MediaOption generated by
// googleapi.ChunkSize. The chunk size defaults to
// googleapi.DefaultUploadChunkSize.The Content-Type header used in the
// upload request will be determined by sniffing the contents of r,
// unless a MediaOption generated by googleapi.ContentType is
// supplied.
// At most one of Media and ResumableMedia may be set.
func (c *MediaUploadCall) Media(r io.Reader, options ...googleapi.MediaOption) *MediaUploadCall {
	c.mediaInfo_ = gensupport.NewInfoFromMedia(r, options)
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be
// canceled with ctx.
//
// Deprecated: use Media instead.
//
// At most one of Media and ResumableMedia may be set. mediaType
// identifies the MIME media type of the upload, such as "image/png". If
// mediaType is "", it will be auto-detected. The provided ctx will
// supersede any context previously provided to the Context method.
func (c *MediaUploadCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *MediaUploadCall {
	c.ctx_ = ctx
	c.mediaInfo_ = gensupport.NewInfoFromResumableMedia(r, size, mediaType)
	return c
}

// ProgressUpdater provides a callback function that will be called
// after every chunk. It should be a low-latency function in order to
// not slow down the upload operation. This should only be called when
// using ResumableMedia (as opposed to Media).
func (c *MediaUploadCall) ProgressUpdater(pu googleapi.ProgressUpdater) *MediaUploadCall {
	c.mediaInfo_.SetProgressUpdater(pu)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MediaUploadCall) Fields(s ...googleapi.Field) *MediaUploadCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
// This context will supersede any context previously provided to the
// ResumableMedia method.
func (c *MediaUploadCall) Context(ctx context.Context) *MediaUploadCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MediaUploadCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MediaUploadCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.media)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/media/{+resourceName}")
	if c.mediaInfo_ != nil {
		urls = strings.Replace(urls, "https://www.googleapis.com/", "https://www.googleapis.com/upload/", 1)
		c.urlParams_.Set("uploadType", c.mediaInfo_.UploadType())
	}
	if body == nil {
		body = new(bytes.Buffer)
		reqHeaders.Set("Content-Type", "application/json")
	}
	body, getBody, cleanup := c.mediaInfo_.UploadRequest(reqHeaders, body)
	defer cleanup()
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	gensupport.SetGetBody(req, getBody)
	googleapi.Expand(req.URL, map[string]string{
		"resourceName": c.resourceName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.media.upload" call.
// Exactly one of *Media or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Media.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *MediaUploadCall) Do(opts ...googleapi.CallOption) (*Media, error) {
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
	rx := c.mediaInfo_.ResumableUpload(res.Header.Get("Location"))
	if rx != nil {
		rx.Client = c.s.client
		rx.UserAgent = c.s.userAgent()
		ctx := c.ctx_
		if ctx == nil {
			ctx = context.TODO()
		}
		res, err = rx.Upload(ctx)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		if err := googleapi.CheckResponse(res); err != nil {
			return nil, err
		}
	}
	ret := &Media{
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
	//   "description": "Uploads media for indexing.\n\nThe upload endpoint supports direct and resumable upload protocols and\nis intended for large items that can not be inlined during index requests. To\nindex large content:\n\n1. Call upload to begin\n   a session and get the item reference.\n1. Upload the content using the item reference's resource name.\n1. Call index with the item\n   reference as the content.\n\nFor additional information, see\n[Create a content connector using the REST API](https://developers.google.com/cloud-search/docs/guides/content-connector#rest).",
	//   "flatPath": "v1/media/{mediaId}",
	//   "httpMethod": "POST",
	//   "id": "cloudsearch.media.upload",
	//   "mediaUpload": {
	//     "accept": [
	//       "*/*"
	//     ],
	//     "protocols": {
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/v1/media/{+resourceName}"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "resourceName"
	//   ],
	//   "parameters": {
	//     "resourceName": {
	//       "description": "Name of the media that is being downloaded.  See\nReadRequest.resource_name.",
	//       "location": "path",
	//       "pattern": "^.+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/media/{+resourceName}",
	//   "request": {
	//     "$ref": "Media"
	//   },
	//   "response": {
	//     "$ref": "Media"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.indexing"
	//   ],
	//   "supportsMediaUpload": true
	// }

}

// method id "cloudsearch.operations.get":

type OperationsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets the latest state of a long-running operation.  Clients can
// use this
// method to poll the operation result at intervals as recommended by
// the API
// service.
func (r *OperationsService) Get(name string) *OperationsGetCall {
	c := &OperationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsGetCall) Fields(s ...googleapi.Field) *OperationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *OperationsGetCall) IfNoneMatch(entityTag string) *OperationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OperationsGetCall) Context(ctx context.Context) *OperationsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OperationsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OperationsGetCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.operations.get" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *OperationsGetCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Gets the latest state of a long-running operation.  Clients can use this\nmethod to poll the operation result at intervals as recommended by the API\nservice.",
	//   "flatPath": "v1/operations/{operationsId}",
	//   "httpMethod": "GET",
	//   "id": "cloudsearch.operations.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource.",
	//       "location": "path",
	//       "pattern": "^operations/.+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.debug",
	//     "https://www.googleapis.com/auth/cloud_search.indexing",
	//     "https://www.googleapis.com/auth/cloud_search.settings",
	//     "https://www.googleapis.com/auth/cloud_search.settings.indexing",
	//     "https://www.googleapis.com/auth/cloud_search.settings.query"
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/query/search")
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
	//   "flatPath": "v1/query/search",
	//   "httpMethod": "POST",
	//   "id": "cloudsearch.query.search",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1/query/search",
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

// Suggest: Provides suggestions for autocompleting the query.
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/query/suggest")
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
	//   "description": "Provides suggestions for autocompleting the query.",
	//   "flatPath": "v1/query/suggest",
	//   "httpMethod": "POST",
	//   "id": "cloudsearch.query.suggest",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1/query/suggest",
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

// method id "cloudsearch.query.sources.list":

type QuerySourcesListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns list of sources that user can use for Search and
// Suggest APIs.
func (r *QuerySourcesService) List() *QuerySourcesListCall {
	c := &QuerySourcesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// PageToken sets the optional parameter "pageToken": Number of sources
// to return in the response.
func (c *QuerySourcesListCall) PageToken(pageToken string) *QuerySourcesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// RequestOptionsDebugOptionsEnableDebugging sets the optional parameter
// "requestOptions.debugOptions.enableDebugging": If set, the request
// will enable debugging features of Cloud Search.
// Only turn on this field, if asked by Google to help with debugging.
func (c *QuerySourcesListCall) RequestOptionsDebugOptionsEnableDebugging(requestOptionsDebugOptionsEnableDebugging bool) *QuerySourcesListCall {
	c.urlParams_.Set("requestOptions.debugOptions.enableDebugging", fmt.Sprint(requestOptionsDebugOptionsEnableDebugging))
	return c
}

// RequestOptionsLanguageCode sets the optional parameter
// "requestOptions.languageCode": The BCP-47 language code, such as
// "en-US" or "sr-Latn".
// For more information,
// see
// http://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
// Fo
// r translations.
func (c *QuerySourcesListCall) RequestOptionsLanguageCode(requestOptionsLanguageCode string) *QuerySourcesListCall {
	c.urlParams_.Set("requestOptions.languageCode", requestOptionsLanguageCode)
	return c
}

// RequestOptionsSearchApplicationId sets the optional parameter
// "requestOptions.searchApplicationId": Id of the application created
// using SearchApplicationsService.
func (c *QuerySourcesListCall) RequestOptionsSearchApplicationId(requestOptionsSearchApplicationId string) *QuerySourcesListCall {
	c.urlParams_.Set("requestOptions.searchApplicationId", requestOptionsSearchApplicationId)
	return c
}

// RequestOptionsTimeZone sets the optional parameter
// "requestOptions.timeZone": Current user's time zone id, such as
// "America/Los_Angeles" or
// "Australia/Sydney". These IDs are defined by
// [Unicode Common Locale Data Repository
// (CLDR)](http://cldr.unicode.org/)
// project, and currently available in the
// file
// [timezone.xml](http://unicode.org/repos/cldr/trunk/common/bcp47/t
// imezone.xml)
func (c *QuerySourcesListCall) RequestOptionsTimeZone(requestOptionsTimeZone string) *QuerySourcesListCall {
	c.urlParams_.Set("requestOptions.timeZone", requestOptionsTimeZone)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *QuerySourcesListCall) Fields(s ...googleapi.Field) *QuerySourcesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *QuerySourcesListCall) IfNoneMatch(entityTag string) *QuerySourcesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *QuerySourcesListCall) Context(ctx context.Context) *QuerySourcesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *QuerySourcesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *QuerySourcesListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/query/sources")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.query.sources.list" call.
// Exactly one of *ListQuerySourcesResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListQuerySourcesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *QuerySourcesListCall) Do(opts ...googleapi.CallOption) (*ListQuerySourcesResponse, error) {
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
	ret := &ListQuerySourcesResponse{
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
	//   "description": "Returns list of sources that user can use for Search and Suggest APIs.",
	//   "flatPath": "v1/query/sources",
	//   "httpMethod": "GET",
	//   "id": "cloudsearch.query.sources.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "pageToken": {
	//       "description": "Number of sources to return in the response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestOptions.debugOptions.enableDebugging": {
	//       "description": "If set, the request will enable debugging features of Cloud Search.\nOnly turn on this field, if asked by Google to help with debugging.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "requestOptions.languageCode": {
	//       "description": "The BCP-47 language code, such as \"en-US\" or \"sr-Latn\".\nFor more information, see\nhttp://www.unicode.org/reports/tr35/#Unicode_locale_identifier.\nFor translations.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestOptions.searchApplicationId": {
	//       "description": "Id of the application created using SearchApplicationsService.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestOptions.timeZone": {
	//       "description": "Current user's time zone id, such as \"America/Los_Angeles\" or\n\"Australia/Sydney\". These IDs are defined by\n[Unicode Common Locale Data Repository (CLDR)](http://cldr.unicode.org/)\nproject, and currently available in the file\n[timezone.xml](http://unicode.org/repos/cldr/trunk/common/bcp47/timezone.xml)",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/query/sources",
	//   "response": {
	//     "$ref": "ListQuerySourcesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.query"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *QuerySourcesListCall) Pages(ctx context.Context, f func(*ListQuerySourcesResponse) error) error {
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

// method id "cloudsearch.settings.datasources.create":

type SettingsDatasourcesCreateCall struct {
	s          *Service
	datasource *DataSource
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates data source.
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/settings/datasources")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.settings.datasources.create" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *SettingsDatasourcesCreateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Creates data source.",
	//   "flatPath": "v1/settings/datasources",
	//   "httpMethod": "POST",
	//   "id": "cloudsearch.settings.datasources.create",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1/settings/datasources",
	//   "request": {
	//     "$ref": "DataSource"
	//   },
	//   "response": {
	//     "$ref": "Operation"
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

// Delete: Deletes a data source.
func (r *SettingsDatasourcesService) Delete(name string) *SettingsDatasourcesDeleteCall {
	c := &SettingsDatasourcesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// DebugOptionsEnableDebugging sets the optional parameter
// "debugOptions.enableDebugging": If set, the request will enable
// debugging features of Cloud Search.
// Only turn on this field, if asked by Google to help with debugging.
func (c *SettingsDatasourcesDeleteCall) DebugOptionsEnableDebugging(debugOptionsEnableDebugging bool) *SettingsDatasourcesDeleteCall {
	c.urlParams_.Set("debugOptions.enableDebugging", fmt.Sprint(debugOptionsEnableDebugging))
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/settings/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.settings.datasources.delete" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *SettingsDatasourcesDeleteCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Deletes a data source.",
	//   "flatPath": "v1/settings/datasources/{datasourcesId}",
	//   "httpMethod": "DELETE",
	//   "id": "cloudsearch.settings.datasources.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "debugOptions.enableDebugging": {
	//       "description": "If set, the request will enable debugging features of Cloud Search.\nOnly turn on this field, if asked by Google to help with debugging.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "name": {
	//       "description": "Name of the data source.\nFormat: datasources/{source_id}.",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/settings/{+name}",
	//   "response": {
	//     "$ref": "Operation"
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

// Get: Gets a data source.
func (r *SettingsDatasourcesService) Get(name string) *SettingsDatasourcesGetCall {
	c := &SettingsDatasourcesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// DebugOptionsEnableDebugging sets the optional parameter
// "debugOptions.enableDebugging": If set, the request will enable
// debugging features of Cloud Search.
// Only turn on this field, if asked by Google to help with debugging.
func (c *SettingsDatasourcesGetCall) DebugOptionsEnableDebugging(debugOptionsEnableDebugging bool) *SettingsDatasourcesGetCall {
	c.urlParams_.Set("debugOptions.enableDebugging", fmt.Sprint(debugOptionsEnableDebugging))
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/settings/{+name}")
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
	//   "description": "Gets a data source.",
	//   "flatPath": "v1/settings/datasources/{datasourcesId}",
	//   "httpMethod": "GET",
	//   "id": "cloudsearch.settings.datasources.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "debugOptions.enableDebugging": {
	//       "description": "If set, the request will enable debugging features of Cloud Search.\nOnly turn on this field, if asked by Google to help with debugging.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "name": {
	//       "description": "Name of the data source resource.\nFormat: datasources/{source_id}.",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/settings/{+name}",
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

// List: Lists data sources.
func (r *SettingsDatasourcesService) List() *SettingsDatasourcesListCall {
	c := &SettingsDatasourcesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// DebugOptionsEnableDebugging sets the optional parameter
// "debugOptions.enableDebugging": If set, the request will enable
// debugging features of Cloud Search.
// Only turn on this field, if asked by Google to help with debugging.
func (c *SettingsDatasourcesListCall) DebugOptionsEnableDebugging(debugOptionsEnableDebugging bool) *SettingsDatasourcesListCall {
	c.urlParams_.Set("debugOptions.enableDebugging", fmt.Sprint(debugOptionsEnableDebugging))
	return c
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// data sources to fetch in a request.
// The max value is 100.
// <br />The default value is 10
func (c *SettingsDatasourcesListCall) PageSize(pageSize int64) *SettingsDatasourcesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/settings/datasources")
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
	//   "description": "Lists data sources.",
	//   "flatPath": "v1/settings/datasources",
	//   "httpMethod": "GET",
	//   "id": "cloudsearch.settings.datasources.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "debugOptions.enableDebugging": {
	//       "description": "If set, the request will enable debugging features of Cloud Search.\nOnly turn on this field, if asked by Google to help with debugging.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "pageSize": {
	//       "description": "Maximum number of data sources to fetch in a request.\nThe max value is 100.\n\u003cbr /\u003eThe default value is 10",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Starting index of the results.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/settings/datasources",
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

// Update: Updates a data source.
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/settings/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.settings.datasources.update" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *SettingsDatasourcesUpdateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Updates a data source.",
	//   "flatPath": "v1/settings/datasources/{datasourcesId}",
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
	//   "path": "v1/settings/{+name}",
	//   "request": {
	//     "$ref": "UpdateDataSourceRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.settings",
	//     "https://www.googleapis.com/auth/cloud_search.settings.indexing"
	//   ]
	// }

}

// method id "cloudsearch.settings.searchapplications.create":

type SettingsSearchapplicationsCreateCall struct {
	s                 *Service
	searchapplication *SearchApplication
	urlParams_        gensupport.URLParams
	ctx_              context.Context
	header_           http.Header
}

// Create: Creates a search application.
func (r *SettingsSearchapplicationsService) Create(searchapplication *SearchApplication) *SettingsSearchapplicationsCreateCall {
	c := &SettingsSearchapplicationsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.searchapplication = searchapplication
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SettingsSearchapplicationsCreateCall) Fields(s ...googleapi.Field) *SettingsSearchapplicationsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SettingsSearchapplicationsCreateCall) Context(ctx context.Context) *SettingsSearchapplicationsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SettingsSearchapplicationsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SettingsSearchapplicationsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.searchapplication)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/settings/searchapplications")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.settings.searchapplications.create" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *SettingsSearchapplicationsCreateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Creates a search application.",
	//   "flatPath": "v1/settings/searchapplications",
	//   "httpMethod": "POST",
	//   "id": "cloudsearch.settings.searchapplications.create",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1/settings/searchapplications",
	//   "request": {
	//     "$ref": "SearchApplication"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.settings",
	//     "https://www.googleapis.com/auth/cloud_search.settings.query"
	//   ]
	// }

}

// method id "cloudsearch.settings.searchapplications.delete":

type SettingsSearchapplicationsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes a search application.
func (r *SettingsSearchapplicationsService) Delete(name string) *SettingsSearchapplicationsDeleteCall {
	c := &SettingsSearchapplicationsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// DebugOptionsEnableDebugging sets the optional parameter
// "debugOptions.enableDebugging": If set, the request will enable
// debugging features of Cloud Search.
// Only turn on this field, if asked by Google to help with debugging.
func (c *SettingsSearchapplicationsDeleteCall) DebugOptionsEnableDebugging(debugOptionsEnableDebugging bool) *SettingsSearchapplicationsDeleteCall {
	c.urlParams_.Set("debugOptions.enableDebugging", fmt.Sprint(debugOptionsEnableDebugging))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SettingsSearchapplicationsDeleteCall) Fields(s ...googleapi.Field) *SettingsSearchapplicationsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SettingsSearchapplicationsDeleteCall) Context(ctx context.Context) *SettingsSearchapplicationsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SettingsSearchapplicationsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SettingsSearchapplicationsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/settings/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.settings.searchapplications.delete" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *SettingsSearchapplicationsDeleteCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Deletes a search application.",
	//   "flatPath": "v1/settings/searchapplications/{searchapplicationsId}",
	//   "httpMethod": "DELETE",
	//   "id": "cloudsearch.settings.searchapplications.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "debugOptions.enableDebugging": {
	//       "description": "If set, the request will enable debugging features of Cloud Search.\nOnly turn on this field, if asked by Google to help with debugging.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "name": {
	//       "description": "The name of the search application to be deleted.\n\u003cbr /\u003eFormat: applications/{application_id}.",
	//       "location": "path",
	//       "pattern": "^searchapplications/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/settings/{+name}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.settings",
	//     "https://www.googleapis.com/auth/cloud_search.settings.query"
	//   ]
	// }

}

// method id "cloudsearch.settings.searchapplications.get":

type SettingsSearchapplicationsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets the specified search application.
func (r *SettingsSearchapplicationsService) Get(name string) *SettingsSearchapplicationsGetCall {
	c := &SettingsSearchapplicationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// DebugOptionsEnableDebugging sets the optional parameter
// "debugOptions.enableDebugging": If set, the request will enable
// debugging features of Cloud Search.
// Only turn on this field, if asked by Google to help with debugging.
func (c *SettingsSearchapplicationsGetCall) DebugOptionsEnableDebugging(debugOptionsEnableDebugging bool) *SettingsSearchapplicationsGetCall {
	c.urlParams_.Set("debugOptions.enableDebugging", fmt.Sprint(debugOptionsEnableDebugging))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SettingsSearchapplicationsGetCall) Fields(s ...googleapi.Field) *SettingsSearchapplicationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *SettingsSearchapplicationsGetCall) IfNoneMatch(entityTag string) *SettingsSearchapplicationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SettingsSearchapplicationsGetCall) Context(ctx context.Context) *SettingsSearchapplicationsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SettingsSearchapplicationsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SettingsSearchapplicationsGetCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/settings/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.settings.searchapplications.get" call.
// Exactly one of *SearchApplication or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *SearchApplication.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *SettingsSearchapplicationsGetCall) Do(opts ...googleapi.CallOption) (*SearchApplication, error) {
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
	ret := &SearchApplication{
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
	//   "description": "Gets the specified search application.",
	//   "flatPath": "v1/settings/searchapplications/{searchapplicationsId}",
	//   "httpMethod": "GET",
	//   "id": "cloudsearch.settings.searchapplications.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "debugOptions.enableDebugging": {
	//       "description": "If set, the request will enable debugging features of Cloud Search.\nOnly turn on this field, if asked by Google to help with debugging.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "name": {
	//       "description": "Name of the search application.\n\u003cbr /\u003eFormat: applications/{application_id}.",
	//       "location": "path",
	//       "pattern": "^searchapplications/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/settings/{+name}",
	//   "response": {
	//     "$ref": "SearchApplication"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.settings",
	//     "https://www.googleapis.com/auth/cloud_search.settings.query"
	//   ]
	// }

}

// method id "cloudsearch.settings.searchapplications.list":

type SettingsSearchapplicationsListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists all search applications.
func (r *SettingsSearchapplicationsService) List() *SettingsSearchapplicationsListCall {
	c := &SettingsSearchapplicationsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// DebugOptionsEnableDebugging sets the optional parameter
// "debugOptions.enableDebugging": If set, the request will enable
// debugging features of Cloud Search.
// Only turn on this field, if asked by Google to help with debugging.
func (c *SettingsSearchapplicationsListCall) DebugOptionsEnableDebugging(debugOptionsEnableDebugging bool) *SettingsSearchapplicationsListCall {
	c.urlParams_.Set("debugOptions.enableDebugging", fmt.Sprint(debugOptionsEnableDebugging))
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return.
func (c *SettingsSearchapplicationsListCall) PageSize(pageSize int64) *SettingsSearchapplicationsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous List request, if
// any.
// <br/> The default value is 10
func (c *SettingsSearchapplicationsListCall) PageToken(pageToken string) *SettingsSearchapplicationsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SettingsSearchapplicationsListCall) Fields(s ...googleapi.Field) *SettingsSearchapplicationsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *SettingsSearchapplicationsListCall) IfNoneMatch(entityTag string) *SettingsSearchapplicationsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SettingsSearchapplicationsListCall) Context(ctx context.Context) *SettingsSearchapplicationsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SettingsSearchapplicationsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SettingsSearchapplicationsListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/settings/searchapplications")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.settings.searchapplications.list" call.
// Exactly one of *ListSearchApplicationsResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListSearchApplicationsResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *SettingsSearchapplicationsListCall) Do(opts ...googleapi.CallOption) (*ListSearchApplicationsResponse, error) {
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
	ret := &ListSearchApplicationsResponse{
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
	//   "description": "Lists all search applications.",
	//   "flatPath": "v1/settings/searchapplications",
	//   "httpMethod": "GET",
	//   "id": "cloudsearch.settings.searchapplications.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "debugOptions.enableDebugging": {
	//       "description": "If set, the request will enable debugging features of Cloud Search.\nOnly turn on this field, if asked by Google to help with debugging.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of items to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The next_page_token value returned from a previous List request, if any.\n\u003cbr/\u003e The default value is 10",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/settings/searchapplications",
	//   "response": {
	//     "$ref": "ListSearchApplicationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.settings",
	//     "https://www.googleapis.com/auth/cloud_search.settings.query"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *SettingsSearchapplicationsListCall) Pages(ctx context.Context, f func(*ListSearchApplicationsResponse) error) error {
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

// method id "cloudsearch.settings.searchapplications.reset":

type SettingsSearchapplicationsResetCall struct {
	s                             *Service
	name                          string
	resetsearchapplicationrequest *ResetSearchApplicationRequest
	urlParams_                    gensupport.URLParams
	ctx_                          context.Context
	header_                       http.Header
}

// Reset: Resets a search application to default settings. This will
// return an empty
// response.
func (r *SettingsSearchapplicationsService) Reset(name string, resetsearchapplicationrequest *ResetSearchApplicationRequest) *SettingsSearchapplicationsResetCall {
	c := &SettingsSearchapplicationsResetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.resetsearchapplicationrequest = resetsearchapplicationrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SettingsSearchapplicationsResetCall) Fields(s ...googleapi.Field) *SettingsSearchapplicationsResetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SettingsSearchapplicationsResetCall) Context(ctx context.Context) *SettingsSearchapplicationsResetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SettingsSearchapplicationsResetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SettingsSearchapplicationsResetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.resetsearchapplicationrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/settings/{+name}:reset")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.settings.searchapplications.reset" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *SettingsSearchapplicationsResetCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Resets a search application to default settings. This will return an empty\nresponse.",
	//   "flatPath": "v1/settings/searchapplications/{searchapplicationsId}:reset",
	//   "httpMethod": "POST",
	//   "id": "cloudsearch.settings.searchapplications.reset",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the search application to be reset.\n\u003cbr /\u003eFormat: applications/{application_id}.",
	//       "location": "path",
	//       "pattern": "^searchapplications/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/settings/{+name}:reset",
	//   "request": {
	//     "$ref": "ResetSearchApplicationRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.settings",
	//     "https://www.googleapis.com/auth/cloud_search.settings.query"
	//   ]
	// }

}

// method id "cloudsearch.settings.searchapplications.update":

type SettingsSearchapplicationsUpdateCall struct {
	s                 *Service
	name              string
	searchapplication *SearchApplication
	urlParams_        gensupport.URLParams
	ctx_              context.Context
	header_           http.Header
}

// Update: Updates a search application.
func (r *SettingsSearchapplicationsService) Update(name string, searchapplication *SearchApplication) *SettingsSearchapplicationsUpdateCall {
	c := &SettingsSearchapplicationsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.searchapplication = searchapplication
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SettingsSearchapplicationsUpdateCall) Fields(s ...googleapi.Field) *SettingsSearchapplicationsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SettingsSearchapplicationsUpdateCall) Context(ctx context.Context) *SettingsSearchapplicationsUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SettingsSearchapplicationsUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SettingsSearchapplicationsUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.searchapplication)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/settings/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.settings.searchapplications.update" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *SettingsSearchapplicationsUpdateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Updates a search application.",
	//   "flatPath": "v1/settings/searchapplications/{searchapplicationsId}",
	//   "httpMethod": "PUT",
	//   "id": "cloudsearch.settings.searchapplications.update",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name of the Search Application.\n\u003cbr /\u003eFormat: searchapplications/{application_id}.",
	//       "location": "path",
	//       "pattern": "^searchapplications/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/settings/{+name}",
	//   "request": {
	//     "$ref": "SearchApplication"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.settings",
	//     "https://www.googleapis.com/auth/cloud_search.settings.query"
	//   ]
	// }

}

// method id "cloudsearch.stats.getIndex":

type StatsGetIndexCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetIndex: Gets indexed item statistics aggreggated across all data
// sources.
func (r *StatsService) GetIndex() *StatsGetIndexCall {
	c := &StatsGetIndexCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// FromDateDay sets the optional parameter "fromDate.day": Day of month.
// Must be from 1 to 31 and valid for the year and month.
func (c *StatsGetIndexCall) FromDateDay(fromDateDay int64) *StatsGetIndexCall {
	c.urlParams_.Set("fromDate.day", fmt.Sprint(fromDateDay))
	return c
}

// FromDateMonth sets the optional parameter "fromDate.month": Month of
// date. Must be from 1 to 12.
func (c *StatsGetIndexCall) FromDateMonth(fromDateMonth int64) *StatsGetIndexCall {
	c.urlParams_.Set("fromDate.month", fmt.Sprint(fromDateMonth))
	return c
}

// FromDateYear sets the optional parameter "fromDate.year": Year of
// date. Must be from 1 to 9999.
func (c *StatsGetIndexCall) FromDateYear(fromDateYear int64) *StatsGetIndexCall {
	c.urlParams_.Set("fromDate.year", fmt.Sprint(fromDateYear))
	return c
}

// ToDateDay sets the optional parameter "toDate.day": Day of month.
// Must be from 1 to 31 and valid for the year and month.
func (c *StatsGetIndexCall) ToDateDay(toDateDay int64) *StatsGetIndexCall {
	c.urlParams_.Set("toDate.day", fmt.Sprint(toDateDay))
	return c
}

// ToDateMonth sets the optional parameter "toDate.month": Month of
// date. Must be from 1 to 12.
func (c *StatsGetIndexCall) ToDateMonth(toDateMonth int64) *StatsGetIndexCall {
	c.urlParams_.Set("toDate.month", fmt.Sprint(toDateMonth))
	return c
}

// ToDateYear sets the optional parameter "toDate.year": Year of date.
// Must be from 1 to 9999.
func (c *StatsGetIndexCall) ToDateYear(toDateYear int64) *StatsGetIndexCall {
	c.urlParams_.Set("toDate.year", fmt.Sprint(toDateYear))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StatsGetIndexCall) Fields(s ...googleapi.Field) *StatsGetIndexCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *StatsGetIndexCall) IfNoneMatch(entityTag string) *StatsGetIndexCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *StatsGetIndexCall) Context(ctx context.Context) *StatsGetIndexCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *StatsGetIndexCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *StatsGetIndexCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/stats/index")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.stats.getIndex" call.
// Exactly one of *GetCustomerIndexStatsResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GetCustomerIndexStatsResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *StatsGetIndexCall) Do(opts ...googleapi.CallOption) (*GetCustomerIndexStatsResponse, error) {
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
	ret := &GetCustomerIndexStatsResponse{
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
	//   "description": "Gets indexed item statistics aggreggated across all data sources.",
	//   "flatPath": "v1/stats/index",
	//   "httpMethod": "GET",
	//   "id": "cloudsearch.stats.getIndex",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "fromDate.day": {
	//       "description": "Day of month. Must be from 1 to 31 and valid for the year and month.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "fromDate.month": {
	//       "description": "Month of date. Must be from 1 to 12.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "fromDate.year": {
	//       "description": "Year of date. Must be from 1 to 9999.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "toDate.day": {
	//       "description": "Day of month. Must be from 1 to 31 and valid for the year and month.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "toDate.month": {
	//       "description": "Month of date. Must be from 1 to 12.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "toDate.year": {
	//       "description": "Year of date. Must be from 1 to 9999.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "v1/stats/index",
	//   "response": {
	//     "$ref": "GetCustomerIndexStatsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.stats",
	//     "https://www.googleapis.com/auth/cloud_search.stats.indexing"
	//   ]
	// }

}

// method id "cloudsearch.stats.index.datasources.get":

type StatsIndexDatasourcesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets indexed item statistics for a single data source.
func (r *StatsIndexDatasourcesService) Get(name string) *StatsIndexDatasourcesGetCall {
	c := &StatsIndexDatasourcesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// FromDateDay sets the optional parameter "fromDate.day": Day of month.
// Must be from 1 to 31 and valid for the year and month.
func (c *StatsIndexDatasourcesGetCall) FromDateDay(fromDateDay int64) *StatsIndexDatasourcesGetCall {
	c.urlParams_.Set("fromDate.day", fmt.Sprint(fromDateDay))
	return c
}

// FromDateMonth sets the optional parameter "fromDate.month": Month of
// date. Must be from 1 to 12.
func (c *StatsIndexDatasourcesGetCall) FromDateMonth(fromDateMonth int64) *StatsIndexDatasourcesGetCall {
	c.urlParams_.Set("fromDate.month", fmt.Sprint(fromDateMonth))
	return c
}

// FromDateYear sets the optional parameter "fromDate.year": Year of
// date. Must be from 1 to 9999.
func (c *StatsIndexDatasourcesGetCall) FromDateYear(fromDateYear int64) *StatsIndexDatasourcesGetCall {
	c.urlParams_.Set("fromDate.year", fmt.Sprint(fromDateYear))
	return c
}

// ToDateDay sets the optional parameter "toDate.day": Day of month.
// Must be from 1 to 31 and valid for the year and month.
func (c *StatsIndexDatasourcesGetCall) ToDateDay(toDateDay int64) *StatsIndexDatasourcesGetCall {
	c.urlParams_.Set("toDate.day", fmt.Sprint(toDateDay))
	return c
}

// ToDateMonth sets the optional parameter "toDate.month": Month of
// date. Must be from 1 to 12.
func (c *StatsIndexDatasourcesGetCall) ToDateMonth(toDateMonth int64) *StatsIndexDatasourcesGetCall {
	c.urlParams_.Set("toDate.month", fmt.Sprint(toDateMonth))
	return c
}

// ToDateYear sets the optional parameter "toDate.year": Year of date.
// Must be from 1 to 9999.
func (c *StatsIndexDatasourcesGetCall) ToDateYear(toDateYear int64) *StatsIndexDatasourcesGetCall {
	c.urlParams_.Set("toDate.year", fmt.Sprint(toDateYear))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StatsIndexDatasourcesGetCall) Fields(s ...googleapi.Field) *StatsIndexDatasourcesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *StatsIndexDatasourcesGetCall) IfNoneMatch(entityTag string) *StatsIndexDatasourcesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *StatsIndexDatasourcesGetCall) Context(ctx context.Context) *StatsIndexDatasourcesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *StatsIndexDatasourcesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *StatsIndexDatasourcesGetCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/stats/index/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudsearch.stats.index.datasources.get" call.
// Exactly one of *GetDataSourceIndexStatsResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GetDataSourceIndexStatsResponse.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *StatsIndexDatasourcesGetCall) Do(opts ...googleapi.CallOption) (*GetDataSourceIndexStatsResponse, error) {
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
	ret := &GetDataSourceIndexStatsResponse{
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
	//   "description": "Gets indexed item statistics for a single data source.",
	//   "flatPath": "v1/stats/index/datasources/{datasourcesId}",
	//   "httpMethod": "GET",
	//   "id": "cloudsearch.stats.index.datasources.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "fromDate.day": {
	//       "description": "Day of month. Must be from 1 to 31 and valid for the year and month.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "fromDate.month": {
	//       "description": "Month of date. Must be from 1 to 12.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "fromDate.year": {
	//       "description": "Year of date. Must be from 1 to 9999.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "name": {
	//       "description": "The resource id of the data source to retrieve statistics for,\nin the following format: \"datasources/{source_id}\"",
	//       "location": "path",
	//       "pattern": "^datasources/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "toDate.day": {
	//       "description": "Day of month. Must be from 1 to 31 and valid for the year and month.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "toDate.month": {
	//       "description": "Month of date. Must be from 1 to 12.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "toDate.year": {
	//       "description": "Year of date. Must be from 1 to 9999.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "v1/stats/index/{+name}",
	//   "response": {
	//     "$ref": "GetDataSourceIndexStatsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud_search",
	//     "https://www.googleapis.com/auth/cloud_search.stats",
	//     "https://www.googleapis.com/auth/cloud_search.stats.indexing"
	//   ]
	// }

}
