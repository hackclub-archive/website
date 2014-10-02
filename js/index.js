var schools = [
  {
    name: "Austin High School",
    lat: 30.27382,
    lng: -97.76745
  },
  {
    name: "Thunderridge High School",
    lat: 39.53532,
    lng: -105.01084
  },
  {
    name: "Highland Park High School",
    lat: 42.193905,
    lng: -87.802505
  },
  {
    name: "El Segundo High School",
    lat: 33.92459,
    lng: -118.41462
  },
  {
    name: "Gig Harbor High School",
    lat: 47.330853,
    lng: -122.60542
  },
  {
    name: "Mundelein High School",
    lat: 42.26986,
    lng: -88.02085
  },
  {
    name: "International Academy East",
    lat: 42.568012,
    lng: -83.12207
  }
]

function initialize() {
  var mapOptions = {
    center: { lat: 39, lng: -101 },
    zoom: 4,
    disableDefaultUI: true,
    scrollwheel: false
  };

  var map = new google.maps.Map(
    document.getElementById('landing-map-jumbotron'), mapOptions);

  for (var i = 0; i < schools.length; i++) {
    var school = schools[i];
    var latLng = new google.maps.LatLng(school.lat, school.lng);
    var marker = new google.maps.Marker({
      position: latLng,
      title: school.name
    });

    marker.setMap(map);
  }
}

google.maps.event.addDomListener(window, 'load', initialize)
