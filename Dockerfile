FROM rails
MAINTAINER Zach Latta <zach@hackedu.us>

WORKDIR /usr/src/app

COPY Gemfile /usr/src/app/
COPY Gemfile.lock /usr/src/app/

RUN bundle install --system

ENV RAILS_ENV production

COPY . /usr/src/app/

RUN bundle exec rake assets:precompile

EXPOSE 8080

ENTRYPOINT ["unicorn", "-c", "config/unicorn.rb", "-E", "production"]
