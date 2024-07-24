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