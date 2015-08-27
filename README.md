<a name="top"></a>
<a href="https://www.irccloud.com/invite?channel=%23hackedu&amp;hostname=irc.freenode.net&amp;port=6697&amp;ssl=1" target="_blank"><img src="https://www.irccloud.com/invite-svg?channel=%23hackedu&amp;hostname=irc.freenode.net&amp;port=6697&amp;ssl=1"  height="18"></a> [![Circle CI](https://circleci.com/gh/hackedu/website.svg?style=svg)](https://circleci.com/gh/hackedu/website)

-------------------------------------------------------------------------------

<p align="center"><img src="https://raw.githubusercontent.com/hackedu/dinosaurs/68ccf2b66be441748ee0639df01deb3ea354cfc7/code_dinosaur.png" alt="hackEDU Website" /></p>
<h1 align="center">hackEDU's Website</h1>

You've stumbled upon hackEDU's website.

## Getting Started

### Prerequisites

*  Ruby 2.0+ with Rails

### Install dependencies

After cloning the repo:

    $ bundle install

### Database migrations

    $ rake db:migrate

### Secrets

Secret keys are handled with [Figaro](https://github.com/laserlemon/figaro).
Create `config/application.yml` containing the following keys and 
corresponding values.

* `SECRET_KEY_BASE`

#### Production Secrets

The following secrets are also required for production.

* `SMTP_ADDRESS`
* `SMTP_PORT`
* `SMTP_USERNAME`
* `SMTP_PASSWORD`
* `SMTP_DOMAIN`

#### Optional Variables

* `REFERRERS_TO_BLOCK`
  * A list of regular expressions of hostnames to block, separated by `|`. This
    is used to prevent
    [referrer spam](https://en.wikipedia.org/wiki/Referer_spam)
    * Examples: `.*\.spam\.com|\w+.spamwow.\w+`, `bad\.referrer\.com`

### Start server

    $ rails serve

The website should now be running on your computer at http://localhost:3000.

## Deployment

If you choose to use the `Procfile` for deployment, you _must_ set the
following environment variables to appropriate values:

* `PORT`
* `RACK_ENV`

## License

The MIT License (MIT)

Copyright (c) 2015 hackEDU

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
