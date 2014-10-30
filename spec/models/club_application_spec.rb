require 'rails_helper'

RSpec.describe ClubApplication, type: :model do
  it { should respond_to :first_name }
  it { should respond_to :last_name }
  it { should respond_to :email }
  it { should respond_to :github }
  it { should respond_to :twitter }
  it { should respond_to :high_school }
  it { should respond_to :interesting_project }
  it { should respond_to :system_hacked }
  it { should respond_to :steps_taken }

  it { should validate_presence_of :first_name }
  it { should validate_presence_of :last_name }
  it { should validate_presence_of :email }
  it { should validate_presence_of :high_school }
  it { should validate_presence_of :interesting_project }
  it { should validate_presence_of :system_hacked }
  it { should validate_presence_of :steps_taken }

  it { create(:club_application); should validate_uniqueness_of :email }
end
