//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type Sessions struct {
	UserID    string `sql:"primary_key"`
	Token     string `sql:"primary_key"`
	CreatedAt time.Time
}
