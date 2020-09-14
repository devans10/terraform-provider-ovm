# Oracle VM Provider

The provider is used to interact with Oracle VM 3.  At this time it is narrowly focused on creating Virtual Machines from templates.  Those template images can be created by Packer and imported into OVM.

Thank you to Bjorn Ahl [(dbgeek)](https://github.com/dbgeek). He started this provider.  I just picked it up once he could no longer support it.

## Example Usage

```sh
terraform {
  required_providers {
    ovm = {
      source  = "devans10/ovm"
      version = "0.3.5"
    }
  }
}

provider "ovm" {
  user       = var.username
  password   = var.password
  entrypoint = var.entrypoint
}
```

## Argument Reference

The following arguments are supported:

+ `entrypoint` - (Optional) This is the URL of the OVM Manager. ex. https://localhost:7002/
+ `user` - (Optional) The username to connect to the OVM Manager.
+ `password` - (Optional) The password used to connect to the OVM Manager

Optionally, the provider can be configured using environment variables `OVM_ENDPOINT`, `OVM_USERNAME`, and `OVM_PASSWORD`