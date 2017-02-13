class ClubApplication < ActiveRecord::Base
  validates_presence_of :first_name, :last_name, :email, :high_school, :year,
    :interesting_project, :system_hacked, :steps_taken, :referer
  validates_uniqueness_of :email
  validates_email_format_of :email, message: 'not a valid email'

  enum year: {
    below_high_school: 4,
    nine: 0,
    ten: 1,
    eleven: 2,
    twelve: 3,
    college_student: 5,
    teacher: 6,
    parent_or_guardian: 7,
    other: 8
  }

  def mail_address
    expected = Mail::Address.new email
    expected.display_name = full_name
    return expected
  end

  def full_name
    "#{first_name} #{last_name}"
  end
end
