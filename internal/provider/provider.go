// resources

package provider

import (
    "context"
    "github.com/hashicorp/terraform-plugin-framework/provider"
    "github.com/hashicorp/terraform-plugin-framework/provider/schema"
    "github.com/hashicorp/terraform-plugin-framework/resource"
    "github.com/hashicorp/terraform-plugin-framework/datasource"
)

type ntfyProvider struct{}

func New() func() provider.Provider {
    return func() provider.Provider {
        return &ntfyProvider{}
    }
}

func (p *ntfyProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
    resp.TypeName = "ntfy"
}

func (p *ntfyProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
    resp.Schema = schema.Schema{}
}

func (p *ntfyProvider) Configure(_ context.Context, _ provider.ConfigureRequest, resp *provider.ConfigureResponse) {}

func (p *ntfyProvider) Resources(_ context.Context) []func() resource.Resource {
    return []func() resource.Resource{
        NewNotificationResource,
    }
}

func (p *ntfyProvider) DataSources(_ context.Context) []func() datasource.DataSource {
    return nil
}