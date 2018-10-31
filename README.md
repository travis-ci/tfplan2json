# tfplan2json

In go tfplan, out go JSON.

Inspired by
[terraform-plan-decoder](https://github.com/pearkes/terraform-plan-decoder).

## installation

:warning: This is specific to the `v11` branch, which uses a vendored copy of
terraform at v0.11.10. The requirements for version-ish pinning enforced by
gopkg.in does not allow differentiation beyond the major release, nor does it
allow leading `0`s.

``` bash
go get gopkg.in/travis-ci/tfplan2json.v11
```
