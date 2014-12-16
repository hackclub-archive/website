class ClubApplication < ActiveRecord::Base
  validates_presence_of :first_name, :last_name, :email, :high_school, :year,
    :interesting_project, :system_hacked, :steps_taken
  validates_uniqueness_of :email
  validates_email_format_of :email, message: 'not a valid email'

  enum year: [:freshman, :sophomore, :junior, :senior]

  def mail_address
    expected = Mail::Address.new email
    expected.display_name = "#{first_name} #{last_name}"
    return expected
  end
end
