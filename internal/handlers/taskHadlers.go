package handlers

import (
	"context"
	"moy_proekt/internal/taskService"
	"moy_proekt/internal/web/tasks"
)

type TaskHandler struct {
	service *taskService.TaskService
}

func NewTaskHandler(service *taskService.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) GetTasks(ctx context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	tasksList, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := make(tasks.GetTasks200JSONResponse, len(tasksList))
	for i, task := range tasksList {
		response[i] = tasks.Task{
			Id:     toIntPtr(int(task.ID)),
			Task:   &task.Task,
			IsDone: &task.IsDone,
		}
	}
	return response, nil
}

func (h *TaskHandler) PostTasks(ctx context.Context, req tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	newTask := taskService.Task{
		Task:   req.Body.Task,
		IsDone: *req.Body.IsDone,
	}
	createdTask, err := h.service.CreateTask(newTask)
	if err != nil {
		return nil, err
	}
	return tasks.PostTasks201JSONResponse{
		Id:     toIntPtr(int(createdTask.ID)),
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}, nil
}

func (h *TaskHandler) PatchTasksId(ctx context.Context, req tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	updatedTask := taskService.Task{
		Task:   *req.Body.Task,
		IsDone: *req.Body.IsDone,
	}
	result, err := h.service.UpdateTaskByID(uint(req.Id), updatedTask)
	if err != nil {
		return nil, err
	}
	return tasks.PatchTasksId200JSONResponse{
		Id:     toIntPtr(int(result.ID)),
		Task:   &result.Task,
		IsDone: &result.IsDone,
	}, nil
}

func (h *TaskHandler) DeleteTasksId(ctx context.Context, req tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	if err := h.service.DeleteTaskByID(uint(req.Id)); err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId204Response{}, nil
}

func (h *TaskHandler) GetTasksId(ctx context.Context, req tasks.GetTasksIdRequestObject) (tasks.GetTasksIdResponseObject, error) {
	task, err := h.service.GetTaskByID(uint(req.Id))
	if err != nil {
		return nil, err
	}
	if task == nil {
		return tasks.GetTasksId404Response{}, nil
	}
	return tasks.GetTasksId200JSONResponse{
		Id:     toIntPtr(int(task.ID)),
		Task:   &task.Task,
		IsDone: &task.IsDone,
	}, nil
}
