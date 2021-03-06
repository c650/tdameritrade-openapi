/*
 * TD Ameritrade API
 *
 * TD Ameritrade API
 *
 * API version: 3.0.1
 * Contact: austin.millan@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CashAccountInitialBalances struct for CashAccountInitialBalances
type CashAccountInitialBalances struct {
	AccountValue float32 `json:"accountValue,omitempty"`
	AccruedInterest float32 `json:"accruedInterest,omitempty"`
	BondValue float32 `json:"bondValue,omitempty"`
	CashAvailableForTrading float32 `json:"cashAvailableForTrading,omitempty"`
	CashAvailableForWithdrawal float32 `json:"cashAvailableForWithdrawal,omitempty"`
	CashBalance float32 `json:"cashBalance,omitempty"`
	CashDebitCallValue float32 `json:"cashDebitCallValue,omitempty"`
	CashReceipts float32 `json:"cashReceipts,omitempty"`
	IsInCall bool `json:"isInCall,omitempty"`
	LiquidationValue float32 `json:"liquidationValue,omitempty"`
	LongOptionMarketValue float32 `json:"longOptionMarketValue,omitempty"`
	LongStockValue float32 `json:"longStockValue,omitempty"`
	MoneyMarketFund float32 `json:"moneyMarketFund,omitempty"`
	MutualFundValue float32 `json:"mutualFundValue,omitempty"`
	PendingDeposits float32 `json:"pendingDeposits,omitempty"`
	ShortOptionMarketValue float32 `json:"shortOptionMarketValue,omitempty"`
	ShortStockValue float32 `json:"shortStockValue,omitempty"`
	UnsettledCash float32 `json:"unsettledCash,omitempty"`
}
