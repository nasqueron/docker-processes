package consoleutilities

import (
	"os"
	"strings"
)

func IsUTF8() bool {
	return strings.Contains(os.Getenv("LC_ALL"), "UTF-8") ||
		strings.Contains(os.Getenv("LANG"), "UTF-8") ||
		strings.Contains(os.Getenv("LC_CTYPE"), "UTF-8")
}
