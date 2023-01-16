package routers

import (
	"github.com/wanton-idol/TO-DO-APP/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/task", controllers.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task/{id}", controllers.TaskCompleted).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", controllers.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", controllers.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/task", controllers.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/deleteAllTask", controllers.DeleteAllTasks).Methods("DELETE", "OPTIONS")

	return router
}
