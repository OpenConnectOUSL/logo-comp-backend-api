CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    firstName CHAR(255) NOT NULL,
    lastName CHAR(255) NOT NULL,
    registrationNumber integer NOT NULL,
    studyProgram CHAR(255) NOT NULL,
    faculty CHAR(255) NOT NULL,
    academicYear integer NOT NULL,
    email citext UNIQUE NOT NULL
);