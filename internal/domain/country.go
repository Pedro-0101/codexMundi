package domain

type Country struct {
	ID         string
	Name       string
	Era        *Era
	Politics   *Politics
	Economy    *Economy
	Population *Population
}
