package guardiand

// Hardcoded telemetry GCP service account. Does not grant any access except for writing logs.
//
// Network operators can opt to send their logs to a shared Cloud Logging project for debugging.
// Logs are available to all network operators. There is no secret data logged anywhere.
//
// We encrypt the service account using a hardcoded symmetric key shared with guardians to
// prevent GitHub credential checkers from freaking out and to stop people from sending gigabytes of "gm".
//
// By using a separate key, we can keep the configuration decoupled from the telemetry backend,
// allowing the key to be replaced or even a different provider to be used without changing the config.

const defaultTelemetryProject = "projects/wormhole-logging"

const defaultTelemetryServiceAccountEnc = `
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
3SWVv2wizHdu
`
