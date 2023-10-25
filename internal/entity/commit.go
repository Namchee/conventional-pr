package entity

// Commit that doesn't follow the conventional commit style
type Commit struct {
	// Commit hash
	Hash string
	// Commit message
	Message string
	// Commit verification status
	Verified bool
}
