# frodo-bot

A Discord bot that does one thing: whenever someone says "secret",
bot replies with "Alright then, keep your secrets."

## Setup

### Pre-Requisite

Only the self hosting is supported atm. You'd need to register your own bot on discord dev portal.

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

### 3. Supported Commands

| Command       | Description                                                         |
| ------------- | ------------------------------------------------------------------- |
| /morgul-blade | Stab Frodo with the cursed Morgul blade to silence him indefinitely |

Press `Ctrl+C` to stop.

## License

[MIT](LICENSE)
