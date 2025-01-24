package main

import(
    "task-manager/routes"
    "task-manager/database"
)


func main(){
    database.Connect()
    router := routes.SetupRouter()
    
    router.Run(":8080")


}
