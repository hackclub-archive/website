FROM ruby:2.1.3-onbuild
MAINTAINER Zach Latta <zach@hackedu.us>
CMD ["bundle", "exec", "shotgun"]
