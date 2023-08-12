package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserPost holds the schema definition for the UserPost entity.
type UserPost struct {
	ent.Schema
}

// Fields of the UserPost.
func (UserPost) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive(),
		field.String("post_text"),
		field.JSON("post_images", map[string]string{}),
		field.Time("create_time").Immutable().Default(time.Now()),
	}
}

// Edges of the UserPost.
func (UserPost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_id", User.Type).Ref("userPosts").Unique(),
	}
}
