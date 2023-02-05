package api

import (
	"a21hc3NpZ25tZW50/entity"
	"a21hc3NpZ25tZW50/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type CategoryAPI interface {
	GetCategory(w http.ResponseWriter, r *http.Request)
	CreateNewCategory(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)
	GetCategoryWithTasks(w http.ResponseWriter, r *http.Request)
}

type categoryAPI struct {
	categoryService service.CategoryService
}

func NewCategoryAPI(categoryService service.CategoryService) *categoryAPI {
	return &categoryAPI{categoryService}
}

func (c *categoryAPI) GetCategory(w http.ResponseWriter, r *http.Request) {
	id := fmt.Sprintf("%s", r.Context().Value("id"))

	if id == "" {
		w.WriteHeader(400)
		status := entity.ErrorResponse{
			Error: "invalid user id",
		}
		jsonStatus, _ := json.Marshal(status)
		w.Write(jsonStatus)
		return
	}

	idInt, _ := strconv.Atoi(id)
	create, err := c.categoryService.GetCategories(r.Context(), idInt)
	if err != nil {
		w.WriteHeader(500)
		status := entity.ErrorResponse{
			Error: "error internal server",
		}
		jsonStatus, _ := json.Marshal(status)
		w.Write(jsonStatus)
		return
	}

	w.WriteHeader(200)
	jsonStatus, _ := json.Marshal(create)
	w.Write(jsonStatus)
	return
	// TODO: answer here
}

func (c *categoryAPI) CreateNewCategory(w http.ResponseWriter, r *http.Request) {
	var category entity.CategoryRequest
	var newCategory entity.Category

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid category request"))
		return
	}

	if category.Type == "" {
		w.WriteHeader(400)
		status := entity.ErrorResponse{
			Error: "invalid category request",
		}
		jsonStatus, _ := json.Marshal(status)
		w.Write(jsonStatus)
		return
	}

	id := fmt.Sprintf("%s", r.Context().Value("id"))

	if id == "" {
		w.WriteHeader(400)
		status := entity.ErrorResponse{
			Error: "invalid user id",
		}
		jsonStatus, _ := json.Marshal(status)
		w.Write(jsonStatus)
		return
	}
	idInt, _ := strconv.Atoi(id)

	newCategory = entity.Category{
		Type:   category.Type,
		UserID: idInt,
	}

	categori, err := c.categoryService.StoreCategory(r.Context(), &newCategory)
	if err != nil {
		w.WriteHeader(500)
		status := entity.ErrorResponse{
			Error: "error internal server",
		}
		jsonStatus, _ := json.Marshal(status)
		w.Write(jsonStatus)
		return
	}

	w.WriteHeader(201)
	status := map[string]interface{}{
		"user_id":     categori.UserID,
		"category_id": categori.ID,
		"message":     "success create new category",
	}
	jsonStatus, _ := json.Marshal(status)
	w.Write(jsonStatus)
	return

	// TODO: answer here
}

func (c *categoryAPI) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := fmt.Sprintf("%s", r.Context().Value("id"))
	categoryID := r.URL.Query().Get("category_id")

	idInt, _ := strconv.Atoi(categoryID)
	idUserInt, _ := strconv.Atoi(id)
	err := c.categoryService.DeleteCategory(r.Context(), idInt)
	if err != nil {
		w.WriteHeader(500)
		status := entity.ErrorResponse{
			Error: "error internal server",
		}
		jsonStatus, _ := json.Marshal(status)
		w.Write(jsonStatus)
		return
	}

	w.WriteHeader(200)
	status := map[string]interface{}{
		"user_id":     idUserInt,
		"category_id": idInt,
		"message":     "success delete category",
	}
	jsonStatus, _ := json.Marshal(status)
	w.Write(jsonStatus)
	return

	// TODO: answer here
}

func (c *categoryAPI) GetCategoryWithTasks(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("id")

	idLogin, err := strconv.Atoi(userId.(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("get category task", err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	categories, err := c.categoryService.GetCategoriesWithTasks(r.Context(), int(idLogin))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)

}
