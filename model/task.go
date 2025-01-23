package model

type task struct{
    ID int `json:"id"`
    Title string `json:"title"`
    Completed bool `json:"completed"`
    UserId int `json:"user_id"`


}
