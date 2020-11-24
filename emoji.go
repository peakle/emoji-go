package emoji

import "regexp"

var regex = regexp.MustCompile(`/([0-9#][\x{20E3}])|` +
	`[\x{00ae}\x{00a9}\x{203C}\x{2047}\x{2048}\x{2049}\x{3030}\x{303D}\x{2139}\x{2122}\x{3297}\x{3299}]|` +
	`[\x{2190}-\x{21FF}]|[\x{FE00}-\x{FEFF}]|` +
	`[\x{2300}-\x{23FF}]|[\x{FE00}-\x{FEFF}]|` +
	`[\x{2460}-\x{24FF}]|[\x{FE00}-\x{FEFF}]|` +
	`[\x{25A0}-\x{25FF}]|[\x{FE00}-\x{FEFF}]|` +
	`[\x{2600}-\x{27BF}]|[\x{FE00}-\x{FEFF}]|` +
	`[\x{2900}-\x{297F}]|[\x{FE00}-\x{FEFF}]|` +
	`[\x{2B00}-\x{2BF0}]|[\x{FE00}-\x{FEFF}]|` +
	`[\x{1F000}-\x{1F6FF}]|[\x{FE00}-\x{FEFF}]|` +
	`[\x{1F600}-\x{1F64F}]|` +
	`[\x{1F680}-\x{1F6FF}]|` +
	`[\x{1F900}-\x{1F9FF}]|` +
	`[\x{1F1E0}-\x{1F1FF}]|` +
	`[\x{1D100}-\x{1D1FF}]|` +
	`[\x{2600}-\x{26FF}]|` +
	`[\x{2700}-\x{27BF}]|` +
	`[\x{10000}-\x{E01EF}]`)

var (
	emptyByteSlice = []byte("")
	emptyString    = ""
)

// RemoveString - removes all emojies from string
func RemoveString(text string) string {
	return regex.ReplaceAllString(text, emptyString)
}

// Remove - removes all emojies from byte slice
func Remove(text []byte) []byte {
	return regex.ReplaceAll(text, emptyByteSlice)
}
