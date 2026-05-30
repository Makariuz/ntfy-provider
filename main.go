// starts the provider binary

package main

import (
    "context"
    "github.com/hashicorp/terraform-plugin-framework/providerserver"
    "terraform-provider-ntfy/internal/provider"
)

func main() {
    providerserver.Serve(context.Background(), provider.New(), providerserver.ServeOpts{
        Address: "registry.terraform.io/<youruser>/ntfy",
    })
}