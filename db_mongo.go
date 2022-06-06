package fmk

import (
	"context"
	"fmt"
	"github.com/naamancurtis/mongo-go-struct-to-bson/mapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
	"time"
)

// mockgen -source=db.go -destination=mocks/mock_db.go  -self_package IDb, IModel

type MDb struct {
	*mongo.Database
}

type MModel struct {
	*mongo.Collection
}

type MPagination struct {
	Limit int64
	Next  string
}

type MSortType int

const (
	ASC  MSortType = 1
	DESC           = -1
)

// NewDb creates Database instance, connects to mongo and complies
// with IMongoDatabase interface
func NewDb(dbUrl string, dbName string) (*MDb, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))
	if err != nil {
		return nil, NewErrorBuilder().
			Category(DBErrorCategory).
			Message(err.Error()).
			Build()
	}

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())

	return &MDb{
		client.Database(dbName),
	}, nil
}

func (d *MDb) Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := d.Client().Disconnect(ctx); err != nil {
		return NewErrorBuilder().
			Category(DBErrorCategory).
			Message(err.Error()).
			Build()
	}
	return nil
}

func (d *MDb) GetModel(colName string) *MModel {
	return &MModel{d.Collection(colName)}
}

func (m *MModel) Insert(doc interface{}) (string, error) {

	insertOptions := options.InsertOne()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	docBson := mapper.ConvertStructToBSONMap(doc, nil)

	singleResult, err := m.InsertOne(ctx, docBson, insertOptions)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", NewErrorBuilder().
				Category(DBErrorCategory).
				StatusCode(http.StatusBadRequest).
				Message("record duplicated").
				Build()
		} else {
			return "", NewErrorBuilder().
				Category(DBErrorCategory).
				StatusCode(http.StatusInternalServerError).
				Message(fmt.Sprintf(err.Error())).
				Build()
		}
	}

	return singleResult.InsertedID.(string), nil
}

func (m *MModel) UpdateByFilter(filter interface{}, toChange interface{}) error {
	filterBson := mapper.ConvertStructToBSONMap(filter, nil)

	update := bson.D{
		{"$set", toChange},
	}
	after := options.After // for returning updated document
	returnOpt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	re := m.FindOneAndUpdate(ctx, filterBson, update, &returnOpt)
	if re.Err() != nil {
		return NewErrorBuilder().
			Category(DBErrorCategory).
			Message(re.Err().Error()).
			StatusCode(http.StatusInternalServerError).
			Build()
	}

	if err := re.Decode(toChange); err != nil {
		return NewErrorBuilder().
			Category(DBErrorCategory).
			StatusCode(http.StatusInternalServerError).
			Message(err.Error()).
			Build()
	}

	return nil
}

func (m *MModel) FindByFilter(filter interface{}, sort interface{}, pagination MPagination, docs interface{}) error {

	filterBson := mapper.ConvertStructToBSONMap(filter, nil)
	if filterBson != nil {
		appendFilter(bson.M{
			"_id": bson.M{"$gt": pagination.Next},
		}, filterBson)
	}

	options := options.Find()
	options.SetSort(sort)
	options.SetLimit(pagination.Limit)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	c, err := m.Find(ctx, filterBson)
	if err != nil {
		return NewErrorBuilder().
			Category(DBErrorCategory).
			Message(err.Error()).
			Build()
	}

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := c.All(ctx, docs); err != nil {
		return NewErrorBuilder().
			Category(DBErrorCategory).
			Message(err.Error()).
			Build()
	}

	return nil
}

func (m *MModel) DeleteByFilter(filter interface{}) error {

	filterBson := mapper.ConvertStructToBSONMap(filter, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := m.DeleteOne(ctx, filterBson)
	if err != nil {
		return NewErrorBuilder().
			Category(DBErrorCategory).
			Message(err.Error()).
			Build()
	}

	return nil
}

func appendFilter(filter bson.M, additionalFilter interface{}) {
	bsonFilter := mapper.ConvertStructToBSONMap(additionalFilter, nil)
	for k, v := range bsonFilter {
		if v != nil {
			filter[k] = v
		}
	}
}
