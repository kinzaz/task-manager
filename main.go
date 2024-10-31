package main

import (
	"fmt"
	"os"
	"strconv"
	"tasks/task"
)

const TASK_FILENAME = "task.json"

var actionMap = map[string]func(){
	"add":    addTask,
	"list":   readTasks,
	"done":   updateStatus,
	"delete": deleteTask,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Необходимо указать команду (add, list, done, delete)")
		return
	}

	command := os.Args[1]
	action, exists := actionMap[command]
	if !exists {
		fmt.Println("Неизвестная команда:", command)
		return
	}

	action()
}

func deleteTask() {
	if len(os.Args) < 3 {
		fmt.Println("Необходимо указать ID задачи для удаления")
		return
	}

	id, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Println("Ошибка конвертации:", err)
		return
	}

	err = task.DeleteTask(TASK_FILENAME, id)
	if err != nil {
		fmt.Println("Ошибка при удалении задачи:", err)
	}
}

func updateStatus() {
	if len(os.Args) < 3 {
		fmt.Println("Необходимо указать ID задачи для обновления")
		return
	}
	id, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Println("Ошибка конвертации:", err)
		return
	}
	err = task.UpdateTaskStatus(TASK_FILENAME, id)
	if err != nil {
		fmt.Println("Ошибка при обновлении задачи:", err)
	}
}

func readTasks() {
	tasks := task.ReadTasks(TASK_FILENAME)
	for _, data := range tasks {
		var isDone string
		if data.Done {
			isDone = "[завершено]"
		} else {
			isDone = "[не завершено]"
		}
		fmt.Printf("%v. %v: %v\n", data.ID, data.Description, isDone)
	}
}

func addTask() {
	if len(os.Args) < 3 {
		fmt.Println("Необходимо указать описание задачи")
		return
	}
	taskDescription := os.Args[2]

	taskId, err := task.AddTask(taskDescription, TASK_FILENAME)
	if err != nil {
		fmt.Println("Не удалось добавить задачу:", err)
		return
	}
	fmt.Printf("Задача успешно добавлена с ID %v\n", taskId)
}
