package day07

import (
	"advent-of-code-2023/utils"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getCardValueJokers(card string) int {
	switch card {
	case "A":
		{
			return 14
		}
	case "K":
		{
			return 13
		}
	case "Q":
		{
			return 12
		}
	case "J":
		{
			return 1
		}
	case "T":
		{
			return 10
		}
	default:
		{
			value, err := strconv.Atoi(card)

			if err != nil {
				fmt.Println("cannot convert card to int", card, err)
				os.Exit(1)
			}

			return value
		}
	}
}

func getHandValueJokers(hand string) int {
	handValue := 0
	cards := strings.Split(hand, "")

	cardCount := map[string]int{}

	jokers := 0
	for i, card := range cards {

		power := 8 - i*2
		multiplier := math.Pow10(power)

		cardValue := getCardValueJokers(card)

		cardPoints := cardValue * int(multiplier)

		handValue += cardPoints

		if card == "J" {
			jokers++
		} else {
			cardCount[card]++
		}
	}

	cardCounts := []int{}

	for _, count := range cardCount {
		cardCounts = append(cardCounts, count)
	}

	slices.SortFunc(cardCounts, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	if len(cardCounts) == 0 {
		cardCounts = append(cardCounts, jokers)
	} else {
		cardCounts[0] += jokers
	}

	handPoints := 0

	if len(cardCounts) == 1 {
		// 5 of a kind
		handPoints = 6
	} else if cardCounts[0] == 4 {
		// 4 of a kind
		handPoints = 5
	} else if len(cardCounts) == 2 {
		// full house
		handPoints = 4
	} else if cardCounts[0] == 3 {
		// 3 of a kind
		handPoints = 3
	} else if cardCounts[1] == 2 {
		// two pair
		handPoints = 2
	} else if len(cardCounts) == 4 {
		// one pair
		handPoints = 1
	}

	handValue += handPoints * 1e10

	return handValue
}

func RunJokers() {
	rawData := utils.ReadFile("07/data.txt")

	rawHands := strings.Split(rawData, "\n")
	hands := []handBid{}

	for _, temp := range rawHands {
		if len(temp) == 0 {
			continue
		}

		temp2 := strings.Split(temp, " ")
		hand := temp2[0]
		bid, err := strconv.Atoi(temp2[1])

		if err != nil {
			fmt.Println("cannot convert bid to int", temp2[1], err)
		}

		handValue := getHandValueJokers(hand)

		hands = append(hands, handBid{
			Hand:      hand,
			HandValue: handValue,
			Bid:       bid,
		})
	}

	slices.SortFunc(hands, func(a, b handBid) int {
		return cmp.Compare(a.HandValue, b.HandValue)
	})

	total2 := 0
	for rank, hand := range hands {
		score := (rank + 1) * hand.Bid

		total2 += score
	}

	fmt.Println("total2:", total2)
}
