package category

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type Handler interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type handler struct {
	Service Service
}

func NewHandler(service Service) Handler {
	return &handler{
		Service: service,
	}
}

func (handler *handler) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	createdRequest := CreateRequest{}
	err := decoder.Decode(&createdRequest)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := handler.Service.Create(request.Context(), createdRequest)
	if err != nil {
		//http.Error(writer, err.Error(), http.StatusInternalServerError)

		handlerResponse := HandlerResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		}

		writer.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(writer)
		err = encoder.Encode(handlerResponse)
		return
	}

	handlerResponse := HandlerResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	}

	writer.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(handlerResponse)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *handler) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	updateRequest := UpdateRequest{}
	err := decoder.Decode(&updateRequest)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	updateRequest.Id = int(categoryId)

	response, err := handler.Service.Update(request.Context(), updateRequest)
	if err != nil {
		//http.Error(writer, err.Error(), http.StatusInternalServerError)

		handlerResponse := HandlerResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		}

		writer.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(writer)
		err = encoder.Encode(handlerResponse)
		return
	}

	handlerResponse := HandlerResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(handlerResponse)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *handler) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.Service.Delete(request.Context(), categoryId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	handlerResponse := HandlerResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(handlerResponse)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *handler) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := handler.Service.FindById(request.Context(), categoryId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	handlerResponse := HandlerResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(handlerResponse)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *handler) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	responses, err := handler.Service.FindAll(request.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	handlerResponse := HandlerResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   responses,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(handlerResponse)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
