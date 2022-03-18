package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/MoritzGruber/gql-with-bytes/graph/generated"
	"github.com/MoritzGruber/gql-with-bytes/graph/model"
)

func (r *queryResolver) GetSomething(ctx context.Context) (*model.GetSomethingResponse, error) {
	type ColorGroup struct {
		ByteSlice     []byte
		SingleByte    byte
		IntSlice      []int
		SingleBoolean bool
	}
	group := ColorGroup{
		ByteSlice:     []byte{0, 0, 0, 1, 2, 3},
		SingleByte:    10,
		IntSlice:      []int{0, 0, 0, 1, 2, 3},
		SingleBoolean: true,
	}
	jsonObj, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}

	var f int = 52452356235      // int
	var s = big.NewInt(int64(f)) // int to big Int
	var bInt64 = s.Bytes()       // big Int to bytes

	jsonInt, _ := json.Marshal(52452356235)
	jsonBool, _ := json.Marshal(true)
	jsonFloat, _ := json.Marshal(6.4)
	jsonString, _ := json.Marshal("something-to-see")

	return &model.GetSomethingResponse{
		SomeAny:         &jsonObj,
		SomeObj:         jsonObj,
		SomeInt:         jsonInt,
		SomeString:      jsonString,
		SomeFloat:       jsonFloat,
		SomeBool:        jsonBool,
		SomeWillFail:    jsonObj[:1],
		SomeWillFailToo: bInt64,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
