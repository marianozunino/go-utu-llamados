// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/marianozunino/go-utu-llamados/ent/joboffer"
	"github.com/marianozunino/go-utu-llamados/ent/predicate"
	"github.com/marianozunino/go-utu-llamados/ent/searchinput"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeJobOffer    = "JobOffer"
	TypeSearchInput = "SearchInput"
)

// JobOfferMutation represents an operation that mutates the JobOffer nodes in the graph.
type JobOfferMutation struct {
	config
	op                   Op
	typ                  string
	id                   *int
	title                *string
	description          *string
	url                  *string
	file_url             *string
	published_at         *time.Time
	updated_at           *time.Time
	clearedFields        map[string]struct{}
	search_inputs        map[int]struct{}
	removedsearch_inputs map[int]struct{}
	clearedsearch_inputs bool
	done                 bool
	oldValue             func(context.Context) (*JobOffer, error)
	predicates           []predicate.JobOffer
}

var _ ent.Mutation = (*JobOfferMutation)(nil)

// jobofferOption allows management of the mutation configuration using functional options.
type jobofferOption func(*JobOfferMutation)

// newJobOfferMutation creates new mutation for the JobOffer entity.
func newJobOfferMutation(c config, op Op, opts ...jobofferOption) *JobOfferMutation {
	m := &JobOfferMutation{
		config:        c,
		op:            op,
		typ:           TypeJobOffer,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withJobOfferID sets the ID field of the mutation.
func withJobOfferID(id int) jobofferOption {
	return func(m *JobOfferMutation) {
		var (
			err   error
			once  sync.Once
			value *JobOffer
		)
		m.oldValue = func(ctx context.Context) (*JobOffer, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().JobOffer.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withJobOffer sets the old JobOffer of the mutation.
func withJobOffer(node *JobOffer) jobofferOption {
	return func(m *JobOfferMutation) {
		m.oldValue = func(context.Context) (*JobOffer, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m JobOfferMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m JobOfferMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *JobOfferMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *JobOfferMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().JobOffer.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTitle sets the "title" field.
func (m *JobOfferMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *JobOfferMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the JobOffer entity.
// If the JobOffer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *JobOfferMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *JobOfferMutation) ResetTitle() {
	m.title = nil
}

// SetDescription sets the "description" field.
func (m *JobOfferMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *JobOfferMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the JobOffer entity.
// If the JobOffer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *JobOfferMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ResetDescription resets all changes to the "description" field.
func (m *JobOfferMutation) ResetDescription() {
	m.description = nil
}

// SetURL sets the "url" field.
func (m *JobOfferMutation) SetURL(s string) {
	m.url = &s
}

// URL returns the value of the "url" field in the mutation.
func (m *JobOfferMutation) URL() (r string, exists bool) {
	v := m.url
	if v == nil {
		return
	}
	return *v, true
}

// OldURL returns the old "url" field's value of the JobOffer entity.
// If the JobOffer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *JobOfferMutation) OldURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldURL: %w", err)
	}
	return oldValue.URL, nil
}

// ResetURL resets all changes to the "url" field.
func (m *JobOfferMutation) ResetURL() {
	m.url = nil
}

// SetFileURL sets the "file_url" field.
func (m *JobOfferMutation) SetFileURL(s string) {
	m.file_url = &s
}

// FileURL returns the value of the "file_url" field in the mutation.
func (m *JobOfferMutation) FileURL() (r string, exists bool) {
	v := m.file_url
	if v == nil {
		return
	}
	return *v, true
}

// OldFileURL returns the old "file_url" field's value of the JobOffer entity.
// If the JobOffer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *JobOfferMutation) OldFileURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFileURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFileURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFileURL: %w", err)
	}
	return oldValue.FileURL, nil
}

// ResetFileURL resets all changes to the "file_url" field.
func (m *JobOfferMutation) ResetFileURL() {
	m.file_url = nil
}

// SetPublishedAt sets the "published_at" field.
func (m *JobOfferMutation) SetPublishedAt(t time.Time) {
	m.published_at = &t
}

// PublishedAt returns the value of the "published_at" field in the mutation.
func (m *JobOfferMutation) PublishedAt() (r time.Time, exists bool) {
	v := m.published_at
	if v == nil {
		return
	}
	return *v, true
}

// OldPublishedAt returns the old "published_at" field's value of the JobOffer entity.
// If the JobOffer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *JobOfferMutation) OldPublishedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPublishedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPublishedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPublishedAt: %w", err)
	}
	return oldValue.PublishedAt, nil
}

// ResetPublishedAt resets all changes to the "published_at" field.
func (m *JobOfferMutation) ResetPublishedAt() {
	m.published_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *JobOfferMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *JobOfferMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the JobOffer entity.
// If the JobOffer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *JobOfferMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *JobOfferMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// AddSearchInputIDs adds the "search_inputs" edge to the SearchInput entity by ids.
func (m *JobOfferMutation) AddSearchInputIDs(ids ...int) {
	if m.search_inputs == nil {
		m.search_inputs = make(map[int]struct{})
	}
	for i := range ids {
		m.search_inputs[ids[i]] = struct{}{}
	}
}

// ClearSearchInputs clears the "search_inputs" edge to the SearchInput entity.
func (m *JobOfferMutation) ClearSearchInputs() {
	m.clearedsearch_inputs = true
}

// SearchInputsCleared reports if the "search_inputs" edge to the SearchInput entity was cleared.
func (m *JobOfferMutation) SearchInputsCleared() bool {
	return m.clearedsearch_inputs
}

// RemoveSearchInputIDs removes the "search_inputs" edge to the SearchInput entity by IDs.
func (m *JobOfferMutation) RemoveSearchInputIDs(ids ...int) {
	if m.removedsearch_inputs == nil {
		m.removedsearch_inputs = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.search_inputs, ids[i])
		m.removedsearch_inputs[ids[i]] = struct{}{}
	}
}

// RemovedSearchInputs returns the removed IDs of the "search_inputs" edge to the SearchInput entity.
func (m *JobOfferMutation) RemovedSearchInputsIDs() (ids []int) {
	for id := range m.removedsearch_inputs {
		ids = append(ids, id)
	}
	return
}

// SearchInputsIDs returns the "search_inputs" edge IDs in the mutation.
func (m *JobOfferMutation) SearchInputsIDs() (ids []int) {
	for id := range m.search_inputs {
		ids = append(ids, id)
	}
	return
}

// ResetSearchInputs resets all changes to the "search_inputs" edge.
func (m *JobOfferMutation) ResetSearchInputs() {
	m.search_inputs = nil
	m.clearedsearch_inputs = false
	m.removedsearch_inputs = nil
}

// Where appends a list predicates to the JobOfferMutation builder.
func (m *JobOfferMutation) Where(ps ...predicate.JobOffer) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the JobOfferMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *JobOfferMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.JobOffer, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *JobOfferMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *JobOfferMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (JobOffer).
func (m *JobOfferMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *JobOfferMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.title != nil {
		fields = append(fields, joboffer.FieldTitle)
	}
	if m.description != nil {
		fields = append(fields, joboffer.FieldDescription)
	}
	if m.url != nil {
		fields = append(fields, joboffer.FieldURL)
	}
	if m.file_url != nil {
		fields = append(fields, joboffer.FieldFileURL)
	}
	if m.published_at != nil {
		fields = append(fields, joboffer.FieldPublishedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, joboffer.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *JobOfferMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case joboffer.FieldTitle:
		return m.Title()
	case joboffer.FieldDescription:
		return m.Description()
	case joboffer.FieldURL:
		return m.URL()
	case joboffer.FieldFileURL:
		return m.FileURL()
	case joboffer.FieldPublishedAt:
		return m.PublishedAt()
	case joboffer.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *JobOfferMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case joboffer.FieldTitle:
		return m.OldTitle(ctx)
	case joboffer.FieldDescription:
		return m.OldDescription(ctx)
	case joboffer.FieldURL:
		return m.OldURL(ctx)
	case joboffer.FieldFileURL:
		return m.OldFileURL(ctx)
	case joboffer.FieldPublishedAt:
		return m.OldPublishedAt(ctx)
	case joboffer.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown JobOffer field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *JobOfferMutation) SetField(name string, value ent.Value) error {
	switch name {
	case joboffer.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case joboffer.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case joboffer.FieldURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetURL(v)
		return nil
	case joboffer.FieldFileURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFileURL(v)
		return nil
	case joboffer.FieldPublishedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPublishedAt(v)
		return nil
	case joboffer.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown JobOffer field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *JobOfferMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *JobOfferMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *JobOfferMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown JobOffer numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *JobOfferMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *JobOfferMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *JobOfferMutation) ClearField(name string) error {
	return fmt.Errorf("unknown JobOffer nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *JobOfferMutation) ResetField(name string) error {
	switch name {
	case joboffer.FieldTitle:
		m.ResetTitle()
		return nil
	case joboffer.FieldDescription:
		m.ResetDescription()
		return nil
	case joboffer.FieldURL:
		m.ResetURL()
		return nil
	case joboffer.FieldFileURL:
		m.ResetFileURL()
		return nil
	case joboffer.FieldPublishedAt:
		m.ResetPublishedAt()
		return nil
	case joboffer.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown JobOffer field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *JobOfferMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.search_inputs != nil {
		edges = append(edges, joboffer.EdgeSearchInputs)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *JobOfferMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case joboffer.EdgeSearchInputs:
		ids := make([]ent.Value, 0, len(m.search_inputs))
		for id := range m.search_inputs {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *JobOfferMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedsearch_inputs != nil {
		edges = append(edges, joboffer.EdgeSearchInputs)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *JobOfferMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case joboffer.EdgeSearchInputs:
		ids := make([]ent.Value, 0, len(m.removedsearch_inputs))
		for id := range m.removedsearch_inputs {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *JobOfferMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedsearch_inputs {
		edges = append(edges, joboffer.EdgeSearchInputs)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *JobOfferMutation) EdgeCleared(name string) bool {
	switch name {
	case joboffer.EdgeSearchInputs:
		return m.clearedsearch_inputs
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *JobOfferMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown JobOffer unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *JobOfferMutation) ResetEdge(name string) error {
	switch name {
	case joboffer.EdgeSearchInputs:
		m.ResetSearchInputs()
		return nil
	}
	return fmt.Errorf("unknown JobOffer edge %s", name)
}

// SearchInputMutation represents an operation that mutates the SearchInput nodes in the graph.
type SearchInputMutation struct {
	config
	op                    Op
	typ                   string
	id                    *int
	search                *string
	clearedFields         map[string]struct{}
	search_results        map[int]struct{}
	removedsearch_results map[int]struct{}
	clearedsearch_results bool
	done                  bool
	oldValue              func(context.Context) (*SearchInput, error)
	predicates            []predicate.SearchInput
}

var _ ent.Mutation = (*SearchInputMutation)(nil)

// searchinputOption allows management of the mutation configuration using functional options.
type searchinputOption func(*SearchInputMutation)

// newSearchInputMutation creates new mutation for the SearchInput entity.
func newSearchInputMutation(c config, op Op, opts ...searchinputOption) *SearchInputMutation {
	m := &SearchInputMutation{
		config:        c,
		op:            op,
		typ:           TypeSearchInput,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withSearchInputID sets the ID field of the mutation.
func withSearchInputID(id int) searchinputOption {
	return func(m *SearchInputMutation) {
		var (
			err   error
			once  sync.Once
			value *SearchInput
		)
		m.oldValue = func(ctx context.Context) (*SearchInput, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().SearchInput.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withSearchInput sets the old SearchInput of the mutation.
func withSearchInput(node *SearchInput) searchinputOption {
	return func(m *SearchInputMutation) {
		m.oldValue = func(context.Context) (*SearchInput, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m SearchInputMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m SearchInputMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *SearchInputMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *SearchInputMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().SearchInput.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetSearch sets the "search" field.
func (m *SearchInputMutation) SetSearch(s string) {
	m.search = &s
}

// Search returns the value of the "search" field in the mutation.
func (m *SearchInputMutation) Search() (r string, exists bool) {
	v := m.search
	if v == nil {
		return
	}
	return *v, true
}

// OldSearch returns the old "search" field's value of the SearchInput entity.
// If the SearchInput object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SearchInputMutation) OldSearch(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSearch is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSearch requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSearch: %w", err)
	}
	return oldValue.Search, nil
}

// ResetSearch resets all changes to the "search" field.
func (m *SearchInputMutation) ResetSearch() {
	m.search = nil
}

// AddSearchResultIDs adds the "search_results" edge to the JobOffer entity by ids.
func (m *SearchInputMutation) AddSearchResultIDs(ids ...int) {
	if m.search_results == nil {
		m.search_results = make(map[int]struct{})
	}
	for i := range ids {
		m.search_results[ids[i]] = struct{}{}
	}
}

// ClearSearchResults clears the "search_results" edge to the JobOffer entity.
func (m *SearchInputMutation) ClearSearchResults() {
	m.clearedsearch_results = true
}

// SearchResultsCleared reports if the "search_results" edge to the JobOffer entity was cleared.
func (m *SearchInputMutation) SearchResultsCleared() bool {
	return m.clearedsearch_results
}

// RemoveSearchResultIDs removes the "search_results" edge to the JobOffer entity by IDs.
func (m *SearchInputMutation) RemoveSearchResultIDs(ids ...int) {
	if m.removedsearch_results == nil {
		m.removedsearch_results = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.search_results, ids[i])
		m.removedsearch_results[ids[i]] = struct{}{}
	}
}

// RemovedSearchResults returns the removed IDs of the "search_results" edge to the JobOffer entity.
func (m *SearchInputMutation) RemovedSearchResultsIDs() (ids []int) {
	for id := range m.removedsearch_results {
		ids = append(ids, id)
	}
	return
}

// SearchResultsIDs returns the "search_results" edge IDs in the mutation.
func (m *SearchInputMutation) SearchResultsIDs() (ids []int) {
	for id := range m.search_results {
		ids = append(ids, id)
	}
	return
}

// ResetSearchResults resets all changes to the "search_results" edge.
func (m *SearchInputMutation) ResetSearchResults() {
	m.search_results = nil
	m.clearedsearch_results = false
	m.removedsearch_results = nil
}

// Where appends a list predicates to the SearchInputMutation builder.
func (m *SearchInputMutation) Where(ps ...predicate.SearchInput) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the SearchInputMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *SearchInputMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.SearchInput, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *SearchInputMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *SearchInputMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (SearchInput).
func (m *SearchInputMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *SearchInputMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.search != nil {
		fields = append(fields, searchinput.FieldSearch)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *SearchInputMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case searchinput.FieldSearch:
		return m.Search()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *SearchInputMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case searchinput.FieldSearch:
		return m.OldSearch(ctx)
	}
	return nil, fmt.Errorf("unknown SearchInput field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SearchInputMutation) SetField(name string, value ent.Value) error {
	switch name {
	case searchinput.FieldSearch:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSearch(v)
		return nil
	}
	return fmt.Errorf("unknown SearchInput field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *SearchInputMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *SearchInputMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SearchInputMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown SearchInput numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *SearchInputMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *SearchInputMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *SearchInputMutation) ClearField(name string) error {
	return fmt.Errorf("unknown SearchInput nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *SearchInputMutation) ResetField(name string) error {
	switch name {
	case searchinput.FieldSearch:
		m.ResetSearch()
		return nil
	}
	return fmt.Errorf("unknown SearchInput field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *SearchInputMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.search_results != nil {
		edges = append(edges, searchinput.EdgeSearchResults)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *SearchInputMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case searchinput.EdgeSearchResults:
		ids := make([]ent.Value, 0, len(m.search_results))
		for id := range m.search_results {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *SearchInputMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedsearch_results != nil {
		edges = append(edges, searchinput.EdgeSearchResults)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *SearchInputMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case searchinput.EdgeSearchResults:
		ids := make([]ent.Value, 0, len(m.removedsearch_results))
		for id := range m.removedsearch_results {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *SearchInputMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedsearch_results {
		edges = append(edges, searchinput.EdgeSearchResults)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *SearchInputMutation) EdgeCleared(name string) bool {
	switch name {
	case searchinput.EdgeSearchResults:
		return m.clearedsearch_results
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *SearchInputMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown SearchInput unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *SearchInputMutation) ResetEdge(name string) error {
	switch name {
	case searchinput.EdgeSearchResults:
		m.ResetSearchResults()
		return nil
	}
	return fmt.Errorf("unknown SearchInput edge %s", name)
}
