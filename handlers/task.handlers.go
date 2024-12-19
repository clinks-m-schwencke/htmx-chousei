package handlers

import (
	"chopitto-task/lang"
	"chopitto-task/services"
	"chopitto-task/views/partials"
	"chopitto-task/views/taskviews"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Handler for Task Views

type TaskService interface {
	CreateTask(task services.Task) (services.Task, error)
	GetAllTasks(createdBy int) ([]services.Task, error)
	GetTaskById(task services.Task) (services.Task, error)
	UpdateTask(task services.Task) (services.Task, error)
	UpdateCompleteTask(task services.Task) (services.Task, error)
	UpdateReviewTask(task services.Task) (services.Task, error)
	DeleteTask(task services.Task) error
	GetAllPersons() ([]services.Person, error)
}

type TaskHandler struct {
	TaskServices TaskService
}

func NewTaskHandler(taskService TaskService) *TaskHandler {
	return &TaskHandler{
		TaskServices: taskService,
	}
}

func (taskHandler *TaskHandler) taskHandler(c echo.Context) error {
	println("Task handler called")
	fromProtected, ok := c.Get("FROMPROTECTED").(bool)
	if !ok {
		return errors.New("invalid type for key FROMPROTECTED")
	}

	l := c.Get(LANG_KEY).(string)
	messages := lang.Messages[l]

	c.Set("ISERROR", false)

	userId := c.Get(USER_ID_KEY).(int)

	tasks, err := taskHandler.TaskServices.GetAllTasks(userId)
	if err != nil {
		println("Error!", err)
		return err
	}

	titlePage := fmt.Sprintf(
		"| %s's Tasks",
		cases.Title(language.English).String(c.Get(USERNAME_KEY).(string)),
	)

	persons, err := taskHandler.TaskServices.GetAllPersons()

	user := services.Person{
		Id:   userId,
		Name: c.Get(USERNAME_KEY).(string),
	}

	pageData := taskviews.TaskPageData{
		TitlePage: "",
		User:      user,
		Tasks:     tasks,
		Members:   persons,
		EditId:    -1,
	}

	return renderView(c, taskviews.TaskIndex(
		titlePage,
		c.Get(USERNAME_KEY).(string),
		time.Now().Format("2006-01-02_15-04-05"),
		fromProtected,
		c.Get("ISERROR").(bool),
		getFlashMessages(c, "error"),
		getFlashMessages(c, "success"),
		taskviews.TaskList(pageData, messages.TaskPageStrings),
		messages.BaseLayoutStrings,
	))

}

func (taskHandler *TaskHandler) taskTableHandler(c echo.Context) error {
	println("Task Table handler called")
	_, ok := c.Get("FROMPROTECTED").(bool)
	if !ok {
		return errors.New("invalid type for key FROMPROTECTED")
	}

	l := c.Get(LANG_KEY).(string)
	messages := lang.Messages[l]

	c.Set("ISERROR", false)

	userId := c.Get(USER_ID_KEY).(int)
	idParam := c.Param("id")
	println(idParam)

	editId, err := strconv.Atoi(idParam)
	if err != nil {
		// clear error
		err = nil
		// No edit id
		editId = -1
	}
	println(editId)

	tasks, err := taskHandler.TaskServices.GetAllTasks(userId)

	if err != nil {
		println("Error!", err)
		return err
	}

	persons, err := taskHandler.TaskServices.GetAllPersons()

	user := services.Person{
		Id:   userId,
		Name: c.Get(USERNAME_KEY).(string),
	}

	pageData := taskviews.TaskPageData{
		TitlePage: "",
		User:      user,
		Tasks:     tasks,
		Members:   persons,
		EditId:    editId,
	}

	return renderView(c, taskviews.TaskTableContent(pageData, messages.TaskPageStrings))
}

func (taskHandler *TaskHandler) createTaskHandler(c echo.Context) error {
	_, ok := c.Get("FROMPROTECTED").(bool)
	if !ok {
		return errors.New("invalid type for key FROMPROTECTED")
	}

	// l := c.Get(LANG_KEY).(string)
	// messages := lang.Messages[l]

	c.Set("ISERROR", false)

	if c.Request().Method == "POST" {
		// Get values from <form>
		userId := c.Get(USER_ID_KEY).(int)

		var errorMessages []string
		var successMessages []string

		assignee, err := strconv.Atoi(c.FormValue("assignee"))
		if err != nil {
			errorMessages = append(errorMessages, "Assigned is invalid!!")
		}
		reviewer, err := strconv.Atoi(c.FormValue("reviewer"))
		if err != nil {
			errorMessages = append(errorMessages, "Reviewer is invalid!!")
		}

		title := strings.Trim(c.FormValue("title"), " ")

		if len(title) == 0 {
			errorMessages = append(errorMessages, "Title cannot be empty!!")
		}

		dueOn := strings.Trim(c.FormValue("dueon"), " ")
		if len(dueOn) != 0 {
			_, err := time.Parse("2006-01-02", dueOn)
			if err != nil {
				errorMessages = append(errorMessages, "Due date invalid!!")
			}
		}

		task := services.Task{
			CreatedBy: userId,
			Title:     title,
			Assignee:  assignee,
			Reviewer:  reviewer,
			DueOn:     dueOn,
			Completed: false,
			Reviewed:  false,
		}

		if len(errorMessages) > 0 {
			return renderView(c, partials.FlashMessages(errorMessages, nil))
		}

		// Create task in db
		_, err = taskHandler.TaskServices.CreateTask(task)
		if err != nil {
			return err
		}

		successMessages = append(successMessages, "Task created successfully!!")

		// Set Header to refresh table
		c.Response().Header().Set("HX-Trigger", "updateTask")

		return renderView(c, partials.FlashMessages(nil, successMessages))

	}

	return nil
}

func (taskHandler *TaskHandler) updateTaskHandler(c echo.Context) error {
	println("Update Task handler called")
	_, ok := c.Get("FROMPROTECTED").(bool)
	if !ok {
		return errors.New("invalid type for key FROMPROTECTED")
	}

	// l := c.Get(LANG_KEY).(string)
	// messages := lang.Messages[l]

	c.Set("ISERROR", false)

	if c.Request().Method == "PATCH" {
		// Get values from <form>
		userId := c.Get(USER_ID_KEY).(int)

		assignee, err := strconv.Atoi(c.FormValue("assignee"))
		if err != nil {
			return err
		}
		reviewer, err := strconv.Atoi(c.FormValue("reviewer"))
		if err != nil {
			return err
		}
		task := services.Task{
			CreatedBy: userId,
			Title:     strings.Trim(c.FormValue("title"), " "),
			Assignee:  assignee,
			Reviewer:  reviewer,
			DueOn:     strings.Trim(c.FormValue("dueon"), " "),
			Completed: false,
			Reviewed:  false,
		}

		// Create person in db
		_, err = taskHandler.TaskServices.UpdateTask(task)
		if err != nil {
			return err
		}

		var successMessages []string
		successMessages = append(successMessages, "Task updated successfully!!")

		// Set Header to refresh table
		c.Response().Header().Set("HX-Trigger", "updateTask")

		return renderView(c, partials.FlashMessages(nil, successMessages))
	}

	return nil

}

func (taskHandler *TaskHandler) completeTaskHandler(c echo.Context) error {
	println("Complete Task handler called")
	c.Set("ISERROR", false)
	_, ok := c.Get("FROMPROTECTED").(bool)
	if !ok {
		return errors.New("invalid type for key 'FROMPROTECTED'")
	}

	l := c.Get(LANG_KEY).(string)
	messages := lang.Messages[l]

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return err
	}

	if c.Request().Method == "PATCH" {
		// Get values from <form>
		userId := c.Get(USER_ID_KEY).(int)

		completed := c.FormValue("complete") == "on"

		task := services.Task{
			Id:        id,
			Assignee:  userId,
			Completed: completed,
		}

		// Create person in db
		task, err = taskHandler.TaskServices.UpdateCompleteTask(task)
		if err != nil {
			return err
		}

		var successMessages []string
		successMessages = append(successMessages, "Task updated successfully!!")

		return renderView(c, partials.OobFlashMessages(taskviews.TaskRow(task, userId, -1, messages.TaskPageStrings), nil, successMessages))
	}

	return nil

}

func (taskHandler *TaskHandler) reviewTaskHandler(c echo.Context) error {
	println("Review Task handler called")
	c.Set("ISERROR", false)
	_, ok := c.Get("FROMPROTECTED").(bool)
	if !ok {
		return errors.New("invalid type for key 'FROMPROTECTED'")
	}

	l := c.Get(LANG_KEY).(string)
	messages := lang.Messages[l]

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return err
	}

	if c.Request().Method == "PATCH" {
		// Get values from <form>
		userId := c.Get(USER_ID_KEY).(int)

		reviewed := c.FormValue("reviewed") == "on"
		task := services.Task{
			Id:       id,
			Assignee: userId,
			Reviewed: reviewed,
		}

		// Update task in db
		task, err = taskHandler.TaskServices.UpdateReviewTask(task)
		if err != nil {
			return err
		}

		var successMessages []string
		successMessages = append(successMessages, "Task updated successfully!!")

		return renderView(c, partials.OobFlashMessages(taskviews.TaskRow(task, userId, -1, messages.TaskPageStrings), nil, successMessages))
	}

	return nil

}

func (taskHandler *TaskHandler) deleteTaskHandler(c echo.Context) error {
	println("Delete Task handler called")
	_, ok := c.Get("FROMPROTECTED").(bool)
	if !ok {
		return errors.New("invalid type for key FROMPROTECTED")
	}

	// l := c.Get(LANG_KEY).(string)
	// messages := lang.Messages[l]

	c.Set("ISERROR", false)

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return err
	}

	if c.Request().Method == "DELETE" {
		// Get values from <form>
		userId := c.Get(USER_ID_KEY).(int)

		task := services.Task{
			Id:        id,
			CreatedBy: userId,
		}

		// Delete task in DB
		err := taskHandler.TaskServices.DeleteTask(task)
		if err != nil {
			return err
		}

		var successMessages []string
		successMessages = append(successMessages, "Task deleted successfully!!")

		// Set Header to refresh table
		c.Response().Header().Set("HX-Trigger", "updateTask")

		return renderView(c, partials.FlashMessages(nil, successMessages))
	}

	return nil

}
