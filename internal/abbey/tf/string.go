package tf

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type (
	String = Attribute[String_, StringType_, string]

	String_     struct{ value string }
	StringType_ struct{}
)

var (
	StringType = AttributeType[String_, StringType_, string]{type_: StringType_{}}

	_ Value[string] = String_{value: ""}
	_ Type[String_] = StringType_{}
)

func (self String_) String() string         { return self.value }
func (self String_) ToBuiltinValue() string { return self.value }

func (self String_) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.value)
}

func (self *String_) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	*self = String_{value: str}

	return nil
}

func (self StringType_) TerraformType(context.Context) tftypes.Type { return tftypes.String }

func (self StringType_) ValueFromTerraform(_ context.Context, value tftypes.Value) (invalid String_, err error) {
	var s string
	if err := value.As(&s); err != nil {
		return invalid, err
	}

	return String_{value: s}, nil
}

func (self StringType_) String() string { return "String" }

func (self StringType_) ApplyTerraform5AttributePathStep(step tftypes.AttributePathStep) (interface{}, error) {
	return nil, fmt.Errorf("cannot apply AttributePathStep %T to %s", step, self.String())
}
