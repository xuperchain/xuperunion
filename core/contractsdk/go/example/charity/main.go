package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/xuperchain/xuperchain/core/contractsdk/go/code"
	"github.com/xuperchain/xuperchain/core/contractsdk/go/driver"
	"github.com/xuperchain/xuperchain/core/contractsdk/go/utils"
	"math/big"
	"strings"
)

type charityDonation struct {
}

const (
	USER_DONATE    = "UserDonate_"
	ALL_DONATE     = "AllDonate_"
	ALL_COST       = "AllCost_"
	TOTAL_RECEIVED = "TotalDonates"
	TOTAL_COSTS    = "TotalCosts"
	BALANCE        = "Balance"
	DONATE_COUNT   = "DonateCount"
	COST_COUNT     = "CostCount"
	ADMIN          = "admin"
	MAX_LIMIT      = 100
)

var (
	ErrLimitExceeded = errors.New("limit exceeded")
)

func (cd *charityDonation) Initialize(ctx code.Context) code.Response {
	admin, err := ctx.GetObject([]byte(ADMIN))
	if err != nil { // 这里的err 什么情况，不存在会是err么？
		return code.Error(err)
	}
	if err := ctx.PutObject([]byte(ADMIN), admin); err != nil {
		return code.Error(err)
	}
	if err := ctx.PutObject([]byte(TOTAL_RECEIVED), []byte("0")); err != nil {
		return code.Error(err)
	}
	if err := ctx.PutObject([]byte(TOTAL_COSTS), []byte("0")); err != nil {
		return code.Error(err)
	}
	if err := ctx.PutObject([]byte(BALANCE), []byte("0")); err != nil {
		return code.Error(err)
	}
	if err := ctx.PutObject([]byte(DONATE_COUNT), []byte("0")); err != nil {
		return code.Error(err)
	}
	if err := ctx.PutObject([]byte(COST_COUNT), []byte("0")); err != nil {
		return code.Error(err)
	}
	return code.OK(nil)
}

func (cd *charityDonation) Donate(ctx code.Context) code.Response {
	err := cd.checkPermission(ctx)
	if err != nil {
		return code.Error(err)
	}
	args := struct {
		Donor     string     `json:"donor",required:"true"`
		Amount    *big.Float `json:"amount",required:"true"`
		Timestamp string     `json:"timestamp",required:"true"`
		Comments  string     `json:"comments",required:"comment"`
	}{}
	if err := utils.Validate(ctx.Args(), &args); err != nil {
		return code.Error(err)
	}
	donateCountByte, err := ctx.GetObject([]byte(DONATE_COUNT))
	if err != nil {
		return code.Error(err)
	}
	donateCount, _ := big.NewFloat(0).SetString(string(donateCountByte)) // TODO
	donateCount = donateCount.Add(donateCount, big.NewFloat(1))

	totalReceivedByte, err := ctx.GetObject([]byte(TOTAL_RECEIVED))
	if err != nil {
		return code.Error(err)
	}
	totalReceived, _ := big.NewFloat(0).SetString(string(totalReceivedByte))
	balanceByte, err := ctx.GetObject([]byte(BALANCE))
	if err != nil {
		return code.Error(err)
	}
	balance, _ := big.NewFloat(0).SetString(string(balanceByte))

	totalReceived = totalReceived.Add(totalReceived, args.Amount)

	balance = balance.Add(balance, args.Amount)

	donateID := fmt.Sprintf("%06d", 12) //TODO
	userDonateKey := USER_DONATE + args.Donor + donateID
	allDonateKey := ALL_DONATE + donateID
	donateDetail := fmt.Sprintf("donor=%s,amount=%s,timestamp=%s,commnets=%s", args.Donor, args.Amount, args.Timestamp, args.Comments)

	if err := ctx.PutObject([]byte(userDonateKey), []byte(donateDetail)); err != nil {
		return code.Error(err)
	}
	if err := ctx.PutObject([]byte(allDonateKey), []byte(donateDetail)); err != nil {
		return code.Error(err)
	}
	if err := ctx.PutObject([]byte(DONATE_COUNT), []byte(donateCount.String())); err != nil {
		return code.Error(err)
	}
	if err := ctx.PutObject([]byte(TOTAL_RECEIVED), []byte(totalReceived.String())); err != nil {
		return code.Error(err)
	}
	if err := ctx.PutObject([]byte(BALANCE), []byte(balance.String())); err != nil {
		return code.Error(err)
	}
	return code.OK([]byte(donateID))

}

func (cd *charityDonation) Cost(ctx code.Context) code.Response {
	if err := cd.checkPermission(ctx); err != nil {
		return code.Error(err)
	}
	args := struct {
		To        string     `json:"to",required:"true"`
		Amount    *big.Float `json:"amount",required:"true"`
		Timestamp string     `json:"timestamp",required:"true"`
		Comments  string     `json:"comments",required:"comment"`
	}{}
	if err := utils.Validate(ctx.Args(), &args); err != nil {
		return code.Error(err)
	}
	costCountByte, err := ctx.GetObject([]byte(COST_COUNT))
	if err != nil {
		return code.Error(err)
	}
	totalCostsByte, err := ctx.GetObject([]byte(TOTAL_COSTS))
	if err != nil {
		return code.Error(err)
	}
	balanceByte, err := ctx.GetObject([]byte(BALANCE))
	if err != nil {
		return code.Error(err)
	}
	costCount, _ := big.NewInt(0).SetString(string(costCountByte), 10)
	totalCost, _ := big.NewFloat(0).SetString(string(totalCostsByte))
	balance, _ := big.NewFloat(0).SetString(string(balanceByte))

	if balance.Cmp(args.Amount) < 0 {

		return code.Error(utils.ErrBalanceLow)
	}

	totalCost = totalCost.Add(totalCost, args.Amount)
	balance = balance.Sub(balance, args.Amount)
	costCount = costCount.Add(costCount, big.NewInt(0))
	costId := "0" //TODO @fengjin
	allCostKey := ALL_COST + costId
	costDetails := fmt.Sprintf(
		"to=%s,amount=%s,timestamp=%s,comments=%s",
		args.To, args.Amount.String(),
		args.Timestamp,
		args.Comments)

	if err := ctx.PutObject([]byte(allCostKey), []byte(costDetails)); err != nil {
		return code.Error(err)
	}
	if err := ctx.PutObject([]byte(COST_COUNT), []byte(costCount.String())); err != nil {
		return code.Error(err)
	}
	if err := ctx.PutObject([]byte(TOTAL_COSTS), []byte(totalCost.String())); err != nil {
		return code.Error(err)
	}

	if err := ctx.PutObject([]byte(BALANCE), []byte(balance.String())); err != nil {
		return code.Error(err)
	}
	return code.OK([]byte(costId))
}

func (cd *charityDonation) Statistics(ctx code.Context) code.Response {
	buf := bytes.NewBuffer([]byte{})
	totalReceived, err := ctx.GetObject([]byte(TOTAL_RECEIVED))
	if err != nil {
		return code.Error(ErrLimitExceeded)
	}
	totalCost, err := ctx.GetObject([]byte(TOTAL_COSTS))
	if err != nil {
		return code.Error(err)
	}
	balance, err := ctx.GetObject([]byte(BALANCE))
	if err != nil {
		return code.Error(err)
	}
	// TODO @chenfengjin buffer size limit
	buf.WriteString("totalDonates=")
	buf.Write(totalReceived)
	buf.WriteString("totalCost=")
	buf.Write(totalCost)
	buf.WriteString("fundBalance=")
	buf.Write(balance)
	return code.OK(buf.Bytes())
}

func (cd *charityDonation) QueryDoner(ctx code.Context) code.Response {
	args := struct {
		Donar string `json:"donar",required:"true"`
	}{}
	if err := utils.Validate(ctx.Args(), &args); err != nil {
		return code.Error(err)
	}
	start := USER_DONATE + args.Donar + "%"
	end := start + "~"
	iter := ctx.NewIterator([]byte(start), []byte(end))
	donateCount := big.NewInt(0)
	defer iter.Close()
	builder := strings.Builder{}
	for iter.Next() {
		donateCount = donateCount.Add(donateCount, big.NewInt(1))
		donateId := iter.Key()[len(start):]
		builder.WriteString("id=")
		builder.Write(donateId)
		builder.WriteString(",")
		builder.Write(iter.Value())
		builder.WriteString("\n")
	}
	return code.OK([]byte("total donate count:" +
		donateCount.String() + "\n" + builder.String()))
}

func (cd *charityDonation) QueryDonates(ctx code.Context) code.Response {
	args := struct {
		StartId string   `json:"startid",required:"true"`
		Limit   *big.Int `json:"limit,"required:"true"`
	}{}
	if err := utils.Validate(ctx.Args(), &args); err != nil {
		return code.Error(err)
	}
	if args.Limit.Cmp(big.NewInt(MAX_LIMIT)) > 0 {
		return code.Error(ErrLimitExceeded)
	}
	donateKey := ALL_DONATE + args.StartId
	start := donateKey
	end := donateKey + "~"
	iter := ctx.NewIterator([]byte(start), []byte(end))
	defer iter.Close()
	selected := int64(0)
	builder := strings.Builder{}
	for iter.Next() {

		if selected < args.Limit.Int64() {
			break
		}
		selected += 1
		donateID := iter.Key()[len([]byte(ALL_DONATE)):]
		builder.WriteString("id=")
		builder.Write(donateID)
		builder.WriteString(",")
		builder.Write(iter.Value())
		builder.WriteString("\n")
	}
	return code.OK([]byte(builder.String()))
}

func (cd *charityDonation) QueryCosts(ctx code.Context) code.Response {
	args := struct {
		StartId string  `json:"startid",required:"true"`
		Limit   big.Int `json:"limit,"required:"true"`
	}{}
	if err := utils.Validate(ctx.Args(), &args); err != nil {
		return code.Error(err)
	}
	if args.Limit.Cmp(big.NewInt(MAX_LIMIT)) <= 0 { //TODO @fengjin
		return code.Error(ErrLimitExceeded)
	}
	costKey := ALL_COST + args.StartId
	start := costKey
	end := costKey + "~"
	iter := ctx.NewIterator([]byte(start), []byte(end))
	defer iter.Close()
	selected := int64(0)
	builder := strings.Builder{}
	for iter.Next() {
		if selected > args.Limit.Int64() { // TODO 边界
			break
		}
		selected += 1
		costId := iter.Key()[len([]byte(ALL_COST)):]
		builder.WriteString("id=")
		builder.Write(costId)
		builder.Write(iter.Value())
		builder.WriteString("\n")
	}
	if err := iter.Error(); err != nil {
		return code.Error(err)
	}
	return code.OK([]byte(builder.String()))

}

func (cd *charityDonation) checkPermission(ctx code.Context) error {
	caller := ctx.Initiator()
	if caller == "" {
		return utils.ErrMissingCaller
	}
	if caller != ADMIN {
		return utils.ErrPermissionDenied
	}
	return nil
}

func main() {
	driver.Serve(new(charityDonation))
}
