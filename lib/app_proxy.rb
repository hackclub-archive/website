class AppProxy < Rack::Proxy
  # See https://github.com/ncr/rack-proxy/issues/50#issuecomment-256563048 for
  # the source of the implementations of call and rewrite_response.
  def call(env)
    @streaming = true
    super
  end

  def rewrite_response(triplet)
    status, headers, body = triplet
    headers.delete "transfer-encoding"
    triplet
  end

  def rewrite_env(env)
    env['HTTP_HOST'] = 'new.hackclub.com'
    env['SERVER_PORT'] = '443'
    env['HTTPS'] = 'on'

    # Always force the proxy to correspond to the "root" of the application.
    #
    # See http://www.rubydoc.info/github/rack/rack/file/SPEC#The_Environment for
    # an explanation on how SCRIPT_NAME works in Rack.
    env['SCRIPT_NAME'] = nil

    # Strip forwarding parameters, so the app doesn't know it's being accessed
    # through a reverse proxy
    env['HTTP_X_FORWARDED_SCHEME'] = nil
    env['HTTP_X_FORWARDED_PROTO'] = nil
    env['HTTP_X_FORWARDED_HOST'] = nil
    env['HTTP_X_FORWARDED_PORT'] = nil
    env['HTTP_X_FORWARDED_SSL'] = nil

    env
  end
end
