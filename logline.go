package main

import (
	"bytes"
	"encoding/json"
	"strconv"
	"time"
)

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
	CreatedTimestamp  time.Time `json:"createdTs"`
	SettlementDate    string    `json:"settlementDate,omitempty"`

	Amount float64 `json:"amount"`
}

func getTimestamp(cts, ctz []byte) string {
	v := make([]byte, 0, len(cts)+len(ctz)+1)
	v = append(v, cts...)
	v = append(v, ' ')
	v = append(v, ctz...)
	return string(v)
}
