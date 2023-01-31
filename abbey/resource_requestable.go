package abbey

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Requestable struct {
	Id              string       `json:"id"`
	CreatedAt       time.Time    `json:"created_at"`
	Name            string       `json:"name"`
	Manage          string       `json:"manage,omitempty"`
	Generate        string       `json:"generate"`
	InputJsonSchema *interface{} `json:"input_json_schema"`
}

func resourceRequestable() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRequestableCreate,
		ReadContext:   resourceRequestableRead,
		UpdateContext: resourceRequestableUpdate,
		DeleteContext: resourceRequestableDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"manage": {
				Type:     schema.TypeString,
				Required: true,
			},
			"generate": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"input_json_schema_json": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceRequestableCreate(
	ctx context.Context,
	d *schema.ResourceData,
	m interface{},
) diag.Diagnostics {
	type RequestBody struct {
		Name            string      `json:"name"`
		Manage          string      `json:"manage"`
		Generate        string      `json:"generate"`
		InputJsonSchema interface{} `json:"input_json_schema"`
	}

	var diags diag.Diagnostics

	requestBody := RequestBody{
		Name:            d.Get("name").(string),
		Manage:          d.Get("manage").(string),
		Generate:        d.Get("generate").(string),
		InputJsonSchema: json.RawMessage(d.Get("input_json_schema_json").(string)),
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/v1/requestables", m.(Meta).Host),
		bytes.NewReader(body),
	)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(http.CanonicalHeaderKey("Authorization"), fmt.Sprintf("Bearer %s", m.(Meta).Token))

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	var requestable Requestable

	err = json.NewDecoder(response.Body).Decode(&requestable)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.SetId(requestable.Id)

	return diags
}

func resourceRequestableRead(
	ctx context.Context,
	d *schema.ResourceData,
	m interface{},
) diag.Diagnostics {
	var diags diag.Diagnostics

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/v1/requestables/%s", m.(Meta).Host, d.Id()),
		nil,
	)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	req.Header.Set(http.CanonicalHeaderKey("Authorization"), fmt.Sprintf("Bearer %s", m.(Meta).Token))

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	if response.StatusCode == http.StatusNotFound {
		d.SetId("")
		return diags
	}

	if response.StatusCode != http.StatusCreated {
		diags = append(diags, diag.FromErr(fmt.Errorf("unexpected status %s", response.Status))...)
		return diags
	}

	var requestable Requestable

	err = json.NewDecoder(response.Body).Decode(&requestable)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	inputJsonSchemaJson := new(strings.Builder)

	err = json.NewEncoder(inputJsonSchemaJson).Encode(requestable.InputJsonSchema)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.Set("name", requestable.Name)
	d.Set("manage", requestable.Manage)
	d.Set("generate", requestable.Generate)
	d.Set("input_json_schema_json", inputJsonSchemaJson)

	return diags
}

func resourceRequestableUpdate(
	ctx context.Context,
	d *schema.ResourceData,
	m interface{},
) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}

func resourceRequestableDelete(
	ctx context.Context,
	d *schema.ResourceData,
	m interface{},
) diag.Diagnostics {
	var diags diag.Diagnostics

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("%s/v1/requestables/%s", m.(Meta).Host, d.Id()),
		nil,
	)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	req.Header.Set(http.CanonicalHeaderKey("Authorization"), fmt.Sprintf("Bearer %s", m.(Meta).Token))

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	if response.StatusCode != http.StatusNoContent {
		diags = append(diags, diag.FromErr(fmt.Errorf("unexpected status %s", response.Status))...)
		return diags
	}

	return diags
}
