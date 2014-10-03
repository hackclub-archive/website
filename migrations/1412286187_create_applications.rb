Sequel.migration do
  up do
    create_table :applications do
      primary_key :id
      String :first_name, null: false
      String :last_name, null: false
      String :email, null: false, unique: true
      String :github
      String :twitter
      String :password, null: false
      String :high_school, null: false
      String :interesting_project, null: false
      String :system_hacked, null: false
      String :steps_taken, null: false
    end
  end

  down do
    drop_table :applications
  end
end
