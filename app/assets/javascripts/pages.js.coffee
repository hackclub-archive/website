@['pages#home'] = (data) ->
  # The markers: clusterer: undefined argument makes it so items on the map don't cluster.
  #
  # See docs at https://github.com/apneadiving/Google-Maps-for-Rails/wiki/Customized-Cluster-Icons#disable-clusterer-icons
  handler = Gmaps.build('Google', markers: clusterer: undefined)
  handler.buildMap
    provider:
      center:
        lat: 40
        lng: -95
      zoom: 4
      scrollwheel: false
      styles:
        [
          {
            'featureType': 'administrative'
            'elementType': 'all'
            'stylers': [
              { 'hue': '#000000' }
              { 'lightness': -100 }
              { 'visibility': 'off' }
            ]
          }
          {
            'featureType': 'landscape'
            'elementType': 'geometry'
            'stylers': [
              { 'hue': '#dddddd' }
              { 'saturation': -100 }
              { 'lightness': -3 }
              { 'visibility': 'on' }
            ]
          }
          {
            'featureType': 'landscape'
            'elementType': 'labels'
            'stylers': [
              { 'hue': '#000000' }
              { 'saturation': -100 }
              { 'lightness': -100 }
              { 'visibility': 'off' }
            ]
          }
          {
            'featureType': 'poi'
            'elementType': 'all'
            'stylers': [
              { 'hue': '#000000' }
              { 'saturation': -100 }
              { 'lightness': -100 }
              { 'visibility': 'off' }
            ]
          }
          {
            'featureType': 'road'
            'elementType': 'geometry'
            'stylers': [
              { 'hue': '#bbbbbb' }
              { 'saturation': -100 }
              { 'lightness': 26 }
              { 'visibility': 'on' }
            ]
          }
          {
            'featureType': 'road'
            'elementType': 'labels'
            'stylers': [
              { 'hue': '#ffffff' }
              { 'saturation': -100 }
              { 'lightness': 100 }
              { 'visibility': 'off' }
            ]
          }
          {
            'featureType': 'road.local'
            'elementType': 'all'
            'stylers': [
              { 'hue': '#ffffff' }
              { 'saturation': -100 }
              { 'lightness': 100 }
              { 'visibility': 'on' }
            ]
          }
          {
            'featureType': 'transit'
            'elementType': 'labels'
            'stylers': [
              { 'hue': '#000000' }
              { 'lightness': -100 }
              { 'visibility': 'off' }
            ]
          }
          {
            'featureType': 'water'
            'elementType': 'geometry'
            'stylers': [
              { 'hue': '#ffffff' }
              { 'saturation': -100 }
              { 'lightness': 100 }
              { 'visibility': 'on' }
            ]
          }
          {
            'featureType': 'water'
            'elementType': 'labels'
            'stylers': [
              { 'hue': '#000000' }
              { 'saturation': -100 }
              { 'lightness': -100 }
              { 'visibility': 'off' }
            ]
          }
        ]
    internal:
      id: 'home-map-jumbotron'
    ->
      handler.addMarkers markersJSON

      google.maps.event.addDomListener document.getElementById('focus-map-usa'), 'click', ->
        handler.getMap().setCenter new (google.maps.LatLng)(40, -95)
        handler.getMap().setZoom 4

      google.maps.event.addDomListener document.getElementById('focus-map-international'), 'click', ->
        handler.getMap().setCenter new (google.maps.LatLng)(32, 0)
        handler.getMap().setZoom 2

  $ ->
    $('a[href*="#"]:not([href="#"]):not([href="#map-usa"]):not([href="#map-international"])').click ->
      target = undefined
      if location.pathname.replace(/^\//, '') == @pathname.replace(/^\//, '') and location.hostname == @hostname
        target = $(@hash)
        target = if target.length then target else $('[name=' + @hash.slice(1) + ']')
        if target.length
          $('html, body').animate { scrollTop: target.offset().top }, 1000
          return false
      return
    return

@['pages#sponsor'] = (data) ->
  handler = Gmaps.build('Google')
  handler.buildMap
    provider:
      center:
        lat: 39
        lng: -101
      zoom: 3
      scrollwheel: false
    internal:
      id: 'sponsor-clubs-map'
    ->
      handler.addMarkers markersJSON
