// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-blog/app/content/service/internal/data/ent/photo"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PhotoCreate is the builder for creating a Photo entity.
type PhotoCreate struct {
	config
	mutation *PhotoMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (pc *PhotoCreate) SetCreateTime(i int64) *PhotoCreate {
	pc.mutation.SetCreateTime(i)
	return pc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pc *PhotoCreate) SetNillableCreateTime(i *int64) *PhotoCreate {
	if i != nil {
		pc.SetCreateTime(*i)
	}
	return pc
}

// SetUpdateTime sets the "update_time" field.
func (pc *PhotoCreate) SetUpdateTime(i int64) *PhotoCreate {
	pc.mutation.SetUpdateTime(i)
	return pc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pc *PhotoCreate) SetNillableUpdateTime(i *int64) *PhotoCreate {
	if i != nil {
		pc.SetUpdateTime(*i)
	}
	return pc
}

// SetDeleteTime sets the "delete_time" field.
func (pc *PhotoCreate) SetDeleteTime(i int64) *PhotoCreate {
	pc.mutation.SetDeleteTime(i)
	return pc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (pc *PhotoCreate) SetNillableDeleteTime(i *int64) *PhotoCreate {
	if i != nil {
		pc.SetDeleteTime(*i)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *PhotoCreate) SetName(s string) *PhotoCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pc *PhotoCreate) SetNillableName(s *string) *PhotoCreate {
	if s != nil {
		pc.SetName(*s)
	}
	return pc
}

// SetThumbnail sets the "thumbnail" field.
func (pc *PhotoCreate) SetThumbnail(s string) *PhotoCreate {
	pc.mutation.SetThumbnail(s)
	return pc
}

// SetNillableThumbnail sets the "thumbnail" field if the given value is not nil.
func (pc *PhotoCreate) SetNillableThumbnail(s *string) *PhotoCreate {
	if s != nil {
		pc.SetThumbnail(*s)
	}
	return pc
}

// SetTakeTime sets the "take_time" field.
func (pc *PhotoCreate) SetTakeTime(i int64) *PhotoCreate {
	pc.mutation.SetTakeTime(i)
	return pc
}

// SetNillableTakeTime sets the "take_time" field if the given value is not nil.
func (pc *PhotoCreate) SetNillableTakeTime(i *int64) *PhotoCreate {
	if i != nil {
		pc.SetTakeTime(*i)
	}
	return pc
}

// SetURL sets the "url" field.
func (pc *PhotoCreate) SetURL(s string) *PhotoCreate {
	pc.mutation.SetURL(s)
	return pc
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (pc *PhotoCreate) SetNillableURL(s *string) *PhotoCreate {
	if s != nil {
		pc.SetURL(*s)
	}
	return pc
}

// SetTeam sets the "team" field.
func (pc *PhotoCreate) SetTeam(s string) *PhotoCreate {
	pc.mutation.SetTeam(s)
	return pc
}

// SetNillableTeam sets the "team" field if the given value is not nil.
func (pc *PhotoCreate) SetNillableTeam(s *string) *PhotoCreate {
	if s != nil {
		pc.SetTeam(*s)
	}
	return pc
}

// SetLocation sets the "location" field.
func (pc *PhotoCreate) SetLocation(s string) *PhotoCreate {
	pc.mutation.SetLocation(s)
	return pc
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (pc *PhotoCreate) SetNillableLocation(s *string) *PhotoCreate {
	if s != nil {
		pc.SetLocation(*s)
	}
	return pc
}

// SetDescription sets the "description" field.
func (pc *PhotoCreate) SetDescription(s string) *PhotoCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pc *PhotoCreate) SetNillableDescription(s *string) *PhotoCreate {
	if s != nil {
		pc.SetDescription(*s)
	}
	return pc
}

// SetLikes sets the "likes" field.
func (pc *PhotoCreate) SetLikes(i int32) *PhotoCreate {
	pc.mutation.SetLikes(i)
	return pc
}

// SetNillableLikes sets the "likes" field if the given value is not nil.
func (pc *PhotoCreate) SetNillableLikes(i *int32) *PhotoCreate {
	if i != nil {
		pc.SetLikes(*i)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PhotoCreate) SetID(u uint32) *PhotoCreate {
	pc.mutation.SetID(u)
	return pc
}

// Mutation returns the PhotoMutation object of the builder.
func (pc *PhotoCreate) Mutation() *PhotoMutation {
	return pc.mutation
}

// Save creates the Photo in the database.
func (pc *PhotoCreate) Save(ctx context.Context) (*Photo, error) {
	var (
		err  error
		node *Photo
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PhotoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Photo)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PhotoMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PhotoCreate) SaveX(ctx context.Context) *Photo {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PhotoCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PhotoCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PhotoCreate) defaults() {
	if _, ok := pc.mutation.CreateTime(); !ok {
		v := photo.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PhotoCreate) check() error {
	if v, ok := pc.mutation.ID(); ok {
		if err := photo.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Photo.id": %w`, err)}
		}
	}
	return nil
}

func (pc *PhotoCreate) sqlSave(ctx context.Context) (*Photo, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
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

func (pc *PhotoCreate) createSpec() (*Photo, *sqlgraph.CreateSpec) {
	var (
		_node = &Photo{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: photo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: photo.FieldID,
			},
		}
	)
	_spec.OnConflict = pc.conflict
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: photo.FieldCreateTime,
		})
		_node.CreateTime = &value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: photo.FieldUpdateTime,
		})
		_node.UpdateTime = &value
	}
	if value, ok := pc.mutation.DeleteTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: photo.FieldDeleteTime,
		})
		_node.DeleteTime = &value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: photo.FieldName,
		})
		_node.Name = &value
	}
	if value, ok := pc.mutation.Thumbnail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: photo.FieldThumbnail,
		})
		_node.Thumbnail = &value
	}
	if value, ok := pc.mutation.TakeTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: photo.FieldTakeTime,
		})
		_node.TakeTime = &value
	}
	if value, ok := pc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: photo.FieldURL,
		})
		_node.URL = &value
	}
	if value, ok := pc.mutation.Team(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: photo.FieldTeam,
		})
		_node.Team = &value
	}
	if value, ok := pc.mutation.Location(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: photo.FieldLocation,
		})
		_node.Location = &value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: photo.FieldDescription,
		})
		_node.Description = &value
	}
	if value, ok := pc.mutation.Likes(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: photo.FieldLikes,
		})
		_node.Likes = &value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Photo.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PhotoUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (pc *PhotoCreate) OnConflict(opts ...sql.ConflictOption) *PhotoUpsertOne {
	pc.conflict = opts
	return &PhotoUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Photo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *PhotoCreate) OnConflictColumns(columns ...string) *PhotoUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PhotoUpsertOne{
		create: pc,
	}
}

type (
	// PhotoUpsertOne is the builder for "upsert"-ing
	//  one Photo node.
	PhotoUpsertOne struct {
		create *PhotoCreate
	}

	// PhotoUpsert is the "OnConflict" setter.
	PhotoUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *PhotoUpsert) SetUpdateTime(v int64) *PhotoUpsert {
	u.Set(photo.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *PhotoUpsert) UpdateUpdateTime() *PhotoUpsert {
	u.SetExcluded(photo.FieldUpdateTime)
	return u
}

// AddUpdateTime adds v to the "update_time" field.
func (u *PhotoUpsert) AddUpdateTime(v int64) *PhotoUpsert {
	u.Add(photo.FieldUpdateTime, v)
	return u
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *PhotoUpsert) ClearUpdateTime() *PhotoUpsert {
	u.SetNull(photo.FieldUpdateTime)
	return u
}

// SetDeleteTime sets the "delete_time" field.
func (u *PhotoUpsert) SetDeleteTime(v int64) *PhotoUpsert {
	u.Set(photo.FieldDeleteTime, v)
	return u
}

// UpdateDeleteTime sets the "delete_time" field to the value that was provided on create.
func (u *PhotoUpsert) UpdateDeleteTime() *PhotoUpsert {
	u.SetExcluded(photo.FieldDeleteTime)
	return u
}

// AddDeleteTime adds v to the "delete_time" field.
func (u *PhotoUpsert) AddDeleteTime(v int64) *PhotoUpsert {
	u.Add(photo.FieldDeleteTime, v)
	return u
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (u *PhotoUpsert) ClearDeleteTime() *PhotoUpsert {
	u.SetNull(photo.FieldDeleteTime)
	return u
}

// SetName sets the "name" field.
func (u *PhotoUpsert) SetName(v string) *PhotoUpsert {
	u.Set(photo.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PhotoUpsert) UpdateName() *PhotoUpsert {
	u.SetExcluded(photo.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *PhotoUpsert) ClearName() *PhotoUpsert {
	u.SetNull(photo.FieldName)
	return u
}

// SetThumbnail sets the "thumbnail" field.
func (u *PhotoUpsert) SetThumbnail(v string) *PhotoUpsert {
	u.Set(photo.FieldThumbnail, v)
	return u
}

// UpdateThumbnail sets the "thumbnail" field to the value that was provided on create.
func (u *PhotoUpsert) UpdateThumbnail() *PhotoUpsert {
	u.SetExcluded(photo.FieldThumbnail)
	return u
}

// ClearThumbnail clears the value of the "thumbnail" field.
func (u *PhotoUpsert) ClearThumbnail() *PhotoUpsert {
	u.SetNull(photo.FieldThumbnail)
	return u
}

// SetTakeTime sets the "take_time" field.
func (u *PhotoUpsert) SetTakeTime(v int64) *PhotoUpsert {
	u.Set(photo.FieldTakeTime, v)
	return u
}

// UpdateTakeTime sets the "take_time" field to the value that was provided on create.
func (u *PhotoUpsert) UpdateTakeTime() *PhotoUpsert {
	u.SetExcluded(photo.FieldTakeTime)
	return u
}

// AddTakeTime adds v to the "take_time" field.
func (u *PhotoUpsert) AddTakeTime(v int64) *PhotoUpsert {
	u.Add(photo.FieldTakeTime, v)
	return u
}

// ClearTakeTime clears the value of the "take_time" field.
func (u *PhotoUpsert) ClearTakeTime() *PhotoUpsert {
	u.SetNull(photo.FieldTakeTime)
	return u
}

// SetURL sets the "url" field.
func (u *PhotoUpsert) SetURL(v string) *PhotoUpsert {
	u.Set(photo.FieldURL, v)
	return u
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *PhotoUpsert) UpdateURL() *PhotoUpsert {
	u.SetExcluded(photo.FieldURL)
	return u
}

// ClearURL clears the value of the "url" field.
func (u *PhotoUpsert) ClearURL() *PhotoUpsert {
	u.SetNull(photo.FieldURL)
	return u
}

// SetTeam sets the "team" field.
func (u *PhotoUpsert) SetTeam(v string) *PhotoUpsert {
	u.Set(photo.FieldTeam, v)
	return u
}

// UpdateTeam sets the "team" field to the value that was provided on create.
func (u *PhotoUpsert) UpdateTeam() *PhotoUpsert {
	u.SetExcluded(photo.FieldTeam)
	return u
}

// ClearTeam clears the value of the "team" field.
func (u *PhotoUpsert) ClearTeam() *PhotoUpsert {
	u.SetNull(photo.FieldTeam)
	return u
}

// SetLocation sets the "location" field.
func (u *PhotoUpsert) SetLocation(v string) *PhotoUpsert {
	u.Set(photo.FieldLocation, v)
	return u
}

// UpdateLocation sets the "location" field to the value that was provided on create.
func (u *PhotoUpsert) UpdateLocation() *PhotoUpsert {
	u.SetExcluded(photo.FieldLocation)
	return u
}

// ClearLocation clears the value of the "location" field.
func (u *PhotoUpsert) ClearLocation() *PhotoUpsert {
	u.SetNull(photo.FieldLocation)
	return u
}

// SetDescription sets the "description" field.
func (u *PhotoUpsert) SetDescription(v string) *PhotoUpsert {
	u.Set(photo.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *PhotoUpsert) UpdateDescription() *PhotoUpsert {
	u.SetExcluded(photo.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *PhotoUpsert) ClearDescription() *PhotoUpsert {
	u.SetNull(photo.FieldDescription)
	return u
}

// SetLikes sets the "likes" field.
func (u *PhotoUpsert) SetLikes(v int32) *PhotoUpsert {
	u.Set(photo.FieldLikes, v)
	return u
}

// UpdateLikes sets the "likes" field to the value that was provided on create.
func (u *PhotoUpsert) UpdateLikes() *PhotoUpsert {
	u.SetExcluded(photo.FieldLikes)
	return u
}

// AddLikes adds v to the "likes" field.
func (u *PhotoUpsert) AddLikes(v int32) *PhotoUpsert {
	u.Add(photo.FieldLikes, v)
	return u
}

// ClearLikes clears the value of the "likes" field.
func (u *PhotoUpsert) ClearLikes() *PhotoUpsert {
	u.SetNull(photo.FieldLikes)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Photo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(photo.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PhotoUpsertOne) UpdateNewValues() *PhotoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(photo.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(photo.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Photo.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PhotoUpsertOne) Ignore() *PhotoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PhotoUpsertOne) DoNothing() *PhotoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PhotoCreate.OnConflict
// documentation for more info.
func (u *PhotoUpsertOne) Update(set func(*PhotoUpsert)) *PhotoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PhotoUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *PhotoUpsertOne) SetUpdateTime(v int64) *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.SetUpdateTime(v)
	})
}

// AddUpdateTime adds v to the "update_time" field.
func (u *PhotoUpsertOne) AddUpdateTime(v int64) *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.AddUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *PhotoUpsertOne) UpdateUpdateTime() *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.UpdateUpdateTime()
	})
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *PhotoUpsertOne) ClearUpdateTime() *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.ClearUpdateTime()
	})
}

// SetDeleteTime sets the "delete_time" field.
func (u *PhotoUpsertOne) SetDeleteTime(v int64) *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.SetDeleteTime(v)
	})
}

// AddDeleteTime adds v to the "delete_time" field.
func (u *PhotoUpsertOne) AddDeleteTime(v int64) *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.AddDeleteTime(v)
	})
}

// UpdateDeleteTime sets the "delete_time" field to the value that was provided on create.
func (u *PhotoUpsertOne) UpdateDeleteTime() *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.UpdateDeleteTime()
	})
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (u *PhotoUpsertOne) ClearDeleteTime() *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.ClearDeleteTime()
	})
}

// SetName sets the "name" field.
func (u *PhotoUpsertOne) SetName(v string) *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PhotoUpsertOne) UpdateName() *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *PhotoUpsertOne) ClearName() *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.ClearName()
	})
}

// SetThumbnail sets the "thumbnail" field.
func (u *PhotoUpsertOne) SetThumbnail(v string) *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.SetThumbnail(v)
	})
}

// UpdateThumbnail sets the "thumbnail" field to the value that was provided on create.
func (u *PhotoUpsertOne) UpdateThumbnail() *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.UpdateThumbnail()
	})
}

// ClearThumbnail clears the value of the "thumbnail" field.
func (u *PhotoUpsertOne) ClearThumbnail() *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.ClearThumbnail()
	})
}

// SetTakeTime sets the "take_time" field.
func (u *PhotoUpsertOne) SetTakeTime(v int64) *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.SetTakeTime(v)
	})
}

// AddTakeTime adds v to the "take_time" field.
func (u *PhotoUpsertOne) AddTakeTime(v int64) *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.AddTakeTime(v)
	})
}

// UpdateTakeTime sets the "take_time" field to the value that was provided on create.
func (u *PhotoUpsertOne) UpdateTakeTime() *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.UpdateTakeTime()
	})
}

// ClearTakeTime clears the value of the "take_time" field.
func (u *PhotoUpsertOne) ClearTakeTime() *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.ClearTakeTime()
	})
}

// SetURL sets the "url" field.
func (u *PhotoUpsertOne) SetURL(v string) *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *PhotoUpsertOne) UpdateURL() *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.UpdateURL()
	})
}

// ClearURL clears the value of the "url" field.
func (u *PhotoUpsertOne) ClearURL() *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.ClearURL()
	})
}

// SetTeam sets the "team" field.
func (u *PhotoUpsertOne) SetTeam(v string) *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.SetTeam(v)
	})
}

// UpdateTeam sets the "team" field to the value that was provided on create.
func (u *PhotoUpsertOne) UpdateTeam() *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.UpdateTeam()
	})
}

// ClearTeam clears the value of the "team" field.
func (u *PhotoUpsertOne) ClearTeam() *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.ClearTeam()
	})
}

// SetLocation sets the "location" field.
func (u *PhotoUpsertOne) SetLocation(v string) *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.SetLocation(v)
	})
}

// UpdateLocation sets the "location" field to the value that was provided on create.
func (u *PhotoUpsertOne) UpdateLocation() *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.UpdateLocation()
	})
}

// ClearLocation clears the value of the "location" field.
func (u *PhotoUpsertOne) ClearLocation() *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.ClearLocation()
	})
}

// SetDescription sets the "description" field.
func (u *PhotoUpsertOne) SetDescription(v string) *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *PhotoUpsertOne) UpdateDescription() *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *PhotoUpsertOne) ClearDescription() *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.ClearDescription()
	})
}

// SetLikes sets the "likes" field.
func (u *PhotoUpsertOne) SetLikes(v int32) *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.SetLikes(v)
	})
}

// AddLikes adds v to the "likes" field.
func (u *PhotoUpsertOne) AddLikes(v int32) *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.AddLikes(v)
	})
}

// UpdateLikes sets the "likes" field to the value that was provided on create.
func (u *PhotoUpsertOne) UpdateLikes() *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.UpdateLikes()
	})
}

// ClearLikes clears the value of the "likes" field.
func (u *PhotoUpsertOne) ClearLikes() *PhotoUpsertOne {
	return u.Update(func(s *PhotoUpsert) {
		s.ClearLikes()
	})
}

// Exec executes the query.
func (u *PhotoUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PhotoCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PhotoUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PhotoUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PhotoUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PhotoCreateBulk is the builder for creating many Photo entities in bulk.
type PhotoCreateBulk struct {
	config
	builders []*PhotoCreate
	conflict []sql.ConflictOption
}

// Save creates the Photo entities in the database.
func (pcb *PhotoCreateBulk) Save(ctx context.Context) ([]*Photo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Photo, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PhotoMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PhotoCreateBulk) SaveX(ctx context.Context) []*Photo {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PhotoCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PhotoCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Photo.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PhotoUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (pcb *PhotoCreateBulk) OnConflict(opts ...sql.ConflictOption) *PhotoUpsertBulk {
	pcb.conflict = opts
	return &PhotoUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Photo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *PhotoCreateBulk) OnConflictColumns(columns ...string) *PhotoUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PhotoUpsertBulk{
		create: pcb,
	}
}

// PhotoUpsertBulk is the builder for "upsert"-ing
// a bulk of Photo nodes.
type PhotoUpsertBulk struct {
	create *PhotoCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Photo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(photo.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PhotoUpsertBulk) UpdateNewValues() *PhotoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(photo.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(photo.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Photo.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PhotoUpsertBulk) Ignore() *PhotoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PhotoUpsertBulk) DoNothing() *PhotoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PhotoCreateBulk.OnConflict
// documentation for more info.
func (u *PhotoUpsertBulk) Update(set func(*PhotoUpsert)) *PhotoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PhotoUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *PhotoUpsertBulk) SetUpdateTime(v int64) *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.SetUpdateTime(v)
	})
}

// AddUpdateTime adds v to the "update_time" field.
func (u *PhotoUpsertBulk) AddUpdateTime(v int64) *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.AddUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *PhotoUpsertBulk) UpdateUpdateTime() *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.UpdateUpdateTime()
	})
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *PhotoUpsertBulk) ClearUpdateTime() *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.ClearUpdateTime()
	})
}

// SetDeleteTime sets the "delete_time" field.
func (u *PhotoUpsertBulk) SetDeleteTime(v int64) *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.SetDeleteTime(v)
	})
}

// AddDeleteTime adds v to the "delete_time" field.
func (u *PhotoUpsertBulk) AddDeleteTime(v int64) *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.AddDeleteTime(v)
	})
}

// UpdateDeleteTime sets the "delete_time" field to the value that was provided on create.
func (u *PhotoUpsertBulk) UpdateDeleteTime() *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.UpdateDeleteTime()
	})
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (u *PhotoUpsertBulk) ClearDeleteTime() *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.ClearDeleteTime()
	})
}

// SetName sets the "name" field.
func (u *PhotoUpsertBulk) SetName(v string) *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PhotoUpsertBulk) UpdateName() *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *PhotoUpsertBulk) ClearName() *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.ClearName()
	})
}

// SetThumbnail sets the "thumbnail" field.
func (u *PhotoUpsertBulk) SetThumbnail(v string) *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.SetThumbnail(v)
	})
}

// UpdateThumbnail sets the "thumbnail" field to the value that was provided on create.
func (u *PhotoUpsertBulk) UpdateThumbnail() *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.UpdateThumbnail()
	})
}

// ClearThumbnail clears the value of the "thumbnail" field.
func (u *PhotoUpsertBulk) ClearThumbnail() *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.ClearThumbnail()
	})
}

// SetTakeTime sets the "take_time" field.
func (u *PhotoUpsertBulk) SetTakeTime(v int64) *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.SetTakeTime(v)
	})
}

// AddTakeTime adds v to the "take_time" field.
func (u *PhotoUpsertBulk) AddTakeTime(v int64) *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.AddTakeTime(v)
	})
}

// UpdateTakeTime sets the "take_time" field to the value that was provided on create.
func (u *PhotoUpsertBulk) UpdateTakeTime() *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.UpdateTakeTime()
	})
}

// ClearTakeTime clears the value of the "take_time" field.
func (u *PhotoUpsertBulk) ClearTakeTime() *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.ClearTakeTime()
	})
}

// SetURL sets the "url" field.
func (u *PhotoUpsertBulk) SetURL(v string) *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *PhotoUpsertBulk) UpdateURL() *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.UpdateURL()
	})
}

// ClearURL clears the value of the "url" field.
func (u *PhotoUpsertBulk) ClearURL() *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.ClearURL()
	})
}

// SetTeam sets the "team" field.
func (u *PhotoUpsertBulk) SetTeam(v string) *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.SetTeam(v)
	})
}

// UpdateTeam sets the "team" field to the value that was provided on create.
func (u *PhotoUpsertBulk) UpdateTeam() *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.UpdateTeam()
	})
}

// ClearTeam clears the value of the "team" field.
func (u *PhotoUpsertBulk) ClearTeam() *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.ClearTeam()
	})
}

// SetLocation sets the "location" field.
func (u *PhotoUpsertBulk) SetLocation(v string) *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.SetLocation(v)
	})
}

// UpdateLocation sets the "location" field to the value that was provided on create.
func (u *PhotoUpsertBulk) UpdateLocation() *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.UpdateLocation()
	})
}

// ClearLocation clears the value of the "location" field.
func (u *PhotoUpsertBulk) ClearLocation() *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.ClearLocation()
	})
}

// SetDescription sets the "description" field.
func (u *PhotoUpsertBulk) SetDescription(v string) *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *PhotoUpsertBulk) UpdateDescription() *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *PhotoUpsertBulk) ClearDescription() *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.ClearDescription()
	})
}

// SetLikes sets the "likes" field.
func (u *PhotoUpsertBulk) SetLikes(v int32) *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.SetLikes(v)
	})
}

// AddLikes adds v to the "likes" field.
func (u *PhotoUpsertBulk) AddLikes(v int32) *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.AddLikes(v)
	})
}

// UpdateLikes sets the "likes" field to the value that was provided on create.
func (u *PhotoUpsertBulk) UpdateLikes() *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.UpdateLikes()
	})
}

// ClearLikes clears the value of the "likes" field.
func (u *PhotoUpsertBulk) ClearLikes() *PhotoUpsertBulk {
	return u.Update(func(s *PhotoUpsert) {
		s.ClearLikes()
	})
}

// Exec executes the query.
func (u *PhotoUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PhotoCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PhotoCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PhotoUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
