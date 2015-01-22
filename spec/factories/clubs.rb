FactoryGirl.define do
  factory :club do
    school { "#{ Faker::Address.city } High School" }
    latitude { Faker::Address.longitude }
    longitude { Faker::Address.latitude }
  end
end
