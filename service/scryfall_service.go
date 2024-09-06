package service

import (
	"app/model"
	"context"
	"log"

	"github.com/BlueMonday/go-scryfall"
)

func SetCardImagesForCards(cards []*model.Card) error {
	client, err := scryfall.NewClient()
	if err != nil {
		log.Default().Println("failed to initialize scryfall client")
		return err
	}

	idMap := make(map[string][]*model.Card, 0)

	for _, card := range cards {
		_, ok := idMap[card.ScryfallId]
		if !ok {
			idMap[card.ScryfallId] = []*model.Card{card}
		} else {
			idMap[card.ScryfallId] = append(idMap[card.ScryfallId], card)
		}

	}

	identifiers := make([]scryfall.CardIdentifier, 0)
	for k := range idMap {
		identifier := scryfall.CardIdentifier{ID: k}
		identifiers = append(identifiers, identifier)
	}

	response, err := client.GetCardsByIdentifiers(context.Background(), identifiers)
	if err != nil {
		log.Default().Println("failed to get response from scryfall client")
		return err
	}

	for _, i := range response.Data {
		for _, card := range idMap[i.ID] {
			card.ImageUrl = i.ImageURIs.Small
		}

	}
	return nil
}
