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