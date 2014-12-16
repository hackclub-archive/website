class AddYearToClubApplications < ActiveRecord::Migration
  def change
    add_column :club_applications, :year, :integer
  end
end
