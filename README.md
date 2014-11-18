# hackedu [![Build Status](https://travis-ci.org/hackedu/hackedu.svg?branch=master)](https://travis-ci.org/hackedu/hackedu)

<img src="http://i.imgur.com/zHbn6N2.png" alt="hackEDU logo" height="60"
  align="right">
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
For Grasp to run on your system, you must create `config/application.yml`
containing the following keys and corresponding values.

* `SECRET_KEY_BASE`

#### Production Secrets

The following secrets are also required for production.

* `SMTP_ADDRESS`
* `SMTP_PORT`
* `SMTP_USERNAME`
* `SMTP_PASSWORD`
* `SMTP_DOMAIN`

### Start server

    $ rails serve

The website should now be running on your computer at http://localhost:3000.

## License

The MIT License (MIT)

Copyright (c) 2014 hackEDU

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
