package entity

type Actor struct {
	Name string
	Type string
}

type PullRequest struct {
	Title string
	Body string
	Branch string
	IsDraft bool
	Changes int
	Author Actor
}

type Issue struct {

}