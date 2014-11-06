class CreateClubApplications < ActiveRecord::Migration
  def change
    create_table :club_applications do |t|
      t.string :first_name, null: false
      t.string :last_name, null: false
      t.string :email, null: false
      t.index :email, unique: true
      t.string :github
      t.string :twitter
      t.string :high_school, null: false
      t.text :interesting_project, null: false
      t.text :system_hacked, null: false
      t.text :steps_taken, null: false

      t.timestamps
    end
  end
end
