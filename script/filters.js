// Affiche la valeur des inputs range
    function updateValue(val, id) {
    document.getElementById(id).innerText = val;
}


// Active automatiquement le filtrage lorsque une case est coché
    function submitForm() {
    document.getElementById("filterForm").submit();
}


// Ajoute toutes les cases
    var checkboxContainer = document.getElementById('checkboxContainer');
    for (var i = 1; i <= 8; i++) {
    var checkbox = document.createElement('input');
    checkbox.type = 'checkbox';
    checkbox.id = 'members' + i;
    checkbox.name = 'members';
    checkbox.value = i.toString();
    checkbox.onchange = submitForm;

    var label = document.createElement('label');
    label.htmlFor = 'members' + i;
    label.innerText = i.toString();

    checkboxContainer.appendChild(checkbox);
    checkboxContainer.appendChild(label);
}

// Ajoute toutes les options
    var locations = ["all", "usa", "japan", "new_zealand", "mexico", "french_polynesia", "new_caledonia", "uk", "switzerland", "france", "australia", "indonesia", "slovakia", "hungary", "belarus", "brazil", "argentina", "sweden", "belgium", "portugal", "spain", "colombia", "germany", "denmark", "qatar", "india", "united_arab_emirates", "netherlands", "canada", "finland", "chile", "peru", "costa_rica", "austria", "greece", "italy", "south_korea", "china", "taiwan", "thailand", "philippines", "poland", "norway"];
    var select = document.getElementById('location');

    for (var i = 0; i < locations.length; i++) {
    var option = document.createElement('option');
    option.value = locations[i];
    option.text = locations[i];
    select.appendChild(option);
}

var timeout = null;

function submitForm() {
    clearTimeout(timeout);
    timeout = setTimeout(function () {
        document.getElementById("filterForm").submit();
    }, 500);
}

// si le filtre n'existe pas on affiche un message

// Fonction pour afficher ou masquer l'image "Artiste non trouvé"
function toggleNotFoundImage() {
    var filteredContainer = document.getElementById("filteredArtistContainer");
    var notFoundImage = document.getElementById("notFoundImage");
    if (filteredContainer.children.length === 0) {
        notFoundImage.style.display = "block"; // Afficher l'image
    } else {
        notFoundImage.style.display = "none"; // Masquer l'image
    }
}

// Fonction pour soumettre le formulaire
function submitForm() {
    // Soumission du formulaire (vous devez implémenter cette partie en fonction de votre logique)
    var form = document.getElementById("filterForm");
    form.submit(); // Soumettre le formulaire

    // Appeler la fonction pour afficher l'image "Artiste non trouvé" après la soumission du formulaire
    toggleNotFoundImage();
}

// Appel initial pour afficher l'image "Artiste non trouvé" au chargement de la page
window.onload = toggleNotFoundImage;


// le bouton pour ouvrir le system de filtre

function toggleForm() {
    var form = document.getElementById("filterForm");
    var button = document.getElementById("toggleFormButton");
    if (form.style.display === "none") {
        form.style.display = "block";
        button.textContent = "Masquer les filtres";
    } else {
        form.style.display = "none";
        button.textContent = "Afficher les filtres";
    }
}
