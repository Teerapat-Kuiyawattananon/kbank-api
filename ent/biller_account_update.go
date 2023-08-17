// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kapi/ent/bill"
	"kapi/ent/biller_account"
	"kapi/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BillerAccountUpdate is the builder for updating Biller_account entities.
type BillerAccountUpdate struct {
	config
	hooks    []Hook
	mutation *BillerAccountMutation
}

// Where appends a list predicates to the BillerAccountUpdate builder.
func (bau *BillerAccountUpdate) Where(ps ...predicate.Biller_account) *BillerAccountUpdate {
	bau.mutation.Where(ps...)
	return bau
}

// SetName sets the "name" field.
func (bau *BillerAccountUpdate) SetName(s string) *BillerAccountUpdate {
	bau.mutation.SetName(s)
	return bau
}

// SetNillableName sets the "name" field if the given value is not nil.
func (bau *BillerAccountUpdate) SetNillableName(s *string) *BillerAccountUpdate {
	if s != nil {
		bau.SetName(*s)
	}
	return bau
}

// SetServiceName sets the "service_name" field.
func (bau *BillerAccountUpdate) SetServiceName(s string) *BillerAccountUpdate {
	bau.mutation.SetServiceName(s)
	return bau
}

// SetNillableServiceName sets the "service_name" field if the given value is not nil.
func (bau *BillerAccountUpdate) SetNillableServiceName(s *string) *BillerAccountUpdate {
	if s != nil {
		bau.SetServiceName(*s)
	}
	return bau
}

// AddBillIDs adds the "bills" edge to the Bill entity by IDs.
func (bau *BillerAccountUpdate) AddBillIDs(ids ...int) *BillerAccountUpdate {
	bau.mutation.AddBillIDs(ids...)
	return bau
}

// AddBills adds the "bills" edges to the Bill entity.
func (bau *BillerAccountUpdate) AddBills(b ...*Bill) *BillerAccountUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bau.AddBillIDs(ids...)
}

// Mutation returns the BillerAccountMutation object of the builder.
func (bau *BillerAccountUpdate) Mutation() *BillerAccountMutation {
	return bau.mutation
}

// ClearBills clears all "bills" edges to the Bill entity.
func (bau *BillerAccountUpdate) ClearBills() *BillerAccountUpdate {
	bau.mutation.ClearBills()
	return bau
}

// RemoveBillIDs removes the "bills" edge to Bill entities by IDs.
func (bau *BillerAccountUpdate) RemoveBillIDs(ids ...int) *BillerAccountUpdate {
	bau.mutation.RemoveBillIDs(ids...)
	return bau
}

// RemoveBills removes "bills" edges to Bill entities.
func (bau *BillerAccountUpdate) RemoveBills(b ...*Bill) *BillerAccountUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bau.RemoveBillIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bau *BillerAccountUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bau.sqlSave, bau.mutation, bau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bau *BillerAccountUpdate) SaveX(ctx context.Context) int {
	affected, err := bau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bau *BillerAccountUpdate) Exec(ctx context.Context) error {
	_, err := bau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bau *BillerAccountUpdate) ExecX(ctx context.Context) {
	if err := bau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bau *BillerAccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(biller_account.Table, biller_account.Columns, sqlgraph.NewFieldSpec(biller_account.FieldID, field.TypeInt))
	if ps := bau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bau.mutation.Name(); ok {
		_spec.SetField(biller_account.FieldName, field.TypeString, value)
	}
	if value, ok := bau.mutation.ServiceName(); ok {
		_spec.SetField(biller_account.FieldServiceName, field.TypeString, value)
	}
	if bau.mutation.BillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   biller_account.BillsTable,
			Columns: []string{biller_account.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bau.mutation.RemovedBillsIDs(); len(nodes) > 0 && !bau.mutation.BillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   biller_account.BillsTable,
			Columns: []string{biller_account.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bau.mutation.BillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   biller_account.BillsTable,
			Columns: []string{biller_account.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{biller_account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bau.mutation.done = true
	return n, nil
}

// BillerAccountUpdateOne is the builder for updating a single Biller_account entity.
type BillerAccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BillerAccountMutation
}

// SetName sets the "name" field.
func (bauo *BillerAccountUpdateOne) SetName(s string) *BillerAccountUpdateOne {
	bauo.mutation.SetName(s)
	return bauo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (bauo *BillerAccountUpdateOne) SetNillableName(s *string) *BillerAccountUpdateOne {
	if s != nil {
		bauo.SetName(*s)
	}
	return bauo
}

// SetServiceName sets the "service_name" field.
func (bauo *BillerAccountUpdateOne) SetServiceName(s string) *BillerAccountUpdateOne {
	bauo.mutation.SetServiceName(s)
	return bauo
}

// SetNillableServiceName sets the "service_name" field if the given value is not nil.
func (bauo *BillerAccountUpdateOne) SetNillableServiceName(s *string) *BillerAccountUpdateOne {
	if s != nil {
		bauo.SetServiceName(*s)
	}
	return bauo
}

// AddBillIDs adds the "bills" edge to the Bill entity by IDs.
func (bauo *BillerAccountUpdateOne) AddBillIDs(ids ...int) *BillerAccountUpdateOne {
	bauo.mutation.AddBillIDs(ids...)
	return bauo
}

// AddBills adds the "bills" edges to the Bill entity.
func (bauo *BillerAccountUpdateOne) AddBills(b ...*Bill) *BillerAccountUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bauo.AddBillIDs(ids...)
}

// Mutation returns the BillerAccountMutation object of the builder.
func (bauo *BillerAccountUpdateOne) Mutation() *BillerAccountMutation {
	return bauo.mutation
}

// ClearBills clears all "bills" edges to the Bill entity.
func (bauo *BillerAccountUpdateOne) ClearBills() *BillerAccountUpdateOne {
	bauo.mutation.ClearBills()
	return bauo
}

// RemoveBillIDs removes the "bills" edge to Bill entities by IDs.
func (bauo *BillerAccountUpdateOne) RemoveBillIDs(ids ...int) *BillerAccountUpdateOne {
	bauo.mutation.RemoveBillIDs(ids...)
	return bauo
}

// RemoveBills removes "bills" edges to Bill entities.
func (bauo *BillerAccountUpdateOne) RemoveBills(b ...*Bill) *BillerAccountUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bauo.RemoveBillIDs(ids...)
}

// Where appends a list predicates to the BillerAccountUpdate builder.
func (bauo *BillerAccountUpdateOne) Where(ps ...predicate.Biller_account) *BillerAccountUpdateOne {
	bauo.mutation.Where(ps...)
	return bauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bauo *BillerAccountUpdateOne) Select(field string, fields ...string) *BillerAccountUpdateOne {
	bauo.fields = append([]string{field}, fields...)
	return bauo
}

// Save executes the query and returns the updated Biller_account entity.
func (bauo *BillerAccountUpdateOne) Save(ctx context.Context) (*Biller_account, error) {
	return withHooks(ctx, bauo.sqlSave, bauo.mutation, bauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bauo *BillerAccountUpdateOne) SaveX(ctx context.Context) *Biller_account {
	node, err := bauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bauo *BillerAccountUpdateOne) Exec(ctx context.Context) error {
	_, err := bauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bauo *BillerAccountUpdateOne) ExecX(ctx context.Context) {
	if err := bauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bauo *BillerAccountUpdateOne) sqlSave(ctx context.Context) (_node *Biller_account, err error) {
	_spec := sqlgraph.NewUpdateSpec(biller_account.Table, biller_account.Columns, sqlgraph.NewFieldSpec(biller_account.FieldID, field.TypeInt))
	id, ok := bauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Biller_account.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, biller_account.FieldID)
		for _, f := range fields {
			if !biller_account.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != biller_account.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bauo.mutation.Name(); ok {
		_spec.SetField(biller_account.FieldName, field.TypeString, value)
	}
	if value, ok := bauo.mutation.ServiceName(); ok {
		_spec.SetField(biller_account.FieldServiceName, field.TypeString, value)
	}
	if bauo.mutation.BillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   biller_account.BillsTable,
			Columns: []string{biller_account.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bauo.mutation.RemovedBillsIDs(); len(nodes) > 0 && !bauo.mutation.BillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   biller_account.BillsTable,
			Columns: []string{biller_account.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bauo.mutation.BillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   biller_account.BillsTable,
			Columns: []string{biller_account.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Biller_account{config: bauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{biller_account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bauo.mutation.done = true
	return _node, nil
}