class ClubApplicationMailer < ActionMailer::Base
  default from: 'Hack Club Team <team@hackclub.com>'

  def applicant_confirmation(application)
    @application = application

    to = Mail::Address.new @application.email
    to.display_name = @application.full_name
    mail(to: to.format, subject: 'Application Confirmation')
  end

  def admin_notification(application)
    @application = application

    to = Mail::Address.new 'team@hackclub.com'
    to.display_name = 'Hack Club Team'

    subject = "Hack Club Application (#{@application.full_name}, #{@application.high_school})"

    mail(to: to.format, reply_to: @application.mail_address.format,
         subject: subject)
  end
end
