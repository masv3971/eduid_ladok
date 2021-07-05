package model

import (
	"github.com/elimity-com/scim/optional"
	"github.com/elimity-com/scim/schema"
)

// UserSchema scim
var UserSchema = schema.Schema{
	ID:          "urn:ietf:params:scim:schemas:core:2.0:User",
	Name:        optional.NewString("User"),
	Description: optional.NewString("User Account"),
	Attributes: []schema.CoreAttribute{
		schema.SimpleCoreAttribute(schema.SimpleStringParams(schema.StringParams{
			Name:       "userName",
			Required:   true,
			Uniqueness: schema.AttributeUniquenessServer(),
		})),
	},
}

// UserExtension scim extention
var UserExtension = schema.Schema{
	ID:          "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
	Name:        optional.NewString("EnterpriseUser"),
	Description: optional.NewString("Enterprise User"),
	Attributes: []schema.CoreAttribute{
		schema.SimpleCoreAttribute(schema.SimpleStringParams(schema.StringParams{
			Name: "employeeNumber",
		})),
		schema.SimpleCoreAttribute(schema.SimpleStringParams(schema.StringParams{
			Name: "organization",
		})),
	},
}
