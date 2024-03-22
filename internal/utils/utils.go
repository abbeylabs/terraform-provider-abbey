package utils

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"reflect"
)

func Pointer[T any](v T) *T {
	return &v
}

func NullableString(s *string) types.String {
	if s != nil {
		return types.StringValue(*s)
	}
	return types.StringNull()
}

func NullableBool(b *bool) types.Bool {
	if b != nil {
		return types.BoolValue(*b)
	}
	return types.BoolNull()
}

func NullableInt64(i *int64) types.Int64 {
	if i != nil {
		return types.Int64Value(*i)
	}
	return types.Int64Null()
}

func NullableTfStateObject[T any, R any](source *T, fn func(t *T) R) *R {
	if source != nil {
		r := fn(source)
		return &r
	}
	return nil
}

func NullableObject[T any, R any](source *T, value R) *R {
	if source != nil {
		return &value
	}

	return nil
}

func NullableFloat64(f *float64) types.Float64 {
	if f != nil {
		return types.Float64Value(*f)
	}
	return types.Float64Null()
}

func MapList[T, R any](from []T, f func(T) R) []R {
	if from == nil {
		return nil
	}
	to := make([]R, len(from))
	for i, v := range from {
		to[i] = f(v)
	}
	return to
}

func ToList(ctx context.Context, from any, toType attr.Type, diagnostics *diag.Diagnostics) types.List {
	result, err := types.ListValueFrom(ctx, toType, from)
	if err != nil {
		diagnostics.Append(err.Warnings()...)
		diagnostics.Append(err.Errors()...)
		return types.ListUnknown(toType)
	}

	return result
}

func MapArrayQueryParams[T any](ctx context.Context, from types.List, diagnostics *diag.Diagnostics) []T {
	params := make([]T, 0)

	diags := from.ElementsAs(ctx, params, false)

	if diags.HasError() {
		diagnostics.Append(diags...)
	}

	return params
}

// ToMap toType expects a types.MapType
func ToMap(ctx context.Context, from any, toType attr.Type, diagnostics *diag.Diagnostics) types.Map {
	mapType, ok := toType.(types.MapType)

	if !ok {
		diagnostics.AddError("Cannot convert to Map", fmt.Sprintf("Error converting %v", toType))
		return types.MapUnknown(toType)
	}

	result, err := types.MapValueFrom(ctx, mapType.ElemType, from)
	if err != nil {
		diagnostics.Append(err.Warnings()...)
		diagnostics.Append(err.Errors()...)
		return types.MapUnknown(toType)
	}

	return result
}

func FromListToPrimitiveSlice[T any](ctx context.Context, from types.List, diagnostics *diag.Diagnostics) []T {
	elements := from.Elements()
	result := make([]T, len(elements))
	for i, elem := range elements {
		conversionMethod, err := getConversionMethodName(from.ElementType(ctx))
		if err != nil {
			diagnostics.Append(diag.NewErrorDiagnostic("conversion error", err.Error()))
			return nil
		}

		res := reflect.ValueOf(elem).MethodByName(conversionMethod).Call([]reflect.Value{})
		result[i] = res[0].Interface().(T)

	}
	return result
}

func FromTypesMapToMap[T any](ctx context.Context, from types.Map, diagnostics *diag.Diagnostics) map[string]T {
	m := make(map[string]T)

	diags := from.ElementsAs(ctx, &m, false)

	if diags.HasError() {
		diagnostics.Append(diags...)
	}

	return m
}

// PopulateModelData populates target interface with data from plan, replacing null and unknown with empty
// planGetterFn is a function which retrieves data from the request Plan, State or Config. Examples: req.State.Get, req.Plan.Get
// target should be a pointer
func PopulateModelData(ctx context.Context, target interface{}, diagnostics diag.Diagnostics, planGetterFn func(ctx context.Context, target interface{}) diag.Diagnostics) {
	var obj types.Object

	diagnostics.Append(planGetterFn(ctx, &obj)...)

	diagnostics.Append(obj.As(ctx, target, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
}

func TypeAtPath(ctx context.Context, p path.Path, state tfsdk.State, diagnostics *diag.Diagnostics) attr.Type {
	targetType, diags := state.Schema.TypeAtPath(ctx, p)

	if diags.HasError() {
		diagnostics.Append(diags...)
	}

	return targetType
}

func getConversionMethodName(t attr.Type) (string, error) {
	if t.Equal(types.StringType) {
		return "ValueString", nil
	} else if t.Equal(types.BoolType) {
		return "ValueBool", nil
	} else if t.Equal(types.Float64Type) {
		return "ValueFloat64", nil
	} else if t.Equal(types.Int64Type) {
		return "ValueInt64", nil
	} else {
		return "", fmt.Errorf("unsupported type %s", t.String())
	}
}
