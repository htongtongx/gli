package mongo

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Filter struct {
	D bson.D
}

func NewFilter() (f *Filter) {
	f = &Filter{}
	f.D = bson.D{}
	return
}

// pass的参数表示无论当前值是否是0都加入筛选
func (f *Filter) Append(k string, v interface{}, pass ...bool) *Filter {
	onePass := len(pass) > 0 && pass[0]
	canA := f.checkPass(v, onePass)

	if canA {
		f.D = append(f.D, bson.E{Key: k, Value: v})
	}
	return f
}

func (f *Filter) checkPass(v interface{}, pass bool) bool {
	var canA = false
	onePass := pass
	switch v.(type) {
	case string:
		canA = v.(string) != "" || onePass
	case int8:
		canA = v.(int8) > 0 || onePass
	case int:
		canA = v.(int) > 0 || onePass
	case int32:
		canA = v.(int32) > 0 || onePass
	case int64:
		canA = v.(int64) > 0 || onePass
	case bool:
		canA = true
	}
	return canA
}

func (f *Filter) Range(k string, min, max interface{}, contain ...bool) *Filter {
	if !(f.checkPass(min, false) && f.checkPass(max, false)) {
		return f
	}
	if len(contain) > 0 && contain[0] {
		f.D = append(f.D, bson.E{Key: k, Value: bson.M{"$gte": min, "$lte": max}})
	} else {
		f.D = append(f.D, bson.E{Key: k, Value: bson.M{"$gt": min, "$lt": max}})
	}
	return f
}

//pass等于true表示包含0
func (f *Filter) GT(k string, v interface{}, pass ...bool) *Filter {
	can := f.checkPass(v, len(pass) > 0 && pass[0])
	if can {
		f.D = append(f.D, bson.E{Key: k, Value: bson.M{"$gt": v}})
	}
	return f
}

func (f *Filter) GTE(k string, v interface{}, pass ...bool) *Filter {
	can := f.checkPass(v, len(pass) > 0 && pass[0])
	if can {
		f.D = append(f.D, bson.E{Key: k, Value: bson.M{"$gte": v}})
	}
	return f
}

func (f *Filter) LT(k string, v interface{}, pass ...bool) *Filter {
	can := f.checkPass(v, len(pass) > 0 && pass[0])
	if can {
		f.D = append(f.D, bson.E{Key: k, Value: bson.M{"$lt": v}})
	}
	return f
}

func (f *Filter) LTE(k string, v interface{}, pass ...bool) *Filter {
	can := f.checkPass(v, len(pass) > 0 && pass[0])
	if can {
		f.D = append(f.D, bson.E{Key: k, Value: bson.M{"$lte": v}})
	}
	return f
}

//不等于
func (f *Filter) NE(k string, v interface{}, pass ...bool) *Filter {
	can := f.checkPass(v, len(pass) > 0 && pass[0])
	if can {
		f.D = append(f.D, bson.E{Key: k, Value: bson.M{"$ne": v}})
	}
	return f

}

func (f *Filter) In(key string, l []interface{}) *Filter {
	if len(l) == 0 {
		return f
	}
	ba := bson.A{}
	for _, v := range l {
		ba = append(ba, v)
	}
	// f.D = append(f.D, bson.E{Key: "$in", Value: l})
	f.D = append(f.D, bson.E{Key: key, Value: bson.M{"$in": ba}})
	return f
}

func (f *Filter) Nin(key string, l []interface{}) *Filter {
	if len(l) == 0 {
		return f
	}
	ba := bson.A{}
	for _, v := range l {
		ba = append(ba, v)
	}
	// f.D = append(f.D, bson.E{Key: key, Value: bson.D{bson.M{"$nin": ba}}})
	f.D = append(f.D, bson.E{Key: key, Value: bson.M{"$nin": ba}})
	// f.D = append(f.D, bson.E{Key: key, Value: bson.M{"$nin": ba}})
	// f.D = append(f.D, bson.E{Key: "$nin", Value: l})
	return f
}

func (f *Filter) Search(k, regex string) *Filter {
	if regex == "" {
		return f
	}
	f.D = append(f.D, bson.E{Key: k, Value: bson.M{"$regex": regex, "$options": "$i"}})
	return f
}

//date的参数格式请务必为“2006-01”
func (f *Filter) AddMonth(k, date string) *Filter {
	if date == "" {
		return f
	}
	curMon, err := time.Parse("2006-01", date)
	if err != nil {
		log.Println("AddMonth to fail:", date)
		return f
	}

	nextMon := curMon.AddDate(0, 1, 0)
	return f.GT(k, curMon.Unix()).LT(k, nextMon.Unix())
}
