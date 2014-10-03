require 'sequel'
require_relative '../lib/config'

module HackEDU
  DB = Sequel.connect(HackEDU::Config.database_url)
  DB << "SET CLIENT_ENCODING TO 'UTF8';"
end
