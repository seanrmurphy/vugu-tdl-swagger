package main

import (
	"context"
	"log"
	"net/url"
	"strings"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/vugu/vugu"

	"github.com/seanrmurphy/go-vecty-swagger/client"
	"github.com/seanrmurphy/go-vecty-swagger/client/developers"
	"github.com/seanrmurphy/go-vecty-swagger/models"
)

type Todo struct {
	Id        string
	Title     string
	Completed bool
}

var restEndpoint = "https://glm3dpf2yi.execute-api.eu-west-1.amazonaws.com/prod"

func createClient() *client.SimpleTodo {
	url, _ := url.Parse(restEndpoint)
	conf := client.Config{
		URL: url,
	}
	c := client.New(conf)
	return c
}

//func (c *ToDoList) updateItem(i *model.Item) {
//c := createClient()

//t := models.Todo{
//Completed:    i.BackEndModel.Completed,
//ID:           i.BackEndModel.ID,
//Title:        i.BackEndModel.Title,
//CreationDate: strfmt.DateTime(time.Now()),
//}
//params := developers.NewUpdateTodoParams()
//params.Todo = &t
//params.Todoid = i.BackEndModel.ID.String()
//ctx := context.TODO()

//if _, err := c.Developers.UpdateTodo(ctx, params); err != nil {
//log.Printf("Error updating item on backend - error %v\n", err)
//return
//}
//}

//func (c *ToDoList) postItemToBackend(i model.Item) {
//c := createClient()

//t := models.Todo{
//Completed:    i.BackEndModel.Completed,
//ID:           i.BackEndModel.ID,
//Title:        i.BackEndModel.Title,
//CreationDate: strfmt.DateTime(time.Now()),
//}
//params := developers.NewAddTodoParams()
//params.Todo = &t
//ctx := context.TODO()

//if _, err := c.Developers.AddTodo(ctx, params); err != nil {
//log.Printf("Error pusting new item on backend - error %v\n", err)
//return
//}
//}

//func (c *ToDoList) destroyItemOnBackend(i *model.Item) {
//c := createClient()

//params := developers.NewDeleteTodoParams()
//params.Todoid = i.BackEndModel.ID.String()

//ctx := context.TODO()

//if _, err := c.Developers.DeleteTodo(ctx, params); err != nil {
//log.Printf("Error deleting item on backend - error %v\n", err)
//return
//}
//}

func (c *ToDoList) getTodosFromBackend() []*models.Todo {
	url, _ := url.Parse(restEndpoint)
	conf := client.Config{
		URL: url,
	}
	backend := client.New(conf)

	p := developers.NewGetAllTodosParams()
	ctx := context.TODO()
	todos, err := backend.Developers.GetAllTodos(ctx, p)

	if err != nil {
		log.Printf("Error obtaining items from backend - error %v\n", err)
	}

	return todos.Payload
}

func (c *ToDoList) BeforeBuild() {
	// get the latest data from the backend...could be expensive to keep calling this
	if len(c.Todos) == 0 {
		log.Printf("Initializing todos...")
		todos := c.getTodosFromBackend()

		c.Todos = make(map[string]models.Todo)
		c.Index = []string{}
		for _, v := range todos {
			idString := v.ID.String()
			c.Todos[idString] = *v
			c.Index = append(c.Index, idString)
		}
		//c.Todos = map[string]Todo{
		//"1": Todo{Id: "1", Title: "Todo1", Completed: true},
		//"2": Todo{Id: "2", Title: "Todo2", Completed: false},
		//"3": Todo{Id: "3", Title: "Todo3", Completed: true},
		//"4": Todo{Id: "4", Title: "Todo4", Completed: false},
		//"5": Todo{Id: "5", Title: "Todo5", Completed: true},
		//}
		//c.Index = []string{"1", "2", "3", "4", "5"}
	}
}

func (c *ToDoList) getTodoId(s interface{}) (o, id string) {
	slice := strings.SplitN(s.(string), "-", 2)
	o = slice[0]
	id = slice[1]
	return
}

func (c *ToDoList) Done(e vugu.DOMEvent) {
	_, id := c.getTodoId(e.Prop("target", "id"))
	t := c.Todos[id]
	t.Completed = !t.Completed
	c.Todos[id] = t
}

func (c *ToDoList) Delete(e vugu.DOMEvent) {
	_, id := c.getTodoId(e.Prop("target", "id"))
	log.Printf("Delete - id = %v", id)
	// remove from index
	found := false
	for i, v := range c.Index {
		if v == id {
			switch i {
			case 0:
				found = true
				c.Index = c.Index[i+1:]
				break
			case len(c.Index) - 1:
				found = true
				c.Index = c.Index[:i]
				break
			default:
				found = true
				c.Index = append(c.Index[:i], c.Index[i+1:]...)
				break
			}
		}
	}
	if found == false {
		log.Printf("Unable to remove item from slice")
	}

	// remove from map
	delete(c.Todos, id)
}

func (c *ToDoList) AddTodo(t models.Todo) {
	c.Todos[t.ID.String()] = t
	c.Index = append(c.Index, t.ID.String())
}

func (c *ToDoList) Keypress(e vugu.DOMEvent) {
	keyCode := e.PropFloat64("keyCode")
	// when enter is pressed...
	if keyCode == 13 {
		todoString := e.PropString("target", "value")
		t := models.Todo{ID: strfmt.UUID(uuid.New().String()), Title: &todoString, Completed: false}
		c.AddTodo(t)
	}
}
