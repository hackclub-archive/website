class AppProxy < Rack::Proxy
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
