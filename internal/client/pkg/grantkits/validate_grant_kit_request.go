package grantkits

import (
	"github.com/go-provider-sdk/internal/marshal"
	"github.com/go-provider-sdk/internal/unmarshal"
)

type ValidateGrantKitRequest struct {
	GrantKitCreateParams *GrantKitCreateParams
	GrantKitUpdateParams *GrantKitUpdateParams
}

func (v *ValidateGrantKitRequest) SetGrantKitCreateParams(grantKitCreateParams GrantKitCreateParams) {
	v.GrantKitCreateParams = &grantKitCreateParams
}

func (v *ValidateGrantKitRequest) GetGrantKitCreateParams() *GrantKitCreateParams {
	if v == nil {
		return nil
	}
	return v.GrantKitCreateParams
}

func (v *ValidateGrantKitRequest) SetGrantKitUpdateParams(grantKitUpdateParams GrantKitUpdateParams) {
	v.GrantKitUpdateParams = &grantKitUpdateParams
}

func (v *ValidateGrantKitRequest) GetGrantKitUpdateParams() *GrantKitUpdateParams {
	if v == nil {
		return nil
	}
	return v.GrantKitUpdateParams
}

func (v ValidateGrantKitRequest) MarshalJSON() ([]byte, error) {
	return marshal.FromComplexObject(v)
}

func (v *ValidateGrantKitRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.ToComplexObject[ValidateGrantKitRequest](data, v)
}
