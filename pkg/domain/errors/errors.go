package errors

import "errors"

var (
	ErrWrongEmailFormat       = errors.New("WRONG_EMAIL_FORMAT")
	ErrEmptyUsername          = errors.New("USERNAME_CANNOT_BE_EMPTY")
	ErrWrongPasswordFormat    = errors.New("WRONG_PASSWORD_FORMAT")
	ErrAgeRestriction         = errors.New("MUST_BE_MORE_THAN_8_YEAR_OLD")
	ErrBadRequest             = errors.New("BAD_REQUEST")
	ErrInternalServerError    = errors.New("INTERNAL_SERVER_ERROR")
	ErrEmailNotFound          = errors.New("EMAIL_NOT_FOUND")
	ErrEmailUsed              = errors.New("EMAIL_USED")
	ErrUnauthorizhedReq       = errors.New("UNAUTHORIZED_REQUEST")
	ErrDobEmpty               = errors.New("DATE_OF_BIRTH_CAN'T_BE_EMPTY")
	ErrUserNameUsed           = errors.New("USERNAME_IS_USED")
	ErrWrongPassword          = errors.New("WRONG_PASSWORD")
	ErrUsernameNotFound       = errors.New("USERNAME_NOT_FOUND")
	ErrUserNotFound           = errors.New("USER_NOT_FOUND")
	ErrNameEmpty              = errors.New("NAME_CANNOT_BE_EMPTY")
	ErrUrlEmpty               = errors.New("URL_CANNOT_BE_EMPTY")
	ErrWrongEmailFormatMsg    = errors.New("wrong email format")
	ErrEmptyUsernameMsg       = errors.New("username cannot be empty")
	ErrWrongPasswordFormatMsg = errors.New("password must be more than 6 characters")
	ErrAgeRestrictionMsg      = errors.New("must be more than 8 year old")
	ErrBindPayload            = errors.New("failed to bind payload")
	ErrInternalServerErrorMsg = errors.New("server error")
	ErrEmailNotFoundMsg       = errors.New("email not found")
	ErrDobEmptyMsg            = errors.New("date of birth cannot be empty")
	ErrEmailUsedMsg           = errors.New("email is already registered")
	ErrUserNameUsedMsg        = errors.New("username is already taken")
	ErrUnauthorizhedReqMsg    = errors.New("unauthorized request")
	ErrWrongPasswordMsg       = errors.New("password mismatch")
	ErrUsernameNotFoundMsg    = errors.New("can't find username")
	ErrUserNotFoundMsg        = errors.New("user not found")
	ErrNameEmptyMsg           = errors.New("name cannot be empty")
	ErrUrlEmptyMsg            = errors.New("url cannot be empty")
)
