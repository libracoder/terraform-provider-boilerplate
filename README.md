<p align="center">
    <img alt="terraform-provider-boilerplate Logo" src="https://www.freelogovectors.net/wp-content/uploads/2022/01/terra-form-logo-freelogovectors.net_.png" width="150" />
    <h1 align="center">Terraform Provider Boilerplate.</h1>
    <p align="center">Start your terraform provider project in seconds
with a professional template for building fast, robust, and adaptable terraform providers with a focus on performance and best practices</p>
    <p align="center">
        <a href="https://godoc.org/github.com/libracoder/terraform-provider-boilerplate"><img src="https://godoc.org/github.com/libracoder/terraform-provider-boilerplate?status.svg"></a>
        <a href="https://godoc.org/github.com/libracoder/terraform-provider-boilerplate"><img src="https://github.com/libracoder/terraform-provider-boilerplate/actions/workflows/release.yaml/badge.svg"></a>
        <a href="https://github.com/libracoder/terraform-provider-boilerplate/releases"><img src="https://img.shields.io/badge/Version-0.0.1alpha-red.svg"></a>
        <a href="https://goreportcard.com/report/github.com/libracoder/terraform-provider-boilerplate"><img src="https://goreportcard.com/badge/github.com/libracoder/terraform-provider-boilerplate"></a>
        <a href="https://github.com/libracoder/terraform-provider-boilerplate/blob/master/LICENSE"><img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg"></a>
    </p>




```bash
$ git clone https://github.com/libracoder/terraform-provider-boilerplate.git

# Build the provider
$ make ARGS="terraform-provider-boilerplate" build
// OR
$ go build -o terraform-provider-boilerplate

# Initialize a working directory containing Terraform configuration files
$ terraform init

# Create an execution plan.
$ terraform plan

# Apply the changes required to reach the desired state of the configuration
$ terraform apply

# Revert changes
$ terraform destroy
```


## General Rules

- Terraform provider should always consume an independent client library or sdk which implements the core logic for communicating with the upstream. You should consider moving the `/sdk` to be a separate project.
- Data sources are a special subset of resources which are read-only. They are resolved earlier than regular resources and can be used as part of Terraform's interpolation.


## Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, terraform-provider-boilerplate is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our GitHub project](https://github.com/libracoder/terraform-provider-boilerplate/releases) for changelogs for each release version of terraform-provider-boilerplate. It contains summaries of the most noteworthy changes made in each release.


## Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/libracoder/terraform-provider-boilerplate/issues


## Security Issues

If you discover a security vulnerability within terraform-provider-boilerplate, please send an email to [libracoder@gmail.com](mailto:libracoder@gmail.com)


## Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.


## License

© 2022, libracoder. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**terraform-provider-boilerplate** is authored and maintained by [@libracoder](http://github.com/libracoder).
