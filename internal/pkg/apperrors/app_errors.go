package apperrors

// このファイルで独自エラーを定義したけど使わないかも

// BadRequestError リクエストが不正
type BadRequestError struct {
	Message string
}

func (e *BadRequestError) Error() string {
	return e.Message
}
func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{Message: message}
}

// NotFoundError リソースが見つからない
type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}
func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{Message: message}
}

// ForbiddenError アクセス権限がない
type ForbiddenError struct {
	Message string
}

func (e *ForbiddenError) Error() string {
	return e.Message
}
func NewForbiddenError(message string) *ForbiddenError {
	return &ForbiddenError{Message: message}
}

// UnauthorizedError 認証が必要
type UnauthorizedError struct {
	Message string
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}
func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{Message: message}
}

// InternalServerError 内部サーバーエラー
type InternalServerError struct {
	Message string
}

func (e *InternalServerError) Error() string {
	return e.Message
}
func NewInternalServerError(message string) *InternalServerError {
	return &InternalServerError{Message: message}
}

// SessionTokenNotFoundError SessionToken cookieが見つからない
type SessionTokenNotFoundError struct{}

func (e *SessionTokenNotFoundError) Error() string {
	return "SessionToken cookie not found"
}
func NewSessionTokenNotFoundError() error {
	return &SessionTokenNotFoundError{}
}

// AlreadyInGuildError 既にギルドに所属している
type AlreadyInGuildError struct{}

func (e *AlreadyInGuildError) Error() string {
	return "Already in the guild"
}

// NewAlreadyInGuildError 既にギルドに所属しているエラーを返す
func NewAlreadyInGuildError() error {
	return &AlreadyInGuildError{}
}
