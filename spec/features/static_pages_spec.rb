require 'rails_helper'

describe 'Static pages' do
  subject { page }

  describe 'Home page' do
    before { visit root_path }

    it { should have_content('Start a Hack Club') }
    it { should have_title(full_title('')) }
  end

  describe 'Team page' do
    before { visit team_path }

    it { should have_content('Team') }
    it { should have_title(full_title('Team')) }
  end
end
