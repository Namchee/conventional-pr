# Conventional PR

[![Go Version Badge](https://img.shields.io/github/go-mod/go-version/namchee/conventional-pr)](https://github.com/Namchee/conventional-pr) [![Go Report Card](https://goreportcard.com/badge/github.com/Namchee/conventional-pr)](https://goreportcard.com/report/github.com/Namchee/conventional-pr) [![codecov](https://codecov.io/gh/Namchee/conventional-pr/branch/master/graph/badge.svg)](https://codecov.io/gh/Namchee/conventional-pr) [![All Contributors](https://img.shields.io/badge/all_contributors-2-orange.svg?style=flat)](#contributors-)

Conventional PR is a GitHub Action that validates all pull requests sent to a GitHub-hosted repository.

Conventional PR aims to ease your burden in managing your GitHub-hosted repository by validating, marking, even moderating low-quality attempts of pull request just by integrating Conventional PR to your existing CI/CD workflow.

## Features

- ✨ Configurable, tune Conventional PR easily to suit your needs.
- 💡 Sensible defaults, validates pull request with out-of-the-box sensibility.
- ♿ Whitelisting, validates pull request that actually matters.
- 📈 Transparent reporting, see what Conventional PR is actually doing.

## Usage

> To use Conventional PR, you'll need to prepare a GitHub access token. Please refer to this [article](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) on how to generate an access token.

You can integrate Conventional PR to your existing CI/CD workflow by using `Namchee/conventional-pr@<version>` in one of your jobs using the YAML syntax.

Below is the example of creating a new Conventional PR workflow in your CI/CD workflow.

```yml
on:
  pull_request:
    # You can add specific pull request types too.
    # If this is omitted, the workflow will only executed on opened, synchronize, and reopened
    # which is enough for generic use cases.
    # Below is the list of supported pull request sub-events
    # types: [opened, reopened, ready_for_review, unlocked, synchronize]

jobs:
  cpr:
    runs-on: ubuntu-latest
    steps:
      - name: Validates the pull request
        uses: Namchee/conventional-pr@v(version)
        with:
          access_token: YOUR_GITHUB_ACCESS_TOKEN_HERE
```

Please refer to [GitHub workflow syntax](https://docs.github.com/en/free-pro-team@latest/actions/reference/workflow-syntax-for-github-actions#about-yaml-syntax-for-workflows) for more advanced usage.

> Access token is **required**. Please generate one or use `${{ secrets.GITHUB_TOKEN }}` as your access token and the `github-actions` bot will run the job for you.

## Whitelist

Whitelist is one of the features of Conventional PR that allows you to bypass pull request validaton if the pull request satisfies **one or more** enabled whitelisting criteria.

Currently, there are 3 whitelisting criteria that are available in Conventional PR:

1. Pull request status is a `draft`.
2. Pull request is submitted by a bot.
3. Pull request is submitted by a user with high administrative privileges.

All whitelists are configurable. Please refer to the [inputs](#inputs) section on how to configure whitelists.

## Validator

Validator is the core feature of Conventional PR. Validator will validate pull request that triggers the workflow according to a specified criteria.

A pull request is considered to be valid if it satisfies **all** enabled validation flow.

Currently, there are 6 validation flow that are available in Conventional PR:

1. Pull request has a valid title.
2. Pull request has a non-empty body.
3. Pull request mentioned one or more issue.
4. Pull request does not introduce too many file changes.
5. All commits in the pull request have valid messages.
6. Pull request has a valid branch name.

All validators are configurable. Please refer to the [inputs](#inputs) section on how to configure whitelists.

## Inputs

You can customize this actions with these following options (fill it on `with` section):

| **Name**              | **Required?** | **Default Value**                       | **Description**                                                                                                                                                                                                                                                                                                            |
| --------------------- | ------------- | --------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `access_token`        | `true`        |                                         | [GitHub access token](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token) to interact with the GitHub API. It is recommended to store this token with [GitHub Secrets](https://docs.github.com/en/free-pro-team@latest/actions/reference/encrypted-secrets). |
| `close`               | `false`       | `true`                                  | Determine whether `conventional pr` should attempt to automatically close invalid pull requests.                                                                                                                                                                                                                           |
| `template`            | `false`       | ``                                      | Comment template to use when commenting on invalid pull requests. Fill with an empty string to disable further comments.                                                                                                                                                                                                   |
| `label`               | `false`       | `conventional pr:invalid`               | Label to use when marking invalid pull requests. If it doesn't exist, this action will automatically create it. Fill with an empty string to disable labeling.                                                                                                                                                             |
| `draft`               | `false`       | `true`                                  | Determine whether `conventional pr` should skip validating draft pull requests.                                                                                                                                                                                                                                            |
| `strict`              | `false`       | `true`                                  | Determine whether the restrictions should apply to repository administrators.                                                                                                                                                                                                                                              |
| `bot`                 | `false`       | `true`                                  | Determines whether checks should be skipped on PRs that is created by bots. Useful when relying on bots like [dependabot](https://github.com/dependabot)                                                                                                                                                                   |
| `title_pattern`       | `false`       | `([\w\-]+)(\([\w\-]+\))?!?: [\w\s:\-]+` | Valid pull request title regex pattern in Perl syntax. Defaults to the [conventional commit style](https://www.conventionalcommits.org/en/v1.0.0/) commit messages. Fill with an empty string to disabled pull request title validation.                                                                                   |
| `commit_pattern`      | `false`       | ``                                      | Valid pull request commit messages regex pattern in Perl syntax. Fill with an empty string to disabled commit message validation.                                                                                                                                                                                      |
| `branch_pattern`      | `false`       | ``                                      | Valid pull request branch name regex pattern in Perl syntax. Fill with an empty string to disabled branch name validation.                                                                                                                                                                                          |
| `body`                | `false`       | `true`                                  | Determine whether a valid pull request should always have a non-empty body.                                                                                                                                                                                                                                                |
| `issue`               | `false`       | `true`                                  | Determine whether a valid pull request should always have an issue or pull requests references on it.                                                                                                                                                                                                                      |
| `maximum_file_change` | `false`       | `0`                                     | Limits how many file can be changed per one pull request. Fill with zero to disable this feature.                                                                                                                                                                                                                          |
| `verified_commits` | `false` | `false` | Require all commits on the pull request to be [signed commits](https://docs.github.com/en/authentication/managing-commit-signature-verification/signing-commits) |

## Supported Events

Ideally, Conventional PR workflow should only triggered when an event related to pull requests is fired. However, Conventional PR is only able to function properly with some pull request sub-events and will ignore the rest. Below is the list of supported [GitHub Action pull request sub-events](https://docs.github.com/en/actions/reference/events-that-trigger-workflows):

| **Name**           | **Triggered On**                                                                           |
| ------------------ | ------------------------------------------------------------------------------------------ |
| `opened`           | When a new pull request is submitted.                                                      |
| `reopened`         | When a closed pull request is re-opened, either by the original author or by someone else. |
| `ready_for_review` | When a draft pull request is finished and transformed into normal pull request.            |
| `sychronize`       | When a pull request has changes on its history, such as pushing new commits to the branch. |
| `unlocked`         | When a pull request is unlocked.                                                           |

> Omitting `types` is enough for generic use-cases since omitting it will cause the workflow to be triggered on
`opened`, `reopened`, and `synchronize` events.

## Caveats

- If the issues are linked manually and are not mentioned in the pull request body, the pull request is still considered to be invalid.

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):
<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://github.com/Namchee"><img src="https://avatars.githubusercontent.com/u/32661241?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Cristopher</b></sub></a><br /><a href="https://github.com/Namchee/conventional-pr/commits?author=Namchee" title="Code">💻</a> <a href="https://github.com/Namchee/conventional-pr/issues?q=author%3ANamchee" title="Bug reports">🐛</a> <a href="https://github.com/Namchee/conventional-pr/commits?author=Namchee" title="Documentation">📖</a> <a href="#example-Namchee" title="Examples">💡</a> <a href="#ideas-Namchee" title="Ideas, Planning, & Feedback">🤔</a></td>
    <td align="center"><a href="https://github.com/smutel"><img src="https://avatars.githubusercontent.com/u/12967891?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Samuel Mutel</b></sub></a><br /><a href="https://github.com/Namchee/conventional-pr/issues?q=author%3Asmutel" title="Bug reports">🐛</a></td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://allcontributors.org) specification.
Contributions of any kind are welcome!

## License

This project is licensed under the [MIT License](./LICENSE)
