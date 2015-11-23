require 'rails_helper'

RSpec.describe ClubApplicationsController, type: :controller do
  describe 'POST #create' do
    let(:params) { attributes_for(:club_application) }
    let(:invalid_params) {
      attributes_for(:club_application_with_no_high_school) }

    context 'with valid fields' do
      it 'saves the new application to the database' do
        expect {
          post :create, club_application: params
        }.to change(ClubApplication, :count).by 1
      end

      it 'delivers a confirmation email to me and applicant' do
        expect {
          post :create, club_application: params
        }.to change(ActionMailer::Base.deliveries, :size).by 2

        emails = ActionMailer::Base.deliveries.last(2)
        my_email = 'team@hackclub.io'

        # make sure that one email was sent to the applicant and another to me
        expect(emails.find { |d| d.to.include? params[:email] }).to_not be nil
        expect(emails.find { |d| d.to.include? my_email }).to_not be nil
      end
    end

    context 'with invalid fields' do
      it 'does not save the application to the database' do
        expect {
          post :create, club_application: invalid_params
        }.to change(ClubApplication, :count).by 0
      end

      it 'renders a page with an error' do
        post :create, club_application: invalid_params

        expect(assigns(:application).errors.empty?).to_not be true
      end
    end
  end
end
