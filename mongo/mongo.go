package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/htongtongx/gli/conf"
)

var Client *Mongo

type Mongo struct {
	Cli  *mongo.Client
	Node string
	Pwd  string
	User string
	c    *conf.MongoConf
}

type DataFileVersion struct {
	Major int
	Minor int
}

type MonDbStatInfo struct {
	Db          string //当前数据库
	Collections int32  //当前数据库多少表
	Objects     int64  //当前数据库所有表多少条数据
	AvgObjSize  int64  //每条数据的平均大小
	DataSize    int64  //所有数据的总大小
	StorageSize int64  //所有数据占的磁盘大小
	NumExtents  int
	Indexes     int64 //索引数
	IndexSize   int64 //索引大小
	FileSize    int64 //预分配给数据库的文件大小
	Ok          int16
}

func NewMongo(c *conf.MongoConf) (m *Mongo, err error) {
	if !c.Verify() {
		return
	}
	url := fmt.Sprintf("mongodb://%s:%s@%s/?authSource=%s", c.User, c.Pwd, c.Node, c.Auth)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	m = new(Mongo)
	m.c = c
	m.Cli, err = mongo.Connect(ctx, options.Client().ApplyURI(url).SetMaxPoolSize(20))
	return
}
func (m *Mongo) DB() string {
	return m.c.DB
}

func (m *Mongo) SelectOne(dbname, collname string, filter interface{}, v interface{}) (hasResult bool, err error) {
	hasResult = true
	collection, ctx, cancel := m.getCollection(dbname, collname)
	defer cancel()
	err = collection.FindOne(ctx, filter).Decode(v)
	if err == mongo.ErrNoDocuments {
		hasResult = false
		err = nil
	}
	return
}

// Parameters:
func (m *Mongo) InsertOne(dbname, collname string, v interface{}) (err error) {
	collection, ctx, cancel := m.getCollection(dbname, collname)
	defer cancel()
	_, err = collection.InsertOne(ctx, v)
	return
}

func (m *Mongo) SaveMulit(dbname, collname string, v []interface{}) (err error) {
	collection, ctx, cancel := m.getCollection(dbname, collname)
	defer cancel()

	// dbcollection := fmt.Sprintf("%s.%s", dbname, collname)
	// db.RunCommand(context, bson.M{"enablesharding": dbname})
	// db.RunCommand(context, bson.M{"shardcollection": dbcollection, "key": "_id"})

	_, err = collection.InsertMany(ctx, v)
	return
}

func (m *Mongo) getCollection(dbname, collname string) (collection *mongo.Collection, ctx context.Context, cancel context.CancelFunc) {
	db, ctx, cancel := m.getDB(dbname)
	collection = db.Collection(collname)
	return
}

func (m *Mongo) GetCollection(dbname, collname string) (collection *mongo.Collection, ctx context.Context, cancel context.CancelFunc) {
	return m.getCollection(dbname, collname)
}

func (m *Mongo) getDB(dbname string) (db *mongo.Database, ctx context.Context, cancel context.CancelFunc) {
	db = m.Cli.Database(dbname)
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	return
}

func (m *Mongo) UpdateOne(dbname, collname string, filter interface{}, v interface{}, upsertA ...bool) (ur *mongo.UpdateResult, err error) {
	collection, ctx, cancel := m.getCollection(dbname, collname)
	defer cancel()

	update := bson.M{
		"$set": v,
	}
	upsert := true
	if len(upsertA) > 0 {
		upsert = upsertA[0]
	}
	ur, err = collection.UpdateOne(ctx, filter, update, &options.UpdateOptions{Upsert: &upsert})
	return
}

func (m *Mongo) DeleteOne(dbname, collname string, filter interface{}) (dr *mongo.DeleteResult, err error) {
	collection, ctx, cancel := m.getCollection(dbname, collname)
	defer cancel()

	dr, err = collection.DeleteOne(ctx, filter)
	return
}

func (m *Mongo) UpdateMutil(dbname, collname string, filter interface{}, v interface{}, upsertA ...bool) (ur *mongo.UpdateResult, err error) {
	collection, ctx, cancel := m.getCollection(dbname, collname)
	defer cancel()
	upsert := true
	if len(upsertA) > 0 {
		upsert = upsertA[0]
	}
	ur, err = collection.UpdateMany(ctx, filter, v, &options.UpdateOptions{Upsert: &upsert})
	return
}

// Parameters:
//  - page和limite都为0时获取全部数据
func (m *Mongo) SelectMutil(dbname, collname string, filter interface{}, sort bson.M, page, limit int64, results interface{}) (err error) {
	collection, ctx, cancel := m.getCollection(dbname, collname)
	defer cancel()
	findOptions := options.Find()
	if sort != nil {
		findOptions = findOptions.SetSort(sort)
	}
	if !(page == 0 && limit == 0) {
		skip := (page - 1) * limit
		findOptions = findOptions.SetLimit(limit).SetSkip(skip)
	}

	c, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return
	}
	err = c.All(ctx, results)
	return
}

func (m *Mongo) RowsCount(dbname, collname string, filter interface{}) (count int64, err error) {
	collection, ctx, cancel := m.getCollection(dbname, collname)
	defer cancel()
	count, err = collection.CountDocuments(ctx, filter)
	return
}

func (m *Mongo) GetDbStatus(dbname string) (ret *MonDbStatInfo, err error) {
	db, ctx, cancel := m.getDB(dbname)
	defer cancel()

	result := db.RunCommand(ctx, bson.M{"dbStats": 1})
	err = result.Err()
	if result.Err() != nil {
		return
	}
	err = result.Decode(&ret)
	return
}

func (m *Mongo) GroupBy(dbname, collname string, pipeline interface{}) (*[]map[string]interface{}, error) {
	// pipeline := []bson.M{bson.M{"$group": bson.M{"_id": bson.M{"postuid": "$postuid", "createdby": "$createdby"},
	// "createdon": bson.M{"$max": "$createdon.t"}, "uid": bson.M{"$max": "$_id"}, "readtime": bson.M{"$sum": "$readtime"}}}}
	// pipeline := []bson.M{bson.M{"$group": bson.M{"_id": bson.M{"postuid": "$postuid", "createdby": "$createdby"}}}}
	collection, ctx, cancel := m.getCollection(dbname, collname)
	defer cancel()
	cur, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	// results := ]interface{}
	results := []map[string]interface{}{}
	for cur.Next(ctx) {
		line := make(map[string]interface{})
		cur.Decode(&line)
		results = append(results, line)
	}
	cur.Close(context.TODO())
	return &results, nil
}

func (m *Mongo) FindOneAndUpdate(dbname, collname string, filter interface{}, update bson.M) (count map[string]interface{}, err error) {
	data := make(map[string]interface{})
	collection, ctx, cancel := m.getCollection(dbname, collname)
	defer cancel()
	err = collection.FindOneAndUpdate(ctx, filter, update,
		options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(true)).Decode(data)
	return data, err
}
