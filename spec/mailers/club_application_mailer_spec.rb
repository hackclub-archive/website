require 'rails_helper'

RSpec.describe ClubApplicationMailer, type: :mailer do
  include EmailSpec::Helpers
  include EmailSpec::Matchers
  include ActionMailer::MailHelper 

  shared_examples 'a club application email' do
    it 'contains all of the fields in the application' do
      application.attributes.each_pair do |name, field|
        unless ['id', 'created_at', 'updated_at'].include? name
          expect(subject).to have_body_text(/#{block_format field}/) 
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
    it { should deliver_from 'Zach Latta <zach@hackedu.us>' }

    it_behaves_like 'a club application email'
  end

  describe 'admin notification' do
    subject { ClubApplicationMailer.admin_notification(application) }

    it { should have_subject 'hackEDU Application' }
    it { should deliver_to 'Zach Latta <zach@hackedu.us>' }
    it { should deliver_from application.mail_address.format }

    it_behaves_like 'a club application email'
  end
end
