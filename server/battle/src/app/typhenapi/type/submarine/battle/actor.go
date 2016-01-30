// This file was generated by typhen-api

package battle

import (
	"app/typhenapi/core"
	"errors"
)

var _ = errors.New

// Actor is a kind of TyphenAPI type.
type Actor struct {
	Id       int64      `codec:"id"`
	UserId   int64      `codec:"user_id"`
	Type     *ActorType `codec:"type"`
	Position *Vector    `codec:"position"`
}

// Coerce the fields.
func (t *Actor) Coerce() error {
	if t.Type == nil {
		return errors.New("Type should not be empty")
	}
	if err := t.Type.Coerce(); err != nil {
		return err
	}
	if t.Position == nil {
		return errors.New("Position should not be empty")
	}
	return nil
}

// Bytes creates the byte array.
func (t *Actor) Bytes(serializer typhenapi.Serializer) ([]byte, error) {
	if err := t.Coerce(); err != nil {
		return nil, err
	}

	data, err := serializer.Serialize(t)
	if err != nil {
		return nil, err
	}

	return data, nil
}