package bsonx

import (
	"go/types"

	"github.com/colorfuzzy/yo"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// ErrConvertType indicate a convert error
var ErrConvertType = "error convert type"

// D convert map[string]interface{} to bson.D type
func D(data yo.M) bson.D {
	result := bson.D{}

	for key, value := range data {
		switch value.(type) {
		case types.Map:
			valueMap, ok := value.(yo.M)
			if !ok {
				panic(ErrConvertType)
			}

			result = append(result, primitive.E{Key: key, Value: D(valueMap)})
		case types.Slice:
			valueSlice, ok := value.([]yo.Any)
			if !ok {
				panic(ErrConvertType)
			}

			result = append(result, primitive.E{Key: key, Value: sliceToArray(valueSlice)})
		default:
			result = append(result, primitive.E{Key: key, Value: value})
		}
	}

	return result
}

func sliceToArray(data []yo.Any) bson.A {
	result := bson.A{}
	for _, item := range data {
		switch item.(type) {
		case types.Map:
			itemMap, ok := item.(yo.M)
			if !ok {
				panic(ErrConvertType)
			}

			result = append(result, D(itemMap))
		case types.Slice:
			itemSlice, ok := item.([]yo.Any)
			if !ok {
				panic(ErrConvertType)
			}

			result = append(result, sliceToArray(itemSlice))
		default:
			result = append(result, item)
		}
	}
	return result
}
