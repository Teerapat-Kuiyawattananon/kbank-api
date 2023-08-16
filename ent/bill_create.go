// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"kapi/ent/bill"
	"kapi/ent/billdetail"
	"kapi/ent/customer"
	"kapi/ent/store"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BillCreate is the builder for creating a Bill entity.
type BillCreate struct {
	config
	mutation *BillMutation
	hooks    []Hook
}

// SetBillerID sets the "biller_id" field.
func (bc *BillCreate) SetBillerID(i int) *BillCreate {
	bc.mutation.SetBillerID(i)
	return bc
}

// SetNillableBillerID sets the "biller_id" field if the given value is not nil.
func (bc *BillCreate) SetNillableBillerID(i *int) *BillCreate {
	if i != nil {
		bc.SetBillerID(*i)
	}
	return bc
}

// SetRef1 sets the "ref_1" field.
func (bc *BillCreate) SetRef1(i int) *BillCreate {
	bc.mutation.SetRef1(i)
	return bc
}

// SetNillableRef1 sets the "ref_1" field if the given value is not nil.
func (bc *BillCreate) SetNillableRef1(i *int) *BillCreate {
	if i != nil {
		bc.SetRef1(*i)
	}
	return bc
}

// SetRef2 sets the "ref_2" field.
func (bc *BillCreate) SetRef2(i int) *BillCreate {
	bc.mutation.SetRef2(i)
	return bc
}

// SetNillableRef2 sets the "ref_2" field if the given value is not nil.
func (bc *BillCreate) SetNillableRef2(i *int) *BillCreate {
	if i != nil {
		bc.SetRef2(*i)
	}
	return bc
}

// SetID sets the "id" field.
func (bc *BillCreate) SetID(i int) *BillCreate {
	bc.mutation.SetID(i)
	return bc
}

// SetStoreID sets the "store" edge to the Store entity by ID.
func (bc *BillCreate) SetStoreID(id int) *BillCreate {
	bc.mutation.SetStoreID(id)
	return bc
}

// SetNillableStoreID sets the "store" edge to the Store entity by ID if the given value is not nil.
func (bc *BillCreate) SetNillableStoreID(id *int) *BillCreate {
	if id != nil {
		bc = bc.SetStoreID(*id)
	}
	return bc
}

// SetStore sets the "store" edge to the Store entity.
func (bc *BillCreate) SetStore(s *Store) *BillCreate {
	return bc.SetStoreID(s.ID)
}

// SetCustomersID sets the "customers" edge to the Customer entity by ID.
func (bc *BillCreate) SetCustomersID(id int) *BillCreate {
	bc.mutation.SetCustomersID(id)
	return bc
}

// SetNillableCustomersID sets the "customers" edge to the Customer entity by ID if the given value is not nil.
func (bc *BillCreate) SetNillableCustomersID(id *int) *BillCreate {
	if id != nil {
		bc = bc.SetCustomersID(*id)
	}
	return bc
}

// SetCustomers sets the "customers" edge to the Customer entity.
func (bc *BillCreate) SetCustomers(c *Customer) *BillCreate {
	return bc.SetCustomersID(c.ID)
}

// SetBillDetailID sets the "bill_detail" edge to the BillDetail entity by ID.
func (bc *BillCreate) SetBillDetailID(id int) *BillCreate {
	bc.mutation.SetBillDetailID(id)
	return bc
}

// SetNillableBillDetailID sets the "bill_detail" edge to the BillDetail entity by ID if the given value is not nil.
func (bc *BillCreate) SetNillableBillDetailID(id *int) *BillCreate {
	if id != nil {
		bc = bc.SetBillDetailID(*id)
	}
	return bc
}

// SetBillDetail sets the "bill_detail" edge to the BillDetail entity.
func (bc *BillCreate) SetBillDetail(b *BillDetail) *BillCreate {
	return bc.SetBillDetailID(b.ID)
}

// Mutation returns the BillMutation object of the builder.
func (bc *BillCreate) Mutation() *BillMutation {
	return bc.mutation
}

// Save creates the Bill in the database.
func (bc *BillCreate) Save(ctx context.Context) (*Bill, error) {
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BillCreate) SaveX(ctx context.Context) *Bill {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BillCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BillCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BillCreate) check() error {
	return nil
}

func (bc *BillCreate) sqlSave(ctx context.Context) (*Bill, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BillCreate) createSpec() (*Bill, *sqlgraph.CreateSpec) {
	var (
		_node = &Bill{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(bill.Table, sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt))
	)
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bc.mutation.Ref2(); ok {
		_spec.SetField(bill.FieldRef2, field.TypeInt, value)
		_node.Ref2 = value
	}
	if nodes := bc.mutation.StoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bill.StoreTable,
			Columns: []string{bill.StoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(store.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.BillerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.CustomersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bill.CustomersTable,
			Columns: []string{bill.CustomersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.Ref1 = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.BillDetailIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   bill.BillDetailTable,
			Columns: []string{bill.BillDetailColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(billdetail.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BillCreateBulk is the builder for creating many Bill entities in bulk.
type BillCreateBulk struct {
	config
	builders []*BillCreate
}

// Save creates the Bill entities in the database.
func (bcb *BillCreateBulk) Save(ctx context.Context) ([]*Bill, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Bill, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BillMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BillCreateBulk) SaveX(ctx context.Context) []*Bill {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BillCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BillCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
