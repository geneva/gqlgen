// Code generated by github.com/geneva/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/geneva/gqlgen/_examples/scalars/external"
)

type Address struct {
	ID       external.ObjectID `json:"id"`
	Location *Point            `json:"location,omitempty"`
}
