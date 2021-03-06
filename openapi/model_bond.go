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
// Bond struct for Bond
type Bond struct {
	AssetType AssetType `json:"assetType,omitempty"`
	BondPrice float32 `json:"bondPrice,omitempty"`
	Cusip string `json:"cusip,omitempty"`
	Description string `json:"description,omitempty"`
	Exchange string `json:"exchange,omitempty"`
	Symbol string `json:"symbol,omitempty"`
}
