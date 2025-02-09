package main

import (
 "bufio"
 "fmt"
 "os"
 "strconv"
 "strings"
)

type Task struct {
 Title string
}

var tasks []Task

func main() {
 scanner := bufio.NewScanner(os.Stdin)
 for {
  fmt.Println("To-Do List:")
  printTasks()

  fmt.Println("Choose an action: [add/remove/exit]")
  scanner.Scan()
  action := scanner.Text()

  switch strings.ToLower(action) {
  case "add":
   fmt.Println("Enter task title:")
   scanner.Scan()
   taskTitle := scanner.Text()
   addTask(taskTitle)
   fmt.Println("Task added.")
  case "remove":
   fmt.Println("Enter task number to remove:")
   scanner.Scan()
   taskNumber := scanner.Text()
   removeTask(taskNumber)
  case "exit":
   fmt.Println("Exiting...")
   return
  default:
   fmt.Println("Unknown action, please try again.")
  }
 }
}

func addTask(title string) {
 tasks = append(tasks, Task{Title: title})
}

func removeTask(taskNumber string) {
 index, err := strconv.Atoi(taskNumber) 
 if err != nil || index < 1 || index > len(tasks) {
  fmt.Println("Invalid task number.")
  return
 }
 index-- 
 tasks = append(tasks[:index], tasks[index+1:]...)
 fmt.Println("Task removed.")
}

func printTasks() {
 if len(tasks) == 0 {
  fmt.Println("No tasks available.")
  return
 }
 for i, task := range tasks {
  fmt.Printf("%d: %s\n", i+1, task.Title)
 }
}