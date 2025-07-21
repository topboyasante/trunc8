# Environment Variables

For environment variables, I'm using the `os` package. I have a config struct to "model" how my env variables would look like.

I use a **grouped config approach** with separate structs for `ServerConfig` and `DatabaseConfig` to keep related settings organized.

I have helper functions:

- `getRequiredEnv()` - gets required env vars and returns an error if missing (used for DATABASE_URL)
- `getEnvWithDefault()` - gets optional env vars with fallback defaults (used for SERVER_PORT with default "8080")

My `LoadConfig()` function uses these helpers to load environment variables and populate the config struct, then returns a pointer to it. The function fails fast with descriptive errors for missing required variables.

**Environment variables used:**

- `DATABASE_URL` (required) - PostgreSQL connection string
- `SERVER_PORT` (optional, defaults to "8080") - HTTP server port
