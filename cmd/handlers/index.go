package handlers

import (
	"fmt"
	"log"
	"net/http"

	"chopitto-task/cmd/lib/data"
	"chopitto-task/cmd/lib/types"
	"chopitto-task/views"
	"github.com/labstack/echo/v4"
)

func HandleIndexGet(c echo.Context) error {
	// Get data from database
	query := `
		SELECT t.task_id, t.title, t.author, t.assigned, t.due_on, author.name AS author_name, assigned.name as assigned_name
		FROM task t
		LEFT JOIN person author ON author.person_id = t.author
		LEFT JOIN person assigned ON assigned.person_id = t.assigned;
	`
	rows, err := lib.Db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	var tasks []types.Task

	for rows.Next() {
		var task types.Task
		err := rows.Scan(
			&task.TaskId,
			&task.Title,
			&task.Author,
			&task.Assigned,
			&task.DueOn,
			&task.AuthorName,
			&task.AssignedName,
		)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Task ID, Title, Author, Assigned, DueOn, AuthorName, AssignedName")
	for _, task := range tasks {
		fmt.Printf("%d %s %d %d %s %s %s\n", task.TaskId, task.Title, task.Author, task.Assigned, task.DueOn, task.AuthorName, task.AssignedName)
	}

	// hand it to template
	return Render(c, http.StatusOK, views.Index(tasks))
}
