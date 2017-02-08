class AddPhoneNumberToClubApplication < ActiveRecord::Migration
  def change
    add_column :club_applications, :phone_number, :string
  end
end
