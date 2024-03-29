// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/ent/article"
	"blog/ent/category"
	"blog/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	uuid "github.com/gofrs/uuid/v5"
)

// CategoryUpdate is the builder for updating Category entities.
type CategoryUpdate struct {
	config
	hooks    []Hook
	mutation *CategoryMutation
}

// Where appends a list predicates to the CategoryUpdate builder.
func (cu *CategoryUpdate) Where(ps ...predicate.Category) *CategoryUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CategoryUpdate) SetUpdatedAt(t time.Time) *CategoryUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetTitle sets the "title" field.
func (cu *CategoryUpdate) SetTitle(s string) *CategoryUpdate {
	cu.mutation.SetTitle(s)
	return cu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableTitle(s *string) *CategoryUpdate {
	if s != nil {
		cu.SetTitle(*s)
	}
	return cu
}

// SetRemark sets the "remark" field.
func (cu *CategoryUpdate) SetRemark(s string) *CategoryUpdate {
	cu.mutation.SetRemark(s)
	return cu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableRemark(s *string) *CategoryUpdate {
	if s != nil {
		cu.SetRemark(*s)
	}
	return cu
}

// AddArticleIDs adds the "articles" edge to the Article entity by IDs.
func (cu *CategoryUpdate) AddArticleIDs(ids ...uuid.UUID) *CategoryUpdate {
	cu.mutation.AddArticleIDs(ids...)
	return cu
}

// AddArticles adds the "articles" edges to the Article entity.
func (cu *CategoryUpdate) AddArticles(a ...*Article) *CategoryUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cu.AddArticleIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cu *CategoryUpdate) Mutation() *CategoryMutation {
	return cu.mutation
}

// ClearArticles clears all "articles" edges to the Article entity.
func (cu *CategoryUpdate) ClearArticles() *CategoryUpdate {
	cu.mutation.ClearArticles()
	return cu
}

// RemoveArticleIDs removes the "articles" edge to Article entities by IDs.
func (cu *CategoryUpdate) RemoveArticleIDs(ids ...uuid.UUID) *CategoryUpdate {
	cu.mutation.RemoveArticleIDs(ids...)
	return cu
}

// RemoveArticles removes "articles" edges to Article entities.
func (cu *CategoryUpdate) RemoveArticles(a ...*Article) *CategoryUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cu.RemoveArticleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CategoryUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CategoryUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CategoryUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CategoryUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := category.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

func (cu *CategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(category.Table, category.Columns, sqlgraph.NewFieldSpec(category.FieldID, field.TypeUint64))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(category.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Title(); ok {
		_spec.SetField(category.FieldTitle, field.TypeString, value)
	}
	if value, ok := cu.mutation.Remark(); ok {
		_spec.SetField(category.FieldRemark, field.TypeString, value)
	}
	if cu.mutation.ArticlesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   category.ArticlesTable,
			Columns: []string{category.ArticlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(article.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedArticlesIDs(); len(nodes) > 0 && !cu.mutation.ArticlesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   category.ArticlesTable,
			Columns: []string{category.ArticlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(article.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ArticlesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   category.ArticlesTable,
			Columns: []string{category.ArticlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(article.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CategoryUpdateOne is the builder for updating a single Category entity.
type CategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CategoryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CategoryUpdateOne) SetUpdatedAt(t time.Time) *CategoryUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetTitle sets the "title" field.
func (cuo *CategoryUpdateOne) SetTitle(s string) *CategoryUpdateOne {
	cuo.mutation.SetTitle(s)
	return cuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableTitle(s *string) *CategoryUpdateOne {
	if s != nil {
		cuo.SetTitle(*s)
	}
	return cuo
}

// SetRemark sets the "remark" field.
func (cuo *CategoryUpdateOne) SetRemark(s string) *CategoryUpdateOne {
	cuo.mutation.SetRemark(s)
	return cuo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableRemark(s *string) *CategoryUpdateOne {
	if s != nil {
		cuo.SetRemark(*s)
	}
	return cuo
}

// AddArticleIDs adds the "articles" edge to the Article entity by IDs.
func (cuo *CategoryUpdateOne) AddArticleIDs(ids ...uuid.UUID) *CategoryUpdateOne {
	cuo.mutation.AddArticleIDs(ids...)
	return cuo
}

// AddArticles adds the "articles" edges to the Article entity.
func (cuo *CategoryUpdateOne) AddArticles(a ...*Article) *CategoryUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cuo.AddArticleIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cuo *CategoryUpdateOne) Mutation() *CategoryMutation {
	return cuo.mutation
}

// ClearArticles clears all "articles" edges to the Article entity.
func (cuo *CategoryUpdateOne) ClearArticles() *CategoryUpdateOne {
	cuo.mutation.ClearArticles()
	return cuo
}

// RemoveArticleIDs removes the "articles" edge to Article entities by IDs.
func (cuo *CategoryUpdateOne) RemoveArticleIDs(ids ...uuid.UUID) *CategoryUpdateOne {
	cuo.mutation.RemoveArticleIDs(ids...)
	return cuo
}

// RemoveArticles removes "articles" edges to Article entities.
func (cuo *CategoryUpdateOne) RemoveArticles(a ...*Article) *CategoryUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cuo.RemoveArticleIDs(ids...)
}

// Where appends a list predicates to the CategoryUpdate builder.
func (cuo *CategoryUpdateOne) Where(ps ...predicate.Category) *CategoryUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CategoryUpdateOne) Select(field string, fields ...string) *CategoryUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Category entity.
func (cuo *CategoryUpdateOne) Save(ctx context.Context) (*Category, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CategoryUpdateOne) SaveX(ctx context.Context) *Category {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CategoryUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CategoryUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := category.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

func (cuo *CategoryUpdateOne) sqlSave(ctx context.Context) (_node *Category, err error) {
	_spec := sqlgraph.NewUpdateSpec(category.Table, category.Columns, sqlgraph.NewFieldSpec(category.FieldID, field.TypeUint64))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Category.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, category.FieldID)
		for _, f := range fields {
			if !category.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != category.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(category.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Title(); ok {
		_spec.SetField(category.FieldTitle, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Remark(); ok {
		_spec.SetField(category.FieldRemark, field.TypeString, value)
	}
	if cuo.mutation.ArticlesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   category.ArticlesTable,
			Columns: []string{category.ArticlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(article.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedArticlesIDs(); len(nodes) > 0 && !cuo.mutation.ArticlesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   category.ArticlesTable,
			Columns: []string{category.ArticlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(article.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ArticlesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   category.ArticlesTable,
			Columns: []string{category.ArticlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(article.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Category{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
