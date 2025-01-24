package routes

import(
    "net/http"
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
    router := gin.Default()
    //routes for tasks
    router.GET("/tasks", func(c * gin.Context){
        c.JSON(http.StatusOK, gin.H{"message": "list of tasks"})
    })

    router.POST("/tasks", func(c * gin.Context){
        c.JSON(http.StatusOK, gin.H{"message": "task created"})

    })

    router.PUT("/tasks/:id", func(c * gin.Context){
        c.JSON(http.StatusOK, gin.H{"message": "task updated"})
    })

    router.DELETE("/tasks/:id", func( c* gin.Context){
        c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
    })
    
    //routes for users

    router.POST("/register", func(c * gin.Context){
        c.JSON(http.StatusOK, gin.H{"message" : "user created"})
    })
    
    router.POST("/login", func(c * gin.Context){
        c.JSON(http.StatusOK, gin.H{"message": "user logged in"})
    })
    

    return router
}
