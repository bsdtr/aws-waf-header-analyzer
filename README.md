<p align="center">
  <h3 align="center">AWS WaF HTTP header analyzer</h3>
  <p align="center">
    <a href="https://twitter.com/nulldutra">
      <img src="https://img.shields.io/badge/twitter-@nulldutra-blue.svg">
    </a>
    <a href="https://github.com/nulldutra/aws-waf-header-analyzer/actions/workflows/release.yaml">
      <img src="https://github.com/nulldutra/aws-waf-header-analyzer/actions/workflows/release.yaml/badge.svg">
    </a>
  </p>
</p>

<hr>

## Building and install

Golang is a dependencie to build the binary. See the documentation to install: https://go.dev/doc/install

```sh
make
sudo make install
```

## Rules

The rules configuration is very simple, for example, the threshold is the limited of the requests in X time.
It's possible to monitoring multiples headers, but, the header needs to be in HTTP Request header log.

### Config example

```yaml
rules:
  header:
    x-api-id: # The header name in HTTP Request header
      threshold: 100

    token:
      threshold: 1000
```

## Notifications

It's possible send notifications to Slack and Telegram. To configure slack notifications, you needs create a webhook configuration, see the slack documentation:
https://api.slack.com/messaging/webhooks

Telegram bot father:
https://t.me/botfather

### Config example

```yaml
notifications:
  slack:
    webhook-url: https://hooks.slack.com/services/DA2DA13QS/LW5DALDSMFDT5/qazqqd4f5Qph7LgXdZaHesXs

  telegram:
    bot-token: "123456789:NNDa2tbpq97izQx_invU6cox6uarhrlZDfa"
    chat-id:  "-4128833322"
```

## AWS

### Credentials

To set up AWS credentials, it's advisable to export them as environment variables. Here's a recommended approach:

```sh
export AWS_ACCESS_KEY_ID=".."
export AWS_SECRET_ACCESS_KEY=".."
export AWS_REGION="us-east-1"
```

### Log group

retrive-logs-minutes-ago is the time range you want to fetch the logs, in this example, logs from 1 hour ago.

```yaml
aws:
  waf-log-group-name: aws-waf-logs-cloudwatch-cloudfront
  region: us-east-1
  retrive-logs-minutes-ago: 60
```
