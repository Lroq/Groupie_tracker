/* search bar sur index.html*/

function searchArtist() {
    var input = document.getElementsByClassName('input')[0].value.toLowerCase();
    var cards = document.getElementsByClassName('card');

    var found = false;

    for (var i = 0; i < cards.length; i++) {
        var artistName = cards[i].getAttribute('data-artist').toLowerCase();
        if (artistName.includes(input)) {
            cards[i].style.display = '';
            found = true;
        } else {
            cards[i].style.display = 'none';
        }
    }

    var errorMessage = document.getElementById('error-message');
    if (!found) {
        errorMessage.style.display = 'block';
    } else {
        errorMessage.style.display = 'none';
    }
}
