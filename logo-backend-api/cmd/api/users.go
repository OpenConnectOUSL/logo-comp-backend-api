package main

import (
	"fmt"
	"net/http"
	"time"

	"oc.api.org/internal/data"
	"oc.api.org/internal/validator"
)

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	
	var input struct {
		FirstName string `json:"firstname"`
		LastName string `json:"lastname"`
		RegistrationNumber data.RegistrationNumber `json:"registrationNumber"`
		StudyProgram string `json:"studyprogram"`
		Faculty string `json:"faculty"`
		Email string `json:"email"`
		AcademicYear int64 `json:"academicyear"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &data.User{
		FirstName: input.FirstName,
		LastName: input.LastName,
		RegistrationNumber: input.RegistrationNumber,
		StudyProgram: input.StudyProgram,
		Faculty: input.Faculty,
		Email: input.Email,
		AcademicYear : input.AcademicYear,
	}

	v := validator.New()


	if data.ValidateUser(v, user);!v.Valid() {
		app.faildValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	user := data.User{
		ID: id,
		CreatedAt: time.Now(),
		FirstName: "Mayura",
		LastName: "Andrew", 
		RegistrationNumber: 831133,
		StudyProgram: "Bachelor Of Software Engineering",
		Faculty: "Faculty of Engineering Technology",
		Email: "mayura@gmail.com",
		AcademicYear: 2023,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
