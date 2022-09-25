# Bot Whitelist

Default | Type | Configuration
------- | ---- | -------------
**ACTIVE** | `boolean` | `bot`

The bot whitelist is a whitelist that allows any pull requests that have been submitted by a bot account to be marked as a valid pull request.

Do note that this whitelist only works on GitHub-recognized bot accounts. For usage with user-bot account, please use the [user whitelist](./user.md).

## Example

`conventional-pr` is executed with the following inputs.

```yml
on:
  pull_request:

jobs:
  cpr:
    runs-on: ubuntu-latest
    steps:
      - name: Validates the pull request
        uses: Namchee/conventional-pr@master
        with:
          access_token: access_token
          bot: true
```

Bot account `allcontributors` submitted the following pull request.

![Pull request by allcontributors](./bot.png)

Since `allcontributors` is marked as a bot account GitHub, the pull request will automatically be marked as a valid pull request.