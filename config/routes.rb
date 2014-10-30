Rails.application.routes.draw do
  root to: 'pages#home'

  %w[home contact attributions].each do |page|
    get page, controller: 'pages', action: page
  end
end
