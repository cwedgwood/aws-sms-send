# aws-sms-send

A minimalist command-line tool to send an SMS message via the AWS SNS/SMS API.

> [!WARNING]
> **This project is archived and no longer maintained.**
>
> When it was written, you could send an SMS to any number straight from AWS
> SNS with almost no setup. That era is over. AWS now requires a **registered
> origination identity** — 10DLC, a toll-free number, or a short code — plus an
> attached resource policy before you can send application-to-person (A2P) SMS
> to US recipients, and many other countries require sender-ID or number
> registration too. Unregistered long-code sending has effectively been
> retired. For casual or hobby use this tool no longer makes practical sense.
>
> **Consider a modern alternative.** For pushing notifications to browsers and
> devices, [Web Push](https://www.w3.org/TR/push-api/) (the W3C Push API) is a
> standards-based, no-per-message-cost option that mostly works across major
> platforms today.
>
> The code is left here for historical reference.

## Usage

```console
aws-sms-send -verbose +15551231234 "your pizza is ready"

aws-sms-send -transactional +15551231234 "ALERT: Thromdibulator malfunction, rebooting."
```

## Building

With Go installed, clone and build:

```console
git clone https://github.com/cwedgwood/aws-sms-send.git
cd aws-sms-send
make
# or
make install   # installs to /usr/local/bin/
```

## Exit codes

| Code | Meaning                    |
| ---- | -------------------------- |
| 1    | Usage (bad arguments)      |
| 2    | AWS API error              |
| 3    | Local IP resolution error  |

## Regions

See the AWS list of
[supported SNS SMS regions and countries](https://docs.aws.amazon.com/sns/latest/dg/sns-supported-regions-countries.html).

## License

GPL-3.0 — see the license header in the source and
<https://www.gnu.org/licenses/gpl-3.0.en.html>.
