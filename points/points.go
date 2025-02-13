package points

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"receipt-processor/models"
)

// returns: integer value of the points tally for a given receipt
func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// 1 point for every alphanumeric character
	alnumRegex := regexp.MustCompile(`[a-zA-Z0-9]`)
	points += len(alnumRegex.FindAllString(receipt.Retailer, -1))

	// 50 points for round total with no cents
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == math.Floor(total) {
		points += 50
	}

	// 25 points if total is a multiple of 0.25
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// 5 points for every 2 items
	points += (len(receipt.Items) / 2) * 5

	// for each item
	for _, item := range receipt.Items {
		descLen := len(strings.TrimSpace(item.ShortDescription))

		// if trimmed description length is a multiple of 3
		if descLen%3 == 0 {
			itemPrice, _ := strconv.ParseFloat(item.Price, 64)

			// multiply the price by 0.2 and round up to nearest integer
			// add to points tally
			points += int(math.Ceil(itemPrice * 0.2))
		}
	}

	// 6 points if the purchase date is odd
	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if purchaseDate.Day()%2 == 1 {
		points += 6
	}

	// 10 points if purchase time is between 2:00pm and 4:00pm
	// I have understood this wording as the accepted range being 2:01pm to 3:59pm
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() == 14 && purchaseTime.Minute() > 0 || purchaseTime.Hour() == 15 {
		points += 10
	}

	return points
}
