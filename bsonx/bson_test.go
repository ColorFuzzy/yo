package bsonx

import (
	"testing"
	"time"

	"github.com/colorfuzzy/zo"
)

func TestToBson(t *testing.T) {
	D(zo.Map{
		"sign": "sign",
		"status": 1,
		"create_at": time.Now(),
	})

	D(zo.Map{
		"status": zo.Map{
			"$in": []int{1, 2},
		},
	})

	D(zo.Map{
		"$set": zo.Map{
			"status": 1,
		},
		"$inc": zo.Map{
			"err_count": 1,
		},
	})
}
