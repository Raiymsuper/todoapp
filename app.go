package main

import (
	"context"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Task представляет задачу
type Task struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

// App структура для хранения задач
type App struct {
	db  *gorm.DB
	ctx context.Context
}

// NewApp создаёт новый экземпляр приложения
func NewApp() *App {
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Println("❌ Не удалось подключиться к базе данных:", err)
		return nil
	}

	// Автоматическая миграция схемы
	if err := db.AutoMigrate(&Task{}); err != nil {
		log.Println("❌ Ошибка миграции:", err)
		return nil
	}

	log.Println("✅ База данных успешно подключена")
	return &App{db: db}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	log.Println("🚀 Приложение запущено!")
}

// AddTask добавляет новую задачу
func (a *App) AddTask(text string) (*Task, error) {
	if a.db == nil {
		return nil, fmt.Errorf("❌ database is not initialized")
	}

	task := Task{Text: text, Done: false}
	if err := a.db.Create(&task).Error; err != nil {
		log.Println("❌ Ошибка при добавлении задачи:", err)
		return nil, err
	}

	log.Println("✅ Задача добавлена:", task)
	return &task, nil
}

// GetTasks возвращает все задачи
func (a *App) GetTasks() ([]Task, error) {
	if a.db == nil {
		return nil, fmt.Errorf("❌ database is not initialized")
	}

	var tasks []Task
	if err := a.db.Find(&tasks).Error; err != nil {
		log.Println("❌ Ошибка при получении задач:", err)
		return nil, err
	}

	log.Println("📋 Задачи получены:", tasks)
	return tasks, nil
}

// ToggleTaskStatus меняет статус выполнения задачи
func (a *App) ToggleTaskStatus(id int) error {
	if a.db == nil {
		return fmt.Errorf("❌ database is not initialized")
	}

	var task Task
	if err := a.db.First(&task, id).Error; err != nil {
		log.Println("❌ Задача не найдена:", err)
		return err
	}

	task.Done = !task.Done
	if err := a.db.Save(&task).Error; err != nil {
		log.Println("❌ Ошибка при обновлении задачи:", err)
		return err
	}

	log.Println("✅ Статус задачи изменен:", task)
	return nil
}

// DeleteTask удаляет задачу по ID
func (a *App) DeleteTask(id int) error {
	if a.db == nil {
		return fmt.Errorf("❌ database is not initialized")
	}

	var task Task
	if err := a.db.First(&task, id).Error; err != nil {
		log.Println("❌ Задача не найдена:", err)
		return err
	}

	if err := a.db.Delete(&task).Error; err != nil {
		log.Println("❌ Ошибка при удалении задачи:", err)
		return err
	}

	log.Println("🗑️ Задача удалена:", task)
	return nil
}
