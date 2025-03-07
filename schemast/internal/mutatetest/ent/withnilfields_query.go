// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/contrib/schemast/internal/mutatetest/ent/predicate"
	"entgo.io/contrib/schemast/internal/mutatetest/ent/withnilfields"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WithNilFieldsQuery is the builder for querying WithNilFields entities.
type WithNilFieldsQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.WithNilFields
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WithNilFieldsQuery builder.
func (wnfq *WithNilFieldsQuery) Where(ps ...predicate.WithNilFields) *WithNilFieldsQuery {
	wnfq.predicates = append(wnfq.predicates, ps...)
	return wnfq
}

// Limit adds a limit step to the query.
func (wnfq *WithNilFieldsQuery) Limit(limit int) *WithNilFieldsQuery {
	wnfq.limit = &limit
	return wnfq
}

// Offset adds an offset step to the query.
func (wnfq *WithNilFieldsQuery) Offset(offset int) *WithNilFieldsQuery {
	wnfq.offset = &offset
	return wnfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wnfq *WithNilFieldsQuery) Unique(unique bool) *WithNilFieldsQuery {
	wnfq.unique = &unique
	return wnfq
}

// Order adds an order step to the query.
func (wnfq *WithNilFieldsQuery) Order(o ...OrderFunc) *WithNilFieldsQuery {
	wnfq.order = append(wnfq.order, o...)
	return wnfq
}

// First returns the first WithNilFields entity from the query.
// Returns a *NotFoundError when no WithNilFields was found.
func (wnfq *WithNilFieldsQuery) First(ctx context.Context) (*WithNilFields, error) {
	nodes, err := wnfq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{withnilfields.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wnfq *WithNilFieldsQuery) FirstX(ctx context.Context) *WithNilFields {
	node, err := wnfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WithNilFields ID from the query.
// Returns a *NotFoundError when no WithNilFields ID was found.
func (wnfq *WithNilFieldsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wnfq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{withnilfields.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wnfq *WithNilFieldsQuery) FirstIDX(ctx context.Context) int {
	id, err := wnfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WithNilFields entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WithNilFields entity is found.
// Returns a *NotFoundError when no WithNilFields entities are found.
func (wnfq *WithNilFieldsQuery) Only(ctx context.Context) (*WithNilFields, error) {
	nodes, err := wnfq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{withnilfields.Label}
	default:
		return nil, &NotSingularError{withnilfields.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wnfq *WithNilFieldsQuery) OnlyX(ctx context.Context) *WithNilFields {
	node, err := wnfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WithNilFields ID in the query.
// Returns a *NotSingularError when more than one WithNilFields ID is found.
// Returns a *NotFoundError when no entities are found.
func (wnfq *WithNilFieldsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wnfq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{withnilfields.Label}
	default:
		err = &NotSingularError{withnilfields.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wnfq *WithNilFieldsQuery) OnlyIDX(ctx context.Context) int {
	id, err := wnfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WithNilFieldsSlice.
func (wnfq *WithNilFieldsQuery) All(ctx context.Context) ([]*WithNilFields, error) {
	if err := wnfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wnfq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wnfq *WithNilFieldsQuery) AllX(ctx context.Context) []*WithNilFields {
	nodes, err := wnfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WithNilFields IDs.
func (wnfq *WithNilFieldsQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := wnfq.Select(withnilfields.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wnfq *WithNilFieldsQuery) IDsX(ctx context.Context) []int {
	ids, err := wnfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wnfq *WithNilFieldsQuery) Count(ctx context.Context) (int, error) {
	if err := wnfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wnfq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wnfq *WithNilFieldsQuery) CountX(ctx context.Context) int {
	count, err := wnfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wnfq *WithNilFieldsQuery) Exist(ctx context.Context) (bool, error) {
	if err := wnfq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wnfq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wnfq *WithNilFieldsQuery) ExistX(ctx context.Context) bool {
	exist, err := wnfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WithNilFieldsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wnfq *WithNilFieldsQuery) Clone() *WithNilFieldsQuery {
	if wnfq == nil {
		return nil
	}
	return &WithNilFieldsQuery{
		config:     wnfq.config,
		limit:      wnfq.limit,
		offset:     wnfq.offset,
		order:      append([]OrderFunc{}, wnfq.order...),
		predicates: append([]predicate.WithNilFields{}, wnfq.predicates...),
		// clone intermediate query.
		sql:    wnfq.sql.Clone(),
		path:   wnfq.path,
		unique: wnfq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (wnfq *WithNilFieldsQuery) GroupBy(field string, fields ...string) *WithNilFieldsGroupBy {
	grbuild := &WithNilFieldsGroupBy{config: wnfq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wnfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wnfq.sqlQuery(ctx), nil
	}
	grbuild.label = withnilfields.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (wnfq *WithNilFieldsQuery) Select(fields ...string) *WithNilFieldsSelect {
	wnfq.fields = append(wnfq.fields, fields...)
	selbuild := &WithNilFieldsSelect{WithNilFieldsQuery: wnfq}
	selbuild.label = withnilfields.Label
	selbuild.flds, selbuild.scan = &wnfq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a WithNilFieldsSelect configured with the given aggregations.
func (wnfq *WithNilFieldsQuery) Aggregate(fns ...AggregateFunc) *WithNilFieldsSelect {
	return wnfq.Select().Aggregate(fns...)
}

func (wnfq *WithNilFieldsQuery) prepareQuery(ctx context.Context) error {
	for _, f := range wnfq.fields {
		if !withnilfields.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wnfq.path != nil {
		prev, err := wnfq.path(ctx)
		if err != nil {
			return err
		}
		wnfq.sql = prev
	}
	return nil
}

func (wnfq *WithNilFieldsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WithNilFields, error) {
	var (
		nodes = []*WithNilFields{}
		_spec = wnfq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WithNilFields).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WithNilFields{config: wnfq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wnfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (wnfq *WithNilFieldsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wnfq.querySpec()
	_spec.Node.Columns = wnfq.fields
	if len(wnfq.fields) > 0 {
		_spec.Unique = wnfq.unique != nil && *wnfq.unique
	}
	return sqlgraph.CountNodes(ctx, wnfq.driver, _spec)
}

func (wnfq *WithNilFieldsQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := wnfq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (wnfq *WithNilFieldsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   withnilfields.Table,
			Columns: withnilfields.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: withnilfields.FieldID,
			},
		},
		From:   wnfq.sql,
		Unique: true,
	}
	if unique := wnfq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := wnfq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, withnilfields.FieldID)
		for i := range fields {
			if fields[i] != withnilfields.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wnfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wnfq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wnfq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wnfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wnfq *WithNilFieldsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wnfq.driver.Dialect())
	t1 := builder.Table(withnilfields.Table)
	columns := wnfq.fields
	if len(columns) == 0 {
		columns = withnilfields.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wnfq.sql != nil {
		selector = wnfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wnfq.unique != nil && *wnfq.unique {
		selector.Distinct()
	}
	for _, p := range wnfq.predicates {
		p(selector)
	}
	for _, p := range wnfq.order {
		p(selector)
	}
	if offset := wnfq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wnfq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNilFieldsGroupBy is the group-by builder for WithNilFields entities.
type WithNilFieldsGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wnfgb *WithNilFieldsGroupBy) Aggregate(fns ...AggregateFunc) *WithNilFieldsGroupBy {
	wnfgb.fns = append(wnfgb.fns, fns...)
	return wnfgb
}

// Scan applies the group-by query and scans the result into the given value.
func (wnfgb *WithNilFieldsGroupBy) Scan(ctx context.Context, v any) error {
	query, err := wnfgb.path(ctx)
	if err != nil {
		return err
	}
	wnfgb.sql = query
	return wnfgb.sqlScan(ctx, v)
}

func (wnfgb *WithNilFieldsGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range wnfgb.fields {
		if !withnilfields.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := wnfgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wnfgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wnfgb *WithNilFieldsGroupBy) sqlQuery() *sql.Selector {
	selector := wnfgb.sql.Select()
	aggregation := make([]string, 0, len(wnfgb.fns))
	for _, fn := range wnfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(wnfgb.fields)+len(wnfgb.fns))
		for _, f := range wnfgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(wnfgb.fields...)...)
}

// WithNilFieldsSelect is the builder for selecting fields of WithNilFields entities.
type WithNilFieldsSelect struct {
	*WithNilFieldsQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wnfs *WithNilFieldsSelect) Aggregate(fns ...AggregateFunc) *WithNilFieldsSelect {
	wnfs.fns = append(wnfs.fns, fns...)
	return wnfs
}

// Scan applies the selector query and scans the result into the given value.
func (wnfs *WithNilFieldsSelect) Scan(ctx context.Context, v any) error {
	if err := wnfs.prepareQuery(ctx); err != nil {
		return err
	}
	wnfs.sql = wnfs.WithNilFieldsQuery.sqlQuery(ctx)
	return wnfs.sqlScan(ctx, v)
}

func (wnfs *WithNilFieldsSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(wnfs.fns))
	for _, fn := range wnfs.fns {
		aggregation = append(aggregation, fn(wnfs.sql))
	}
	switch n := len(*wnfs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		wnfs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		wnfs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := wnfs.sql.Query()
	if err := wnfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
