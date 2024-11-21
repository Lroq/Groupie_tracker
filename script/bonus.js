// Ecoutez l'événement de changement sur l'élément de case à cocher
document.querySelector('.switch input').addEventListener('change', function() {
  // Changez la couleur de fond lorsque l'état de la case à cocher change
  document.body.style.backgroundColor = this.checked ? '' : '#2F303D';

  // Sauvegardez l'état de la case à cocher dans un cookie
  document.cookie = 'switchState=' + this.checked + '; path=/';
});

// recupere l'etat de la page grace au cookie
window.addEventListener('load', function() {
  var matches = document.cookie.match(/switchState=(true|false)/);
  if (matches) {
    var switchState = matches[1] === 'true' ? true : false;
    document.querySelector('.switch input').checked = switchState;
    document.body.style.backgroundColor = switchState ? '' : '#2F303D';
  }
});
