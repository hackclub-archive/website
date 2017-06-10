# Hack Club's Website [![Circle CI](https://circleci.com/gh/hackclub/website.svg?style=svg)](https://circleci.com/gh/hackclub/website)

Congratulations, you've stumbled upon Hack Club's website. Achievement get!

## Getting Started

Our website is developed using [Docker Compose](https://docs.docker.com/compose/). If you don't already have it installed, do that now.

Once you've cloned the repo, run the following commands to get the basics set up:

    $ docker-compose build
    $ docker-compose run web bundle install
    $ docker-compose run web rake db:create db:migrate

### Secrets

Secret keys are handled with [Figaro](https://github.com/laserlemon/figaro). Create `config/application.yml` containing the following keys and corresponding values.

- `SECRET_KEY_BASE`
- `STRIPE_PUBLISHABLE_KEY` - Publishable key for [Stripe](https://stripe.com/)
- `STRIPE_SECRET_KEY` - Secret key for [Stripe](https://stripe.com/)
- `GOOGLE_MAPS_KEY` - JavaScript API key for Google Maps

#### Production Secrets

The following secrets are also required for production.

- `SMTP_ADDRESS`
- `SMTP_PORT`
- `SMTP_USERNAME`
- `SMTP_PASSWORD`
- `SMTP_DOMAIN`
- `SENTRY_DSN` - DSN for Sentry (https://sentry.io)

#### Optional Variables

- `REFERRERS_TO_BLOCK`
  - A list of regular expressions of hostnames to block, separated by `|`. This is used to prevent [referrer spam](https://en.wikipedia.org/wiki/Referer_spam)
    - Examples: `.*\.spam\.com|\w+.spamwow.\w+`, `bad\.referrer\.com`

### Start server

    $ docker-compose up

The website should now be up and running!

## Deployment

If you choose to use the `Procfile` for deployment, you _must_ set the following environment variables to appropriate values:

- `PORT`
- `RAILS_ENV`

## License

```
The MIT License (MIT)

Copyright (c) 2015 - 2016 Hack Club

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
