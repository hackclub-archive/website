class PagesController < ApplicationController
  def home
    @clubs_hash = clubs_markers(Club.all)
  end

  def sponsor
    @clubs_hash = clubs_markers(Club.all)
  end

  def how_it_works
  end

  def team
  end

  private

  def clubs_markers(clubs)
    Gmaps4rails.build_markers(clubs) do |club, marker|
      marker.lat club.latitude
      marker.lng club.longitude
    end
  end
end
