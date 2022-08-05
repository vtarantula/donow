package config

/*
Linux error codes are from 1 to 131
Windows error codes are from 1 to 123 and 1450
Keep all the error codes at and above 256 for safety
Use these codes for logging purposes only. Linux programs don't honor any integers > 128. They will be treated as 1
*/
const (
	ERR_UNKNOWN     uint = iota + 256
	ERR_PLATFORM    uint = iota + ERR_UNKNOWN
	ERR_APP_GENERIC uint = 300
)

const (
	EXIT_SUCCESS_CODE int = iota
	EXIT_ERROR_CODE
)
