package domain

import "fmt"

type Country struct {
	ID         string
	Name       string
	Politics   *Politics
	Economy    *Economy
	Population *Population
}

func NewCountry(name string, politics *Politics, economy *Economy, population *Population) *Country {
	return &Country{
		Name:       name,
		Politics:   politics,
		Economy:    economy,
		Population: population,
	}
}

func (c *Country) Update() string {
	return fmt.Sprintf("País %s atualizado", c.Name)
}
