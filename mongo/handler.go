package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SortType string

const (
	ASC  SortType = "ASC"
	DESC SortType = "DESC"
)

func (m *Mongo) Add(iMongo IMongo) (err error) {
	if iMongo.IsAI() {
		var id int64
		id, err = m.NextID(m.c.DB, m.c.Sequence, iMongo.Table())
		if err != nil {
			return
		}
		iMongo.SetID(id)
	}
	err = m.InsertOne(m.c.DB, iMongo.Table(), iMongo)
	return
}

func (m *Mongo) SimpleGet(iMongo IMongo, k string, v interface{}) (has bool, err error) {
	filter := bson.M{k: v}
	return m.SelectOne(m.c.DB, iMongo.Table(), filter, iMongo)
}

func (m *Mongo) Get(iMongo IMongo, filter interface{}) (has bool, err error) {
	has, err = m.SelectOne(m.c.DB, iMongo.Table(), filter, iMongo)
	return
}

//在使用此简化函数时，请确保实体的UniqueKey()和UniqueValue()有值
func (m *Mongo) SimpleUpdate(iMongo IMongo, k string, v interface{}, upsert ...bool) (ur *mongo.UpdateResult, err error) {
	var vData = make(map[string]interface{})
	vData[k] = v
	filter := bson.M{iMongo.UniqueKey(): iMongo.UniqueValue()}
	ur, err = m.UpdateOne(m.c.DB, iMongo.Table(), filter, vData, upsert...)
	return
}

func (m *Mongo) Update(iMongo IMongo, filter bson.M, v map[string]interface{}, upsertA ...bool) (ur *mongo.UpdateResult, err error) {
	ur, err = m.UpdateOne(m.c.DB, iMongo.Table(), filter, v, upsertA...)
	return
}

func (m *Mongo) UpdateSelf(iMongo IMongo, upsert ...bool) (ur *mongo.UpdateResult, err error) {
	filter := bson.M{iMongo.UniqueKey(): iMongo.UniqueValue()}
	ur, err = m.UpdateOne(m.c.DB, iMongo.Table(), filter, iMongo, upsert...)
	return
}

func (m *Mongo) UpdateAll(iMongo IMongo, filter interface{}, v interface{}, upsertA ...bool) (ur *mongo.UpdateResult, err error) {
	ur, err = m.UpdateMutil(m.c.DB, iMongo.Table(), filter, v, upsertA...)
	return
}

//使用当前请对UniqueKey和UniqueValue赋值
func (m *Mongo) SimpleDel(iMongo IMongo) (dr *mongo.DeleteResult, err error) {
	filter := bson.M{iMongo.UniqueKey(): iMongo.UniqueValue()}
	dr, err = m.DeleteOne(m.c.DB, iMongo.Table(), filter)
	return
}

func (m *Mongo) Del(iMongo IMongo, filter bson.M) (dr *mongo.DeleteResult, err error) {
	dr, err = m.DeleteOne(m.c.DB, iMongo.Table(), filter)
	return
}

func (m *Mongo) GetList(iMongo IMongo, filter interface{}, sort bson.M, page, limit int64, results interface{}, backTotal bool) (total int64, err error) {
	//其中 1 为升序排列，而-1是用于降序排列
	// sort := bson.M{"_id": -1}
	if filter == nil {
		filter = bson.M{}
	}
	if err = m.SelectMutil(m.c.DB, iMongo.Table(), filter, sort, page, limit, results); err != nil {
		return
	}
	if backTotal {
		total, err = m.Count(iMongo, filter)
	}
	return
}

func (m *Mongo) Count(iMongo IMongo, filter interface{}) (int64, error) {
	return m.RowsCount(m.c.DB, iMongo.Table(), filter)
}

//值自增
//field要自增的字段，和自增的值
func (m *Mongo) Inc(iMongo IMongo, filter interface{}, field string, value int) (newValue interface{}, err error) {
	update := bson.M{"$inc": bson.M{field: value}}
	data, err := m.FindOneAndUpdate(m.c.DB, iMongo.Table(), filter, update)
	return data[field], err
}
