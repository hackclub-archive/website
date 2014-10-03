# hackEDU

hackEDU is a nonprofit initiative that helps high school students start and
lead programming clubs at their schools. This is our website.

## Requirements

* [Docker](https://github.com/docker/docker)
* [Fig](https://github.com/docker/fig)

## Usage

After you've cloned the repo and gone into the project directory:

    $ fig up

Now run the database migrations

    $ fig run web bundle exec rake db:migrate

And you're done! The app is now accessible at http://localhost:9393.

#### Migrations

##### Generate Migration

Run the following rake task:

    $ bundle exec rake generate:migration[migration_name]

Example:

    $ bundle exec rake generate:migration[create_applications]

##### Run Migrations

    $ bundle exec rake db:migrate

##### Nuke Database (Drop All Tables)

    $ bundle exec rake db:nuke

## License

The MIT License (MIT)

Copyright (c) 2014 hackEDU

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
