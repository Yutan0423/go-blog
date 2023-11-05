package apperrors

import (
	"encoding/json"
	"errors"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	var appErr *MyAppError
	// 引数のerrをMyAppError型のappErrに変換
	if !errors.As(err, &appErr) {
		// もし変換に失敗したらUnKnownエラーを変数appErrに手動で格納
		appErr = &MyAppError{
			ErrCode: UnKnown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	var statusCode int

	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
