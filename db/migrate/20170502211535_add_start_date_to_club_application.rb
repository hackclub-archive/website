class AddStartDateToClubApplication < ActiveRecord::Migration
  def change
    add_column :club_applications, :start_date, :text
  end
end
