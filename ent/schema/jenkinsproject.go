package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// JenkinsProject holds the schema definition for the JenkinsProject entity.
type JenkinsProject struct {
	ent.Schema
}

// Fields of the JenkinsProject.
func (JenkinsProject) Fields() []ent.Field {
	return []ent.Field{
		field.String("project_name"),  // project_name+environment
		field.String("project_value"), // project_name
	}
}

// Edges of the JenkinsProject.
func (JenkinsProject) Edges() []ent.Edge {
	return nil
}
