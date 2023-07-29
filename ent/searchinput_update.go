// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/marianozunino/go-utu-llamados/ent/joboffer"
	"github.com/marianozunino/go-utu-llamados/ent/predicate"
	"github.com/marianozunino/go-utu-llamados/ent/searchinput"
)

// SearchInputUpdate is the builder for updating SearchInput entities.
type SearchInputUpdate struct {
	config
	hooks    []Hook
	mutation *SearchInputMutation
}

// Where appends a list predicates to the SearchInputUpdate builder.
func (siu *SearchInputUpdate) Where(ps ...predicate.SearchInput) *SearchInputUpdate {
	siu.mutation.Where(ps...)
	return siu
}

// SetSearch sets the "search" field.
func (siu *SearchInputUpdate) SetSearch(s string) *SearchInputUpdate {
	siu.mutation.SetSearch(s)
	return siu
}

// AddSearchResultIDs adds the "search_results" edge to the JobOffer entity by IDs.
func (siu *SearchInputUpdate) AddSearchResultIDs(ids ...int) *SearchInputUpdate {
	siu.mutation.AddSearchResultIDs(ids...)
	return siu
}

// AddSearchResults adds the "search_results" edges to the JobOffer entity.
func (siu *SearchInputUpdate) AddSearchResults(j ...*JobOffer) *SearchInputUpdate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return siu.AddSearchResultIDs(ids...)
}

// Mutation returns the SearchInputMutation object of the builder.
func (siu *SearchInputUpdate) Mutation() *SearchInputMutation {
	return siu.mutation
}

// ClearSearchResults clears all "search_results" edges to the JobOffer entity.
func (siu *SearchInputUpdate) ClearSearchResults() *SearchInputUpdate {
	siu.mutation.ClearSearchResults()
	return siu
}

// RemoveSearchResultIDs removes the "search_results" edge to JobOffer entities by IDs.
func (siu *SearchInputUpdate) RemoveSearchResultIDs(ids ...int) *SearchInputUpdate {
	siu.mutation.RemoveSearchResultIDs(ids...)
	return siu
}

// RemoveSearchResults removes "search_results" edges to JobOffer entities.
func (siu *SearchInputUpdate) RemoveSearchResults(j ...*JobOffer) *SearchInputUpdate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return siu.RemoveSearchResultIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (siu *SearchInputUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, siu.sqlSave, siu.mutation, siu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (siu *SearchInputUpdate) SaveX(ctx context.Context) int {
	affected, err := siu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (siu *SearchInputUpdate) Exec(ctx context.Context) error {
	_, err := siu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (siu *SearchInputUpdate) ExecX(ctx context.Context) {
	if err := siu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (siu *SearchInputUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(searchinput.Table, searchinput.Columns, sqlgraph.NewFieldSpec(searchinput.FieldID, field.TypeInt))
	if ps := siu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := siu.mutation.Search(); ok {
		_spec.SetField(searchinput.FieldSearch, field.TypeString, value)
	}
	if siu.mutation.SearchResultsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   searchinput.SearchResultsTable,
			Columns: searchinput.SearchResultsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(joboffer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := siu.mutation.RemovedSearchResultsIDs(); len(nodes) > 0 && !siu.mutation.SearchResultsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   searchinput.SearchResultsTable,
			Columns: searchinput.SearchResultsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(joboffer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := siu.mutation.SearchResultsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   searchinput.SearchResultsTable,
			Columns: searchinput.SearchResultsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(joboffer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, siu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{searchinput.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	siu.mutation.done = true
	return n, nil
}

// SearchInputUpdateOne is the builder for updating a single SearchInput entity.
type SearchInputUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SearchInputMutation
}

// SetSearch sets the "search" field.
func (siuo *SearchInputUpdateOne) SetSearch(s string) *SearchInputUpdateOne {
	siuo.mutation.SetSearch(s)
	return siuo
}

// AddSearchResultIDs adds the "search_results" edge to the JobOffer entity by IDs.
func (siuo *SearchInputUpdateOne) AddSearchResultIDs(ids ...int) *SearchInputUpdateOne {
	siuo.mutation.AddSearchResultIDs(ids...)
	return siuo
}

// AddSearchResults adds the "search_results" edges to the JobOffer entity.
func (siuo *SearchInputUpdateOne) AddSearchResults(j ...*JobOffer) *SearchInputUpdateOne {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return siuo.AddSearchResultIDs(ids...)
}

// Mutation returns the SearchInputMutation object of the builder.
func (siuo *SearchInputUpdateOne) Mutation() *SearchInputMutation {
	return siuo.mutation
}

// ClearSearchResults clears all "search_results" edges to the JobOffer entity.
func (siuo *SearchInputUpdateOne) ClearSearchResults() *SearchInputUpdateOne {
	siuo.mutation.ClearSearchResults()
	return siuo
}

// RemoveSearchResultIDs removes the "search_results" edge to JobOffer entities by IDs.
func (siuo *SearchInputUpdateOne) RemoveSearchResultIDs(ids ...int) *SearchInputUpdateOne {
	siuo.mutation.RemoveSearchResultIDs(ids...)
	return siuo
}

// RemoveSearchResults removes "search_results" edges to JobOffer entities.
func (siuo *SearchInputUpdateOne) RemoveSearchResults(j ...*JobOffer) *SearchInputUpdateOne {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return siuo.RemoveSearchResultIDs(ids...)
}

// Where appends a list predicates to the SearchInputUpdate builder.
func (siuo *SearchInputUpdateOne) Where(ps ...predicate.SearchInput) *SearchInputUpdateOne {
	siuo.mutation.Where(ps...)
	return siuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (siuo *SearchInputUpdateOne) Select(field string, fields ...string) *SearchInputUpdateOne {
	siuo.fields = append([]string{field}, fields...)
	return siuo
}

// Save executes the query and returns the updated SearchInput entity.
func (siuo *SearchInputUpdateOne) Save(ctx context.Context) (*SearchInput, error) {
	return withHooks(ctx, siuo.sqlSave, siuo.mutation, siuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (siuo *SearchInputUpdateOne) SaveX(ctx context.Context) *SearchInput {
	node, err := siuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (siuo *SearchInputUpdateOne) Exec(ctx context.Context) error {
	_, err := siuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (siuo *SearchInputUpdateOne) ExecX(ctx context.Context) {
	if err := siuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (siuo *SearchInputUpdateOne) sqlSave(ctx context.Context) (_node *SearchInput, err error) {
	_spec := sqlgraph.NewUpdateSpec(searchinput.Table, searchinput.Columns, sqlgraph.NewFieldSpec(searchinput.FieldID, field.TypeInt))
	id, ok := siuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SearchInput.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := siuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, searchinput.FieldID)
		for _, f := range fields {
			if !searchinput.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != searchinput.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := siuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := siuo.mutation.Search(); ok {
		_spec.SetField(searchinput.FieldSearch, field.TypeString, value)
	}
	if siuo.mutation.SearchResultsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   searchinput.SearchResultsTable,
			Columns: searchinput.SearchResultsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(joboffer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := siuo.mutation.RemovedSearchResultsIDs(); len(nodes) > 0 && !siuo.mutation.SearchResultsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   searchinput.SearchResultsTable,
			Columns: searchinput.SearchResultsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(joboffer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := siuo.mutation.SearchResultsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   searchinput.SearchResultsTable,
			Columns: searchinput.SearchResultsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(joboffer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SearchInput{config: siuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, siuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{searchinput.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	siuo.mutation.done = true
	return _node, nil
}