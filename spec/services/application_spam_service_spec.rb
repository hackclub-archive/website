require 'rails_helper'

RSpec.describe ApplicationSpamService, type: :model do
  SHORT_ANSWER_LENGTH = 5

  let(:base_app) { create(:club_application) }
  let(:short_answer) { Faker::Lorem.words(SHORT_ANSWER_LENGTH).join(' ') }

  describe '#is_spam?' do
    it 'returns false with a legit application' do
      expect(ApplicationSpamService.new.is_spam?(base_app)).to eq(false)
    end

    it 'returns true with a spam application' do
      base_app.system_hacked = short_answer
      base_app.interesting_project = short_answer

      expect(ApplicationSpamService.new.is_spam?(base_app)).to eq(true)
    end
  end
end
