class ClubApplicationMailer < ActionMailer::Base
  default from: 'hackEDU Team <team@hackedu.us>'

  def applicant_confirmation(application)
    @application = application

    to = Mail::Address.new @application.email
    to.display_name = "#{@application.first_name} #{@application.last_name}"
    mail(to: to.format, subject: 'Application Confirmation')
  end

  def admin_notification(application)
    @application = application

    to = Mail::Address.new 'team@hackedu.us'
    to.display_name = 'hackEDU Team'
    mail(to: to.format, from: @application.mail_address.format,
         subject: 'hackEDU Application')
  end
end
