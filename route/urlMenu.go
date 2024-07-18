// route/url.go

package route

import (
    "github.com/cerdas-buatan/be/module"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(db *mongo.Database) *gin.Engine {
    r := gin.Default()

    menuService := module.NewMenuService(db)

    r.POST("/menus", func(c *gin.Context) {
        var menu module.Menu
        if err := c.ShouldBindJSON(&menu); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        createdMenu, err := menuService.CreateMenu(c.Request.Context(), menu)
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        c.JSON(201, createdMenu)
    })

    r.GET("/menus/:id", func(c *gin.Context) {
        id, err := primitive.ObjectIDFromHex(c.Param("id"))
        if err != nil {
            c.JSON(400, gin.H{"error": "Invalid ID"})
            return
        }
        menu, err := menuService.GetMenu(c.Request.Context(), id)
        if err != nil {
            c.JSON(404, gin.H{"error": "Menu not found"})
            return
        }
        c.JSON(200, menu)
    })

    r.PUT("/menus/:id", func(c *gin.Context) {
        id, err := primitive.ObjectIDFromHex(c.Param("id"))
        if err != nil {
            c.JSON(400, gin.H{"error": "Invalid ID"})
            return
        }
        var menu module.Menu
        if err := c.ShouldBindJSON(&menu); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        err = menuService.UpdateMenu(c.Request.Context(), id, menu)
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, gin.H{"message": "Menu updated successfully"})
    })

    r.DELETE("/menus/:id", func(c *gin.Context) {
        id, err := primitive.ObjectIDFromHex(c.Param("id"))
        if err != nil {
            c.JSON(400, gin.H{"error": "Invalid ID"})
            return
        }
        err = menuService.DeleteMenu(c.Request.Context(), id)
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, gin.H{"message": "Menu deleted successfully"})
    })

    r.GET("/menus", func(c *gin.Context) {
        menus, err := menuService.ListMenus(c.Request.Context())
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, menus)
    })

    return r
}
