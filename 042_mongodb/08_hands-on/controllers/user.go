package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/cannahum/golang-web-dev/042_mongodb/08_hands-on/models"
	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
)

type UserController struct {
	session map[string]models.User
}

func NewUserController(m map[string]models.User) *UserController {
	return &UserController{m}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Retrieve user
	u := uc.session[id]

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	// create ID
	newID, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	u.Id = newID.String()

	// store the user
	uc.session[u.Id] = u

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Marshalling error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = uc.writeSessionsToFile()
	if err != nil {
		fmt.Println("[create] writeSessionsToFile error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	delete(uc.session, id)

	err := uc.writeSessionsToFile()
	if err != nil {
		fmt.Println("[delete] writeSessionsToFile error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", id, "\n")
}

func (uc UserController) writeSessionsToFile() error {
	newFile, err := os.Create(models.UsersFileName)
	if err != nil {
		return err
	}
	return json.NewEncoder(newFile).Encode(uc.session)
}
