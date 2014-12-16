class ClubApplicationsController < ApplicationController
  def new
    @application = ClubApplication.new
  end

  def create
    @application = ClubApplication.new(club_application_params)
    if @application.save
      ClubApplicationMailer.applicant_confirmation(@application).deliver
      ClubApplicationMailer.admin_notification(@application).deliver

      flash[:success] = 'Application submitted successfully!'
      redirect_to apply_path
    else
      render :new
    end
  end

  private

  def club_application_params
    params.require(:club_application).permit(:first_name, :last_name, :email,
                                             :github, :twitter, :high_school,
                                             :year, :interesting_project,
                                             :system_hacked, :steps_taken)
  end
end
