package types

import "github.com/99designs/gqlgen/graphql"

type ID string

func NewID(id string) *ID {
	if id == "" {
		return nil
	}

	v := ID(id)
	return &v
}

func MarshalID(v ID) graphql.Marshaler {
	return graphql.MarshalString(string(v))
}

func UnmarshalID(v interface{}) (ID, error) {
	s, err := graphql.UnmarshalString(v)
	if err != nil {
		return "", err
	}
	return ID(s), err
}
