package api

import (
	"a21hc3NpZ25tZW50/entity"
	"a21hc3NpZ25tZW50/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type TaskAPI interface {
	GetTask(w http.ResponseWriter, r *http.Request)
	CreateNewTask(w http.ResponseWriter, r *http.Request)
	UpdateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
	UpdateTaskCategory(w http.ResponseWriter, r *http.Request)
}

type taskAPI struct {
	taskService service.TaskService
}

func NewTaskAPI(taskService service.TaskService) *taskAPI {
	return &taskAPI{taskService}
}

func (t *taskAPI) GetTask(w http.ResponseWriter, r *http.Request) {
	id := fmt.Sprintf("%s", r.Context().Value("id"))
	if id == "" {
		w.WriteHeader(400)
		status := entity.ErrorResponse{
			Error: "invalid user id",
		}
		jsonStatus, _ := json.Marshal(status)
		w.Write(jsonStatus)
		return
	}

	taskID := r.URL.Query().Get("task_id")
	idIntUser, _ := strconv.Atoi(id)
	idIntTask, _ := strconv.Atoi(taskID)

	if taskID == "" {
		tasks, err := t.taskService.GetTasks(r.Context(), idIntUser)
		if err != nil {
			w.WriteHeader(500)
			status := entity.ErrorResponse{
				Error: "error internal server",
			}
			jsonStatus, _ := json.Marshal(status)
			w.Write(jsonStatus)
			return
		}

		w.WriteHeader(200)
		jsonStatus, _ := json.Marshal(tasks)
		w.Write(jsonStatus)
		return
	}

	getTask, err := t.taskService.GetTaskByID(r.Context(), idIntTask)
	if err != nil {
		w.WriteHeader(500)
		status := entity.ErrorResponse{
			Error: "error internal server",
		}
		jsonStatus, _ := json.Marshal(status)
		w.Write(jsonStatus)
		return
	}

	w.WriteHeader(200)
	jsonStatus, _ := json.Marshal(getTask)
	w.Write(jsonStatus)
	return
	// TODO: answer here
}

func (t *taskAPI) CreateNewTask(w http.ResponseWriter, r *http.Request) {
	var task entity.TaskRequest
	var newTask entity.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid task request"))
		return
	}

	if task.Title == "" || task.Description == "" || task.CategoryID == 0 {
		w.WriteHeader(400)
		status := entity.ErrorResponse{
			Error: "invalid task request",
		}
		jsonStatus, _ := json.Marshal(status)
		w.Write(jsonStatus)
		return
	}

	id := fmt.Sprintf("%s", r.Context().Value("id"))
	if id == "" {
		w.WriteHeader(400)
		status := entity.ErrorResponse{
			Error: "invalid user id",
		}
		jsonStatus, _ := json.Marshal(status)
		w.Write(jsonStatus)
		return
	}

	idInt, _ := strconv.Atoi(id)

	newTask = entity.Task{
		Title:       task.Title,
		Description: task.Description,
		CategoryID:  task.CategoryID,
		UserID:      idInt,
	}

	create, err := t.taskService.StoreTask(r.Context(), &newTask)
	if err != nil {
		w.WriteHeader(500)
		status := entity.ErrorResponse{
			Error: "error internal server",
		}
		jsonStatus, _ := json.Marshal(status)
		w.Write(jsonStatus)
		return
	}

	w.WriteHeader(201)
	status := map[string]interface{}{
		"user_id": create.UserID,
		"task_id": create.ID,
		"message": "success create new task",
	}
	jsonStatus, _ := json.Marshal(status)
	w.Write(jsonStatus)
	return

	// TODO: answer here
}

func (t *taskAPI) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := fmt.Sprintf("%s", r.Context().Value("id"))
	if id == "" {
		w.WriteHeader(400)
		status := entity.ErrorResponse{
			Error: "invalid user id",
		}
		jsonStatus, _ := json.Marshal(status)
		w.Write(jsonStatus)
		return
	}

	taskID := r.URL.Query().Get("task_id")
	idIntTask, _ := strconv.Atoi(taskID)
	idIntUser, _ := strconv.Atoi(id)

	err := t.taskService.DeleteTask(r.Context(), idIntTask)
	if err != nil {
		w.WriteHeader(500)
		status := entity.ErrorResponse{
			Error: "error internal server",
		}
		jsonStatus, _ := json.Marshal(status)
		w.Write(jsonStatus)
		return
	}

	w.WriteHeader(200)
	status := map[string]interface{}{
		"user_id": idIntUser,
		"task_id": idIntTask,
		"message": "success delete task",
	}
	jsonStatus, _ := json.Marshal(status)
	w.Write(jsonStatus)
	return
	// TODO: answer here
}

func (t *taskAPI) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task entity.TaskRequest
	var updtTask entity.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid decode json"))
		return
	}

	id := fmt.Sprintf("%s", r.Context().Value("id"))
	if id == "" {
		w.WriteHeader(400)
		status := entity.ErrorResponse{
			Error: "invalid user id",
		}
		jsonStatus, _ := json.Marshal(status)
		w.Write(jsonStatus)
		return
	}

	idInt, _ := strconv.Atoi(id)

	updtTask = entity.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		CategoryID:  task.CategoryID,
		UserID:      idInt,
	}

	update, err := t.taskService.UpdateTask(r.Context(), &updtTask)
	if err != nil {
		w.WriteHeader(500)
		status := entity.ErrorResponse{
			Error: "error internal server",
		}
		jsonStatus, _ := json.Marshal(status)
		w.Write(jsonStatus)
		return
	}

	w.WriteHeader(200)
	status := map[string]interface{}{
		"user_id": update.UserID,
		"task_id": update.ID,
		"message": "success update task",
	}
	jsonStatus, _ := json.Marshal(status)
	w.Write(jsonStatus)
	return

	// TODO: answer here
}

func (t *taskAPI) UpdateTaskCategory(w http.ResponseWriter, r *http.Request) {
	var task entity.TaskCategoryRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid decode json"))
		return
	}

	userId := r.Context().Value("id")

	idLogin, err := strconv.Atoi(userId.(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	var updateTask = entity.Task{
		ID:         task.ID,
		CategoryID: task.CategoryID,
		UserID:     int(idLogin),
	}

	_, err = t.taskService.UpdateTask(r.Context(), &updateTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": userId,
		"task_id": task.ID,
		"message": "success update task category",
	})
}
