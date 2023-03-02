package tf

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type (
	Attribute[V ComparableValue[B], T Type[V], B Builtin] struct {
		inner attribute[V]
		type_ T
	}

	AttributeType[V ComparableValue[B], T Type[V], B Builtin] struct{ type_ T }

	attribute[V any] interface{ visit(Visitor[V]) }

	Visitor[V any] struct {
		Null    func()
		Unknown func()
		Valued  func(value V)
	}

	null[_ any]    struct{}
	unknown[_ any] struct{}
	valued[V any]  struct{ value V }

	ComparableValue[B Builtin] interface {
		comparable
		Value[B]
	}

	Builtin interface {
		string | bool
	}

	Value[B Builtin] interface {
		fmt.Stringer

		ToBuiltinValue() B
	}

	Type[V any] interface {
		TerraformType(context.Context) tftypes.Type
		ValueFromTerraform(context.Context, tftypes.Value) (V, error)

		fmt.Stringer
		tftypes.AttributePathStepper
	}
)

var (
	_ attribute[any] = (*null[any])(nil)
	_ attribute[any] = (*unknown[any])(nil)
	_ attribute[any] = (*valued[any])(nil)

	_ attr.Value               = (*Attribute[Value[string], Type[Value[string]], string])(nil)
	_ basetypes.StringValuable = (*Attribute[Value[string], Type[Value[string]], string])(nil)
	_ basetypes.StringTypable  = (*AttributeType[Value[string], Type[Value[string]], string])(nil)
)

func (self null[V]) visit(visitor Visitor[V])    { visitor.Null() }             //nolint:unused
func (self unknown[V]) visit(visitor Visitor[V]) { visitor.Unknown() }          //nolint:unused
func (self valued[V]) visit(visitor Visitor[V])  { visitor.Valued(self.value) } //nolint:unused

func (self Attribute[V, T, B]) Visit(visitor Visitor[V]) { self.inner.visit(visitor) }

func (self Attribute[V, T, B]) Type(context.Context) attr.Type {
	return AttributeType[V, T, B]{type_: self.type_}
}

func (self Attribute[V, T, B]) ToTerraformValue(ctx context.Context) (value tftypes.Value, err error) {
	self.inner.visit(Visitor[V]{
		Null: func() {
			value = tftypes.NewValue(self.type_.TerraformType(ctx), nil)
		},
		Unknown: func() {
			value = tftypes.NewValue(self.type_.TerraformType(ctx), tftypes.UnknownValue)
		},
		Valued: func(v V) {
			value = tftypes.NewValue(self.type_.TerraformType(ctx), v.ToBuiltinValue())
		},
	})

	return value, nil
}

func (self Attribute[V, T, B]) Equal(value attr.Value) (equal bool) {
	other, ok := value.(Attribute[V, T, B])
	if !ok {
		return false
	}

	self.Visit(Visitor[V]{
		Null:    func() { equal = value.IsNull() },
		Unknown: func() {},
		Valued: func(this V) {
			other.Visit(Visitor[V]{
				Null:    func() {},
				Unknown: func() {},
				Valued: func(that V) {
					equal = this == that
				},
			})
		},
	})

	return equal
}

func (self Attribute[V, T, B]) IsNull() (null bool) {
	self.Visit(Visitor[V]{
		Null:    func() { null = true },
		Unknown: func() {},
		Valued:  func(V) {},
	})

	return null
}

func (self Attribute[V, T, B]) IsUnknown() (unknown bool) {
	self.Visit(Visitor[V]{
		Null:    func() {},
		Unknown: func() { unknown = true },
		Valued:  func(V) {},
	})

	return unknown
}

func (self Attribute[V, T, B]) String() (str string) {
	self.Visit(Visitor[V]{
		Null:    func() { str = attr.NullValueString },
		Unknown: func() { str = attr.UnknownValueString },
		Valued:  func(v V) { str = v.String() },
	})

	return str
}

func (self Attribute[V, T, B]) ToStringValue(ctx context.Context) (value basetypes.StringValue, diags diag.Diagnostics) {
	if !self.type_.TerraformType(ctx).Equal(tftypes.String) {
		diags.AddError("Wrong Type", "The attribute is not a string type.")
		return value, diags
	}

	self.Visit(Visitor[V]{
		Null:    func() { value = basetypes.NewStringNull() },
		Unknown: func() { value = basetypes.NewStringUnknown() },
		Valued: func(v V) {
			s, ok := any(v.ToBuiltinValue()).(string)
			if !ok {
				diags.AddError("Wrong Type", "The attribute is not a string type.")
				return
			}

			value = basetypes.NewStringValue(s)
		},
	})

	return value, diags
}

func (self *Attribute[V, T, B]) UnmarshalJSON(data []byte) error {
	var value *V
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if value == nil {
		*self = Attribute[V, T, B]{inner: null[V]{}, type_: self.type_}
		return nil
	}

	*self = Attribute[V, T, B]{inner: valued[V]{value: *value}, type_: self.type_}

	return nil
}

func (self AttributeType[V, T, B]) TerraformType(ctx context.Context) tftypes.Type {
	return self.type_.TerraformType(ctx)
}

func (self AttributeType[V, T, B]) ValueFromTerraform(ctx context.Context, value tftypes.Value) (attr.Value, error) {
	if value.IsNull() {
		return Attribute[V, T, B]{inner: null[V]{}, type_: self.type_}, nil
	}

	if !value.IsKnown() {
		return Attribute[V, T, B]{inner: unknown[V]{}, type_: self.type_}, nil
	}

	val, err := self.type_.ValueFromTerraform(ctx, value)
	if err != nil {
		return nil, err
	}

	return Attribute[V, T, B]{inner: valued[V]{value: val}, type_: self.type_}, nil
}

func (self AttributeType[V, T, B]) ValueType(context.Context) attr.Value {
	return Attribute[V, T, B]{inner: null[V]{}, type_: self.type_}
}

func (self AttributeType[V, T, B]) Equal(t attr.Type) bool {
	_, equal := t.(AttributeType[V, T, B])
	return equal
}

func (self AttributeType[V, T, B]) String() string {
	return self.type_.String()
}

func (self AttributeType[V, T, B]) ApplyTerraform5AttributePathStep(step tftypes.AttributePathStep) (interface{}, error) {
	return self.type_.ApplyTerraform5AttributePathStep(step)
}

func (self AttributeType[V, T, B]) ValueFromString(
	ctx context.Context,
	value basetypes.StringValue,
) (invalid basetypes.StringValuable, diags diag.Diagnostics) {
	if value.IsUnknown() {
		return Attribute[V, T, B]{inner: unknown[V]{}, type_: self.type_}, nil
	}

	if value.IsNull() {
		return Attribute[V, T, B]{inner: null[V]{}, type_: self.type_}, nil
	}

	val, err := value.ToTerraformValue(ctx)
	if err != nil {
		diags.AddError("Unknown", "Failed to convert to tftypes.Value.")
		return invalid, diags
	}

	v, err := self.type_.ValueFromTerraform(ctx, val)
	if err != nil {
		diags.AddError("Unknown", "Failed to create value from Terraform value.")
		return invalid, diags
	}

	return Attribute[V, T, B]{inner: valued[V]{value: v}, type_: self.type_}, nil
}
