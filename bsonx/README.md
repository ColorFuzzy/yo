**github.com/mongodb/mongo-go-driver/mongo** is fucking hard to use!

# Easy library to use bson.D

`bsonx.ToBson` convert `map[string]interface{}` to `bson.D`.

## Before yo

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
filter := bsonx.D(yo.M{
    "status": yo.M{
        "$in": []int{StatusNew, StatusRenew},
    },
})
err := c.FindOne(ctx, filter).Decode(&msg)

// case 2
filter := bsonx.D(yo.M{"sign": msg.Sign})
update := bsonx.D(yo.M{
    "$set": yo.M{"status": status},
    "$inc": yo.M{"err_count": 1},
})
_, err := c.UpdateOne(ctx, filter, update)
```