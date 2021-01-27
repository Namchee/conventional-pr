# Super Duper Conventional PR

[![Go Version Badge](https://img.shields.io/github/go-mod/go-version/namchee/conventional-pr)](https://github.com/Namchee/conventional-pr)

Conventional PR is a simple GitHub Action that attempts to ease your burden in managing your GitHub-hosted projects by automatically moderating low-quality pull requests (PR).

A PR is considered to be low quality if it satisfies at least one of these criterion:

- Doesn't have a description on it (empty body).
- Doesn't attempt to solve an issue (reference an issue).
- Doesn't follow [conventional commit](https://www.conventionalcommits.org/en/v1.0.0/) style
- Affects too many file at once. Smaller pull requests are easier to be reviewed, thus keeping the review quality high.

> You can disable checking on draft pull requests, [see below options](#Usage)

## Features

1. Automatically label and comment on low quality PRs.
2. Enforce commit style to everyone, even on administrators.
3. Automatically close low quality PRs.

> **Disclaimer**: This project is really new. If you want to suggest a new feature, just drop it to [issues section](https://github.com/Namchee/conventional-pr/issues) and label your request with the `enhancement`.

## Usage

Simply add this action on one of your GitHub workflows job:

```yml
on: pull_request # Run these every new pull request

jobs:
  hqprs:
    runs-on: ubuntu-latest
    steps:
      - name: hqprs
        uses: Namchee/conventional-pr@v(version)
        with:
          access_token: YOUR_GITHUB_ACCESS_TOKEN_HERE
```

Please refer to [GitHub workflow syntax](https://docs.github.com/en/free-pro-team@latest/actions/reference/workflow-syntax-for-github-actions#about-yaml-syntax-for-workflows) for more advanced usage.

See it in action [here](https://github.com/Namchee/conventional-pr/blob/master/.github/workflows/conventional-pr.yml)

## Inputs

You can customize this actions with these following options (fill it on `with` section):

Name | Required? | Default | Description
---- | --------- | ------- | -----------
`access_token` | `true` | | [GitHub access token](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token) to interact with GitHub API. It is recommended to store this token with [GitHub Secrets](https://docs.github.com/en/free-pro-team@latest/actions/reference/encrypted-secrets)
`close` | `false` | `true` | Determine whether the action should automatically close low quality pull requests.
`template` | `false` | [See The Metadata](./action.yml) | Comment template to use when commenting on low quality pull requests
`strict` | `false` | `true` | Determine whether the restrictions should apply to repository administrators.
`label` | `false` | `Nonconventional` | Label to use when marking low quality pull requests. If it doesn't exist, the action will automatically create it.
`allowed_types` | `false` | `''` | Allowed types according to [conventional commit](https://docs.github.com/en/free-pro-team@latest/actions/reference/encrypted-secrets) style. Fill it with empty string to allow all types.
`allowed_scopes` | `false` | `''` | Allowed scopes according to [conventional commit](https://docs.github.com/en/free-pro-team@latest/actions/reference/encrypted-secrets) style. Fill it with empty string to allow all scopes.
`check_draft` | `false` | `false` | Determine whether quality checks should also check draft pull requests.
`maximum_file_change` | `false` | `0` | Limits how many file can be changed per one pull request. Set it to zero to disable this feature.

## Caveats

- If the issues are linked manually and are not mentioned in the PR body, the PR is still considered to be invalid.

## License

This project is licensed under the [MIT License](./LICENSE)
