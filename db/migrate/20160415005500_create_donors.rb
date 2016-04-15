class CreateDonors < ActiveRecord::Migration
  def change
    create_table :donors do |t|
      t.string :email, index: true
      t.string :stripe_id, index: true

      t.timestamps
    end
  end
end
