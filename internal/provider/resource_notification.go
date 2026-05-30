// actual service

package provider

import (
    "context"
    "net/http"
    "strings"

    "github.com/hashicorp/terraform-plugin-framework/resource"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema"
    "github.com/hashicorp/terraform-plugin-framework/types"
)

type notificationResource struct{}

type notificationModel struct {
    Topic   types.String `tfsdk:"topic"`
    Message types.String `tfsdk:"message"`
}

func NewNotificationResource() resource.Resource {
    return &notificationResource{}
}

func (r *notificationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_notification"
}

func (r *notificationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
            "topic":   schema.StringAttribute{Required: true},
            "message": schema.StringAttribute{Required: true},
        },
    }
}

func (r *notificationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
    var data notificationModel
    req.Plan.Get(ctx, &data)

    http.Post(
        "https://ntfy.sh/"+data.Topic.ValueString(),
        "text/plain",
        strings.NewReader(data.Message.ValueString()),
    )

    resp.State.Set(ctx, &data)
}

func (r *notificationResource) Read(_ context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {}
func (r *notificationResource) Update(_ context.Context, _ resource.UpdateRequest, _ *resource.UpdateResponse) {}
func (r *notificationResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {}