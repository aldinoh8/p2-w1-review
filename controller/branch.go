package controller

import (
	"demo/entity"
	"demo/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type Branch struct {
	Repository repository.Branch
}

func (b Branch) Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	branches := b.Repository.FindAll()

	WriteResponseJson(w, http.StatusOK, branches)
}

func (b Branch) Detail(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	idStr := p.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		WriteResponseJson(w, http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	branch, err := b.Repository.FindById(id)
	if err != nil {
		WriteResponseJson(w, http.StatusNotFound, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	WriteResponseJson(w, http.StatusOK, branch)
}

func (b Branch) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	reqBody := entity.Branch{}
	err := decoder.Decode(&reqBody)

	if err != nil {
		WriteResponseJson(w, http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	if len(reqBody.Name) == 0 || len(reqBody.Location) == 0 {
		WriteResponseJson(w, http.StatusBadRequest, ErrorResponse{
			Message: "Name and Location required",
		})
		return
	}

	newBranch := b.Repository.Create(reqBody.Name, reqBody.Location)

	WriteResponseJson(w, http.StatusCreated, map[string]interface{}{
		"message": "success create",
		"data":    newBranch,
	})
}

func (b Branch) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Update")
}

func (b Branch) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Delete")
}
