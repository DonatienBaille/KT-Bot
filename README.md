# KT Bot

KT Bot (Kitty Bot) is a little Discord Bot to manage containers directly from Discord. It was developed to allow my friends to restart game servers while I was away.

KT Bot allows starting, stopping, restarting, and updating containers. The bot only discovers specific containers from Docker Daemon (see configuration below).

## Dependencies

The main dependencies are:

- [DiscordGo](https://github.com/bwmarrin/discordgo): Used to communicate with Discord API. Thanks bwmarrin.
- [Moby](https://github.com/moby/moby): Used to communicate with Docker Daemon API (official dependency of Docker CLI).
- [Watchtower](https://github.com/containrrr/watchtower): Used to update Docker images for containers.
- [Viper](https://github.com/spf13/viper): Used for configuration of environment variables.
- [lo](https://github.com/samber/lo): Used for slice manipulation.

There are also some other packages from the Go community. See go.mod for the full list.

## Get started

You can pull the image directly [from Docker Hub](https://hub.docker.com/r/donatienbaille/ktbot) or build it yourself and host it where you prefer.

### Discord Bot

The creation of the Discord bot is not described here. There is already a lot of documentation on the web. I recommend following [this guide](https://discordjs.guide/preparations/setting-up-a-bot-application.html#creating-your-bot).

Required permissions:

- "Send messages"
- "Manage messages"

### Watchtower

To enable container images updates, you have to run Watchtower and [enable HTTP API](https://containrrr.dev/watchtower/http-api-mode/).

All of this is already well described in the [documentation of Watchtower](https://containrrr.dev/watchtower/).

### Configuration

This image rely on environment variables for configuration:

| Environment variable    | Default value          | Description                                                                                                                              |
| ----------------------- | ---------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| KT_GUILD_ID             |                        | ID of the Discord server to send messages                                                                                                |
| KT_CHANNEL_ID           |                        | ID of the channel to send container statuses                                                                                             |
| KT_BOT_DISCORD_TOKEN    |                        | Discord Bot Token (from developer portal) prefixed by "Bot " (more details [here](https://pkg.go.dev/github.com/bwmarrin/discordgo#New)) |
| KT_WATCHTOWER_API_URL   | http://127.0.0.1:8080  | URL of Watchtower                                                                                                                        |
| KT_WATCHTOWER_API_TOKEN |                        | API Token configured for Watchtower                                                                                                      |
| DOCKER_HOST             | `/var/run/docker.sock` | URL of the Docker server                                                                                                                 |

Guild ID and Channel ID can be obtained by using developer mode in Discord (Settings -> Advanced -> Developer mode). After that, you can right click on a server or a channel and copy the ID.

For Docker environment variables, you can also override other variables as described [here](https://pkg.go.dev/github.com/docker/docker/client#FromEnv).

Now, if you start the bot, you should see nothing. That's because only specific containers are exposed to Discord. To add a container, you have to add a new label on the container "visibility" with the value "bot-discord".

## Local testing

You can test the bot locally. For that, you need to link the Docker daemon :

- MacOS: on Docker Desktop, go to Settings -> Advanced -> Allow the default Docker socket to be used (`/var/run/docker.sock`).
- Windows: on Docker Desktop, go to Settings -> General -> Expose daemon on `tcp://localhost:2375` without TLS.

After that, you can just configure the DOCKER_HOST. If this doesn't work, you'll have to use the "Host" Network for the container in order to access the local url.

## Getting help

If you have questions, concerns, bug reports, etc, please file an issue in this repository's Issue Tracker.

## Getting involved

This project can have some bugs. Feel free to create issues with logs et steps to reproduce or create a pull request.

There are currently no planned developments (maybe display logs on demand or add RCON for live status?), but don't hesitate to contact me if you have any use cases by creating an issue!

If you want to add some features, feel free to pull request.
