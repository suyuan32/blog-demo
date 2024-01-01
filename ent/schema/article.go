package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment("文章标题").Annotations(entsql.WithComments(true)),
		field.String("content").Comment("文章内容").Annotations(entsql.WithComments(true)),
		field.String("keyword").Comment("关键字").Optional().Annotations(entsql.WithComments(true)),
		field.Int("visit").Comment("浏览量").Default(0).Annotations(entsql.WithComments(true)),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("category", Category.Type).Unique(),
	}
}

// Mixin of the Article.
func (Article) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
	}
}

// Indexes of the Article.
func (Article) Indexes() []ent.Index {
	return nil
}

// Annotations of the Article
func (Article) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("blog_article"),
	}
}
