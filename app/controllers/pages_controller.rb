class PagesController < ApplicationController
  def home
    @clubs_hash = Gmaps4rails.build_markers(Club.all) do |club, marker|
      marker.lat club.latitude
      marker.lng club.longitude
    end
  end

  def contact
  end

  def sponsor
  end

  def attributions
    @icons = [
      {
        name: 'Books',
        url: 'http://thenounproject.com/term/books/21509/',
        author: 'Piotrek Chuchla',
        author_url: 'http://www.piotrekchuchla.com'
      },
      {
        name: 'Community',
        url: 'http://thenounproject.com/term/community/5040/',
        author: 'Dmitry Baranovskiy',
        author_url: 'http://dmitry.baranovskiy.com'
      },
      {
        name: 'Open Source',
        url: 'http://thenounproject.com/term/open-source/8233/',
        author: 'Oriol Carbonell',
        author_url: 'http://www.hiperic.com'
      }
    ]
    @icons.map! { |i| OpenStruct.new(i) }
  end
end
