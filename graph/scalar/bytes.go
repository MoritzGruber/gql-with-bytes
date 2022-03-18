package scalar

import (
	"io"
)

type Bytes []byte

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (value *Bytes) UnmarshalGQL(v interface{}) error {
	// we do not need this
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (value Bytes) MarshalGQL(w io.Writer) {
	w.Write(value)
}
