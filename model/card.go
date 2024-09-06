package model

type CardGrading int

const (
	NearMint CardGrading = iota
	LightlyPlayed
	ModeratelyPlayed
	HeavilyPlayed
	Damaged
)

type Card struct {
	Uuid       string
	Oracle     string
	Name       string
	ImageUrl   string
	ScryfallId string
	Stock      int
}

func NewCard(name, uuid, oracle string, price float64) Card {
	return Card{Name: name, Uuid: uuid, Oracle: oracle}
}

func (c *Card) GetID() string {
	return c.Uuid
}

func (c *Card) GetDisplayName() string {
	return c.Name
}

func (c *Card) GetDescription() string {
	return c.Oracle
}

func (c *Card) GetImageUrl() string {
	return c.ImageUrl
}

func (c *Card) GetUuid() string {
	return c.Uuid
}

func (c *Card) GetStock() int {
	return c.Stock
}
