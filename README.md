# frodo-bot

A Discord bot that does one thing: whenever someone says "secret",
bot replies with "Alright then, keep your secrets."

## Add Frodo to your server

**[Invite Frodo](https://discord.com/oauth2/authorize?client_id=1561964364746195014&permissions=84992&integration_type=0&scope=bot+applications.commands)**

This is a hosted instance that always runs the latest `main` of this repo, so there's nothing to set up.
If you'd rather run your own, see [Self-hosting](#self-hosting).

## Commands

| Command       | Description                                                         |
| ------------- | ------------------------------------------------------------------- |
| /morgul-blade | Stab Frodo with the cursed Morgul blade to silence him indefinitely |

## Self-hosting

### Pre-Requisite

Register your own bot on the [Discord Developer Portal](https://discord.com/developers/applications),
and enable **Message Content Intent** in its Bot tab. Invite it with the `bot` and
`applications.commands` scopes.

### 1. Configure

```bash
cp .env.example .env
```

Then fill in `.env`:

| Variable            | Required | Description                              |
| ------------------- | -------- | ---------------------------------------- |
| `DISCORD_TOKEN`     | Yes      | Your bot token                           |
| `KEEP_YOUR_SECRETS` | Yes      | Direct link to the GIF the bot will post |

### 2. Run

```bash
go run ./cmd/frodo-bot
```

Or with Docker:

```bash
docker build -t frodo-bot .
docker run --rm --env-file .env frodo-bot
```

Press `Ctrl+C` to stop.

## License

[MIT](LICENSE)
