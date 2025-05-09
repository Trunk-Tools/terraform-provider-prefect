<p align="center"><img src="https://github.com/PrefectHQ/prefect/assets/3407835/c654cbc6-63e8-4ada-a92a-efd2f8f24b85" width=1000></p>

<p align="center">
    <a href="https://prefect.io/slack" alt="Slack">
        <img src="https://img.shields.io/badge/slack-join_community-red.svg?color=0052FF&labelColor=090422&logo=slack" /></a>
    <a href="https://discourse.prefect.io/" alt="GitHub Discussions">
        <img src="https://img.shields.io/github/discussions/prefecthq/prefect"></a>
</p>

<a href="https://terraform.io">
    <img src=".github/tf.png" alt="Terraform logo" title="Terraform" align="left" height="50" />
</a>

# Terraform Provider for Prefect Cloud
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/prefecthq/terraform-provider-prefect?label=release)](https://github.com/prefecthq/terraform-provider-prefect/releases) ![Acceptance tests](https://github.com/PrefectHQ/terraform-provider-prefect/actions/workflows/acceptance-tests.yaml/badge.svg) ![Provider Release](https://github.com/PrefectHQ/terraform-provider-prefect/actions/workflows/provider-release.yaml/badge.svg)

[Prefect](https://www.prefect.io/) is a powerful tool for creating workflow applications. The Terraform Prefect provider is a plugin that allows Terraform to manage resources on [Prefect Cloud](https://app.prefect.cloud) and [Prefect OSS](https://github.com/prefecthq/prefect). This provider is maintained by the Platform team at Prefect.

- [Documentation](https://registry.terraform.io/providers/PrefectHQ/prefect/latest/docs)
- [Examples](https://github.com/PrefectHQ/terraform-provider-prefect/tree/main/examples)

## Supported objects

The supported objects are listed in the [provider documentation](https://registry.terraform.io/providers/PrefectHQ/prefect/latest/docs).

We maintain an [API parity document](https://github.com/PrefectHQ/terraform-provider-prefect/wiki/API-Parity) to keep track of
features available in the [Prefect API](https://app.prefect.cloud/api/docs) along with their current level of support in the Terraform provider.

If you notice any missing features, see the [contributing section](#contributing) to file an issue in our bug tracker.

## Contributing

We appreciate your interest in the Prefect provider! If you find any issues or have ideas for improvement, you can always:

- [Check out our contributing guide](/_about/CONTRIBUTING.md)
- [File an issue](https://github.com/PrefectHQ/terraform-provider-prefect/issues/new?assignees=&labels=bug&projects=&template=bug.md)
- [Send us a note](mailto:security@prefect.io) for any potential security issues
- Reach out in the [Prefect community Slack workspace](https://communityinviter.com/apps/prefect-community/prefect-community)
