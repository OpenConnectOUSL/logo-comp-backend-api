package main

import (
	"errors"
	"fmt"
	"net/http"
	"oc.api.org/internal/data"
	"oc.api.org/internal/validator"
)

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		FirstName          string                  `json:"firstname"`
		LastName           string                  `json:"lastname"`
		RegistrationNumber data.RegistrationNumber `json:"registrationNumber"`
		StudyProgram       string                  `json:"studyprogram"`
		Faculty            string                  `json:"faculty"`
		Email              string                  `json:"email"`
		AcademicYear       int64                   `json:"academicyear"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &data.User{
		FirstName:          input.FirstName,
		LastName:           input.LastName,
		RegistrationNumber: input.RegistrationNumber,
		StudyProgram:       input.StudyProgram,
		Faculty:            input.Faculty,
		Email:              input.Email,
		AcademicYear:       input.AcademicYear,
	}

	v := validator.New()

	if data.ValidateUser(v, user); !v.Valid() {
		app.faildValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Users.Insert(user)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprint("/v1/users/", user.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"user": user}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) showUserHandler(w http.ResponseWriter, r *http.Request) {
	
	user, err := app.models.Users.Get()
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
