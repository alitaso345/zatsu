Rails.application.routes.draw do
  resources :entries do
    get :more, on: :collection
    patch :like, on: :member
  end
  resources :users
  # Define your application routes per the DSL in https://guides.rubyonrails.org/routing.html

  # Defines the root path route ("/")
  # root "articles#index"
end
