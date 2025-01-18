package handlers

import (
	"context"
	"moy_proekt/internal/taskService"
	"moy_proekt/internal/web/tasks"
)

type TaskHandler struct {
	service *taskService.TaskService
}

func NewHandler(service *taskService.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	tasksList, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}
	response := tasks.GetTasks200JSONResponse{}
	for _, t := range tasksList {
		response = append(response, tasks.Task{
			Id: func(id uint) *uint32 {
				v := uint32(id)
				return &v
			}(t.ID),
			Task:   &t.Task,
			IsDone: &t.IsDone,
		})
	}
	return response, nil
}

func (h *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	newTask := taskService.Task{
		Task:   *request.Body.Task,
		IsDone: *request.Body.IsDone,
	}
	createdTask, err := h.service.CreateTask(newTask)
	if err != nil {
		return nil, err
	}
	response := tasks.PostTasks201JSONResponse{
		Id: func(id uint) *uint32 {
			v := uint32(id)
			return &v
		}(createdTask.ID),
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	return response, nil
}

func (h *TaskHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	err := h.service.DeleteTaskByID(uint(request.Id))
	if err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId204Response{}, nil
}

func (h *TaskHandler) GetTasksId(ctx context.Context, request tasks.GetTasksIdRequestObject) (tasks.GetTasksIdResponseObject, error) {
	task, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}
	if len(task) == 0 {
		return nil, nil
	}
	response := tasks.GetTasksId200JSONResponse{
		Id: func(id uint) *uint32 {
			v := uint32(task[0].ID)
			return &v
		}(task[0].ID),
		Task:   &task[0].Task,
		IsDone: &task[0].IsDone,
	}
	return response, nil
}

func (h *TaskHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	updatedTask := taskService.Task{
		Task:   *request.Body.Task,
		IsDone: *request.Body.IsDone,
	}
	result, err := h.service.UpdateTaskByID(uint(request.Id), updatedTask)
	if err != nil {
		return nil, err
	}
	response := tasks.PatchTasksId200JSONResponse{
		Id: func(id uint) *uint32 {
			v := uint32(id)
			return &v
		}(result.ID),
		Task:   &result.Task,
		IsDone: &result.IsDone,
	}
	return response, nil
}
