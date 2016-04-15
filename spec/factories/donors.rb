FactoryGirl.define do
  factory :donor do
    email { Faker::Internet.email }
    stripe_id { "cus_#{SecureRandom.hex(7)}" }
  end
end
