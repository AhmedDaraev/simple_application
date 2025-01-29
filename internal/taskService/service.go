package taskService

// TaskService предоставляет бизнес-логику для работы с задачами
type TaskService struct {
	repo TaskRepository // Интерфейс репозитория для работы с задачами
}

// NewService создает новый экземпляр TaskService
func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

// CreateTask создает новую задачу
func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}

// GetAllTasks возвращает все задачи
func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

// UpdateTaskByID обновляет задачу по ID
func (s *TaskService) UpdateTaskByID(id uint, task Task) (Task, error) {
	return s.repo.UpdateTaskByID(id, task)
}

// DeleteTaskByID удаляет задачу по ID
func (s *TaskService) DeleteTaskByID(id uint) error {
	return s.repo.DeleteTaskByID(id)
}

// GetTaskByID возвращает задачу по ID
func (s *TaskService) GetTaskByID(id uint) (*Task, error) {
	return s.repo.GetTaskByID(id)
}
