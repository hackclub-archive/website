require 'sinatra/base'
require 'sinatra/assetpack'
require 'sinatra/partial'
require 'better_errors'
require 'sass'

configure :development do
  use BetterErrors::Middleware
  BetterErrors.application_root = File.expand_path('..', __FILE__)
end

class HackEDU < Sinatra::Base
  register Sinatra::AssetPack
  register Sinatra::Partial

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

  get '/' do
    sponsors = [
      {
        name: 'Test',
        logo: '/images/open_source.svg'
      },
      {
        name: 'Test 2',
        logo: '/images/open_source.svg'
      },
      {
        name: 'Test 3',
        logo: '/images/open_source.svg'
      }
    ]
    erb :index, locals: { sponsors: sponsors }
  end

  get '/contact' do
    erb :contact
  end

  get '/attributions' do
    icons = [
      {
        name: 'Books',
        url: 'http://thenounproject.com/term/books/21509/',
        author: 'Piotrek Chuchla',
        author_url: 'http://www.piotrekchuchla.com'
      },
      {
        name: 'Community',
        url: 'http://thenounproject.com/term/community/5040/',
        author: 'Dmitry Baranovskiy',
        author_url: 'http://dmitry.baranovskiy.com'
      },
      {
        name: 'Open Source',
        url: 'http://thenounproject.com/term/open-source/8233/',
        author: 'Oriol Carbonell',
        author_url: 'http://www.hiperic.com'
      }
    ]
    erb :attributions, locals: { icons: icons }
  end
end
