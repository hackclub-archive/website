FROM rails:4.1.1
MAINTAINER Zach Latta <zach@hackedu.us>

WORKDIR /usr/src/app

COPY Gemfile /usr/src/app/
COPY Gemfile.lock /usr/src/app/

RUN bundle install --system

ENV RAILS_ENV production

COPY . /usr/src/app/

RUN bundle exec rake assets:precompile

EXPOSE 8080

CMD ["unicorn", "-c", "config/unicorn.rb", "-E", "production"]
