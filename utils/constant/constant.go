package constant

// Success
const (
	SUCCESS_LOGIN     = "logged in successfully"
	SUCCESS_CREATED   = "data created successfully"
	SUCCESS_DELETED   = "data deleted successfully"
	SUCCESS_RETRIEVED = "data retrieved successfully"
)

// Error
const (
	ERROR_ID_NOTFOUND        = "id not found"
	ERROR_ID_INVALID         = "invalid id"
	ERROR_EMAIL_NOTFOUND     = "email not found"
	ERROR_EMAIL_FORMAT       = "invalid email format"
	ERROR_EMAIL_EXIST        = "email already exists"
	ERROR_EMAIL_UNREGISTERED = "email not registered"
	ERROR_LOGIN              = "incorrect email pr password"
	ERROR_PASSWORD_INVALID   = "invalid password"
	ERROR_PASSWORD_HASH      = "error hashing password"
	ERROR_DATA_NOTFOUND      = "data not found"
	ERROR_DATA_EMPTY         = "data is empty"
	ERROR_DATA_EXIST         = "data already exists"
	ERROR_DATA_TYPE          = "data type unsupported"
	ERROR_DATA_INVALID       = "invalid data. allowed data: "
	ERROR_FILE_EMPTY         = "file is empty"
	ERROR_DATE_FORMAT        = "invalid date format. expected format: '2001-12-30'"
	ERROR_MIN_LENGTH         = "minimum length is %d characters"
	ERROR_MAX_LENGTH         = "maximum length is %d characters"
	ERROR_TOKEN_INVALID      = "invalid token"
	ERROR_TOKEN_GENERATE     = "generate token failed"
	ERROR_TEMPLATE_FILE      = "invalid template file"
)
