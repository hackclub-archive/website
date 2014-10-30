FactoryGirl.define do
  factory :club_application do
    first_name { Faker::Name.first_name }
    last_name { Faker::Name.last_name }
    email { Faker::Internet.email }
    github { Faker::Internet.user_name }
    twitter { Faker::Internet.user_name }
    high_school { "#{ Faker::Address.city } High School" }
    interesting_project { Faker::Lorem.paragraph }
    system_hacked { Faker::Lorem.paragraph }
    steps_taken { Faker::Lorem.paragraph }
  end
end
