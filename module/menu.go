package module

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/cerdas-buatan/be/model"
)

type MenuService struct {
    collection *mongo.Collection
}

func NewMenuService(db *mongo.Database) *MenuService {
    return &MenuService{
        collection: db.Collection("menus"),
    }
}

// CreateMenu creates a new menu item and stores it in the database
func (s *MenuService) CreateMenu(ctx context.Context, menu model.Menu) (model.Menu, error) {
    menu.ID = primitive.NewObjectID()
    _, err := s.collection.InsertOne(ctx, menu)
    return menu, err
}

// GetMenu retrieves a menu item by its ID
func (s *MenuService) GetMenu(ctx context.Context, id primitive.ObjectID) (model.Menu, error) {
    var menu model.Menu
    err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&menu)
    return menu, err
}

// UpdateMenu updates the details of an existing menu item
func (s *MenuService) UpdateMenu(ctx context.Context, id primitive.ObjectID, updateFields map[string]interface{}) error {
    update := bson.M{"$set": updateFields}
    _, err := s.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
    return err
}

// DeleteMenu removes a menu item from the database
func (s *MenuService) DeleteMenu(ctx context.Context, id primitive.ObjectID) error {
    _, err := s.collection.DeleteOne(ctx, bson.M{"_id": id})
    return err
}

// ListMenus retrieves all menu items from the database
func (s *MenuService) ListMenus(ctx context.Context) ([]model.Menu, error) {
    var menus []model.Menu
    cursor, err := s.collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)
    for cursor.Next(ctx) {
        var menu model.Menu
        if err := cursor.Decode(&menu); err != nil {
            return nil, err
        }
        menus = append(menus, menu)
    }
    return menus, nil
}

// AddChatHistory creates a new chat history entry
func (s *MenuService) AddChatHistory(ctx context.Context, chat model.ChatHistory) (model.ChatHistory, error) {
    chat.ID = primitive.NewObjectID()
    _, err := s.collection.InsertOne(ctx, chat)
    return chat, err
}

// DeleteChatHistory removes a chat history entry by its ID
func (s *MenuService) DeleteChatHistory(ctx context.Context, id primitive.ObjectID) error {
    _, err := s.collection.DeleteOne(ctx, bson.M{"_id": id})
    return err
}

// UpdateChatHistory updates a specific chat history entry
func (s *MenuService) UpdateChatHistory(ctx context.Context, id primitive.ObjectID, updateFields map[string]interface{}) error {
    update := bson.M{"$set": updateFields}
    _, err := s.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
    return err
}
