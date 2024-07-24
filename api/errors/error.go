package errors

import (
	"errors"
)

// Define your error codes
const (
	ErrNotFound           = "ErrNotFound"
	ErrInvalidInput       = "ErrInvalidInput"
	ErrUnauthorized       = "ErrUnauthorized"
	ErrInternal           = "ErrInternal"
	ErrExist              = "ErrExist"
	ErrForbidden          = "ErrForbidden"
	ErrTimeout            = "ErrTimeout"
	ErrConflict           = "ErrConflict"
	ErrBadRequest         = "ErrBadRequest"
	ErrNotImplemented     = "ErrNotImplemented"
	ErrServiceUnavailable = "ErrServiceUnavailable"
	ErrTooManyRequests    = "ErrTooManyRequests"
	// Add more error codes as needed
)

// Define error messages in English
var errorMessagesEn = map[string]error{
	ErrNotFound:           errors.New("resource not found"),
	ErrInvalidInput:       errors.New("invalid input provided"),
	ErrUnauthorized:       errors.New("unauthorized access"),
	ErrInternal:           errors.New("internal server error"),
	ErrExist:              errors.New("resource already exists"),
	ErrForbidden:          errors.New("forbidden access"),
	ErrTimeout:            errors.New("request timed out"),
	ErrConflict:           errors.New("conflict with current state"),
	ErrBadRequest:         errors.New("bad request"),
	ErrNotImplemented:     errors.New("not implemented"),
	ErrServiceUnavailable: errors.New("service unavailable"),
	ErrTooManyRequests:    errors.New("too many requests"),
	// Add more error messages as needed
}

// Define error messages in French
var errorMessagesFr = map[string]error{
	ErrNotFound:           errors.New("ressource non trouvée"),
	ErrInvalidInput:       errors.New("entrée invalide fournie"),
	ErrUnauthorized:       errors.New("accès non autorisé"),
	ErrInternal:           errors.New("erreur interne du serveur"),
	ErrExist:              errors.New("ressource déjà existante"),
	ErrForbidden:          errors.New("accès interdit"),
	ErrTimeout:            errors.New("temps de la demande écoulé"),
	ErrConflict:           errors.New("conflit avec l'état actuel"),
	ErrBadRequest:         errors.New("mauvaise requête"),
	ErrNotImplemented:     errors.New("non implémenté"),
	ErrServiceUnavailable: errors.New("service indisponible"),
	ErrTooManyRequests:    errors.New("trop de demandes"),
	// Add more error messages as needed
}

func (e *Err) GetErrorMessage(code string) error {
	switch e.Lang {
	case "fr":
		if msg, exists := errorMessagesFr[code]; exists {
			return msg
		}
	default:
		if msg, exists := errorMessagesEn[code]; exists {
			return msg
		}
	}
	return nil
}
func (e *Err) ErrInternal() error {
	return e.GetErrorMessage(ErrInternal)
}

func (e *Err) ErrNotFound() error {
	return e.GetErrorMessage(ErrNotFound)
}
func (e *Err) New(err string) error {
	return errors.New(err)
}

type Err struct {
	Lang string
}
