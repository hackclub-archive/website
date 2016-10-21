class PagesController < ApplicationController
  def home
    @clubs = get_clubs
    @clubs_hash = clubs_markers(@clubs)
  end

  def sponsor
    @clubs_hash = clubs_markers(get_clubs)
  end

  def how_it_works
  end

  def team
  end

  private

  def get_clubs
    conn = Faraday.new(url: "https://api.hackclub.com")
    resp = conn.get "/v1/clubs"

    JSON.parse(resp.body, symbolize_names: true)
  end

  def clubs_markers(clubs)
    Gmaps4rails.build_markers(clubs) do |club, marker|
      marker.lat club[:latitude]
      marker.lng club[:longitude]
    end
  end
end
