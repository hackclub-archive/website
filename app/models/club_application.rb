class ClubApplication < ActiveRecord::Base
  validates_presence_of :first_name, :last_name, :email, :high_school,
    :interesting_project, :system_hacked, :steps_taken
  validates_uniqueness_of :email
end
