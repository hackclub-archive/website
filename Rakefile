APP_FILE  = 'hackedu.rb'
APP_CLASS = 'HackEDU'

require 'sinatra/assetpack/rake'

namespace :generate do
  desc 'Generate a timestamped, empty Sequel migration.'
  task :migration, :name do |_, args|
    if args[:name].nil?
      puts 'You must specify a migration name (e.g. rake '\
        'generate:migration[create_events])!'
      exit false
    end

    content = "Sequel.migration do\n  up do\n  end\n\n  down do\n  "\
      "end\nend\n"
    timestamp = Time.now.to_i
    filename = File.join(File.dirname(__FILE__), 'migrations',
                         "#{timestamp}_#{args[:name]}.rb")

    File.open(filename, 'w') do |f|
      f.puts content
    end

    puts "Created the migration #{filename}"
  end
end

namespace :db do
  require 'sequel'
  require 'sequel/extensions/migration'
  require_relative 'models/init'

  desc 'Run database migrations'
  task :migrate, :env do |cmd, args|
    Sequel::Migrator.apply(HackEDU::DB, 'migrations')
  end
   
  desc 'Nuke the database (drop all tables)'
  task :nuke, :env do |cmd, args|
    HackEDU::DB.tables.each do |table|
      HackEDU::DB.run("DROP TABLE #{table}")
    end
  end
   
  desc 'Reset the database'
  task :reset, [:env] => [:nuke, :migrate]
end
