FROM rails:4.2.4
MAINTAINER Zach Latta <zach@hackedu.us>

WORKDIR /usr/src/app

COPY Gemfile /usr/src/app/
COPY Gemfile.lock /usr/src/app/

RUN bundle install -j4 --system

COPY . /usr/src/app/

RUN RAILS_ENV=production bundle exec rake assets:precompile

EXPOSE 3000

CMD ["bundle", "exec", "puma", "-C", "config/puma.rb"]
