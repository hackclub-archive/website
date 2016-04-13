require 'rails_helper'

describe 'Donation page' do
  subject { page }

  before { visit donate_path }

  it { should have_content('Donate') }
  it { should have_title(full_title('Donate')) }
end
