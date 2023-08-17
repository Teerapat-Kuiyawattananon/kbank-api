// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"kapi/ent/bill"
	"kapi/ent/biller_account"
	"kapi/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BillerAccountQuery is the builder for querying Biller_account entities.
type BillerAccountQuery struct {
	config
	ctx        *QueryContext
	order      []biller_account.OrderOption
	inters     []Interceptor
	predicates []predicate.Biller_account
	withBills  *BillQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BillerAccountQuery builder.
func (baq *BillerAccountQuery) Where(ps ...predicate.Biller_account) *BillerAccountQuery {
	baq.predicates = append(baq.predicates, ps...)
	return baq
}

// Limit the number of records to be returned by this query.
func (baq *BillerAccountQuery) Limit(limit int) *BillerAccountQuery {
	baq.ctx.Limit = &limit
	return baq
}

// Offset to start from.
func (baq *BillerAccountQuery) Offset(offset int) *BillerAccountQuery {
	baq.ctx.Offset = &offset
	return baq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (baq *BillerAccountQuery) Unique(unique bool) *BillerAccountQuery {
	baq.ctx.Unique = &unique
	return baq
}

// Order specifies how the records should be ordered.
func (baq *BillerAccountQuery) Order(o ...biller_account.OrderOption) *BillerAccountQuery {
	baq.order = append(baq.order, o...)
	return baq
}

// QueryBills chains the current query on the "bills" edge.
func (baq *BillerAccountQuery) QueryBills() *BillQuery {
	query := (&BillClient{config: baq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := baq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := baq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(biller_account.Table, biller_account.FieldID, selector),
			sqlgraph.To(bill.Table, bill.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, biller_account.BillsTable, biller_account.BillsColumn),
		)
		fromU = sqlgraph.SetNeighbors(baq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Biller_account entity from the query.
// Returns a *NotFoundError when no Biller_account was found.
func (baq *BillerAccountQuery) First(ctx context.Context) (*Biller_account, error) {
	nodes, err := baq.Limit(1).All(setContextOp(ctx, baq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{biller_account.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (baq *BillerAccountQuery) FirstX(ctx context.Context) *Biller_account {
	node, err := baq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Biller_account ID from the query.
// Returns a *NotFoundError when no Biller_account ID was found.
func (baq *BillerAccountQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = baq.Limit(1).IDs(setContextOp(ctx, baq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{biller_account.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (baq *BillerAccountQuery) FirstIDX(ctx context.Context) int {
	id, err := baq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Biller_account entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Biller_account entity is found.
// Returns a *NotFoundError when no Biller_account entities are found.
func (baq *BillerAccountQuery) Only(ctx context.Context) (*Biller_account, error) {
	nodes, err := baq.Limit(2).All(setContextOp(ctx, baq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{biller_account.Label}
	default:
		return nil, &NotSingularError{biller_account.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (baq *BillerAccountQuery) OnlyX(ctx context.Context) *Biller_account {
	node, err := baq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Biller_account ID in the query.
// Returns a *NotSingularError when more than one Biller_account ID is found.
// Returns a *NotFoundError when no entities are found.
func (baq *BillerAccountQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = baq.Limit(2).IDs(setContextOp(ctx, baq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{biller_account.Label}
	default:
		err = &NotSingularError{biller_account.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (baq *BillerAccountQuery) OnlyIDX(ctx context.Context) int {
	id, err := baq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Biller_accounts.
func (baq *BillerAccountQuery) All(ctx context.Context) ([]*Biller_account, error) {
	ctx = setContextOp(ctx, baq.ctx, "All")
	if err := baq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Biller_account, *BillerAccountQuery]()
	return withInterceptors[[]*Biller_account](ctx, baq, qr, baq.inters)
}

// AllX is like All, but panics if an error occurs.
func (baq *BillerAccountQuery) AllX(ctx context.Context) []*Biller_account {
	nodes, err := baq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Biller_account IDs.
func (baq *BillerAccountQuery) IDs(ctx context.Context) (ids []int, err error) {
	if baq.ctx.Unique == nil && baq.path != nil {
		baq.Unique(true)
	}
	ctx = setContextOp(ctx, baq.ctx, "IDs")
	if err = baq.Select(biller_account.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (baq *BillerAccountQuery) IDsX(ctx context.Context) []int {
	ids, err := baq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (baq *BillerAccountQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, baq.ctx, "Count")
	if err := baq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, baq, querierCount[*BillerAccountQuery](), baq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (baq *BillerAccountQuery) CountX(ctx context.Context) int {
	count, err := baq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (baq *BillerAccountQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, baq.ctx, "Exist")
	switch _, err := baq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (baq *BillerAccountQuery) ExistX(ctx context.Context) bool {
	exist, err := baq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BillerAccountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (baq *BillerAccountQuery) Clone() *BillerAccountQuery {
	if baq == nil {
		return nil
	}
	return &BillerAccountQuery{
		config:     baq.config,
		ctx:        baq.ctx.Clone(),
		order:      append([]biller_account.OrderOption{}, baq.order...),
		inters:     append([]Interceptor{}, baq.inters...),
		predicates: append([]predicate.Biller_account{}, baq.predicates...),
		withBills:  baq.withBills.Clone(),
		// clone intermediate query.
		sql:  baq.sql.Clone(),
		path: baq.path,
	}
}

// WithBills tells the query-builder to eager-load the nodes that are connected to
// the "bills" edge. The optional arguments are used to configure the query builder of the edge.
func (baq *BillerAccountQuery) WithBills(opts ...func(*BillQuery)) *BillerAccountQuery {
	query := (&BillClient{config: baq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	baq.withBills = query
	return baq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BillerAccount.Query().
//		GroupBy(biller_account.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (baq *BillerAccountQuery) GroupBy(field string, fields ...string) *BillerAccountGroupBy {
	baq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BillerAccountGroupBy{build: baq}
	grbuild.flds = &baq.ctx.Fields
	grbuild.label = biller_account.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.BillerAccount.Query().
//		Select(biller_account.FieldName).
//		Scan(ctx, &v)
func (baq *BillerAccountQuery) Select(fields ...string) *BillerAccountSelect {
	baq.ctx.Fields = append(baq.ctx.Fields, fields...)
	sbuild := &BillerAccountSelect{BillerAccountQuery: baq}
	sbuild.label = biller_account.Label
	sbuild.flds, sbuild.scan = &baq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BillerAccountSelect configured with the given aggregations.
func (baq *BillerAccountQuery) Aggregate(fns ...AggregateFunc) *BillerAccountSelect {
	return baq.Select().Aggregate(fns...)
}

func (baq *BillerAccountQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range baq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, baq); err != nil {
				return err
			}
		}
	}
	for _, f := range baq.ctx.Fields {
		if !biller_account.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if baq.path != nil {
		prev, err := baq.path(ctx)
		if err != nil {
			return err
		}
		baq.sql = prev
	}
	return nil
}

func (baq *BillerAccountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Biller_account, error) {
	var (
		nodes       = []*Biller_account{}
		_spec       = baq.querySpec()
		loadedTypes = [1]bool{
			baq.withBills != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Biller_account).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Biller_account{config: baq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, baq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := baq.withBills; query != nil {
		if err := baq.loadBills(ctx, query, nodes,
			func(n *Biller_account) { n.Edges.Bills = []*Bill{} },
			func(n *Biller_account, e *Bill) { n.Edges.Bills = append(n.Edges.Bills, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (baq *BillerAccountQuery) loadBills(ctx context.Context, query *BillQuery, nodes []*Biller_account, init func(*Biller_account), assign func(*Biller_account, *Bill)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Biller_account)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(bill.FieldBillerID)
	}
	query.Where(predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(biller_account.BillsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.BillerID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "biller_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (baq *BillerAccountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := baq.querySpec()
	_spec.Node.Columns = baq.ctx.Fields
	if len(baq.ctx.Fields) > 0 {
		_spec.Unique = baq.ctx.Unique != nil && *baq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, baq.driver, _spec)
}

func (baq *BillerAccountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(biller_account.Table, biller_account.Columns, sqlgraph.NewFieldSpec(biller_account.FieldID, field.TypeInt))
	_spec.From = baq.sql
	if unique := baq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if baq.path != nil {
		_spec.Unique = true
	}
	if fields := baq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, biller_account.FieldID)
		for i := range fields {
			if fields[i] != biller_account.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := baq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := baq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := baq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := baq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (baq *BillerAccountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(baq.driver.Dialect())
	t1 := builder.Table(biller_account.Table)
	columns := baq.ctx.Fields
	if len(columns) == 0 {
		columns = biller_account.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if baq.sql != nil {
		selector = baq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if baq.ctx.Unique != nil && *baq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range baq.predicates {
		p(selector)
	}
	for _, p := range baq.order {
		p(selector)
	}
	if offset := baq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := baq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BillerAccountGroupBy is the group-by builder for Biller_account entities.
type BillerAccountGroupBy struct {
	selector
	build *BillerAccountQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bagb *BillerAccountGroupBy) Aggregate(fns ...AggregateFunc) *BillerAccountGroupBy {
	bagb.fns = append(bagb.fns, fns...)
	return bagb
}

// Scan applies the selector query and scans the result into the given value.
func (bagb *BillerAccountGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bagb.build.ctx, "GroupBy")
	if err := bagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BillerAccountQuery, *BillerAccountGroupBy](ctx, bagb.build, bagb, bagb.build.inters, v)
}

func (bagb *BillerAccountGroupBy) sqlScan(ctx context.Context, root *BillerAccountQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bagb.fns))
	for _, fn := range bagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bagb.flds)+len(bagb.fns))
		for _, f := range *bagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BillerAccountSelect is the builder for selecting fields of BillerAccount entities.
type BillerAccountSelect struct {
	*BillerAccountQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bas *BillerAccountSelect) Aggregate(fns ...AggregateFunc) *BillerAccountSelect {
	bas.fns = append(bas.fns, fns...)
	return bas
}

// Scan applies the selector query and scans the result into the given value.
func (bas *BillerAccountSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bas.ctx, "Select")
	if err := bas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BillerAccountQuery, *BillerAccountSelect](ctx, bas.BillerAccountQuery, bas, bas.inters, v)
}

func (bas *BillerAccountSelect) sqlScan(ctx context.Context, root *BillerAccountQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bas.fns))
	for _, fn := range bas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}