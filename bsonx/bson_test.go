package bsonx

import (
	"testing"
	"time"

	"github.com/colorfuzzy/zo"
)

func TestToBson(t *testing.T) {
	ToBson(zo.Map{
		"sign": "sign",
		"status": 1,
		"create_at": time.Now(),
	})

	ToBson(zo.Map{
		"status": zo.Map{
			"$in": []int{1, 2},
		},
	})

	ToBson(zo.Map{
		"$set": zo.Map{
			"status": 1,
		},
		"$inc": zo.Map{
			"err_count": 1,
		},
	})
}
