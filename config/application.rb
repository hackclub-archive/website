require File.expand_path('../boot', __FILE__)

require 'rails/all'

# Require the gems listed in Gemfile, including any gems
# you've limited to :test, :development, or :production.
Bundler.require(*Rails.groups)

module HackClub
  class Application < Rails::Application
    # Redirect HTTP to HTTPS in production
    config.middleware.use Rack::SslEnforcer, only_environments: 'production'

    # Block referer spam
    config.middleware.use Rack::Attack

    # Compress Rails-generated responses
    config.middleware.use Rack::Deflater

    # Also use Rack::Deflater for runtime asset compression
    config.middleware.insert_before ActionDispatch::Static, Rack::Deflater

    # Autoload files from lib/
    config.autoload_paths << Rails.root.join('lib')

    # Settings in config/environments/* take precedence over those specified here.
    # Application configuration should go into files in config/initializers
    # -- all .rb files in that directory are automatically loaded.

    # Set Time.zone default to the specified zone and make Active Record auto-convert to this zone.
    # Run "rake -D time" for a list of tasks for finding time zone names. Default is UTC.
    # config.time_zone = 'Central Time (US & Canada)'

    # The default locale is :en and all translations from config/locales/*.rb,yml are auto loaded.
    # config.i18n.load_path += Dir[Rails.root.join('my', 'locales', '*.{rb,yml}').to_s]
    # config.i18n.default_locale = :de
  end
end
