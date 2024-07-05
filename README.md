# TO DO LIST

This is a simple To-Do List application built with Golang and Redis, hosted using Docker. The application allows users to add tasks, delete tasks, view all tasks, and mark tasks as done through the terminal.

# Features
1. Add a task
2. View all tasks
3. Delete a task
4. Mark a task as done

# Prerequisites
1. Golang
2. Docker - compose
3. Redis

After you have installed golang in your device, build and run the docker containers using docker-compose. Then run the main file and you can proceed to have a working to-do list. If you have docker-compose installed, then you can also view the changes there.

# Project files
1. main.go :  The main application code
2. redis-docker.yaml : configuration file for docker-compose
3. opr1 : Contains the AddItem.go file that is required for adding a task to the list
4. opr2 : Contains the ViewAll.go file that will show the list of all tasks
5. opr3 : Contains the DeleteItem.go file that contains the code to delete a task
6. opr4 : Contains the MarkAsDone.go file that will mark a task as done


