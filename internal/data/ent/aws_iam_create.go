// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/golifez/zkit/internal/data/ent/aws_iam"
)

// AWSIamCreate is the builder for creating a Aws_iam entity.
type AWSIamCreate struct {
	config
	mutation *AWSIamMutation
	hooks    []Hook
}

// SetUID sets the "uid" field.
func (aic *AWSIamCreate) SetUID(s string) *AWSIamCreate {
	aic.mutation.SetUID(s)
	return aic
}

// SetAccountID sets the "account_id" field.
func (aic *AWSIamCreate) SetAccountID(s string) *AWSIamCreate {
	aic.mutation.SetAccountID(s)
	return aic
}

// SetIamName sets the "iam_name" field.
func (aic *AWSIamCreate) SetIamName(s string) *AWSIamCreate {
	aic.mutation.SetIamName(s)
	return aic
}

// SetAccessKey sets the "access_key" field.
func (aic *AWSIamCreate) SetAccessKey(s string) *AWSIamCreate {
	aic.mutation.SetAccessKey(s)
	return aic
}

// SetSecretKey sets the "secret_key" field.
func (aic *AWSIamCreate) SetSecretKey(s string) *AWSIamCreate {
	aic.mutation.SetSecretKey(s)
	return aic
}

// SetCreatedAt sets the "created_at" field.
func (aic *AWSIamCreate) SetCreatedAt(t time.Time) *AWSIamCreate {
	aic.mutation.SetCreatedAt(t)
	return aic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aic *AWSIamCreate) SetNillableCreatedAt(t *time.Time) *AWSIamCreate {
	if t != nil {
		aic.SetCreatedAt(*t)
	}
	return aic
}

// SetUpdatedAt sets the "updated_at" field.
func (aic *AWSIamCreate) SetUpdatedAt(t time.Time) *AWSIamCreate {
	aic.mutation.SetUpdatedAt(t)
	return aic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (aic *AWSIamCreate) SetNillableUpdatedAt(t *time.Time) *AWSIamCreate {
	if t != nil {
		aic.SetUpdatedAt(*t)
	}
	return aic
}

// Mutation returns the AWSIamMutation object of the builder.
func (aic *AWSIamCreate) Mutation() *AWSIamMutation {
	return aic.mutation
}

// Save creates the Aws_iam in the database.
func (aic *AWSIamCreate) Save(ctx context.Context) (*Aws_iam, error) {
	aic.defaults()
	return withHooks(ctx, aic.sqlSave, aic.mutation, aic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (aic *AWSIamCreate) SaveX(ctx context.Context) *Aws_iam {
	v, err := aic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aic *AWSIamCreate) Exec(ctx context.Context) error {
	_, err := aic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aic *AWSIamCreate) ExecX(ctx context.Context) {
	if err := aic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aic *AWSIamCreate) defaults() {
	if _, ok := aic.mutation.CreatedAt(); !ok {
		v := aws_iam.DefaultCreatedAt()
		aic.mutation.SetCreatedAt(v)
	}
	if _, ok := aic.mutation.UpdatedAt(); !ok {
		v := aws_iam.DefaultUpdatedAt()
		aic.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aic *AWSIamCreate) check() error {
	if _, ok := aic.mutation.UID(); !ok {
		return &ValidationError{Name: "uid", err: errors.New(`ent: missing required field "Aws_iam.uid"`)}
	}
	if _, ok := aic.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account_id", err: errors.New(`ent: missing required field "Aws_iam.account_id"`)}
	}
	if _, ok := aic.mutation.IamName(); !ok {
		return &ValidationError{Name: "iam_name", err: errors.New(`ent: missing required field "Aws_iam.iam_name"`)}
	}
	if _, ok := aic.mutation.AccessKey(); !ok {
		return &ValidationError{Name: "access_key", err: errors.New(`ent: missing required field "Aws_iam.access_key"`)}
	}
	if _, ok := aic.mutation.SecretKey(); !ok {
		return &ValidationError{Name: "secret_key", err: errors.New(`ent: missing required field "Aws_iam.secret_key"`)}
	}
	if _, ok := aic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Aws_iam.created_at"`)}
	}
	if _, ok := aic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Aws_iam.updated_at"`)}
	}
	return nil
}

func (aic *AWSIamCreate) sqlSave(ctx context.Context) (*Aws_iam, error) {
	if err := aic.check(); err != nil {
		return nil, err
	}
	_node, _spec := aic.createSpec()
	if err := sqlgraph.CreateNode(ctx, aic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	aic.mutation.id = &_node.ID
	aic.mutation.done = true
	return _node, nil
}

func (aic *AWSIamCreate) createSpec() (*Aws_iam, *sqlgraph.CreateSpec) {
	var (
		_node = &Aws_iam{config: aic.config}
		_spec = sqlgraph.NewCreateSpec(aws_iam.Table, sqlgraph.NewFieldSpec(aws_iam.FieldID, field.TypeInt))
	)
	if value, ok := aic.mutation.UID(); ok {
		_spec.SetField(aws_iam.FieldUID, field.TypeString, value)
		_node.UID = value
	}
	if value, ok := aic.mutation.AccountID(); ok {
		_spec.SetField(aws_iam.FieldAccountID, field.TypeString, value)
		_node.AccountID = value
	}
	if value, ok := aic.mutation.IamName(); ok {
		_spec.SetField(aws_iam.FieldIamName, field.TypeString, value)
		_node.IamName = value
	}
	if value, ok := aic.mutation.AccessKey(); ok {
		_spec.SetField(aws_iam.FieldAccessKey, field.TypeString, value)
		_node.AccessKey = value
	}
	if value, ok := aic.mutation.SecretKey(); ok {
		_spec.SetField(aws_iam.FieldSecretKey, field.TypeString, value)
		_node.SecretKey = value
	}
	if value, ok := aic.mutation.CreatedAt(); ok {
		_spec.SetField(aws_iam.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := aic.mutation.UpdatedAt(); ok {
		_spec.SetField(aws_iam.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// AWSIamCreateBulk is the builder for creating many Aws_iam entities in bulk.
type AWSIamCreateBulk struct {
	config
	err      error
	builders []*AWSIamCreate
}

// Save creates the Aws_iam entities in the database.
func (aicb *AWSIamCreateBulk) Save(ctx context.Context) ([]*Aws_iam, error) {
	if aicb.err != nil {
		return nil, aicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(aicb.builders))
	nodes := make([]*Aws_iam, len(aicb.builders))
	mutators := make([]Mutator, len(aicb.builders))
	for i := range aicb.builders {
		func(i int, root context.Context) {
			builder := aicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AWSIamMutation)
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
					_, err = mutators[i+1].Mutate(root, aicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
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
		if _, err := mutators[0].Mutate(ctx, aicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aicb *AWSIamCreateBulk) SaveX(ctx context.Context) []*Aws_iam {
	v, err := aicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aicb *AWSIamCreateBulk) Exec(ctx context.Context) error {
	_, err := aicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aicb *AWSIamCreateBulk) ExecX(ctx context.Context) {
	if err := aicb.Exec(ctx); err != nil {
		panic(err)
	}
}
