package handler

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/budougumi0617/go_todo_app/entity"
	"github.com/budougumi0617/go_todo_app/testutil"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func TestChangeTaskStatus(t *testing.T) {
	type want struct {
		status  int
		rspFile string
	}

	task := entity.Task{ID: 1, Title: "go to bed", Status: entity.TaskStatusTodo}
	tests := map[string]struct {
		reqFile string
		want    want
	}{
		"ok": {
			reqFile: "testdata/change_task_status/ok_req.json.golden",
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/change_task_status/ok_rsp.json.golden",
			},
		},
		"ng": {
			reqFile: "testdata/change_task_status/ng_req.json.golden",
			want: want{
				status:  http.StatusBadRequest,
				rspFile: "testdata/change_task_status/ng_rsp.json.golden",
			},
		},
	}
	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPatch,
				fmt.Sprintf("/tasks/%d", task.ID),
				bytes.NewReader(testutil.LoadFile(t, tt.reqFile)),
			)
			moq := &ChangeTaskStatusServiceMock{}
			moq.ChangeTaskStatusFunc = func(
				ctx context.Context, id entity.TaskID, status entity.TaskStatus,
			) (*entity.Task, error) {
				if tt.want.status == http.StatusOK {
					new := task
					new.Status = status
					return &new, nil
				}
				return nil, errors.New("error from mock")
			}

			sut := ChangeTaskStatus{
				Service:   moq,
				Validator: validator.New(),
			}
			router := chi.NewRouter()
			router.Patch("/tasks/{id}", sut.ServeHTTP)
			router.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(t,
				resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile),
			)
		})
	}
}
