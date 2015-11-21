class ClubApplicationMailer < ActionMailer::Base
  default from: 'Zach Latta <team@hackclub.io>'

  def applicant_confirmation(application)
    @application = application

    to = Mail::Address.new @application.email
    to.display_name = "#{@application.first_name} #{@application.last_name}"
    mail(to: to.format, subject: 'Application Confirmation')
  end

  def admin_notification(application)
    @application = application

    to = Mail::Address.new 'team@hackclub.io'
    to.display_name = 'Zach Latta'
    mail(to: to.format, from: @application.mail_address.format,
         subject: 'Hack Club Application')
  end
end
