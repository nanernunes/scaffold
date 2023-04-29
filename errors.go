package scaffold

type ResponseError struct {
	// Message with an error
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
} //	@name	Error
