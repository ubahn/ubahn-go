package ubahn

import (
	"fmt"
	"time"
)

// HistoryItem describes historical piece of a conversation, where you can see what was the input, the output,
// the flow and when it was recorded.
type HistoryItem struct {
	inputName  string
	outputName string
	flowName   string
	timestamp  time.Time
}

// NewHistoryItem create a new instance of a history item.
func NewHistoryItem(input IInput, nextOutput IOutput, conv IConversation) IHistoryItem {
	return &HistoryItem{
		inputName:  input.Name(),
		outputName: nextOutput.Name(),
		flowName:   conv.FlowName(),
		timestamp:  time.Now().UTC()}
}

// InputName returns input name.
func (item *HistoryItem) InputName() string {
	return item.inputName
}

// OutputName returns output name.
func (item *HistoryItem) OutputName() string {
	return item.outputName
}

// FlowName returns flow name.
func (item *HistoryItem) FlowName() string {
	return item.flowName
}

// Timestamp returns timestamp.
func (item *HistoryItem) Timestamp() time.Time {
	return item.timestamp
}

// String returns string view of the history items
func (item *HistoryItem) String() string {
	return fmt.Sprintf(
		"[%s] %s -> %s",
		item.Timestamp().String(),
		item.InputName(),
		item.OutputName())
}
