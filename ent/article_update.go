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
)

// ArticleUpdate is the builder for updating Article entities.
type ArticleUpdate struct {
	config
	hooks    []Hook
	mutation *ArticleMutation
}

// Where appends a list predicates to the ArticleUpdate builder.
func (au *ArticleUpdate) Where(ps ...predicate.Article) *ArticleUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *ArticleUpdate) SetUpdatedAt(t time.Time) *ArticleUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetStatus sets the "status" field.
func (au *ArticleUpdate) SetStatus(u uint8) *ArticleUpdate {
	au.mutation.ResetStatus()
	au.mutation.SetStatus(u)
	return au
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableStatus(u *uint8) *ArticleUpdate {
	if u != nil {
		au.SetStatus(*u)
	}
	return au
}

// AddStatus adds u to the "status" field.
func (au *ArticleUpdate) AddStatus(u int8) *ArticleUpdate {
	au.mutation.AddStatus(u)
	return au
}

// ClearStatus clears the value of the "status" field.
func (au *ArticleUpdate) ClearStatus() *ArticleUpdate {
	au.mutation.ClearStatus()
	return au
}

// SetTitle sets the "title" field.
func (au *ArticleUpdate) SetTitle(s string) *ArticleUpdate {
	au.mutation.SetTitle(s)
	return au
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableTitle(s *string) *ArticleUpdate {
	if s != nil {
		au.SetTitle(*s)
	}
	return au
}

// SetContent sets the "content" field.
func (au *ArticleUpdate) SetContent(s string) *ArticleUpdate {
	au.mutation.SetContent(s)
	return au
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableContent(s *string) *ArticleUpdate {
	if s != nil {
		au.SetContent(*s)
	}
	return au
}

// SetKeyword sets the "keyword" field.
func (au *ArticleUpdate) SetKeyword(s string) *ArticleUpdate {
	au.mutation.SetKeyword(s)
	return au
}

// SetNillableKeyword sets the "keyword" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableKeyword(s *string) *ArticleUpdate {
	if s != nil {
		au.SetKeyword(*s)
	}
	return au
}

// ClearKeyword clears the value of the "keyword" field.
func (au *ArticleUpdate) ClearKeyword() *ArticleUpdate {
	au.mutation.ClearKeyword()
	return au
}

// SetVisit sets the "visit" field.
func (au *ArticleUpdate) SetVisit(i int) *ArticleUpdate {
	au.mutation.ResetVisit()
	au.mutation.SetVisit(i)
	return au
}

// SetNillableVisit sets the "visit" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableVisit(i *int) *ArticleUpdate {
	if i != nil {
		au.SetVisit(*i)
	}
	return au
}

// AddVisit adds i to the "visit" field.
func (au *ArticleUpdate) AddVisit(i int) *ArticleUpdate {
	au.mutation.AddVisit(i)
	return au
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (au *ArticleUpdate) SetCategoryID(id uint64) *ArticleUpdate {
	au.mutation.SetCategoryID(id)
	return au
}

// SetNillableCategoryID sets the "category" edge to the Category entity by ID if the given value is not nil.
func (au *ArticleUpdate) SetNillableCategoryID(id *uint64) *ArticleUpdate {
	if id != nil {
		au = au.SetCategoryID(*id)
	}
	return au
}

// SetCategory sets the "category" edge to the Category entity.
func (au *ArticleUpdate) SetCategory(c *Category) *ArticleUpdate {
	return au.SetCategoryID(c.ID)
}

// Mutation returns the ArticleMutation object of the builder.
func (au *ArticleUpdate) Mutation() *ArticleMutation {
	return au.mutation
}

// ClearCategory clears the "category" edge to the Category entity.
func (au *ArticleUpdate) ClearCategory() *ArticleUpdate {
	au.mutation.ClearCategory()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ArticleUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ArticleUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ArticleUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ArticleUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *ArticleUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := article.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

func (au *ArticleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(article.Table, article.Columns, sqlgraph.NewFieldSpec(article.FieldID, field.TypeUUID))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(article.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.Status(); ok {
		_spec.SetField(article.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := au.mutation.AddedStatus(); ok {
		_spec.AddField(article.FieldStatus, field.TypeUint8, value)
	}
	if au.mutation.StatusCleared() {
		_spec.ClearField(article.FieldStatus, field.TypeUint8)
	}
	if value, ok := au.mutation.Title(); ok {
		_spec.SetField(article.FieldTitle, field.TypeString, value)
	}
	if value, ok := au.mutation.Content(); ok {
		_spec.SetField(article.FieldContent, field.TypeString, value)
	}
	if value, ok := au.mutation.Keyword(); ok {
		_spec.SetField(article.FieldKeyword, field.TypeString, value)
	}
	if au.mutation.KeywordCleared() {
		_spec.ClearField(article.FieldKeyword, field.TypeString)
	}
	if value, ok := au.mutation.Visit(); ok {
		_spec.SetField(article.FieldVisit, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedVisit(); ok {
		_spec.AddField(article.FieldVisit, field.TypeInt, value)
	}
	if au.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   article.CategoryTable,
			Columns: []string{article.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   article.CategoryTable,
			Columns: []string{article.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{article.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// ArticleUpdateOne is the builder for updating a single Article entity.
type ArticleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArticleMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *ArticleUpdateOne) SetUpdatedAt(t time.Time) *ArticleUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetStatus sets the "status" field.
func (auo *ArticleUpdateOne) SetStatus(u uint8) *ArticleUpdateOne {
	auo.mutation.ResetStatus()
	auo.mutation.SetStatus(u)
	return auo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableStatus(u *uint8) *ArticleUpdateOne {
	if u != nil {
		auo.SetStatus(*u)
	}
	return auo
}

// AddStatus adds u to the "status" field.
func (auo *ArticleUpdateOne) AddStatus(u int8) *ArticleUpdateOne {
	auo.mutation.AddStatus(u)
	return auo
}

// ClearStatus clears the value of the "status" field.
func (auo *ArticleUpdateOne) ClearStatus() *ArticleUpdateOne {
	auo.mutation.ClearStatus()
	return auo
}

// SetTitle sets the "title" field.
func (auo *ArticleUpdateOne) SetTitle(s string) *ArticleUpdateOne {
	auo.mutation.SetTitle(s)
	return auo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableTitle(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetTitle(*s)
	}
	return auo
}

// SetContent sets the "content" field.
func (auo *ArticleUpdateOne) SetContent(s string) *ArticleUpdateOne {
	auo.mutation.SetContent(s)
	return auo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableContent(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetContent(*s)
	}
	return auo
}

// SetKeyword sets the "keyword" field.
func (auo *ArticleUpdateOne) SetKeyword(s string) *ArticleUpdateOne {
	auo.mutation.SetKeyword(s)
	return auo
}

// SetNillableKeyword sets the "keyword" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableKeyword(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetKeyword(*s)
	}
	return auo
}

// ClearKeyword clears the value of the "keyword" field.
func (auo *ArticleUpdateOne) ClearKeyword() *ArticleUpdateOne {
	auo.mutation.ClearKeyword()
	return auo
}

// SetVisit sets the "visit" field.
func (auo *ArticleUpdateOne) SetVisit(i int) *ArticleUpdateOne {
	auo.mutation.ResetVisit()
	auo.mutation.SetVisit(i)
	return auo
}

// SetNillableVisit sets the "visit" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableVisit(i *int) *ArticleUpdateOne {
	if i != nil {
		auo.SetVisit(*i)
	}
	return auo
}

// AddVisit adds i to the "visit" field.
func (auo *ArticleUpdateOne) AddVisit(i int) *ArticleUpdateOne {
	auo.mutation.AddVisit(i)
	return auo
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (auo *ArticleUpdateOne) SetCategoryID(id uint64) *ArticleUpdateOne {
	auo.mutation.SetCategoryID(id)
	return auo
}

// SetNillableCategoryID sets the "category" edge to the Category entity by ID if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableCategoryID(id *uint64) *ArticleUpdateOne {
	if id != nil {
		auo = auo.SetCategoryID(*id)
	}
	return auo
}

// SetCategory sets the "category" edge to the Category entity.
func (auo *ArticleUpdateOne) SetCategory(c *Category) *ArticleUpdateOne {
	return auo.SetCategoryID(c.ID)
}

// Mutation returns the ArticleMutation object of the builder.
func (auo *ArticleUpdateOne) Mutation() *ArticleMutation {
	return auo.mutation
}

// ClearCategory clears the "category" edge to the Category entity.
func (auo *ArticleUpdateOne) ClearCategory() *ArticleUpdateOne {
	auo.mutation.ClearCategory()
	return auo
}

// Where appends a list predicates to the ArticleUpdate builder.
func (auo *ArticleUpdateOne) Where(ps ...predicate.Article) *ArticleUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ArticleUpdateOne) Select(field string, fields ...string) *ArticleUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Article entity.
func (auo *ArticleUpdateOne) Save(ctx context.Context) (*Article, error) {
	auo.defaults()
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ArticleUpdateOne) SaveX(ctx context.Context) *Article {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ArticleUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ArticleUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *ArticleUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := article.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

func (auo *ArticleUpdateOne) sqlSave(ctx context.Context) (_node *Article, err error) {
	_spec := sqlgraph.NewUpdateSpec(article.Table, article.Columns, sqlgraph.NewFieldSpec(article.FieldID, field.TypeUUID))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Article.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, article.FieldID)
		for _, f := range fields {
			if !article.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != article.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(article.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.Status(); ok {
		_spec.SetField(article.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := auo.mutation.AddedStatus(); ok {
		_spec.AddField(article.FieldStatus, field.TypeUint8, value)
	}
	if auo.mutation.StatusCleared() {
		_spec.ClearField(article.FieldStatus, field.TypeUint8)
	}
	if value, ok := auo.mutation.Title(); ok {
		_spec.SetField(article.FieldTitle, field.TypeString, value)
	}
	if value, ok := auo.mutation.Content(); ok {
		_spec.SetField(article.FieldContent, field.TypeString, value)
	}
	if value, ok := auo.mutation.Keyword(); ok {
		_spec.SetField(article.FieldKeyword, field.TypeString, value)
	}
	if auo.mutation.KeywordCleared() {
		_spec.ClearField(article.FieldKeyword, field.TypeString)
	}
	if value, ok := auo.mutation.Visit(); ok {
		_spec.SetField(article.FieldVisit, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedVisit(); ok {
		_spec.AddField(article.FieldVisit, field.TypeInt, value)
	}
	if auo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   article.CategoryTable,
			Columns: []string{article.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   article.CategoryTable,
			Columns: []string{article.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Article{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{article.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
