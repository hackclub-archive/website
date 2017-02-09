require 'rails_helper'

RSpec.describe ClubApplicationMailer, type: :mailer do
  include EmailSpec::Helpers
  include EmailSpec::Matchers

  shared_examples 'a club application email' do
    it 'contains all of the fields in the application' do
      application.attributes.each_pair do |name, field|
        unless ['id', 'created_at', 'updated_at', 'year'].include? name
          expect(subject).to have_body_text(field)
        end

        if name == 'year' # check for match for year's corresponding string
          year_text = ClubApplication.years.key(field)
          expect(subject).to have_body_text(/#{year_text}/)
        end
      end
    end

    it 'does not contain the id, created_at, or updated_at fields' do
      [/id:/i, /created at:/i, /updated at:/i].each do |f|
        expect(subject).to_not have_body_text(f)
      end
    end
  end

  let(:application) { create(:club_application) }

  describe 'applicant confirmation' do
    subject { ClubApplicationMailer.applicant_confirmation(application) }

    it { should have_subject 'Application Confirmation' }
    it { should deliver_to application.mail_address.format }
    it { should deliver_from 'Zach Latta <team@hackclub.com>' }

    it_behaves_like 'a club application email'
  end

  describe 'admin notification' do
    subject { ClubApplicationMailer.admin_notification(application) }

    it { should have_subject 'Hack Club Application' }
    it { should deliver_to 'Zach Latta <team@hackclub.com>' }
    it { should deliver_from application.mail_address.format }

    it_behaves_like 'a club application email'
  end
end
