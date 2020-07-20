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

func createClient() *client.SimpleTodoAPISecure {
	url, _ := url.Parse(AuthenticationData.RestEndpoint)
	authenticator := BearerToken(AuthenticationData.LoginData.ResponseParams.AccessToken)

	conf := client.Config{
		URL:      url,
		AuthInfo: authenticator,
	}
	return client.New(conf)
}

func (c *ToDoList) updateItem(t *models.Todo) {
	backend := createClient()

	params := developers.NewUpdateTodoParams()
	params.Todo = t
	params.Todoid = t.ID.String()
	ctx := context.TODO()

	if _, err := backend.Developers.UpdateTodo(ctx, params); err != nil {
		log.Printf("Error updating item on backend - error %v\n", err)
		return
	}
}

func (c *ToDoList) postItemToBackend(t models.Todo) {
	backend := createClient()

	params := developers.NewAddTodoParams()
	params.Todo = &t
	ctx := context.TODO()

	if _, err := backend.Developers.AddTodo(ctx, params); err != nil {
		log.Printf("Error pusting new item on backend - error %v\n", err)
		return
	}
}

func (c *ToDoList) destroyItemOnBackend(t *models.Todo) {
	backend := createClient()

	params := developers.NewDeleteTodoParams()
	params.Todoid = t.ID.String()

	ctx := context.TODO()

	if _, err := backend.Developers.DeleteTodo(ctx, params); err != nil {
		log.Printf("Error deleting item on backend - error %v\n", err)
		return
	}
}

func (c *ToDoList) getTodosFromBackend() ([]*models.Todo, error) {

	backend := createClient()

	p := developers.NewGetAllTodosParams()
	ctx := context.TODO()
	todos, err := backend.Developers.GetAllTodos(ctx, p)

	if err != nil {
		log.Printf("Error obtaining items from backend - error %v\n", err)
		return nil, err
	}

	return todos.Payload, nil
}

func (c *ToDoList) BeforeBuild() {
	// get the latest data from the backend...could be expensive to keep calling this
	if AuthenticationData.LoginData.LoggedIn {
		// this does not handle the case well in which the length of the todolist
		// on the server is 0
		if len(c.Todos) == 0 {
			log.Printf("Retrieving todos from backend...")
			todos, err := c.getTodosFromBackend()

			if err == nil {
				log.Printf("Initializing todo data...")
				c.Todos = make(map[string]models.Todo)
				c.Index = []string{}
				for _, v := range todos {
					idString := v.ID.String()
					c.Todos[idString] = *v
					c.Index = append(c.Index, idString)
				}
			}
		}
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
	go c.updateItem(&t)
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
	t := c.Todos[id]
	delete(c.Todos, id)
	go c.destroyItemOnBackend(&t)
}

func (c *ToDoList) AddTodo(t models.Todo) {
	c.Todos[t.ID.String()] = t
	c.Index = append(c.Index, t.ID.String())
	go c.postItemToBackend(t)
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
