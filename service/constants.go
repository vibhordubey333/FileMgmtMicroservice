package service

const (
	URI = "mongodb://mongodb:27017" // Use container name to connect otherwise error will be received.
	//URI                   = "mongodb://localhost:27017"
	INTERNAL_SERVER_ERROR = "500"
	DEFAULT_UPLOAD_DIR    = "/tmp"
	LIST_FAIL_MSG         = "Request to list files failed."
	GET_FILE_ERROR        = "Request to get file failed."
	DOWNLOAD_PATH         = "/tmp"
	DOWNLOADING_ERROR     = "Error while downloading file. "
)
