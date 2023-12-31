// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/marianozunino/go-utu-llamados/ent/joboffer"
	"github.com/marianozunino/go-utu-llamados/ent/predicate"
	"github.com/marianozunino/go-utu-llamados/ent/searchinput"
)

// JobOfferQuery is the builder for querying JobOffer entities.
type JobOfferQuery struct {
	config
	ctx              *QueryContext
	order            []joboffer.OrderOption
	inters           []Interceptor
	predicates       []predicate.JobOffer
	withSearchInputs *SearchInputQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the JobOfferQuery builder.
func (joq *JobOfferQuery) Where(ps ...predicate.JobOffer) *JobOfferQuery {
	joq.predicates = append(joq.predicates, ps...)
	return joq
}

// Limit the number of records to be returned by this query.
func (joq *JobOfferQuery) Limit(limit int) *JobOfferQuery {
	joq.ctx.Limit = &limit
	return joq
}

// Offset to start from.
func (joq *JobOfferQuery) Offset(offset int) *JobOfferQuery {
	joq.ctx.Offset = &offset
	return joq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (joq *JobOfferQuery) Unique(unique bool) *JobOfferQuery {
	joq.ctx.Unique = &unique
	return joq
}

// Order specifies how the records should be ordered.
func (joq *JobOfferQuery) Order(o ...joboffer.OrderOption) *JobOfferQuery {
	joq.order = append(joq.order, o...)
	return joq
}

// QuerySearchInputs chains the current query on the "search_inputs" edge.
func (joq *JobOfferQuery) QuerySearchInputs() *SearchInputQuery {
	query := (&SearchInputClient{config: joq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := joq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := joq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(joboffer.Table, joboffer.FieldID, selector),
			sqlgraph.To(searchinput.Table, searchinput.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, joboffer.SearchInputsTable, joboffer.SearchInputsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(joq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first JobOffer entity from the query.
// Returns a *NotFoundError when no JobOffer was found.
func (joq *JobOfferQuery) First(ctx context.Context) (*JobOffer, error) {
	nodes, err := joq.Limit(1).All(setContextOp(ctx, joq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{joboffer.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (joq *JobOfferQuery) FirstX(ctx context.Context) *JobOffer {
	node, err := joq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first JobOffer ID from the query.
// Returns a *NotFoundError when no JobOffer ID was found.
func (joq *JobOfferQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = joq.Limit(1).IDs(setContextOp(ctx, joq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{joboffer.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (joq *JobOfferQuery) FirstIDX(ctx context.Context) int {
	id, err := joq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single JobOffer entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one JobOffer entity is found.
// Returns a *NotFoundError when no JobOffer entities are found.
func (joq *JobOfferQuery) Only(ctx context.Context) (*JobOffer, error) {
	nodes, err := joq.Limit(2).All(setContextOp(ctx, joq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{joboffer.Label}
	default:
		return nil, &NotSingularError{joboffer.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (joq *JobOfferQuery) OnlyX(ctx context.Context) *JobOffer {
	node, err := joq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only JobOffer ID in the query.
// Returns a *NotSingularError when more than one JobOffer ID is found.
// Returns a *NotFoundError when no entities are found.
func (joq *JobOfferQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = joq.Limit(2).IDs(setContextOp(ctx, joq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{joboffer.Label}
	default:
		err = &NotSingularError{joboffer.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (joq *JobOfferQuery) OnlyIDX(ctx context.Context) int {
	id, err := joq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of JobOffers.
func (joq *JobOfferQuery) All(ctx context.Context) ([]*JobOffer, error) {
	ctx = setContextOp(ctx, joq.ctx, "All")
	if err := joq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*JobOffer, *JobOfferQuery]()
	return withInterceptors[[]*JobOffer](ctx, joq, qr, joq.inters)
}

// AllX is like All, but panics if an error occurs.
func (joq *JobOfferQuery) AllX(ctx context.Context) []*JobOffer {
	nodes, err := joq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of JobOffer IDs.
func (joq *JobOfferQuery) IDs(ctx context.Context) (ids []int, err error) {
	if joq.ctx.Unique == nil && joq.path != nil {
		joq.Unique(true)
	}
	ctx = setContextOp(ctx, joq.ctx, "IDs")
	if err = joq.Select(joboffer.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (joq *JobOfferQuery) IDsX(ctx context.Context) []int {
	ids, err := joq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (joq *JobOfferQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, joq.ctx, "Count")
	if err := joq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, joq, querierCount[*JobOfferQuery](), joq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (joq *JobOfferQuery) CountX(ctx context.Context) int {
	count, err := joq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (joq *JobOfferQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, joq.ctx, "Exist")
	switch _, err := joq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (joq *JobOfferQuery) ExistX(ctx context.Context) bool {
	exist, err := joq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the JobOfferQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (joq *JobOfferQuery) Clone() *JobOfferQuery {
	if joq == nil {
		return nil
	}
	return &JobOfferQuery{
		config:           joq.config,
		ctx:              joq.ctx.Clone(),
		order:            append([]joboffer.OrderOption{}, joq.order...),
		inters:           append([]Interceptor{}, joq.inters...),
		predicates:       append([]predicate.JobOffer{}, joq.predicates...),
		withSearchInputs: joq.withSearchInputs.Clone(),
		// clone intermediate query.
		sql:  joq.sql.Clone(),
		path: joq.path,
	}
}

// WithSearchInputs tells the query-builder to eager-load the nodes that are connected to
// the "search_inputs" edge. The optional arguments are used to configure the query builder of the edge.
func (joq *JobOfferQuery) WithSearchInputs(opts ...func(*SearchInputQuery)) *JobOfferQuery {
	query := (&SearchInputClient{config: joq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	joq.withSearchInputs = query
	return joq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.JobOffer.Query().
//		GroupBy(joboffer.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (joq *JobOfferQuery) GroupBy(field string, fields ...string) *JobOfferGroupBy {
	joq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &JobOfferGroupBy{build: joq}
	grbuild.flds = &joq.ctx.Fields
	grbuild.label = joboffer.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.JobOffer.Query().
//		Select(joboffer.FieldTitle).
//		Scan(ctx, &v)
func (joq *JobOfferQuery) Select(fields ...string) *JobOfferSelect {
	joq.ctx.Fields = append(joq.ctx.Fields, fields...)
	sbuild := &JobOfferSelect{JobOfferQuery: joq}
	sbuild.label = joboffer.Label
	sbuild.flds, sbuild.scan = &joq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a JobOfferSelect configured with the given aggregations.
func (joq *JobOfferQuery) Aggregate(fns ...AggregateFunc) *JobOfferSelect {
	return joq.Select().Aggregate(fns...)
}

func (joq *JobOfferQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range joq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, joq); err != nil {
				return err
			}
		}
	}
	for _, f := range joq.ctx.Fields {
		if !joboffer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if joq.path != nil {
		prev, err := joq.path(ctx)
		if err != nil {
			return err
		}
		joq.sql = prev
	}
	return nil
}

func (joq *JobOfferQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*JobOffer, error) {
	var (
		nodes       = []*JobOffer{}
		_spec       = joq.querySpec()
		loadedTypes = [1]bool{
			joq.withSearchInputs != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*JobOffer).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &JobOffer{config: joq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, joq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := joq.withSearchInputs; query != nil {
		if err := joq.loadSearchInputs(ctx, query, nodes,
			func(n *JobOffer) { n.Edges.SearchInputs = []*SearchInput{} },
			func(n *JobOffer, e *SearchInput) { n.Edges.SearchInputs = append(n.Edges.SearchInputs, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (joq *JobOfferQuery) loadSearchInputs(ctx context.Context, query *SearchInputQuery, nodes []*JobOffer, init func(*JobOffer), assign func(*JobOffer, *SearchInput)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*JobOffer)
	nids := make(map[int]map[*JobOffer]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(joboffer.SearchInputsTable)
		s.Join(joinT).On(s.C(searchinput.FieldID), joinT.C(joboffer.SearchInputsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(joboffer.SearchInputsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(joboffer.SearchInputsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*JobOffer]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*SearchInput](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "search_inputs" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (joq *JobOfferQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := joq.querySpec()
	_spec.Node.Columns = joq.ctx.Fields
	if len(joq.ctx.Fields) > 0 {
		_spec.Unique = joq.ctx.Unique != nil && *joq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, joq.driver, _spec)
}

func (joq *JobOfferQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(joboffer.Table, joboffer.Columns, sqlgraph.NewFieldSpec(joboffer.FieldID, field.TypeInt))
	_spec.From = joq.sql
	if unique := joq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if joq.path != nil {
		_spec.Unique = true
	}
	if fields := joq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, joboffer.FieldID)
		for i := range fields {
			if fields[i] != joboffer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := joq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := joq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := joq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := joq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (joq *JobOfferQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(joq.driver.Dialect())
	t1 := builder.Table(joboffer.Table)
	columns := joq.ctx.Fields
	if len(columns) == 0 {
		columns = joboffer.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if joq.sql != nil {
		selector = joq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if joq.ctx.Unique != nil && *joq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range joq.predicates {
		p(selector)
	}
	for _, p := range joq.order {
		p(selector)
	}
	if offset := joq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := joq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// JobOfferGroupBy is the group-by builder for JobOffer entities.
type JobOfferGroupBy struct {
	selector
	build *JobOfferQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (jogb *JobOfferGroupBy) Aggregate(fns ...AggregateFunc) *JobOfferGroupBy {
	jogb.fns = append(jogb.fns, fns...)
	return jogb
}

// Scan applies the selector query and scans the result into the given value.
func (jogb *JobOfferGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, jogb.build.ctx, "GroupBy")
	if err := jogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*JobOfferQuery, *JobOfferGroupBy](ctx, jogb.build, jogb, jogb.build.inters, v)
}

func (jogb *JobOfferGroupBy) sqlScan(ctx context.Context, root *JobOfferQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(jogb.fns))
	for _, fn := range jogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*jogb.flds)+len(jogb.fns))
		for _, f := range *jogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*jogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := jogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// JobOfferSelect is the builder for selecting fields of JobOffer entities.
type JobOfferSelect struct {
	*JobOfferQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (jos *JobOfferSelect) Aggregate(fns ...AggregateFunc) *JobOfferSelect {
	jos.fns = append(jos.fns, fns...)
	return jos
}

// Scan applies the selector query and scans the result into the given value.
func (jos *JobOfferSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, jos.ctx, "Select")
	if err := jos.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*JobOfferQuery, *JobOfferSelect](ctx, jos.JobOfferQuery, jos, jos.inters, v)
}

func (jos *JobOfferSelect) sqlScan(ctx context.Context, root *JobOfferQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(jos.fns))
	for _, fn := range jos.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*jos.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := jos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
