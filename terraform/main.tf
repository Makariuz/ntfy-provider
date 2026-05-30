terraform {
  required_providers {
    ntfy = {
      source = "registry.terraform.io/<youruser>/ntfy"
    }
  }
}

provider "ntfy" {}

resource "ntfy_notification" "test" {
  topic   = "<youruser>"
  message = "The provider is live 🔥"
}
