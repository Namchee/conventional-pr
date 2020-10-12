# Conventional PR

Conventional PR is a simple GitHub Action that attempts to ease your burden in managing your GitHub-hosted projects by automatically moderating low-quality pull requests (PR).

A PR is considered to be low quality if it satisfies at least one of these criterion:

- Doesn't have a description on it (empty body).
- Doesn't attempt to solve an issue (reference an issue).
- Doesn't follow [conventional commit](https://www.conventionalcommits.org/en/v1.0.0/) style

> You can disable checking on draft pull requests, [see below options](#Usage)

## Features

1. Automatically label and comment on low quality PRs.
2. Enforce rules to everyone, even on administrators.
3. (Optional) Automatically close low quality PRs.

> **Disclaimer**: This project is really new. If you want to suggest a new feature, just drop it to [issues section](https://github.com/Namchee/conventional-pr/issues) and label your request `feature request`.

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
          access_token: YOUR_GITHUB_ACCESS_TOKEN_HERE # required
          close: true # Determine whether Conventional PR should automatically close low quality PRs or not (default: true)
          template: "Some string here" # Comment template (default: See code)
          strict: true # Force rules on everyone, including administrators (default: true)
          label: 'Some label' # Label on low quality PRs (default: 'Invalid')
          allowedTypes: 'a, b, c' # Allowed conventional commit types. Must be comma-separated string. Fill with empty to allow all types. (default: empty string)
          allowedScopes: 'a, b, c' # Allowed conventional commit scopes. Must be comma-separated string. Fill with empty to allow all scopes. (default: empty string)
          checkDraft: false # Determine whether PR check should also run on draft PRs (default: false)
```

Refer to [GitHub workflow syntax](https://docs.github.com/en/free-pro-team@latest/actions/reference/workflow-syntax-for-github-actions#about-yaml-syntax-for-workflows) for more advanced usage.

See the example [here](https://github.com/Namchee/conventional-pr/blob/master/.github/workflows/conventional-pr.yml)

## Caveats

- If the issues are linked manually and are not mentioned in the PR body, the PR is still considered to be invalid.

## License

This project is licensed under the [MIT License](./LICENSE)
