// Code generated by ent, DO NOT EDIT.

package joboffer

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the joboffer type in the database.
	Label = "job_offer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldFileURL holds the string denoting the file_url field in the database.
	FieldFileURL = "file_url"
	// FieldPublishedAt holds the string denoting the published_at field in the database.
	FieldPublishedAt = "published_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeSearchInputs holds the string denoting the search_inputs edge name in mutations.
	EdgeSearchInputs = "search_inputs"
	// Table holds the table name of the joboffer in the database.
	Table = "job_offers"
	// SearchInputsTable is the table that holds the search_inputs relation/edge. The primary key declared below.
	SearchInputsTable = "search_input_search_results"
	// SearchInputsInverseTable is the table name for the SearchInput entity.
	// It exists in this package in order to avoid circular dependency with the "searchinput" package.
	SearchInputsInverseTable = "search_inputs"
)

// Columns holds all SQL columns for joboffer fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldDescription,
	FieldURL,
	FieldFileURL,
	FieldPublishedAt,
	FieldUpdatedAt,
}

var (
	// SearchInputsPrimaryKey and SearchInputsColumn2 are the table columns denoting the
	// primary key for the search_inputs relation (M2M).
	SearchInputsPrimaryKey = []string{"search_input_id", "job_offer_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the JobOffer queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByFileURL orders the results by the file_url field.
func ByFileURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFileURL, opts...).ToFunc()
}

// ByPublishedAt orders the results by the published_at field.
func ByPublishedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublishedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// BySearchInputsCount orders the results by search_inputs count.
func BySearchInputsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSearchInputsStep(), opts...)
	}
}

// BySearchInputs orders the results by search_inputs terms.
func BySearchInputs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSearchInputsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSearchInputsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SearchInputsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, SearchInputsTable, SearchInputsPrimaryKey...),
	)
}
