package module

import (
    "context"
    "time"

    "github.com/cerdas-buatan/be/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type MenuService struct {
    collection *mongo.Collection
}


func NewMenuService(db *mongo.Database) *MenuService {
    return &MenuService{
        collection: db.Collection("menus"),
    }
}

func (s *MenuService) CreateMenu(ctx context.Context, menu model.Menu) (model.Menu, error) {
    menu.ID = primitive.NewObjectID()
    _, err := s.collection.InsertOne(ctx, menu)
    return menu, err
}

func (s *MenuService) GetMenu(ctx context.Context, id primitive.ObjectID) (model.Menu, error) {
    var menu model.Menu
    err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&menu)
    return menu, err
}


func (s *MenuService) UpdateMenu(ctx context.Context, id primitive.ObjectID, menu model.Menu) error {
    _, err := s.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": menu})
    return err
}

func (s *MenuService) UpdateMenu(ctx context.Context, id primitive.ObjectID, menu model.Menu) error {
    _, err := s.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": menu})
    return err
}
