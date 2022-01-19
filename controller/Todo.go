package controller

import (
	"log"
	"net/http"
	"strconv"
	"todo/db"
	"todo/model"

	"github.com/gin-gonic/gin"
)

type Todo struct{}

func NewTodo() *Todo {
	return &Todo{}
}

func (t *Todo) List(c *gin.Context) {
	db := db.Connection()
	defer db.Close()
	name, _ := c.Get("name") // ログインユーザの取得
	var Todo []model.Todo
	result := db.Where("User =?", name).Find(&Todo)
	//log.Println(result.Value)
	c.HTML(http.StatusOK, "menu", gin.H{
		"name": name,
		"Todo": result.Value,
	})
}

func (t *Todo) CreateTodo(c *gin.Context) {
	db := db.Connection()
	defer db.Close()
	name, _ := c.Get("name")
	todo := c.PostForm("Todo")
	log.Println(name, todo)
	var Todo model.Todo
	//c.BindJSON(&todo)
	user, ok := name.(string)
	if ok == true {
		Todo.User = user
	}
	Todo.Todo = todo
	db.Create(&Todo)
	c.Redirect(302, "/menu/top")
}

func (t *Todo) Get(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	db := db.Connection()
	var todo model.Todo
	result := db.First(&todo, id)
	log.Println(result.Value)
	defer db.Close()
	c.HTML(200, "detail.html", gin.H{
		"todo": result.Value,
	})
}

func (t *Todo) Update(c *gin.Context) {
	db := db.Connection()
	defer db.Close()

	var todo model.Todo
	Todo := c.PostForm("Todo")
	db.First(&todo, c.Param("id"))
	todo.Todo = Todo
	db.Save(&todo)
	c.Redirect(302, "/menu/top")
}

func (t *Todo) Delete(c *gin.Context) {
	db := db.Connection()
	defer db.Close()

	var todo model.Todo
	db.First(&todo, c.Param("id"))
	log.Println(&todo)

	if todo.ID > uint(0) {
		db.Delete(&todo)
	}
	c.Redirect(302, "/menu/top")
}
