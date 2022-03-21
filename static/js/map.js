
// Map 2gis


function toggle(mapMode, defaultMode, artistCard_map, artistCard_content) {

    var mapMode = document.getElementById(mapMode),
    defaultMode = document.getElementById( defaultMode ),
    artistCard_map = document.getElementById( artistCard_map ),
    artistCard_content = document.getElementById(artistCard_content);

    if ( mapMode.style.display !== "none" ) {

        mapMode.style.display = "none";
        defaultMode.style.display = "grid";
        artistCard_content.style.display = "none";
        artistCard_map.style.display = "flex";
        
    } else { // switch back
        mapMode.style.display = "grid";
        defaultMode.style.display = "none";
        artistCard_content.style.display = "flex";
        artistCard_map.style.display = "none";

    }

}