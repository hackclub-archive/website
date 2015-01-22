class Club < ActiveRecord::Base
  validates_presence_of :school, :latitude, :longitude
end
