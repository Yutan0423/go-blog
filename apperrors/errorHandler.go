package apperrors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Yutan0423/go-medium-level/common"
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

	traceID := common.GetTraceID(req.Context())
	log.Printf("[%d]error: %s\n", traceID, appErr)

	var statusCode int

	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed, BadParam:
		// コメント投稿席として指定された記事がなかった（NoTargetData）場合と
		// リクエストボディのjsonでコードに失敗した（ReqBodyDecodeFailed）場合と
		// クエリパラメータの値が不正だった（BadParam）場合は
		// 400番（BadRequest）を使う
		statusCode = http.StatusBadRequest
	case RequiredAuthorizationHeader, Unauthorized:
		statusCode = http.StatusUnauthorized
	default:
		// それ以外の場合には500番（InternalServerError）を使う
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
