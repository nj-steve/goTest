package main

import (
	"fmt"
	"strconv"
	"time"
)

const (
	UnitAttoFil = "AttoFIL"
	UnitNanoFil = "NanoFIL"
	UnitFil     = "FIL"
)

const (
	MaxValueSizeOfAttoFil = 1000000000
	MaxPeerCount          = 432
	FilToAttoFil          = 1000000000000000000
	FilOfGasPrice         = 1000000000
	TransferToNano        = 100000000000000
	FilUseAsBigInt        = 100000000
	FilecoinCountTotal    = 2000000000
	ValueOneMillion       = 1000000
	UnitMillion           = "million"
)

func PriceToString(price uint64) string {
	var priceResult string
	{
		gasPrice := strconv.FormatUint(price, 10)
		fmt.Println("llllll = ", len(gasPrice))
		if len(gasPrice) <= 6 {
			priceResult = gasPrice + " " + UnitAttoFil
		} else if len(gasPrice) <= 15 {
			priceFloat := float64(price) / float64(FilOfGasPrice)
			fmt.Println("priceFloat = ", priceFloat)
			priceStr := strconv.FormatFloat(priceFloat, 'f', -1, 64)
			priceResult = priceStr + " " + UnitNanoFil
		} else {
			priceFloat := float64(price) / float64(FilToAttoFil)
			priceStr := strconv.FormatFloat(priceFloat, 'f', -1, 64)
			priceResult = priceStr + " " + UnitFil
		}

	}
	return priceResult
}

func main() {
	var fee uint64
	var gas uint64
	gas = 1500
	var limit uint64
	limit = 470335
	fee = gas * limit
	fmt.Println("unit64 feeeeee = %v", fee)
	fee_str := PriceToString(fee)
	fmt.Println("feeeeee = %v", fee_str)
	heightList := getHeightAndTimestampList()
	fmt.Println("lllllllllllllllll=", len(heightList))
	for _, v := range heightList {
		fmt.Println("height= %v   time=%v  ", v.height, v.timestamp)
	}
	var nowTimeSS []string
	nowTimeSS = append(nowTimeSS, "AAA")
	nowTimeSS = append(nowTimeSS, "BBB")
	nowTimeSS = append(nowTimeSS, "CCC")
	var timeSS []string
	timeSS = append(timeSS, "ZZZ")
	nowTimeSS = append(timeSS, nowTimeSS...)
	for _, v := range nowTimeSS {
		fmt.Println("nnnnnn", v)
	}
}

type HeightList struct {
	height    int64 `json:"height"`
	timestamp int64 `json:"timestamp"`
}

func getHeightAndTimestampList() []HeightList {
	var initHeigt int64
	var initTimestamp int64
	var heightlist []HeightList
	tn := time.Now().Unix()
	initHeigt = 46
	initTimestamp = 1597338000
	heightlist = append(heightlist, HeightList{height: initHeigt, timestamp: initTimestamp})
	for {
		if initTimestamp < tn {
			initHeigt += 60
			initTimestamp += 1800
			heightlist = append(heightlist, HeightList{height: initHeigt, timestamp: initTimestamp})
		} else {
			break
		}
	}
	return heightlist
}
