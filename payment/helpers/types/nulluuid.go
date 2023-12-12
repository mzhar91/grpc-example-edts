package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	uuid "github.com/satori/go.uuid"
)

// NullUUID nullable uuid with custom functions
type NullUUID struct {
	UUID  uuid.UUID
	Valid bool
}

// UnmarshalJSON from json
func (u *NullUUID) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	if s != "" {
		input, err := uuid.FromString(s)
		if err != nil {
			return errors.New("Could not parse UUID")
		}

		u.UUID = input
		u.Valid = true
	}

	return nil
}

// MarshalJSON to json
func (u NullUUID) MarshalJSON() ([]byte, error) {
	if !u.Valid || u.UUID == uuid.Nil {
		return json.Marshal([]byte(nil))
	}

	return json.Marshal(u.UUID)
}

// Value to get value
func (u NullUUID) Value() (driver.Value, error) {
	if !u.Valid {
		return nil, nil
	}

	return u.UUID.Value()
}

// Scan to scan source
func (u *NullUUID) Scan(src interface{}) error {
	if src == nil {
		u.UUID, u.Valid = uuid.Nil, false
		return nil
	}

	u.Valid = true

	return u.UUID.Scan(src)
}

// FromBytes byte to NullUUID
func FromBytes(input []byte) (u NullUUID, err error) {
	uuidInput, err := uuid.FromBytes(input)
	if err != nil {
		return NullUUID{Valid: false}, err
	}

	return NullUUID{UUID: uuidInput, Valid: true}, nil
}

// FromBytesOrNil byte to NullUUID with Nil UUID
func FromBytesOrNil(input []byte) NullUUID {
	uuidInput, err := uuid.FromBytes(input)
	if err != nil {
		return NullUUID{UUID: uuid.Nil, Valid: false}
	}

	return NullUUID{UUID: uuidInput, Valid: true}
}

// FromString string to NullUUID
func FromString(input string) (u NullUUID, err error) {
	var uuidInput uuid.UUID

	err = uuidInput.UnmarshalText([]byte(input))
	if err != nil {
		return NullUUID{Valid: false}, err
	}

	return NullUUID{UUID: uuidInput, Valid: true}, nil
}

// FromStringOrNil returns UUID parsed from string input.
// Same behavior as FromString, but returns a Nil UUID on error.
func FromStringOrNil(input string) uuid.UUID {
	var uuidInput uuid.UUID

	err := uuidInput.UnmarshalText([]byte(input))
	if err != nil {
		return uuid.Nil
	}

	return uuidInput
}

// FromUUID uuid to NullUUID
func FromUUID(input uuid.UUID) (u NullUUID, err error) {
	_, err = uuid.FromBytes(input.Bytes())
	if err != nil {
		return NullUUID{Valid: false}, errors.New("Could not parse UUID")
	}
	return NullUUID{UUID: input, Valid: true}, nil
}
