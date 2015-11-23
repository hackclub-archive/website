require 'rails_helper'

describe ApplicationHelper do
  describe 'full_title' do
    it 'should include the page title' do
      expect(full_title('foo')).to match(/foo/)
    end

    it 'should inclide the base title' do
      expect(full_title('foo')).to match(/^Hack Club/)
    end

    it 'should not include a pipe for the home page' do
      expect(full_title('')).not_to match(/\|/)
    end
  end
end
