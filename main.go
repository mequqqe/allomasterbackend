package main

import (
	"log"
	"main/controllers"
	"main/middleware"
	"main/models" // Импортируйте ваши модели
	"main/repositories"
	"main/services"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Подключение к базе данных
	dsn := "host=localhost user=postgres password=1234 dbname=allomaster port=5432 sslmode=disable TimeZone=Asia/Almaty"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	// Выполнение автомиграции
	err = db.AutoMigrate(&models.Company{}, &models.Branch{}, &models.Employee{})
	if err != nil {
		log.Fatal("failed to auto migrate models")
	}

	// Инициализация репозиториев, сервисов и контроллеров
	companyRepo := repositories.NewCompanyRepository(db)
	companyService := services.NewCompanyService(companyRepo)
	companyController := &controllers.CompanyController{CompanyService: companyService}

	branchRepo := repositories.NewBranchRepository(db)
	branchService := services.NewBranchService(branchRepo)
	branchesController := controllers.NewBranchesController(branchService)

	employeeRepo := repositories.NewEmployeeRepository(db)
	employeeService := services.NewEmployeeService(employeeRepo)
	employeesController := controllers.NewEmployeesController(employeeService)

	// Настройка роутера
	r := mux.NewRouter()
	r.Use(middleware.CORS)

	r.Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Роуты для публичных эндпоинтов
	r.HandleFunc("/api/company/register", companyController.Register).Methods("POST")
	r.HandleFunc("/api/company/login", companyController.Login).Methods("POST")

	// Защищенные роуты с использованием JWT middleware
	secured := r.PathPrefix("/api").Subrouter()
	secured.Use(middleware.JWTMiddleware)
	secured.HandleFunc("/company/my-company", companyController.GetCompanyInfo).Methods("GET")
	secured.HandleFunc("/branches", branchesController.AddBranch).Methods("POST")
	secured.HandleFunc("/branches", branchesController.GetMyBranches).Methods("GET")
	secured.HandleFunc("/employees", employeesController.AddEmployee).Methods("POST")
	secured.HandleFunc("/employees", employeesController.GetEmployees).Methods("GET")

	// Запуск сервера
	log.Fatal(http.ListenAndServe(":8000", r))
}
