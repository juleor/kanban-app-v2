package api

import (
	"a21hc3NpZ25tZW50/entity"
	"a21hc3NpZ25tZW50/service"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type UserAPI interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)

	Delete(w http.ResponseWriter, r *http.Request)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) *userAPI {
	return &userAPI{userService}
}

func (u *userAPI) Login(w http.ResponseWriter, r *http.Request) {
	var user entity.UserLogin
	var userLogin entity.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid decode json"))
		return
	}

	if user.Email == "" || user.Password == "" {
		w.WriteHeader(400)
		status := entity.ErrorResponse{
			Error: "email or password is empty",
		}
		jsonStatus, _ := json.Marshal(status)
		w.Write(jsonStatus)
		return
	}

	userLogin = entity.User{
		Email:    user.Email,
		Password: user.Password,
	}

	login, err := u.userService.Login(r.Context(), &userLogin)
	if err != nil {
		w.WriteHeader(500)
		status := entity.ErrorResponse{
			Error: "error internal server",
		}
		jsonStatus, _ := json.Marshal(status)
		w.Write(jsonStatus)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "user_id",
		Value: fmt.Sprintf("%d", login),
	})

	w.WriteHeader(200)
	status := map[string]interface{}{
		"user_id": login,
		"message": "login success",
	}
	jsonStatus, _ := json.Marshal(status)
	w.Write(jsonStatus)
	return

	// TODO: answer here
}

func (u *userAPI) Register(w http.ResponseWriter, r *http.Request) {
	var user entity.UserRegister
	var newUser entity.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid decode json"))
		return
	}

	if user.Email == "" || user.Fullname == "" || user.Password == "" {
		w.WriteHeader(400)
		status := entity.ErrorResponse{
			Error: "register data is empty",
		}
		jsonStatus, _ := json.Marshal(status)
		w.Write(jsonStatus)
		return
	}

	newUser = entity.User{
		Fullname: user.Fullname,
		Email:    user.Email,
		Password: user.Password,
	}

	create, err := u.userService.Register(r.Context(), &newUser)
	if err == errors.New("email already exists") {
		w.WriteHeader(500)
		status := entity.ErrorResponse{
			Error: "email already used",
		}
		jsonStatus, _ := json.Marshal(status)
		w.Write(jsonStatus)
		return
	}

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
		"user_id": create.ID,
		"message": "register success",
	}
	jsonStatus, _ := json.Marshal(status)
	w.Write(jsonStatus)
	return

	// TODO: answer here
}

func (u *userAPI) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "user_id",
		Value:   "",
		Expires: time.Now(),
	})

	w.WriteHeader(200)
	status := map[string]interface{}{
		"message": "logout success",
	}
	jsonStatus, _ := json.Marshal(status)
	w.Write(jsonStatus)
	return
	// TODO: answer here
}

func (u *userAPI) Delete(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")

	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("user_id is empty"))
		return
	}

	deleteUserId, _ := strconv.Atoi(userId)

	err := u.userService.Delete(r.Context(), int(deleteUserId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "delete success"})
}
