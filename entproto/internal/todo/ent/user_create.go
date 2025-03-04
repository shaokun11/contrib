// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/contrib/entproto/internal/todo/ent/attachment"
	"entgo.io/contrib/entproto/internal/todo/ent/group"
	"entgo.io/contrib/entproto/internal/todo/ent/pet"
	"entgo.io/contrib/entproto/internal/todo/ent/schema"
	"entgo.io/contrib/entproto/internal/todo/ent/skipedgeexample"
	"entgo.io/contrib/entproto/internal/todo/ent/user"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetUserName sets the "user_name" field.
func (uc *UserCreate) SetUserName(s string) *UserCreate {
	uc.mutation.SetUserName(s)
	return uc
}

// SetJoined sets the "joined" field.
func (uc *UserCreate) SetJoined(t time.Time) *UserCreate {
	uc.mutation.SetJoined(t)
	return uc
}

// SetPoints sets the "points" field.
func (uc *UserCreate) SetPoints(u uint) *UserCreate {
	uc.mutation.SetPoints(u)
	return uc
}

// SetExp sets the "exp" field.
func (uc *UserCreate) SetExp(u uint64) *UserCreate {
	uc.mutation.SetExp(u)
	return uc
}

// SetStatus sets the "status" field.
func (uc *UserCreate) SetStatus(u user.Status) *UserCreate {
	uc.mutation.SetStatus(u)
	return uc
}

// SetExternalID sets the "external_id" field.
func (uc *UserCreate) SetExternalID(i int) *UserCreate {
	uc.mutation.SetExternalID(i)
	return uc
}

// SetCrmID sets the "crm_id" field.
func (uc *UserCreate) SetCrmID(u uuid.UUID) *UserCreate {
	uc.mutation.SetCrmID(u)
	return uc
}

// SetBanned sets the "banned" field.
func (uc *UserCreate) SetBanned(b bool) *UserCreate {
	uc.mutation.SetBanned(b)
	return uc
}

// SetNillableBanned sets the "banned" field if the given value is not nil.
func (uc *UserCreate) SetNillableBanned(b *bool) *UserCreate {
	if b != nil {
		uc.SetBanned(*b)
	}
	return uc
}

// SetCustomPb sets the "custom_pb" field.
func (uc *UserCreate) SetCustomPb(u uint8) *UserCreate {
	uc.mutation.SetCustomPb(u)
	return uc
}

// SetOptNum sets the "opt_num" field.
func (uc *UserCreate) SetOptNum(i int) *UserCreate {
	uc.mutation.SetOptNum(i)
	return uc
}

// SetNillableOptNum sets the "opt_num" field if the given value is not nil.
func (uc *UserCreate) SetNillableOptNum(i *int) *UserCreate {
	if i != nil {
		uc.SetOptNum(*i)
	}
	return uc
}

// SetOptStr sets the "opt_str" field.
func (uc *UserCreate) SetOptStr(s string) *UserCreate {
	uc.mutation.SetOptStr(s)
	return uc
}

// SetNillableOptStr sets the "opt_str" field if the given value is not nil.
func (uc *UserCreate) SetNillableOptStr(s *string) *UserCreate {
	if s != nil {
		uc.SetOptStr(*s)
	}
	return uc
}

// SetOptBool sets the "opt_bool" field.
func (uc *UserCreate) SetOptBool(b bool) *UserCreate {
	uc.mutation.SetOptBool(b)
	return uc
}

// SetNillableOptBool sets the "opt_bool" field if the given value is not nil.
func (uc *UserCreate) SetNillableOptBool(b *bool) *UserCreate {
	if b != nil {
		uc.SetOptBool(*b)
	}
	return uc
}

// SetBigInt sets the "big_int" field.
func (uc *UserCreate) SetBigInt(si schema.BigInt) *UserCreate {
	uc.mutation.SetBigInt(si)
	return uc
}

// SetNillableBigInt sets the "big_int" field if the given value is not nil.
func (uc *UserCreate) SetNillableBigInt(si *schema.BigInt) *UserCreate {
	if si != nil {
		uc.SetBigInt(*si)
	}
	return uc
}

// SetBUser1 sets the "b_user_1" field.
func (uc *UserCreate) SetBUser1(i int) *UserCreate {
	uc.mutation.SetBUser1(i)
	return uc
}

// SetNillableBUser1 sets the "b_user_1" field if the given value is not nil.
func (uc *UserCreate) SetNillableBUser1(i *int) *UserCreate {
	if i != nil {
		uc.SetBUser1(*i)
	}
	return uc
}

// SetHeightInCm sets the "height_in_cm" field.
func (uc *UserCreate) SetHeightInCm(f float32) *UserCreate {
	uc.mutation.SetHeightInCm(f)
	return uc
}

// SetNillableHeightInCm sets the "height_in_cm" field if the given value is not nil.
func (uc *UserCreate) SetNillableHeightInCm(f *float32) *UserCreate {
	if f != nil {
		uc.SetHeightInCm(*f)
	}
	return uc
}

// SetAccountBalance sets the "account_balance" field.
func (uc *UserCreate) SetAccountBalance(f float64) *UserCreate {
	uc.mutation.SetAccountBalance(f)
	return uc
}

// SetNillableAccountBalance sets the "account_balance" field if the given value is not nil.
func (uc *UserCreate) SetNillableAccountBalance(f *float64) *UserCreate {
	if f != nil {
		uc.SetAccountBalance(*f)
	}
	return uc
}

// SetUnnecessary sets the "unnecessary" field.
func (uc *UserCreate) SetUnnecessary(s string) *UserCreate {
	uc.mutation.SetUnnecessary(s)
	return uc
}

// SetNillableUnnecessary sets the "unnecessary" field if the given value is not nil.
func (uc *UserCreate) SetNillableUnnecessary(s *string) *UserCreate {
	if s != nil {
		uc.SetUnnecessary(*s)
	}
	return uc
}

// SetType sets the "type" field.
func (uc *UserCreate) SetType(s string) *UserCreate {
	uc.mutation.SetType(s)
	return uc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (uc *UserCreate) SetNillableType(s *string) *UserCreate {
	if s != nil {
		uc.SetType(*s)
	}
	return uc
}

// SetLabels sets the "labels" field.
func (uc *UserCreate) SetLabels(s []string) *UserCreate {
	uc.mutation.SetLabels(s)
	return uc
}

// SetDeviceType sets the "device_type" field.
func (uc *UserCreate) SetDeviceType(ut user.DeviceType) *UserCreate {
	uc.mutation.SetDeviceType(ut)
	return uc
}

// SetNillableDeviceType sets the "device_type" field if the given value is not nil.
func (uc *UserCreate) SetNillableDeviceType(ut *user.DeviceType) *UserCreate {
	if ut != nil {
		uc.SetDeviceType(*ut)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(u uint32) *UserCreate {
	uc.mutation.SetID(u)
	return uc
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (uc *UserCreate) SetGroupID(id int) *UserCreate {
	uc.mutation.SetGroupID(id)
	return uc
}

// SetNillableGroupID sets the "group" edge to the Group entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableGroupID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetGroupID(*id)
	}
	return uc
}

// SetGroup sets the "group" edge to the Group entity.
func (uc *UserCreate) SetGroup(g *Group) *UserCreate {
	return uc.SetGroupID(g.ID)
}

// SetAttachmentID sets the "attachment" edge to the Attachment entity by ID.
func (uc *UserCreate) SetAttachmentID(id uuid.UUID) *UserCreate {
	uc.mutation.SetAttachmentID(id)
	return uc
}

// SetNillableAttachmentID sets the "attachment" edge to the Attachment entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableAttachmentID(id *uuid.UUID) *UserCreate {
	if id != nil {
		uc = uc.SetAttachmentID(*id)
	}
	return uc
}

// SetAttachment sets the "attachment" edge to the Attachment entity.
func (uc *UserCreate) SetAttachment(a *Attachment) *UserCreate {
	return uc.SetAttachmentID(a.ID)
}

// AddReceived1IDs adds the "received_1" edge to the Attachment entity by IDs.
func (uc *UserCreate) AddReceived1IDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddReceived1IDs(ids...)
	return uc
}

// AddReceived1 adds the "received_1" edges to the Attachment entity.
func (uc *UserCreate) AddReceived1(a ...*Attachment) *UserCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uc.AddReceived1IDs(ids...)
}

// SetPetID sets the "pet" edge to the Pet entity by ID.
func (uc *UserCreate) SetPetID(id int) *UserCreate {
	uc.mutation.SetPetID(id)
	return uc
}

// SetNillablePetID sets the "pet" edge to the Pet entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillablePetID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetPetID(*id)
	}
	return uc
}

// SetPet sets the "pet" edge to the Pet entity.
func (uc *UserCreate) SetPet(p *Pet) *UserCreate {
	return uc.SetPetID(p.ID)
}

// SetSkipEdgeID sets the "skip_edge" edge to the SkipEdgeExample entity by ID.
func (uc *UserCreate) SetSkipEdgeID(id int) *UserCreate {
	uc.mutation.SetSkipEdgeID(id)
	return uc
}

// SetNillableSkipEdgeID sets the "skip_edge" edge to the SkipEdgeExample entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableSkipEdgeID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetSkipEdgeID(*id)
	}
	return uc
}

// SetSkipEdge sets the "skip_edge" edge to the SkipEdgeExample entity.
func (uc *UserCreate) SetSkipEdge(s *SkipEdgeExample) *UserCreate {
	return uc.SetSkipEdgeID(s.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uc.defaults()
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			if node, err = uc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			if uc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*User)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.Banned(); !ok {
		v := user.DefaultBanned
		uc.mutation.SetBanned(v)
	}
	if _, ok := uc.mutation.HeightInCm(); !ok {
		v := user.DefaultHeightInCm
		uc.mutation.SetHeightInCm(v)
	}
	if _, ok := uc.mutation.AccountBalance(); !ok {
		v := user.DefaultAccountBalance
		uc.mutation.SetAccountBalance(v)
	}
	if _, ok := uc.mutation.DeviceType(); !ok {
		v := user.DefaultDeviceType
		uc.mutation.SetDeviceType(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.UserName(); !ok {
		return &ValidationError{Name: "user_name", err: errors.New(`ent: missing required field "User.user_name"`)}
	}
	if _, ok := uc.mutation.Joined(); !ok {
		return &ValidationError{Name: "joined", err: errors.New(`ent: missing required field "User.joined"`)}
	}
	if _, ok := uc.mutation.Points(); !ok {
		return &ValidationError{Name: "points", err: errors.New(`ent: missing required field "User.points"`)}
	}
	if _, ok := uc.mutation.Exp(); !ok {
		return &ValidationError{Name: "exp", err: errors.New(`ent: missing required field "User.exp"`)}
	}
	if _, ok := uc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "User.status"`)}
	}
	if v, ok := uc.mutation.Status(); ok {
		if err := user.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "User.status": %w`, err)}
		}
	}
	if _, ok := uc.mutation.ExternalID(); !ok {
		return &ValidationError{Name: "external_id", err: errors.New(`ent: missing required field "User.external_id"`)}
	}
	if _, ok := uc.mutation.CrmID(); !ok {
		return &ValidationError{Name: "crm_id", err: errors.New(`ent: missing required field "User.crm_id"`)}
	}
	if _, ok := uc.mutation.Banned(); !ok {
		return &ValidationError{Name: "banned", err: errors.New(`ent: missing required field "User.banned"`)}
	}
	if _, ok := uc.mutation.CustomPb(); !ok {
		return &ValidationError{Name: "custom_pb", err: errors.New(`ent: missing required field "User.custom_pb"`)}
	}
	if _, ok := uc.mutation.HeightInCm(); !ok {
		return &ValidationError{Name: "height_in_cm", err: errors.New(`ent: missing required field "User.height_in_cm"`)}
	}
	if _, ok := uc.mutation.AccountBalance(); !ok {
		return &ValidationError{Name: "account_balance", err: errors.New(`ent: missing required field "User.account_balance"`)}
	}
	if _, ok := uc.mutation.DeviceType(); !ok {
		return &ValidationError{Name: "device_type", err: errors.New(`ent: missing required field "User.device_type"`)}
	}
	if v, ok := uc.mutation.DeviceType(); ok {
		if err := user.DeviceTypeValidator(v); err != nil {
			return &ValidationError{Name: "device_type", err: fmt.Errorf(`ent: validator failed for field "User.device_type": %w`, err)}
		}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: user.FieldID,
			},
		}
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.UserName(); ok {
		_spec.SetField(user.FieldUserName, field.TypeString, value)
		_node.UserName = value
	}
	if value, ok := uc.mutation.Joined(); ok {
		_spec.SetField(user.FieldJoined, field.TypeTime, value)
		_node.Joined = value
	}
	if value, ok := uc.mutation.Points(); ok {
		_spec.SetField(user.FieldPoints, field.TypeUint, value)
		_node.Points = value
	}
	if value, ok := uc.mutation.Exp(); ok {
		_spec.SetField(user.FieldExp, field.TypeUint64, value)
		_node.Exp = value
	}
	if value, ok := uc.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := uc.mutation.ExternalID(); ok {
		_spec.SetField(user.FieldExternalID, field.TypeInt, value)
		_node.ExternalID = value
	}
	if value, ok := uc.mutation.CrmID(); ok {
		_spec.SetField(user.FieldCrmID, field.TypeUUID, value)
		_node.CrmID = value
	}
	if value, ok := uc.mutation.Banned(); ok {
		_spec.SetField(user.FieldBanned, field.TypeBool, value)
		_node.Banned = value
	}
	if value, ok := uc.mutation.CustomPb(); ok {
		_spec.SetField(user.FieldCustomPb, field.TypeUint8, value)
		_node.CustomPb = value
	}
	if value, ok := uc.mutation.OptNum(); ok {
		_spec.SetField(user.FieldOptNum, field.TypeInt, value)
		_node.OptNum = value
	}
	if value, ok := uc.mutation.OptStr(); ok {
		_spec.SetField(user.FieldOptStr, field.TypeString, value)
		_node.OptStr = value
	}
	if value, ok := uc.mutation.OptBool(); ok {
		_spec.SetField(user.FieldOptBool, field.TypeBool, value)
		_node.OptBool = value
	}
	if value, ok := uc.mutation.BigInt(); ok {
		_spec.SetField(user.FieldBigInt, field.TypeInt, value)
		_node.BigInt = value
	}
	if value, ok := uc.mutation.BUser1(); ok {
		_spec.SetField(user.FieldBUser1, field.TypeInt, value)
		_node.BUser1 = value
	}
	if value, ok := uc.mutation.HeightInCm(); ok {
		_spec.SetField(user.FieldHeightInCm, field.TypeFloat32, value)
		_node.HeightInCm = value
	}
	if value, ok := uc.mutation.AccountBalance(); ok {
		_spec.SetField(user.FieldAccountBalance, field.TypeFloat64, value)
		_node.AccountBalance = value
	}
	if value, ok := uc.mutation.Unnecessary(); ok {
		_spec.SetField(user.FieldUnnecessary, field.TypeString, value)
		_node.Unnecessary = value
	}
	if value, ok := uc.mutation.GetType(); ok {
		_spec.SetField(user.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := uc.mutation.Labels(); ok {
		_spec.SetField(user.FieldLabels, field.TypeJSON, value)
		_node.Labels = value
	}
	if value, ok := uc.mutation.DeviceType(); ok {
		_spec.SetField(user.FieldDeviceType, field.TypeEnum, value)
		_node.DeviceType = value
	}
	if nodes := uc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.GroupTable,
			Columns: []string{user.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_group = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.AttachmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.AttachmentTable,
			Columns: []string{user.AttachmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attachment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.Received1IDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.Received1Table,
			Columns: user.Received1PrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attachment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.PetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PetTable,
			Columns: []string{user.PetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.SkipEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SkipEdgeTable,
			Columns: []string{user.SkipEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: skipedgeexample.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
					nodes[i].ID = uint32(id)
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
