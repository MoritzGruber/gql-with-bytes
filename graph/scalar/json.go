package scalar

import (
	"io"
)

type JSON []byte

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (value *JSON) UnmarshalGQL(v interface{}) error {
	// we do not need this
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (value JSON) MarshalGQL(w io.Writer) {
	w.Write(value)
}
