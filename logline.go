package main

import (
	"bytes"
	"encoding/json"
	"strconv"
	"time"
)

// Format provided by Braintree, used to parse to time.Time
const tsFmt = `01/02/2006 15:04:05 MST`

func newLogLine(b []byte) (l *LogLine, err error) {
	s := bytes.Split(b, []byte{','})
	l = &LogLine{
		TransactionId:     string(s[col_transactionID]),
		TransactionType:   string(s[col_transactionType]),
		TransactionStatus: string(s[col_transactionStatus]),
		AuthorizationCode: string(s[col_authorizationCode]),
		SettlementDate:    string(s[col_settlementDate]),
	}

	if l.Amount, err = strconv.ParseFloat(string(s[col_amountAuthorized]), 64); err != nil {
		return nil, err
	}

	if l.CreatedTimestamp, err = time.Parse(tsFmt, getTimestamp(
		s[col_createdDatetime],
		s[col_createdTimezone],
	)); err != nil {
		return nil, err
	}

	if b, err = json.Marshal(l); err != nil {
		return nil, err
	}

	return l, nil
}

type LogLine struct {
	TransactionId     string `json:"txnId"`
	TransactionType   string `json:"txnType"`
	TransactionStatus string `json:"txnStatus"`

	AuthorizationCode string    `json:"authCode,omitempty"`
	SettlementDate    string    `json:"settlementDate,omitempty"`
	CreatedTimestamp  time.Time `json:"createdTs"`

	Amount float64 `json:"amount"`
}

// Returns the concatination result of CreatedTimestamp, " ", and CreatedTimezone
func getTimestamp(cts, ctz []byte) string {
	v := make([]byte, 0, len(cts)+len(ctz)+1)
	v = append(v, cts...)
	v = append(v, ' ')
	v = append(v, ctz...)
	return string(v)
}
