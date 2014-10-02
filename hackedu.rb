require 'sinatra/base'
require 'sinatra/assetpack'
require 'sinatra/partial'
require 'better_errors'
require 'sass'

require_relative 'routes/landing'

configure :development do
  use BetterErrors::Middleware
  BetterErrors.application_root = File.expand_path('..', __FILE__)
end

module HackEDU
  class Application < Sinatra::Base
    register Sinatra::AssetPack
    register Sinatra::Partial

    register HackEDU::Routes::Landing

    set :partial_template_engine, :erb

    assets do
      serve '/js', from: 'js'
      serve '/scss', from: 'scss'
      serve '/bower_components', from: 'bower_components'

      js :modernizr, [
        '/bower_components/modernizr/modernizr.js'
      ]

      js :libs, [
        '/bower_components/jquery/dist/jquery.js',
        '/bower_components/foundation/js/foundation.js'
      ]

      js :application, [
        '/js/app.js'
      ]

      js :landing_map_jumbotron, [
        'https://maps.googleapis.com/maps/api/js',
        '/js/index.js'
      ]

      css :application, [
        '/scss/foundation_and_overrides.css',
        '/scss/main.css'
      ]

      js_compression :jsmin
      css_compression :scss
    end
  end
end
