var FLTR = document.getElementById('FLTR'),
    SRCH = document.getElementById("search-choice");
FLTR.addEventListener('change', function() {
    if ( !FLTR.checked ) {
        SRCH.setAttribute("required", "");
    } else {
        SRCH.removeAttribute("required");
    }
});

document.getElementById('search-choice').addEventListener('input', function() {
    var inp = document.getElementById('search-choice').value.split(" -> ");
    if (inp[1] == "artist/band") {
        document.getElementById('categories').selectedIndex = 1;
        document.getElementById('search-choice').value = inp[0];
        return;

    };
    if (inp[1] == "members") {
        document.getElementById('categories').selectedIndex = 2;
        document.getElementById('search-choice').value = inp[0];
        return;

    };
    if (inp[1] == "creation date") {
        document.getElementById('categories').selectedIndex = 3;
        document.getElementById('search-choice').value = inp[0];
        return;

    };
    if (inp[1] == "first album") {
        document.getElementById('categories').selectedIndex = 4;
        document.getElementById('search-choice').value = inp[0];
        return;

    };
    if (inp[1] == "location") {
        document.getElementById('categories').selectedIndex = 5;
        document.getElementById('search-choice').value = inp[0];
        return;
    };
    document.getElementById('categories').selectedIndex = 0;
    document.getElementById('search-choice').value = inp[0];
});
    

function show(cb, filterID) {
    var checkBox = document.getElementById(cb);
    var input = document.getElementById(filterID);

    if (checkBox.checked == true){
        input.style.display = "grid";
    } else {
        input.style.display = "none";
    }
}

function toggle(first, sec, filterID) {

    var firstCD = document.getElementById(first),
    secondCD = document.getElementById( sec ),
    demoCD = document.getElementById(filterID);

    if ( firstCD.style.display === "grid" ) {

        firstCD.style.display = "none";
        secondCD.style.display = "grid";
        demoCD.style.display = "none";
        
    } else { // switch back
        firstCD.style.display = "grid";
        secondCD.style.display = "none";
        demoCD.style.display = "grid";

    }

}
