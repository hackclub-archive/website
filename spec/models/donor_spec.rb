require 'rails_helper'

RSpec.describe Donor, :type => :model do
  let(:donor) { create(:donor) }

  subject { donor }

  it { should respond_to :email }
  it { should respond_to :stripe_id }

  it { should validate_presence_of :email }
  it { should validate_presence_of :stripe_id }

  it { should validate_email_format_of(:email) }

  it { should validate_uniqueness_of :email }
  it { should validate_uniqueness_of :stripe_id }
end
