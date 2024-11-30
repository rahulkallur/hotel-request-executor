package models

// RoomListContainer represents a container that holds a list of RoomContainers,
// each associated with a specific room type, board type, and room code.
type RoomListContainer struct {
	RoomTypeCode   string          // Represents the room type code (e.g., "single", "double").
	BoardCode      string          // Represents the board type code (e.g., "BB" for Bed & Breakfast).
	RoomCode       string          // Represents the unique room code.
	RoomContainers []RoomContainer // A list of RoomContainers associated with the room type.
}

// AddRoomContainer creates a new RoomListContainer and initializes it with a given board code.
// It returns a pointer to the newly created RoomListContainer.
func AddRoomContainer(boardCode string) *RoomListContainer {
	return &RoomListContainer{
		BoardCode:      boardCode,                // Set the board code for the container.
		RoomContainers: make([]RoomContainer, 0), // Initialize the RoomContainers slice as empty.
	}
}

// AddRoomContainer adds a new RoomContainer to the RoomListContainer's RoomContainers list.
// This method appends the provided RoomContainer to the existing list.
func (rlc *RoomListContainer) AddRoomContainer(roomContainer RoomContainer) {
	rlc.RoomContainers = append(rlc.RoomContainers, roomContainer) // Add the roomContainer to the slice.
}
