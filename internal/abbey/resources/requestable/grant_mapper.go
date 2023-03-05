package requestable

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

const (
	grantTypeGenerateTf = "generate"

	generateGrantTypeGithubTf = "github"
)

func GenerateGrantFromTfTypesValue(ctx context.Context, value tftypes.Value) (ret *GenerateGrant, err error) {
	var m map[string]tftypes.Value
	if err := value.As(&m); err != nil {
		return nil, err
	}

	if len(m) == 0 {
		return nil, nil
	}

	var inner GenerateGrantEnum

	for key, val := range m {
		switch key {
		case generateGrantTypeGithubTf:
			inner_, err := GithubGenerateDestinationFromTfTypesValue(ctx, val)
			if err != nil {
				return nil, err
			}
			if inner_ == nil {
				continue
			}

			inner = inner_
		default:
			return nil, fmt.Errorf("unknown key: %s", key)
		}
		if err != nil {
			return nil, err
		}
	}

	return &GenerateGrant{Value: inner}, nil
}

func GithubGenerateDestinationFromTfTypesValue(_ context.Context, value tftypes.Value) (ret *GithubGenerateDestination, err error) {
	var m *map[string]tftypes.Value
	if err := value.As(&m); err != nil {
		return nil, err
	}

	if m == nil {
		return nil, nil
	}

	var (
		repo    string
		path    string
		append_ string
	)

	if err := (*m)["repo"].As(&repo); err != nil {
		return nil, err
	}
	if err := (*m)["path"].As(&path); err != nil {
		return nil, err
	}

	if (*m)["append"].IsKnown() {
		if err := (*m)["append"].As(&append_); err != nil {
			return nil, err
		}
	}

	return &GithubGenerateDestination{
		Repo:   repo,
		Path:   path,
		Append: append_,
	}, nil
}
