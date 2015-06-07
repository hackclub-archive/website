Rails.application.routes.draw do
  root to: 'pages#home'

  get 'apply', to: 'club_applications#new'
  post 'apply', to: 'club_applications#create'

  %w[sponsor attributions team].each do |page|
    get page, controller: 'pages', action: page
  end
end
