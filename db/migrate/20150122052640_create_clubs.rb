class CreateClubs < ActiveRecord::Migration
  def change
    create_table :clubs do |t|
      t.string :school
      t.float :latitude
      t.float :longitude

      t.timestamps
    end
  end
end
