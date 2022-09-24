## Whitelists and Validators

Whitelists and validators are the core of `conventional-pr` that determines the validity of the pull request that triggers the workflow.

### Whitelist

Whitelist is a collection of criteria that allows a pull request validation to be skipped. A pull request that fulfills **at least one** of the enabled whitelists criteria will be considered as a valid pull request. 

Currently, there are 4 available whitelists that can be used

1. [Bot](./whitelist/bot.md)
2. [Draft](./whitelist/draft.md)
3. [Administrator](./whitelist/admin.md)
4. [User](./whitelist/user.md)