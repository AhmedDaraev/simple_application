package taskService

type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) UpdateTaskByID(id int, task Task) (Task, error) {
	uintID := uint(id) // Преобразуем int в uint
	updatedTask, err := s.repo.UpdateTaskByID(uintID, task)
	if err != nil {
		return Task{}, err
	}
	return updatedTask, nil
}

func (s *TaskService) DeleteTaskByID(id int) error {
	uintID := uint(id) // Преобразуем int в uint
	return s.repo.DeleteTaskByID(uintID)
}
