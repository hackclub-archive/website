# encoding: UTF-8
# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# Note that this schema.rb definition is the authoritative source for your
# database schema. If you need to create the application database on another
# system, you should be using db:schema:load, not running all the migrations
# from scratch. The latter is a flawed and unsustainable approach (the more migrations
# you'll amass, the slower it'll run and the greater likelihood for issues).
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 20150530235249) do

  create_table "club_applications", force: true do |t|
    t.string   "first_name",          null: false
    t.string   "last_name",           null: false
    t.string   "email",               null: false
    t.string   "github"
    t.string   "twitter"
    t.string   "high_school",         null: false
    t.text     "interesting_project", null: false
    t.text     "system_hacked",       null: false
    t.text     "steps_taken",         null: false
    t.datetime "created_at"
    t.datetime "updated_at"
    t.integer  "year"
    t.text     "referer"
  end

  add_index "club_applications", ["email"], name: "index_club_applications_on_email", unique: true

  create_table "clubs", force: true do |t|
    t.string   "school"
    t.float    "latitude"
    t.float    "longitude"
    t.datetime "created_at"
    t.datetime "updated_at"
  end

end
