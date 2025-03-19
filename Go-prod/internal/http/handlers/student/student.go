package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/anmol1377/student-api/internal/storage"
	"github.com/anmol1377/student-api/internal/types"
	"github.com/anmol1377/student-api/internal/utils/response"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("khali body h")))
			// response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		// w.Write([]byte("welcome to student api prod!"))
		slog.Info("creating student")

		LastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
			student.CreatedAt,
			student.UpdatedAt,
		)
		slog.Info("user created sucessfyl", slog.String("userid", fmt.Sprintf("%d", LastId)))

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": LastId})

	}

}

func GetID(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")
		slog.Info("getting data of studernt", slog.String("id", id))

		Intid, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("id is not valid")))
			return
		}
		student, err := storage.GetStudent(Intid)
		if err != nil {
			slog.Info("error getting student", slog.String("error", err.Error()))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, student)

	}

}
