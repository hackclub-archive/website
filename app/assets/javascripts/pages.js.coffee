@['pages#home'] = (data) ->
  handler = Gmaps.build('Google')
  handler.buildMap
    provider:
      center:
        lat: 39
        lng: -101
      zoom: 4
      scrollwheel: false
    internal:
      id: 'home-map-jumbotron'
    ->
      handler.addMarkers markersJSON
