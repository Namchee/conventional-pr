package entity

type Actor struct {
	Login string
	Type string
}

type IssueReference struct {
	Meta
	Number int
}

type PullRequest struct {
	Number int
	Title string
	Body string
	Branch string
	IsDraft bool
	Changes int
	Author Actor

	Repository Meta
}

type Issue struct {

}