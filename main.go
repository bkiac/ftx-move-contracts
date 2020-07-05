package main

import (
	"fmt"
	"github.com/bclicn/color"
	"github.com/romanornr/ftx-move-contracts/futures"
	"time"
)

func main() {
	expired := futures.GetExpiredFutures()
	MOVEContractsData := expired.GetDailyMOVEContracts()
	contracts := MOVEContractsData.AverageDailyMOVEContractsThisYear()
	averageYearlyExpirationPrice := MOVEContractsData.AverageDailyMOVEContractsThisYear().AverageExpirationPrice

	days := []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday, time.Sunday}
	months := []time.Month{time.January, time.February, time.March, time.April, time.May, time.June, time.July, time.August, time.September, time.October, time.November, time.December}

	for i := len(contracts.Expired) - 1; i >= 0; i-- {
		if contracts.Expired[i].Mark > averageYearlyExpirationPrice {
			fmt.Printf("%s\tGroup: %s\tDay:%s\tExpiration price: "+color.Green("$%.2f\n"), contracts.Expired[i].Description, contracts.Expired[i].Group, contracts.Expired[i].Expiry.Weekday().String(), contracts.Expired[i].Mark)
		} else {
			fmt.Printf("%s\tGroup: %s\tDay:%s\tExpiration price: "+color.Red("$%.2f\n"), contracts.Expired[i].Description, contracts.Expired[i].Group, contracts.Expired[i].Expiry.Weekday().String(), contracts.Expired[i].Mark)
		}
	}
	fmt.Println()

	for i, month := range months {
		move := contracts.AverageMonth(month)
		if move.Expired[1].Expiry.Month() == time.Now().Month() { // TODO figure out why this works and doesn't index out of range
			break // break out of the loop
		}
		if move.AverageExpirationPrice > averageYearlyExpirationPrice {
			fmt.Printf("Average %s %s MOVE Contract expiration price on %s\t"+color.Green("$%2.f\n"), move.Expired[i].Group, move.Expired[i].UnderlyingDescription, month, move.AverageExpirationPrice)
		} else {
			fmt.Printf("Average %s %s MOVE Contract expiration price on %s\t"+color.Red("$%2.f\n"), move.Expired[i].Group, move.Expired[i].UnderlyingDescription, month, move.AverageExpirationPrice)
		}
	}

	// current month, using this method because index out of range [0] with length 0 by using months
	currentMonth := contracts.CurrentAverageMonth()
	if currentMonth.AverageExpirationPrice > averageYearlyExpirationPrice {
		fmt.Printf("Average %s %s MOVE Contract expiration price on %s\t"+color.Green("$%2.f\n"),  currentMonth.Expired[0].Group, currentMonth.Expired[0].UnderlyingDescription, time.Now().Month(), currentMonth.AverageExpirationPrice)
	} else {
		fmt.Printf("Average %s %s MOVE Contract expiration price on %s\t"+color.Red("$%2.f\n"), currentMonth.Expired[0].Group, currentMonth.Expired[0].UnderlyingDescription, time.Now().Month(), currentMonth.AverageExpirationPrice)

	}

	fmt.Println()

	for _, day := range days {
		move := contracts.AverageDayWeek(day)
		if move.AverageExpirationPrice > averageYearlyExpirationPrice {
			fmt.Printf("Average %s MOVE Contract expiration price on %s\t"+color.Green("$%2.f\n"), move.Expired[0].UnderlyingDescription, day, move.AverageExpirationPrice)
		} else {
			fmt.Printf("Average %s MOVE Contract expiration price on %s\t"+color.Red("$%2.f\n"), move.Expired[0].UnderlyingDescription, day, move.AverageExpirationPrice)
		}
	}

	fmt.Printf("\nCurrent average daily FTX MOVE contracts expiration price for %d: "+color.Green("$%2.f\n\n"), time.Now().Year(), contracts.AverageExpirationPrice)
	fmt.Printf("Signup on FTX with this referral link receive a 10%% discount instead of the regular 5%% discount: %s\n", "https://ftx.com/#a=10percentDiscountOnFees")
}
