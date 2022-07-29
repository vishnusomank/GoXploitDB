package utils

import "github.com/jinzhu/gorm"

//constants for Request Body
const (
	InvalidReqBody  = "Invalid Request Body"
	RequiredReqBody = "Required Request Body Missing"
	WrongReqBody    = "Wrong Request Body : "
)

//constants for Error Messages
const (
	Err400        = "err-400"
	Err500        = "err-500"
	ErrInAddImage = "Error in Add Image"
)

//constants for URL Path
const (
	CVE          = "cve/:cve"
	TYPE         = "type/:type"
	PLATFORM     = "platform/:platform"
	All_PLATFORM = "platforms"
	ALL_TYPE     = "types"
)

const (
	ACTIVE = "Active"
	DELETE = "Delete"
)

var XPLOITDB *gorm.DB

var CURRENT_DIR, GIT_DIR string
