package main

import (
	"log"
	"strings"

	"github.com/vugu/vugu"
)

type Todo struct {
	Id        string
	Title     string
	Completed bool
}

func (c *ToDoList) BeforeBuild() {
	// get the latest data from the backend...could be expensive to keep calling this
	if len(c.Todos) == 0 {
		log.Printf("Initializing todos...")
		c.Todos = map[string]Todo{
			"1": Todo{Id: "1", Title: "Todo1", Completed: true},
			"2": Todo{Id: "2", Title: "Todo2", Completed: false},
			"3": Todo{Id: "3", Title: "Todo3", Completed: true},
			"4": Todo{Id: "4", Title: "Todo4", Completed: false},
			"5": Todo{Id: "5", Title: "Todo5", Completed: true},
		}
		c.Index = []string{"1", "2", "3", "4", "5"}
	}
}

func (c *ToDoList) getTodoId(s interface{}) (o, id string) {
	slice := strings.Split(s.(string), "-")
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
