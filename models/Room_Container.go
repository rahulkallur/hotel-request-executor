package models

// RoomContainer represents a container for a room with related information like
// its allotment, board code, and refundable status.
type RoomContainer struct {
	SortKey       string      `json:"sortKey"`       // A unique key used to sort the room container (e.g., by price or availability).
	Paxes         string      `json:"paxes"`         // Represents the number of people (adults/children) the room can accommodate.
	Allotment     int         `json:"allotment"`     // The number of rooms available in this container.
	BoardCode     string      `json:"boardCode"`     // Code representing the type of board, e.g., "BB" (Bed & Breakfast), "HB" (Half Board).
	Room          RoomDetails `json:"room"`          // Details about the room, including features, rate, and other attributes.
	Refundable    bool        `json:"refundable"`    // Indicates whether the room is refundable or not.
	UsedInMerging bool        `json:"usedInMerging"` // Indicates whether this room container is used in merging processes (e.g., combining offers or availability).
}
