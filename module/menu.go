package module

import (
    "context"
    "github.com/cerdas-buatan/be/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

// Newmodel.MenuService returns a new model.MenuService
func Menu(collection *mongo.Collection) *model.MenuService {
    return &model.MenuService{collection: collection}
}

// RenameMenu renames an existing menu item
func (s *model.MenuService) RenameMenu(ctx context.Context, id primitive.ObjectID, newName string) error {
    update := bson.M{"$set": bson.M{"name": newName}}
    _, err := s.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
    return err
}

// ArchiveMenu moves a menu item to the archive collection
func (s *model.MenuService) ArchiveMenu(ctx context.Context, id primitive.ObjectID) error {
    var menu model.Menu
    err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&menu)
    if err != nil {
        return err
    }
    _, err = s.collection.DeleteOne(ctx, bson.M{"_id": id})
    if err != nil {
        return err
    }
    archiveCollection := s.collection.Database().Collection("archive_menus")
    _, err = archiveCollection.InsertOne(ctx, menu)
    return err
}

// AddMenu adds a new menu item
func (s *model.MenuService) AddMenu(ctx context.Context, menu model.Menu) (model.Menu, error) {
    menu.ID = primitive.NewObjectID()
    _, err := s.collection.InsertOne(ctx, menu)
    return menu, err
}
