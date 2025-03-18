package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Role represents the role schema in MongoDB
type Role struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty"`       // Auto-generated MongoDB ObjectID
	Name      string              `bson:"name" json:"name"`    // Name of the role, required
	Deleted   bool                `bson:"deleted,omitempty"`   // Soft delete flag
	DeletedAt *time.Time          `bson:"deletedAt,omitempty"` // Timestamp when deleted
	DeletedBy *primitive.ObjectID `bson:"deletedBy,omitempty"` // User who deleted it
	CreatedAt time.Time           `bson:"createdAt,omitempty"` // Timestamp when created
	UpdatedAt time.Time           `bson:"updatedAt,omitempty"` // Timestamp when updated
}
