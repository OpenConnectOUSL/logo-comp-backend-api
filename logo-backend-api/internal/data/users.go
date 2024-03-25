package data

import (
	"database/sql"
	"errors"
	"regexp"
	"time"

	"oc.api.org/internal/validator"
)

var EmailRX = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,16}$`)

type User struct {
	ID                 int64              `json:"id"`
	CreatedAt          time.Time          `json:"created_at"`
	FirstName          string             `json:"firstname"`
	LastName           string             `json:"lastname"`
	RegistrationNumber RegistrationNumber `json:"registrationNumber,omitempty"`
	StudyProgram       string             `json:"studyprogram,omitempty"`
	Faculty            string             `json:"faculty,omitempty"`
	Email              string             `json:"email,omitempty"`
	AcademicYear       int64              `json:"academicyear,omitempty"`
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.FirstName != "", "firstname", "must be provided")
	v.Check(len(user.FirstName) <= 500, "firstname", "must not be more than 500 bytes long")
	v.Check(user.LastName != "", "lastname", "must be provided")
	v.Check(len(user.LastName) <= 500, "lastname", "must not be more than 500 bytes long")
	v.Check(user.RegistrationNumber != 0, "registrationnumber", "must be provided")
	v.Check(user.RegistrationNumber > 0, "registrationnumber", "must be a positive integer")
	v.Check(user.StudyProgram != "", "studyprogram", "must be provided")
	v.Check(len(user.StudyProgram) <= 500, "studyprogram", "must not be more than 500 bytes long")
	v.Check(user.Faculty != "", "faculty", "must be provided")
	v.Check(len(user.Faculty) <= 500, "faculty", "must not be more than 500 bytes long")
	v.Check(user.AcademicYear >= 1888, "academicyear", "must be greater than 1888")
	v.Check(user.Email != "", "email", "must be provided")
	v.Check(Matches(user.Email, EmailRX), "email", "must be a valid email address")
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

type UserModel struct {
	DB *sql.DB
}

func (u UserModel) Insert(user *User) error {
	query := `INSERT INTO users (firstname, lastname, registrationnumber, studyprogram, faculty, academicyear, email)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id, created_at`

	args := []interface{}{user.FirstName, user.LastName, user.RegistrationNumber, user.StudyProgram, user.Faculty, user.AcademicYear, user.Email}

	return u.DB.QueryRow(query, args...).Scan(&user.ID, &user.CreatedAt)
}

func (u UserModel) Get() ([]*User, error) {

	// if id < 1 {
	// 	return nil, nil
	// }

	query := `SELECT * FROM users WHERE id > 0`

	rows, err := u.DB.Query(query)
	
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*User
	for rows.Next(){
		var user User
		err := rows.Scan(
			&user.ID,
			&user.CreatedAt,
			&user.FirstName,
			&user.LastName,
			&user.RegistrationNumber,
			&user.StudyProgram,
			&user.Faculty,
			&user.AcademicYear,
			&user.Email,
		)
		users = append(users, &user)

		if err = rows.Err(); err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				return nil, ErrRecordNotFound
			default:
				return nil, err
			}

		}
	}

	if len(users) > 0 {
		return users, nil
	}

	return nil, nil
}
	

