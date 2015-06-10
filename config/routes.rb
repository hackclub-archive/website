Rails.application.routes.draw do
  root to: 'pages#home'

  get 'apply', to: 'club_applications#new'
  post 'apply', to: 'club_applications#create'

  %w[sponsor attributions team].each do |page|
    get page, controller: 'pages', action: page
  end

  # Error page routes
  %w( 404 500 ).each do |code|
    get code, to: "pages#error", code: code
  end

end
