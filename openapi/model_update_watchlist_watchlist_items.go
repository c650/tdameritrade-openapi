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
// UpdateWatchlistWatchlistItems struct for UpdateWatchlistWatchlistItems
type UpdateWatchlistWatchlistItems struct {
	AveragePrice int32 `json:"averagePrice,omitempty"`
	Commission int32 `json:"commission,omitempty"`
	Instrument CreateWatchlistInstrument `json:"instrument,omitempty"`
	PurchasedDate string `json:"purchasedDate,omitempty"`
	Quantity int32 `json:"quantity,omitempty"`
	SequenceId int32 `json:"sequenceId,omitempty"`
}
