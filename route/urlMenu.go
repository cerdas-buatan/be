package module

import (
	"context"

	"github.com/cerdas-buatan/be/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MenuService struct {
	db *mongo.Database
}

func NewMenuService(db *mongo.Database) *MenuService {
	return &MenuService{db: db}
}

func (s *MenuService) CreateMenu(ctx context.Context, menu model.Menu) (model.Menu, error) {
	collection := s.db.Collection("menus")
	res, err := collection.InsertOne(ctx, menu)
	if err != nil {
		return model.Menu{}, err
	}
	menu.ID = res.InsertedID.(primitive.ObjectID)
	return menu, nil
}

func (s *MenuService) GetMenu(ctx context.Context, id primitive.ObjectID) (model.Menu, error) {
	collection := s.db.Collection("menus")
	var menu model.Menu
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&menu)
	if err != nil {
		return model.Menu{}, err
	}
	return menu, nil
}

func (s *MenuService) UpdateMenu(ctx context.Context, id primitive.ObjectID, menu model.Menu) error {
	collection := s.db.Collection("menus")
	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": menu})
	return err
}

func (s *MenuService) DeleteMenu(ctx context.Context, id primitive.ObjectID) error {
	collection := s.db.Collection("menus")
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (s *MenuService) ListMenus(ctx context.Context) ([]model.Menu, error) {
	collection := s.db.Collection("menus")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var menus []model.Menu
	for cur.Next(ctx) {
		var menu model.Menu
		err := cur.Decode(&menu)
		if err != nil {
			return nil, err
		}
		menus = append(menus, menu)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return menus, nil
}
