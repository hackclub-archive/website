class AppProxy
  def self.new(*_args)
    # Construct a Rack app that only runs the reverse proxy middleware and
    # does nothing else.
    Rack::Builder.new do
      use Rack::ReverseProxy do
        reverse_proxy '/', 'https://new.hackclub.com'
        reverse_proxy_options force_ssl: true
      end

      # Don't return for any request that somehow passes through the reverse
      # proxy middleware.
      run proc { |_env| nil }
    end
  end
end
