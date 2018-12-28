**github.com/mongodb/mongo-go-driver/mongo** is fucking hard to use!

# Easy library to use bson.D

`bsonx.ToBson` convert `map[string]interface{}` to `bson.D`.

## Before zo

``` go

// case 1
filter := bson.D{{
    "status",
    bson.D{{
        "$in",
        bson.A{StatusNew, StatusRenew},
    }},
}}
err := c.FindOne(ctx, filter).Decode(&msg)

// case 2
filter := bson.D{{"sign", msg.Sign}}
update := bson.D{
    {"$set", bson.D{{"status", status}}},
    {"$inc", bson.D{{"err_count", 1}}},
}
_, err := c.UpdateOne(ctx, filter, update)
```

## After zo

``` go

// case 1
filter := bsonx.ToBson(zo.Map{
    "status": zo.Map{
        "$in": []int{StatusNew, StatusRenew},
    },
})
err := c.FindOne(ctx, filter).Decode(&msg)

// case 2
filter := bsonx.ToBson(zo.Map{"sign": msg.Sign})
update := bsonx.ToBson(zo.Map{
    "$set": zo.Map{"status": status},
    "$inc": zo.Map{"err_count": 1},
})
_, err := c.UpdateOne(ctx, filter, update)
```