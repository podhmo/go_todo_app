package handler

import (
	"encoding/json"
	"net/http"

	"github.com/budougumi0617/go_todo_app/entity"
	"github.com/go-playground/validator/v10"
)

type ChangeTaskStatus struct {
	Service   ChangeTaskStatusService
	Validator *validator.Validate
}

func (h *ChangeTaskStatus) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		ID     entity.TaskID     `json:"id" validate:"required"`
		Status entity.TaskStatus `json:"status" validate:"required"` // todo: enum
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	if err := h.Validator.Struct(b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	t, err := h.Service.ChangeTaskStatus(ctx, b.ID, b.Status)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := struct {
		*entity.Task
	}{Task: t}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
