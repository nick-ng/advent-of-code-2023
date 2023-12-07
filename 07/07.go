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

func getCardValue(card string) int {
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
			return 11
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

func getHandValue(hand string) int {
	handValue := 0
	cards := strings.Split(hand, "")

	cardCount := map[string]int{}

	for i, card := range cards {
		cardCount[card] += 1

		power := 8 - i*2
		multiplier := math.Pow10(power)

		cardValue := getCardValue(card)

		cardPoints := cardValue * int(multiplier)

		handValue += cardPoints
	}

	cardCounts := []int{}

	for _, count := range cardCount {
		cardCounts = append(cardCounts, count)
	}

	slices.SortFunc(cardCounts, func(a, b int) int {
		return cmp.Compare(b, a)
	})

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

	// ba987654321
	// H0011223344
	handValue += handPoints * 1e10

	fmt.Println(hand, cardCounts, handPoints)

	return handValue
}

type handBid struct {
	Hand      string
	HandValue int
	Bid       int
}

func Run() {
	fmt.Println("day 7")

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

		handValue := getHandValue(hand)

		hands = append(hands, handBid{
			Hand:      hand,
			HandValue: handValue,
			Bid:       bid,
		})

		fmt.Println(hand, handValue)
	}

	slices.SortFunc(hands, func(a, b handBid) int {
		return cmp.Compare(a.HandValue, b.HandValue)
	})

	total1 := 0
	for rank, hand := range hands {
		score := (rank + 1) * hand.Bid

		total1 += score
	}

	fmt.Println("total1:", total1)
}
