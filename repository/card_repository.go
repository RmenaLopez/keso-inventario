package repository

import (
	"app/model"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var cardRepo *CardRepository

type CardRepository struct {
	DB *sql.DB
}

func NewCardRepository() *CardRepository {
	return cardRepo.GetRepository()
}

func (cr *CardRepository) GetRepository() *CardRepository {
	if cardRepo == nil {
		cardRepo = &CardRepository{}
		db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/carddb")
		if err != nil {
			panic(err.Error())
		}
		cardRepo.DB = db
		cr = cardRepo
		fmt.Println("Success!")
	}
	return cr
}

func (cr *CardRepository) GetCardByName(name string) []*model.Card {
	query := "select cards.name, cards.text, cards.uuid, cardIdentifiers.scryfallId from cards " +
		"inner join cardIdentifiers on cards.uuid=cardIdentifiers.uuid where UPPER(name) like UPPER(?)"

	rows, err := cr.DB.Query(query, name)
	if err != nil {
		panic(err)
	}

	cards := make([]*model.Card, 0)

	for rows.Next() {
		var card model.Card
		err = rows.Scan(
			&card.Name,
			&card.Oracle,
			&card.Uuid,
			&card.ScryfallId,
		)

		if err != nil {
			panic(err)
		}
		cards = append(cards, &card)
	}
	return cards
}
