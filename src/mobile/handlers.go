package mobile

/*
These types are exported and need to be implemented and used by the mobile
application.
*/

// CommitHandler handles retreival of the state hash.
// What the purpose of state hash retrieval for the client application?
type CommitHandler interface {
	OnCommit([]byte) []byte
}

// Handles mobile app mobile app exceptions.
type ExceptionHandler interface {
	OnException(string)
}
