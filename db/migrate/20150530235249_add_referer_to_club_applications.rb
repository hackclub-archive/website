class AddRefererToClubApplications < ActiveRecord::Migration
  def change
    add_column :club_applications, :referer, :text
  end
end
