// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/MoritzGruber/gql-with-bytes/graph/scalar"
)

type GetSomethingResponse struct {
	SomeAny         interface{} `json:"someAny"`
	SomeObj         scalar.JSON `json:"someObj"`
	SomeBool        scalar.JSON `json:"someBool"`
	SomeString      scalar.JSON `json:"someString"`
	SomeInt         scalar.JSON `json:"someInt"`
	SomeFloat       scalar.JSON `json:"someFloat"`
	SomeWillFail    scalar.JSON `json:"someWillFail"`
	SomeWillFailToo scalar.JSON `json:"someWillFailToo"`
}

type GetUsageOfBlueprintFieldRequestInput struct {
	ProjectID        *string `json:"projectId"`
	BlueprintFieldID *string `json:"blueprintFieldId"`
	EnvironmentID    *string `json:"environmentId"`
	BlueprintBranch  *string `json:"blueprintBranch"`
}
