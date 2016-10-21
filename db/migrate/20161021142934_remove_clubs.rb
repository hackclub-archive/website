class RemoveClubs < ActiveRecord::Migration
  def up
    drop_table :clubs
  end

  def down
    create_table :clubs do |t|
      t.string :school
      t.float :latitude
      t.float :longitude

      t.timestamps
    end
  end
end
