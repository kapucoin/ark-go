package api

import (
	"github.com/asdine/storm/q"
	"github.com/kristjank/ark-go/cmd/model"
	log "github.com/sirupsen/logrus"
)

func getPayments(offset int, network string) ([]model.PaymentRecord, error) {
	var results []model.PaymentRecord

	query := ArkStatsDB.Select(q.Eq("Network", network)).Reverse().Limit(50).Skip(offset)
	err := query.Find(&results)

	if err != nil {
		log.Error("getPayments ", err.Error())
	}

	return results, err
}

func getPaymentsByDelegate(address string, network string) ([]model.PaymentRecord, error) {
	var results []model.PaymentRecord

	query := ArkStatsDB.Select(q.Eq("Network", network), q.Eq("Delegate", address)).Reverse()
	err := query.Find(&results)

	if err != nil {
		log.Error("getPaymentsByDelegate", err.Error())
	}

	return results, err
}

func getStatistics(network string) (map[string]int, error) {
	var results []model.PaymentRecord
	m := map[string]int{}

	query := ArkStatsDB.Select(q.Eq("Network", network))
	err := query.Find(&results)

	if err != nil {
		log.Error("getStatistics", err.Error())
		return m, err
	}

	for _, el := range results {
		m[el.Delegate]++
	}

	return m, err
}
