package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"chopitto-task/cmd/lib/data"
	"chopitto-task/cmd/lib/types"
	"chopitto-task/views"
	"github.com/labstack/echo/v4"
)

func HandleTaskGet(c echo.Context) error {
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
		task.Done = false
		task.Reviewed = false

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

func HandleTaskPost(c echo.Context) error {
	title := strings.TrimSpace(c.FormValue("title"))
	// := strings.TrimSpace(c.FormValue("author"))
	// TODO: Get from login?
	author := 1
	assigned := strings.TrimSpace(c.FormValue("assigned"))
	dueOn := strings.TrimSpace(c.FormValue("due-on"))

	if title == "" || assigned == "" {
		formData := types.NewFormData()
		formData.Values["title"] = title
		formData.Values["assigned"] = assigned
		formData.Values["dueOn"] = dueOn
		// TODO: Change MeetingForm to sign up
		return Render(c, http.StatusUnprocessableEntity, views.MeetingForm(formData))
	}
	sqlCmd := `
	INSERT INTO task (title, author, assigned, due_on)
	VALUES(?, ?, ?, ?)
	`
	result, err := lib.Db.Exec(sqlCmd, title, author, assigned, dueOn)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	row := lib.Db.QueryRow(`
		SELECT t.task_id, t.title, t.author, t.assigned, t.due_on, author.name AS author_name, assigned.name as assigned_name
		FROM task t 
		LEFT JOIN person author ON author.person_id = t.author
		LEFT JOIN person assigned ON assigned.person_id = t.assigned
		WHERE t.task_id = ?
		LIMIT 1
	`, id)
	var task types.Task
	err = row.Scan(
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

	return Render(c, http.StatusOK, views.TaskPostResponse(task))
}
