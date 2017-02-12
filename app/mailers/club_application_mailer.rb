class ClubApplicationMailer < ActionMailer::Base
  default from: 'Hack Club Team <team@hackclub.com>'

  def applicant_confirmation(application)
    @application = application

    to = Mail::Address.new @application.email
    to.display_name = "#{@application.first_name} #{@application.last_name}"
    mail(to: to.format, subject: 'Application Confirmation')
  end

  def admin_notification(application)
    @application = application

    to = Mail::Address.new 'team@hackclub.com'
    to.display_name = 'Hack Club Team'
    mail(to: to.format, from: @application.mail_address.format,
         subject: 'Hack Club Application')
  end
end
