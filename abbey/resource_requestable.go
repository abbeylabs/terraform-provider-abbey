package abbey

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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
		"POST",
		"http://localhost:8080/v1/requestables",
		bytes.NewReader(body),
	)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	req.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	var x interface{}

	err = json.NewDecoder(response.Body).Decode(&x)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	if errMsg, ok := x.(string); ok {
		tflog.Info(ctx, errMsg)
		return diags
	}

	d.SetId(x.(map[string]interface{})["id"].(string))

	s, _ := json.MarshalIndent(x, "", "    ")
	tflog.Info(ctx, string(s))

	return diags
}

func resourceRequestableRead(
	ctx context.Context,
	d *schema.ResourceData,
	m interface{},
) diag.Diagnostics {
	var diags diag.Diagnostics

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

	return diags
}
