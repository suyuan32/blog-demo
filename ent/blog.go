// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/ent/blog"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Blog is the model entity for the Blog schema.
type Blog struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Blog) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case blog.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Blog fields.
func (b *Blog) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case blog.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Blog.
// This includes values selected through modifiers, order, etc.
func (b *Blog) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// Update returns a builder for updating this Blog.
// Note that you need to call Blog.Unwrap() before calling this method if this Blog
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Blog) Update() *BlogUpdateOne {
	return NewBlogClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Blog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Blog) Unwrap() *Blog {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Blog is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Blog) String() string {
	var builder strings.Builder
	builder.WriteString("Blog(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Blogs is a parsable slice of Blog.
type Blogs []*Blog