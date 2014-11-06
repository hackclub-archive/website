class ClubApplication < ActiveRecord::Base
  validates_presence_of :first_name, :last_name, :email, :high_school,
    :interesting_project, :system_hacked, :steps_taken
  validates_uniqueness_of :email

  def mail_address
    expected = Mail::Address.new email
    expected.display_name = "#{first_name} #{last_name}"
    return expected
  end
end
