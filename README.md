## NTFY notification provider

1. Find where your `go/bin` <PATH> `go env GOPATH` lives.
2. Create `~/.terraformrc` and add this:

```hcl
provider_installation {
  dev_overrides {
    "registry.terraform.io/<youruser>/ntfy" = "<PATH>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

> THIS must match the string in terraform/main.tf as when you run `terraform plan/apply`, `Terraform` will spawn it as a subprocess.

3. Run `go build -o terraform-provider-ntfy .` to build your binary.
