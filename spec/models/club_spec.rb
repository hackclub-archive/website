require 'rails_helper'

RSpec.describe Club, type: :model do
  it { should respond_to :school }
  it { should respond_to :latitude }
  it { should respond_to :longitude }

  it { should validate_presence_of :school }
  it { should validate_presence_of :latitude }
  it { should validate_presence_of :longitude }
end
