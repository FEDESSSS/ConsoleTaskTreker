package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/k0kubun/pp/v3"
)

const file = "task.json"

var tasks []Task
var tpTask []Task
var maxId int

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func main() {

	if len(os.Args) < 2 {
		ListComand()
		return
	}

	if err := runTask(); err != nil {
		log.Fatal(err)
	}
}

func ListComand() {
	fmt.Println("Список команд")
	fmt.Println("")
	fmt.Println("list - просмотр всех задач")
	fmt.Println("add - добавление новой задачи")
	fmt.Println("delete - удаление задачи")
	fmt.Println("update - изменить задачу")
	fmt.Println("mark-done - изменить статус задачи на выполнено")
	fmt.Println("mark-in-progress - изменить статус задачи на выполняемое")
	fmt.Println("list done - просмотр выполненных задач")
	fmt.Println("list todo - просмотр невыполненных задач")
	fmt.Println("list in-progress - просмотр выполняемых задач")
}

func runTask() error {
	args := os.Args[1:]
	cmd := args[0]
	switch cmd {
	case "add":
		if len(args) < 2 {
			return errors.New("неправильная команад, попробйте еще раз")
		}
		return AddTask(args[1])

	case "delete":
		if len(args) < 2 {
			return errors.New("неправильная команад, попробйте еще раз")
		}
		return DeleteTask(args[1])

	case "mark-done":
		if len(args) < 2 {
			return errors.New("неправильная команад, попробйте еще раз")
		}
		return UpdateDone(args[1])

	case "mark-in-progress":
		if len(args) < 2 {
			return errors.New("неправильная команад, попробйте еще раз")
		}
		return UpdateProgress(args[1])

	case "list":
		if len(args) > 2 {
			return errors.New("неправильная команад, попробйте еще раз")
		}
		if len(args) == 2 && args[1] == "done" {
			return ListDone()
		}
		if len(args) == 2 && args[1] == "in-progress" {
			return ListProgress()
		}
		return ListTask()

	case "update":
		if len(args) < 3 {
			return errors.New("неправильная команад, попробйте еще раз")
		}
		return UpdateTask(args[1], args[2])

	default:
		return fmt.Errorf("неизвестная команда: %s", cmd)
	}
}

func AddTask(task string) error {
	maxId++
	data, err := os.ReadFile(file)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}

	if len(data) > 0 {
		json.Unmarshal(data, &tasks)
	}

	maxId = 0
	for _, t := range tasks {
		if t.ID > maxId {
			maxId = t.ID
		}
	}

	newTask := Task{
		ID:          maxId + 1,
		Description: task,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, newTask)

	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка сериализации: %v", err)
	}

	if err := os.WriteFile(file, jsonData, 0644); err != nil {
		return fmt.Errorf("ошибка записи в файл: %v", err)
	}

	fmt.Println("Добавлена задача:")
	fmt.Println("===============================")
	pp.Println(newTask)
	return nil
}

func ListTask() error {
	data, err := os.ReadFile(file)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}

	if len(data) > 0 {
		json.Unmarshal(data, &tasks)
	}

	fmt.Println("Список задач:")
	fmt.Println("===============================")
	pp.Println(tasks)
	return nil

}

func ListDone() error {
	data, err := os.ReadFile(file)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}

	if len(data) > 0 {
		json.Unmarshal(data, &tasks)
	}

	for _, t := range tasks {
		if t.Status == "done" {
			tpTask = append(tpTask, t)
		}
	}

	fmt.Println("Список выполненных задач:")
	fmt.Println("===============================")
	pp.Println(tpTask)
	return nil
}

func ListProgress() error {
	data, err := os.ReadFile(file)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}

	if len(data) > 0 {
		json.Unmarshal(data, &tasks)
	}

	for _, t := range tasks {
		if t.Status == "in-progress" {
			tpTask = append(tpTask, t)
		}
	}

	fmt.Println("Список выполняемых задач:")
	fmt.Println("===============================")
	pp.Println(tpTask)
	return nil
}

func ListTodo() error {
	data, err := os.ReadFile(file)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}

	if len(data) > 0 {
		json.Unmarshal(data, &tasks)
	}

	for _, t := range tasks {
		if t.Status == "todo" {
			tpTask = append(tpTask, t)
		}
	}

	fmt.Println("Список невыполненных задач:")
	fmt.Println("===============================")
	pp.Println(tpTask)
	return nil
}

func DeleteTask(id string) error {
	number, _ := strconv.Atoi(id)
	data, err := os.ReadFile(file)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}

	if len(data) > 0 {
		json.Unmarshal(data, &tasks)
	}

	for i, j := range tasks {
		if j.ID == number {
			k := i + 1
			tasks = append(tasks[:i], tasks[k:]...)
		}
	}

	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка сериализации: %v", err)
	}

	if err := os.WriteFile(file, jsonData, 0644); err != nil {
		return fmt.Errorf("ошибка записи в файл: %v", err)
	}

	fmt.Println("Задача удалена:")
	fmt.Println("===============================")
	pp.Println(tasks)
	return nil
}

func UpdateTask(id string, description string) error {
	number, _ := strconv.Atoi(id)
	data, err := os.ReadFile(file)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}

	if len(data) > 0 {
		json.Unmarshal(data, &tasks)
	}

	for j := range tasks {
		if tasks[j].ID == number {
			tasks[j].Description = description
		}

	}

	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка сериализации: %v", err)
	}

	if err := os.WriteFile(file, jsonData, 0644); err != nil {
		return fmt.Errorf("ошибка сериализации: %v", err)
	}

	fmt.Println("Задача обновлена:")
	fmt.Println("===============================")
	pp.Println(tasks)
	return nil
}

func UpdateDone(id string) error {
	number, _ := strconv.Atoi(id)
	data, err := os.ReadFile(file)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}

	if len(data) > 0 {
		json.Unmarshal(data, &tasks)
	}

	for j := range tasks {
		if tasks[j].ID == number {
			tasks[j].Status = "done"
		}

	}

	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка сериализации: %v", err)
	}

	if err := os.WriteFile(file, jsonData, 0644); err != nil {
		return fmt.Errorf("ошибка сериализации: %v", err)
	}

	fmt.Println("Статус обновлен:")
	fmt.Println("===============================")
	pp.Println(tasks)
	return nil
}

func UpdateProgress(id string) error {
	number, _ := strconv.Atoi(id)
	data, err := os.ReadFile(file)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}

	if len(data) > 0 {
		json.Unmarshal(data, &tasks)
	}

	for j := range tasks {
		if tasks[j].ID == number {
			tasks[j].Status = "in-progress"
		}

	}

	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка сериализации: %v", err)
	}

	if err := os.WriteFile(file, jsonData, 0644); err != nil {
		return fmt.Errorf("ошибка сериализации: %v", err)
	}

	fmt.Println("Статус обновлен:")
	fmt.Println("===============================")
	pp.Println(tasks)
	return nil
}
