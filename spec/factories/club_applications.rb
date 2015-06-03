FactoryGirl.define do
  factory :club_application do
    first_name { Faker::Name.first_name }
    last_name { Faker::Name.last_name }
    email { Faker::Internet.email }
    high_school { "#{ Faker::Address.city } High School" }
    year { ClubApplication.years.keys.sample.to_sym }
    github { Faker::Internet.user_name }
    twitter { Faker::Internet.user_name }
    interesting_project { Faker::Lorem.paragraph }
    system_hacked { Faker::Lorem.paragraph }
    steps_taken { Faker::Lorem.paragraph }
    referer { Faker::Lorem.paragraph }

    factory :club_application_with_no_high_school do
      high_school ''
    end
  end
end
