package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive(),
		field.String("name").NotEmpty().Default("User"),
		field.String("password").NotEmpty().Default("Password"),
		field.String("email").NotEmpty().Default("user@email.com"),
		field.Time("create_time").Default(time.Now()),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("userPosts", UserPost.Type),
	}
}
