require 'rails_helper'
require 'validates_email_format_of/rspec_matcher'

RSpec.describe ClubApplication, type: :model do
  let(:club_application) { create(:club_application) }

  subject { club_application }

  it { should respond_to :first_name }
  it { should respond_to :last_name }
  it { should respond_to :email }
  it { should respond_to :high_school }
  it { should respond_to :year }
  it { should respond_to :github }
  it { should respond_to :twitter }
  it { should respond_to :interesting_project }
  it { should respond_to :system_hacked }
  it { should respond_to :steps_taken }
  it { should respond_to :referer }

  it { should validate_presence_of :first_name }
  it { should validate_presence_of :last_name }
  it { should validate_presence_of :email }
  it { should validate_presence_of :high_school }
  it { should validate_presence_of :year }
  it { should validate_presence_of :interesting_project }
  it { should validate_presence_of :system_hacked }
  it { should validate_presence_of :steps_taken }
  it { should validate_presence_of :referer }

  it { should validate_email_format_of(:email)
                .with_message('not a valid email') }

  it { create(:club_application); should validate_uniqueness_of :email }

  describe '#mail_address' do
    it 'returns a correctly formatted mail address' do
      expected =
        "#{subject.first_name} #{subject.last_name} <#{subject.email}>"
      expect(subject.mail_address.format).to eq expected
    end
  end
end
