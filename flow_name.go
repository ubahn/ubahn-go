package ubahn

import (
	"strings"
)

// NewFlowName extracts a flow name string from a file name.
func NewFlowName(file IConversationFile) string {
	parts := strings.Split(file.FileName(), ".")
	if len(parts) > 0 {
		partsWithoutExt := parts[:len(parts)-1]
		if len(partsWithoutExt) > 0 {
			return strings.Join(partsWithoutExt, ".")
		}
	}
	return file.FileName()
}
