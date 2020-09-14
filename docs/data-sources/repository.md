# Data Source Repository

Data source to return the ID of a Repository object

## Example Usage
```sh
data "ovm_repository" "repo" {
    name = "my-repository"
}
```

## Attributes

The following arguments are supported:

+ `id` - ID of the Repository
+ `value` - Same as ID
+ `uri` - URI of the repository object
+ `name` - Name of the repository
+ `type` - OVM object model