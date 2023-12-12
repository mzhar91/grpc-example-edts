package message

import "sync"

const (
	RequiredField       = 11
	IncorrectFormat     = 12
	UploadFailed        = 13
	IncorrectFormat2    = 15
	UniqueField         = 16
	ExceedMaxCharacter  = 17
	IncorrectFormat3    = 18
	AddressTypeNotFound = 19
	InvalidEmailFormat  = 20
	UnprocessableEntity = 50
	AccessNotGranted    = 51
	ForbiddenAccess     = 52
	ExceedMaxValue      = 53
)

var once sync.Once
var errorDictionary map[int]string

func initErrorMessageMap() {
	errorDictionary = make(map[int]string)
	errorDictionary[RequiredField] = "%s is required field"
	errorDictionary[IncorrectFormat] = "Incorrect format"
	errorDictionary[UploadFailed] = "Upload failed"
	errorDictionary[IncorrectFormat2] = "Alphanumeric, dash, and underscore only."
	errorDictionary[UniqueField] = "Product with this %s " +
		"already exists in the system. Enter a different %s to continue."
	errorDictionary[ExceedMaxCharacter] = "%s maximum character allowed is %s"
	errorDictionary[ExceedMaxValue] = "%s maximum value allowed is %s"
	errorDictionary[UnprocessableEntity] = "UnprocessableEntity"
	errorDictionary[IncorrectFormat3] = "Alphabetic only"
	errorDictionary[AddressTypeNotFound] = "Address type not found"
	errorDictionary[InvalidEmailFormat] = "Invalid email format"
	errorDictionary[AccessNotGranted] = "Access not granted"
	errorDictionary[ForbiddenAccess] = "Forbidden access"
}

// InitErrorMessage return an initiated error to http status map
func InitErrorMessage() map[int]string {
	once.Do(func() {
		initErrorMessageMap()
	})

	return errorDictionary
}
