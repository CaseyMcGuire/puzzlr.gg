// Code generated by ent, DO NOT EDIT.

package codegen

import (
	"puzzlr.gg/src/server/db/ent/codegen/user"
	"puzzlr.gg/src/server/db/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[0].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescHashedPassword is the schema descriptor for hashed_password field.
	userDescHashedPassword := userFields[1].Descriptor()
	// user.HashedPasswordValidator is a validator for the "hashed_password" field. It is called by the builders before save.
	user.HashedPasswordValidator = userDescHashedPassword.Validators[0].(func(string) error)
}
