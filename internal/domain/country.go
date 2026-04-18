package domain

type Country struct {
	ID         string
	Name       string
	Politics   *Politics
	Economy    *Economy
	Population *Population
}
