// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/contrib/entgql/internal/todo/ent/schema/schematype"
	"entgo.io/contrib/entgql/internal/todouuid/ent/category"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Category is the model entity for the Category schema.
type Category struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Status holds the value of the "status" field.
	Status category.Status `json:"status,omitempty"`
	// Config holds the value of the "config" field.
	Config *schematype.CategoryConfig `json:"config,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration time.Duration `json:"duration,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CategoryQuery when eager-loading is set.
	Edges CategoryEdges `json:"edges"`
}

// CategoryEdges holds the relations/edges for other nodes in the graph.
type CategoryEdges struct {
	// Todos holds the value of the todos edge.
	Todos []*Todo `json:"todos,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TodosOrErr returns the Todos value or an error if the edge
// was not loaded in eager-loading.
func (e CategoryEdges) TodosOrErr() ([]*Todo, error) {
	if e.loadedTypes[0] {
		return e.Todos, nil
	}
	return nil, &NotLoadedError{edge: "todos"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Category) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case category.FieldConfig:
			values[i] = new(schematype.CategoryConfig)
		case category.FieldDuration:
			values[i] = new(sql.NullInt64)
		case category.FieldText, category.FieldStatus:
			values[i] = new(sql.NullString)
		case category.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Category", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Category fields.
func (c *Category) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case category.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case category.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				c.Text = value.String
			}
		case category.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				c.Status = category.Status(value.String)
			}
		case category.FieldConfig:
			if value, ok := values[i].(*schematype.CategoryConfig); !ok {
				return fmt.Errorf("unexpected type %T for field config", values[i])
			} else if value != nil {
				c.Config = value
			}
		case category.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				c.Duration = time.Duration(value.Int64)
			}
		}
	}
	return nil
}

// QueryTodos queries the "todos" edge of the Category entity.
func (c *Category) QueryTodos() *TodoQuery {
	return (&CategoryClient{config: c.config}).QueryTodos(c)
}

// Update returns a builder for updating this Category.
// Note that you need to call Category.Unwrap() before calling this method if this Category
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Category) Update() *CategoryUpdateOne {
	return (&CategoryClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Category entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Category) Unwrap() *Category {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Category is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Category) String() string {
	var builder strings.Builder
	builder.WriteString("Category(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", text=")
	builder.WriteString(c.Text)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", c.Status))
	builder.WriteString(", config=")
	builder.WriteString(fmt.Sprintf("%v", c.Config))
	builder.WriteString(", duration=")
	builder.WriteString(fmt.Sprintf("%v", c.Duration))
	builder.WriteByte(')')
	return builder.String()
}

// Categories is a parsable slice of Category.
type Categories []*Category

func (c Categories) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
