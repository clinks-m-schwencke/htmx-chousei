package services

import (
	"errors"
	"fmt"
	"time"

	"chopitto-task/db"
)

// Task service constructor
func NewTaskServices(task Task, taskStore db.Store) *TaskServices {
	return &TaskServices{
		Task:      task,
		TaskStore: taskStore,
	}
}

type Task struct {
	Id            int
	CreatedBy     int
	Title         string
	Assignee      int
	Reviewer      int
	Completed     bool
	Reviewed      bool
	DueOn         string
	CreatedAt     string
	CreatedByName string
	AssigneeName  string
	ReviewerName  string
}

type TaskServices struct {
	Task      Task
	TaskStore db.Store
}

func (ts *TaskServices) GetAllTasks(userId int) ([]Task, error) {
	query := fmt.Sprintf(`
		SELECT t.id, t.created_by, t.title, t.assignee, t.reviewer, t.completed, t.reviewed, t.due_on, 
			t.created_at, c.name as created_by_name, a.name as assignee_name, r.name as reviewer_name
		FROM task t
		LEFT JOIN person c ON c.id = t.created_by
		LEFT JOIN person a ON a.id = t.assignee
		LEFT JOIN person r ON r.id = t.reviewer
		WHERE t.created_by = %d OR t.assignee = %d OR t.reviewer = %d
		ORDER BY created_at DESC;
		`, userId, userId, userId)

	rows, err := ts.TaskStore.Db.Query(query)
	if err != nil {
		println("Get all tasks Query error")
		return []Task{}, err
	}
	// We close the resource
	defer rows.Close()

	todos := []Task{}
	for rows.Next() {
		err = rows.Scan(
			&ts.Task.Id,
			&ts.Task.CreatedBy,
			&ts.Task.Title,
			&ts.Task.Assignee,
			&ts.Task.Reviewer,
			&ts.Task.Completed,
			&ts.Task.Reviewed,
			&ts.Task.DueOn,
			&ts.Task.CreatedAt,
			&ts.Task.CreatedByName,
			&ts.Task.AssigneeName,
			&ts.Task.ReviewerName,
		)

		if err != nil {
			println("scan error")
			return []Task{}, err
		}

		todos = append(todos, ts.Task)
	}

	return todos, nil
}

// Create a new task
func (taskService *TaskServices) CreateTask(task Task) (Task, error) {
	query := `
		INSERT INTO task (created_by, title, assignee, reviewer, due_on)
		VALUES(?, ?, ?, ?, ?) RETURNING *
	`

	stmt, err := taskService.TaskStore.Db.Prepare(query)
	if err != nil {
		return Task{}, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(task.CreatedBy, task.Title, task.Assignee, task.Reviewer, task.DueOn).Scan(
		&taskService.Task.Id,
		&taskService.Task.CreatedBy,
		&taskService.Task.Title,
		&taskService.Task.Assignee,
		&taskService.Task.Reviewer,
		&taskService.Task.Completed,
		&taskService.Task.Reviewed,
		&taskService.Task.DueOn,
		&taskService.Task.CreatedAt,
	)

	if err != nil {
		return Task{}, err
	}

	return taskService.Task, nil
}

// Get a task by ID
// Need to be either an assignee, reviewer or creator to see task
func (taskService *TaskServices) GetTaskById(task Task) (Task, error) {
	query := `
		SELECT t.id, t.created_by, t.title, t.assignee, t.reviewer, t.completed, t.reviewed, t.due_on, 
			t.created_at, c.name as created_by_name, a.name as assignee_name, r.name as reviewer_name
		FROM task t
		LEFT JOIN person c ON c.id = t.created_by
		LEFT JOIN person a ON a.id = t.assignee
		LEFT JOIN person r ON r.id = t.reviewer
		WHERE t.id = ? AND (t.assignee = ? OR t.reviewer = ? OR t.created_by = ?);
	`

	stmt, err := taskService.TaskStore.Db.Prepare(query)
	if err != nil {
		return Task{}, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(
		task.Id,
		task.Assignee,
		task.Reviewer,
		task.CreatedBy,
	).Scan(
		&taskService.Task.Id,
		&taskService.Task.CreatedBy,
		&taskService.Task.Title,
		&taskService.Task.Assignee,
		&taskService.Task.Reviewer,
		&taskService.Task.Completed,
		&taskService.Task.Reviewed,
		&taskService.Task.DueOn,
		&taskService.Task.CreatedAt,
		&taskService.Task.CreatedByName,
		&taskService.Task.AssigneeName,
		&taskService.Task.ReviewerName,
	)

	if err != nil {
		return Task{}, err
	}

	return taskService.Task, nil
}

// Updated an existing task
func (taskService *TaskServices) UpdateTask(task Task) (Task, error) {
	query := `
		UPDATE task SET title = ?, assignee = ?, reviewer = ?, completed = ?, reviewed = ?, due_on = ? FROM task
		WHERE id = ? AND created_by = ?
		RETURNING id, title, assignee, reviewer, completed, reviewed, due_on;
	`

	stmt, err := taskService.TaskStore.Db.Prepare(query)
	if err != nil {
		return Task{}, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(
		task.Title,
		task.Assignee,
		task.Reviewer,
		task.Completed,
		task.Reviewed,
		task.DueOn,
		task.Id,
	).Scan(
		&taskService.Task.Id,
		&taskService.Task.Title,
		&taskService.Task.Assignee,
		&taskService.Task.Reviewer,
		&taskService.Task.Completed,
		&taskService.Task.Reviewed,
		&taskService.Task.DueOn,
	)

	if err != nil {
		return Task{}, err
	}

	return taskService.Task, nil
}

// Updated an existing task
func (taskService *TaskServices) UpdateCompleteTask(task Task) (Task, error) {
	query := `
		UPDATE task SET completed = ?
		WHERE id = ? AND assignee = ?;
	`
	// RETURNING id, title, assignee, reviewer, completed, reviewed, due_on;

	stmt, err := taskService.TaskStore.Db.Prepare(query)
	if err != nil {
		return Task{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		task.Completed,
		task.Id,
		task.Assignee,
	)
	// .Scan(
	// 	&taskService.Task.Id,
	// 	&taskService.Task.Title,
	// 	&taskService.Task.Assignee,
	// 	&taskService.Task.Reviewer,
	// 	&taskService.Task.Completed,
	// 	&taskService.Task.Reviewed,
	// 	&taskService.Task.DueOn,
	// )

	if err != nil {
		return Task{}, err
	}

	taskService.Task, err = taskService.GetTaskById(task)

	if err != nil {
		return Task{}, err
	}

	return taskService.Task, nil
}

// Updated an existing task
func (taskService *TaskServices) UpdateReviewTask(task Task) (Task, error) {
	query := `
		UPDATE task SET reviewed = ?
		WHERE id = ? AND reviewer = ?
		RETURNING id, title, assignee, reviewer, completed, reviewed, due_on;
	`

	stmt, err := taskService.TaskStore.Db.Prepare(query)
	if err != nil {
		return Task{}, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(
		task.Reviewed,
		task.Id,
		task.Assignee,
	).Scan(
		&taskService.Task.Id,
		&taskService.Task.Title,
		&taskService.Task.Assignee,
		&taskService.Task.Reviewer,
		&taskService.Task.Completed,
		&taskService.Task.Reviewed,
		&taskService.Task.DueOn,
	)

	if err != nil {
		return Task{}, err
	}

	return taskService.Task, nil
}

// Delete a task
func (taskService *TaskServices) DeleteTask(task Task) error {
	query := `
		DELETE FROM task
		WHERE id = ? AND created_by = ?;
	`

	stmt, err := taskService.TaskStore.Db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(task.Id, task.CreatedBy)

	if err != nil {
		return err
	}
	i, err := result.RowsAffected()

	if err != nil || i != 1 {
		return errors.New("an affected row was expected")
	}

	return nil
}

func (taskService *TaskServices) GetAllPersons() ([]Person, error) {
	// Query for getting a person from an email
	query := `SELECT id, email, name FROM person`

	// Prepare the query
	rows, err := taskService.TaskStore.Db.Query(query)
	if err != nil {
		// return empty person and error
		println("Get all persons Query error")
		return []Person{}, err
	}

	// Defer the close of the query (Will close after function is complete)
	defer rows.Close()

	persons := []Person{}

	for rows.Next() {
		tmpPerson := Person{}
		err = rows.Scan(
			// Put data into the person service person
			&tmpPerson.Id,
			&tmpPerson.Email,
			&tmpPerson.Name,
		)

		if err != nil {
			// return empty person and error
			println("Get All persons scan error")
			return []Person{}, err
		}

		persons = append(persons, tmpPerson)

	}
	// Return person list
	return persons, nil

}

func ConvertDateTime(tz string, dt time.Time) string {
	loc, _ := time.LoadLocation(tz)

	return dt.In(loc).Format(time.RFC822Z)
}
