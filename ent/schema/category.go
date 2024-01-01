package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment("栏目标题").Annotations(entsql.WithComments(true)),
		field.String("remark").Comment("备注").Annotations(entsql.WithComments(true)),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("articles", Article.Type).Ref("category"),
	}
}

// Mixin of the Category.
func (Category) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

// Indexes of the Category.
func (Category) Indexes() []ent.Index {
	return nil
}

// Annotations of the Category
func (Category) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("blog_category"),
	}
}
