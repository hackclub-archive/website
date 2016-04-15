class Donor < ActiveRecord::Base
  validates_presence_of :email, :stripe_id
  validates_uniqueness_of :email, :stripe_id
  validates_email_format_of :email
end
