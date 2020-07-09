package faspay

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func GetPaymentChannelSignature() string {
	return getSha1(getMD5Hash(FaspayConfig.MerChantUser + FaspayConfig.MerchantPassword))
}

func GetPaymentSignature(billNo string) string {
	return getSha1(getMD5Hash(FaspayConfig.MerChantUser + FaspayConfig.MerchantPassword + billNo))
}

func GetPaymentNotificationSignature(billNo, paymentStatusCode string) string {
	return getSha1(getMD5Hash(FaspayConfig.MerChantUser + FaspayConfig.MerchantPassword + billNo + paymentStatusCode))
}

func GetCheckVaStaticSignature(vaNumber string) string {
	return getSha1(getMD5Hash(FaspayConfig.MerChantUser + FaspayConfig.MerchantPassword + vaNumber))
}

func GetPaymentWithCardSignature(merchantTransID, amount string) string {
	// ##merchantid##txnpassword##trainid##amount##0##
	signFormat := fmt.Sprintf("##%s##%s##%s##%s##0##", FaspayConfig.MerchandID, FaspayConfig.TxnPassword, merchantTransID, amount)
	return getSha1(signFormat)
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func getSha1(str string) (signature string) {
	h := sha1.New()
	h.Write([]byte(str))
	signature = hex.EncodeToString(h.Sum(nil))
	return
}
