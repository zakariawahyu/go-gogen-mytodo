package withmongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"zakariawahyu.com/go-gogen-mytodo/domain_todocore/model/entity"
	"zakariawahyu.com/go-gogen-mytodo/domain_todocore/model/vo"
	"zakariawahyu.com/go-gogen-mytodo/shared/gogen"
	"zakariawahyu.com/go-gogen-mytodo/shared/infrastructure/config"
	"zakariawahyu.com/go-gogen-mytodo/shared/infrastructure/logger"
)

type gateway struct {
	log     logger.Logger
	appData gogen.ApplicationData
	config  *config.Config
	client  *mongo.Client
}

// NewGateway ...
func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *gateway {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(ClientUrl(cfg, appData)))
	if err != nil {
		panic(err)
	}
	return &gateway{
		log:     log,
		appData: appData,
		config:  cfg,
		client:  client,
	}
}

func ClientUrl(cfg *config.Config, appData gogen.ApplicationData) string {
	connPattern := "mongodb://%v:%v@%v:%v"
	if cfg.Servers[appData.AppName].Database.Username == "" {
		connPattern = "mongodb://%s%s%v:%v"
	}
	clientUrl := fmt.Sprintf(connPattern,
		cfg.Servers[appData.AppName].Database.Username,
		cfg.Servers[appData.AppName].Database.Password,
		cfg.Servers[appData.AppName].Database.Host,
		cfg.Servers[appData.AppName].Database.Port,
	)
	return clientUrl
}

func (r *gateway) FindAllTodo(ctx context.Context, page, size int) ([]*entity.Todo, int64, error) {
	r.log.Info(ctx, "called")

	coll := r.client.Database(r.config.Servers[r.appData.AppName].Database.Database).Collection("todo")
	filter := bson.D{{}}

	skip := int64(size * (page - 1))
	limit := int64(size)

	countOpts := options.CountOptions{
		Limit: &limit,
		Skip:  &skip,
	}

	count, err := coll.CountDocuments(context.TODO(), filter, &countOpts)
	if err != nil {
		return nil, 0, err
	}

	findOpts := options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	}

	cursor, err := coll.Find(context.TODO(), filter, &findOpts)
	if err != nil {
		return nil, 0, err
	}

	results := make([]*entity.Todo, 0)
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return nil, 0, err
	}

	return results, count, nil
}

func (r *gateway) FindOneTodoById(ctx context.Context, todoID vo.TodoID) (*entity.Todo, error) {
	r.log.Info(ctx, "called %v", todoID)

	coll := r.client.Database(r.config.Servers[r.appData.AppName].Database.Database).Collection("todo")
	filter := bson.D{{}}
	var result entity.Todo

	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

func (r *gateway) SaveTodo(ctx context.Context, obj *entity.Todo) error {
	r.log.Info(ctx, "called %v", obj)

	coll := r.client.Database(r.config.Servers[r.appData.AppName].Database.Database).Collection("todo")

	filter := bson.D{{"_id", obj.ID}}
	update := bson.D{{"$set", obj}}
	opts := options.Update().SetUpsert(true)
	result, err := coll.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		return err
	}

	r.log.Info(ctx, "inserted with id %v", result.ModifiedCount)

	return nil
}
