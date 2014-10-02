add_import_path "bower_components/foundation/scss"

http_path = "/"
css_dir = "public/stylesheets"
sass_dir = "scss"
images_dir = "images"
javascripts_dir = "js"

if Gem.loaded_specs["sass"].version > Gem::Version.create('3.3')
  warn "You're using Sass 3.4 or higher to compile Foundation. This version causes CSS classes to output incorrectly, so we recommend using Sass 3.3 or 3.2."
  warn "To use the right version of Sass on this project, run \"bundle\" and then use \"bundle exec compass watch\" to compile Foundation."
end
