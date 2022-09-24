# Whitelists and Validators

Whitelists and validators are the core functionality of `conventional-pr` that determines the validity of the pull request that triggers the workflow.

Whitelists and validators can be enabled or disabled from the workflow inputs.

## Whitelist

Whitelist is a collection of criteria that allows a pull request validator to be skipped. A pull request that satisfies **at least one** of the enabled whitelists criteria will be marked as a valid pull request. 

Currently, there are 4 available whitelists that can be used

1. [Bot](./whitelist/bot.md)
2. [Draft](./whitelist/draft.md)
3. [Administrator](./whitelist/admin.md)
4. [User](./whitelist/user.md)

## Validator

Validator is the core feature of Conventional PR. Validator will validate pull request that triggers the workflow according to a specified criteria.

A pull request is considered to be valid if it satisfies **all** enabled validator criteria.