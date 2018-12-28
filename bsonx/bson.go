package bsonx

import (
	"go/types"

	"github.com/colorfuzzy/zo"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

var ErrConvertType = "error convert type"

// ToBson convert map[string]interface{} to bson.D type
func ToBson(data zo.Map) bson.D {
	result := bson.D{}

	for key, value := range data {
		switch value.(type) {
		case types.Map:
			valueMap, ok := value.(zo.Map)
			if !ok {
				panic(ErrConvertType)
			}

			result = append(result, primitive.E{Key: key, Value: ToBson(valueMap)})
		case types.Slice:
			valueSlice, ok := value.([]zo.Any)
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

func sliceToArray(data []zo.Any) bson.A {
	result := bson.A{}
	for _, item := range data {
		switch item.(type) {
		case types.Map:
			itemMap, ok := item.(zo.Map)
			if !ok {
				panic(ErrConvertType)
			}

			result = append(result, ToBson(itemMap))
		case types.Slice:
			itemSlice, ok := item.([]zo.Any)
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
