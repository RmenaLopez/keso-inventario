package service

import (
	"app/model"
	"app/repository"
	"fmt"
)

type CardSearchService struct {
	CardRepo *repository.CardRepository
}

func (c *CardSearchService) GetCardByName(name string) []*model.Card {
	cards := c.CardRepo.GetCardByName(name)
	err := SetCardImagesForCards(cards)
	if err != nil {
		fmt.Println("ERROR GETTING CARD IMAGES")
	}
	return cards
}
