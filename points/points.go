package points

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"receipt-processor/models"
)

func CalculatePoints(receipt models.Receipt) int {
	points := 0

	alnumRegex := regexp.MustCompile(`[a-zA-Z0-9]`)
	points += len(alnumRegex.FindAllString(receipt.Retailer, -1))

	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == math.Floor(total) {
		points += 50
	}

	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	points += (len(receipt.Items) / 2) * 5

	for _, item := range receipt.Items {
		descLen := len(strings.TrimSpace(item.ShortDescription))
		if descLen%3 == 0 {
			itemPrice, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(itemPrice * 0.2))
		}
	}

	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if purchaseDate.Day()%2 == 1 {
		points += 6
	}

	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() == 14 {
		points += 10
	}

	return points
}
