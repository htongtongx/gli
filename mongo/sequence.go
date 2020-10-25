package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Counter struct {
	ID    string `bson:"_id"`
	Value int64  `bson:"value"`
}

func (u *Counter) Table() string {
	return "counter"
}

// 是否自增id
func (u *Counter) IsAI() bool {
	return true
}

func (u *Counter) SetID(id int64) {

}

func (u *Counter) UniqueKey() string {
	return "_id"
}

func (u *Counter) UniqueValue() interface{} {
	return u.ID
}

func (m *Mongo) NextID(dbname, collname, field string) (id int64, err error) {
	var c Counter
	collection, ctx, cancel := m.getCollection(dbname, collname)
	defer cancel()
	filter := bson.M{"_id": field}
	update := bson.M{"$inc": bson.M{"value": 1}}
	err = collection.FindOneAndUpdate(ctx, filter, update,
		options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(true)).
		Decode(&c)
	return c.Value, err
}

// func (m *Mongo) Max(dbname, collname string) (id int, err error) {
// 	id = 0
// 	sort := bson.M{"_id": -1}
// 	results, err := m.SelectMutil(dbname, collname, nil, sort, 1, 1)
// 	if err != nil {
// 		return
// 	}
// 	if len(*results) > 0 {
// 		switch *results[0]["_id"].(type) {
// 		case int:
// 			id = results[0]["_id"].(int)
// 		}
// 	} else {
// 		id = 0
// 	}
// 	return
// }
