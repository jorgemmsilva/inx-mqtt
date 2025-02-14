package main

import (
	"encoding/json"

	inx "github.com/iotaledger/inx/go"
)

// milestoneInfoPayload defines the payload of the milestone latest and confirmed topics
type milestoneInfoPayload struct {
	// The index of the milestone.
	Index uint32 `json:"index"`
	// The unix time of the milestone payload.
	Time uint32 `json:"timestamp"`
	// The ID of the milestone.
	MilestoneID string `json:"milestoneId"`
}

// messageMetadataPayload defines the payload of the message metadata topic
type messageMetadataPayload struct {
	// The hex encoded message ID of the message.
	MessageID string `json:"messageId"`
	// The hex encoded message IDs of the parents the message references.
	Parents []string `json:"parentMessageIds"`
	// Whether the message is solid.
	Solid bool `json:"isSolid"`
	// The milestone index that references this message.
	ReferencedByMilestoneIndex *uint32 `json:"referencedByMilestoneIndex,omitempty"`
	// If this message represents a milestone this is the milestone index
	MilestoneIndex *uint32 `json:"milestoneIndex,omitempty"`
	// The ledger inclusion state of the transaction payload.
	LedgerInclusionState *string `json:"ledgerInclusionState,omitempty"`
	// The reason why this message is marked as conflicting.
	ConflictReason *inx.MessageMetadata_ConflictReason `json:"conflictReason,omitempty"`
	// Whether the message should be promoted.
	ShouldPromote *bool `json:"shouldPromote,omitempty"`
	// Whether the message should be reattached.
	ShouldReattach *bool `json:"shouldReattach,omitempty"`
}

// outputPayload defines the payload of the output topics
type outputPayload struct {
	// The hex encoded message ID of the message.
	MessageID string `json:"messageId"`
	// The hex encoded transaction id from which this output originated.
	TransactionID string `json:"transactionId"`
	// The index of the output.
	OutputIndex uint16 `json:"outputIndex"`
	// Whether this output is spent.
	Spent bool `json:"isSpent"`
	// The milestone index at which this output was spent.
	MilestoneIndexSpent uint32 `json:"milestoneIndexSpent,omitempty"`
	// The milestone timestamp this output was spent.
	MilestoneTimestampSpent uint32 `json:"milestoneTimestampSpent,omitempty"`
	// The transaction this output was spent with.
	TransactionIDSpent string `json:"transactionIdSpent,omitempty"`
	// The milestone index at which this output was booked into the ledger.
	MilestoneIndexBooked uint32 `json:"milestoneIndexBooked"`
	// The milestone timestamp this output was booked in the ledger.
	MilestoneTimestampBooked uint32 `json:"milestoneTimestampBooked"`
	// The ledger index at which this output was available at.
	LedgerIndex uint32 `json:"ledgerIndex"`
	// The output in its serialized form.
	RawOutput *json.RawMessage `json:"output"`
}
