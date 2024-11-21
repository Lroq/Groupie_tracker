/*gerer le affichage bleu sur les sections*/
console.log('artist.js loaded');
window.onload = function() {
    const sections = document.querySelectorAll('section');
    sections.forEach(section => {
        section.addEventListener('mouseover', function() {
            this.style.backgroundColor = '#3498db';
            this.style.color = 'white';
        });
        section.addEventListener('mouseout', function() {
            this.style.backgroundColor = '';
            this.style.color = '';
        });
    });
};

/*gerer le click sur les membres du groupe*/
const membersList = document.getElementById('membersList');
    const members = membersList.getElementsByTagName('li');
    for (let i = 0; i < members.length; i++) {
        const member = members[i];
        const memberName = member.textContent;
        const wikipediaURL = 'https://fr.wikipedia.org/wiki/' + encodeURI(memberName);
        member.innerHTML = '<a href="' + wikipediaURL + '" target="_blank">' + memberName + '</a>';
    };

/* gerer le bouton spotify*/
const artistName = document.getElementById('name').textContent;
const spotifyURL = 'https://open.spotify.com/search/' + encodeURI(artistName);
document.getElementById('spotifyImage').outerHTML = '<a href="' + spotifyURL + '" target="_blank"><img id="spotifyImage" src="../static/assets/spotify.jpg" alt="Spotify"></a>';

/*gerer la map*/
    var map = L.map('map', {
        center: [43.61045120361589, 1.4322067659647386],
        zoom: 1,
        preferCanvas: true
    });
    L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
    }).addTo(map);

    async function geocode(locationNames, callback) {
        try {
            const locations = await Promise.all(locationNames.map(async locationName => {
                const geocodingApiUrl = `https://nominatim.openstreetmap.org/search?format=json&limit=1&q=${encodeURIComponent(locationName)}`;
                const response = await fetch(geocodingApiUrl);
                const data = await response.json();
                if (data.length > 0) {
                    return data[0];
                } else {
                    console.error(`No results for location: ${locationName}`);
                    return null;
                }
            }));

            locations.forEach(location => {
                if (location) {
                    callback(location.lat, location.lon, locationNames[locations.indexOf(location)]);
                }
            });
        } catch (error) {
            console.error(`Error with geocoding API: ${error}`);
        }
    }
    
    // Usage:
    geocode(reflocation, function(lat, lon, city) {
        L.marker([lat, lon]).addTo(map)
            .bindPopup(`A popup for ${city} with coordinates: ${lat}, ${lon}`)
            .openPopup();
            console.log(lat, lon, city);
    });