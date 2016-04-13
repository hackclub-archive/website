Rails.application.routes.draw do
  root to: 'pages#home'

  get 'apply', to: 'club_applications#new'
  post 'apply', to: 'club_applications#create'

  resources :donations, only: [:new, :create]

  %w[sponsor team].each do |page|
    get page, controller: 'pages', action: page
  end
end
