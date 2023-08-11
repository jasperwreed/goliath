package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	ConversationID string `bson:"conversation_id" json:"conversation_id"` // This will correspond to the ID from the relational DB.
	SenderID       string `bson:"sender_id" json:"sender_id"` // This will correspond to the User ID from the relational DB.
	Content        string `bson:"content" json:"content"`
	Timestamp      time.Time `bson:"timestamp" json:"timestamp"`
	Attachments    []string `bson:"attachments" json:"attachments"`
	ReadReceipts   map[string]time.Time `bson:"read_receipts" json:"read_receipts"` // The keys here would be User IDs.
}


// for redis down the road
// type Message struct {
// 	ID             uint      `json:"id"`
// 	ConversationID uint      `json:"conversation_id"`
// 	SenderID       uint      `json:"sender_id"`
// 	Content        string    `json:"content"`
// 	Timestamp      time.Time `json:"timestamp"`
// 	Attachments    []string  `json:"attachments"`
// 	ReadReceipts   map[uint]time.Time `json:"read_receipts"`
// }
