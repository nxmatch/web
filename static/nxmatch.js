var matchTemplate = document.getElementById("match-template"),
    xmlhttp       = new XMLHttpRequest();

xmlhttp.open("GET", "/list", true);

xmlhttp.onreadystatechange = function() {

    if ( xmlhttp.readyState != 4 || xmlhttp.status != 200 ) {
        return;
    }

    var response = JSON.parse(xmlhttp.responseText),
        results  = document.getElementById("results");

    for ( var i = 0 ; i < response.length ; i++ ) {
        var match       = response[i],
            score       = match.Scores[match.Scores.length-1],
            currentNode = matchTemplate.cloneNode(true);

        currentNode.querySelector(".time").innerHTML     = match.Time;
        currentNode.querySelector(".location").innerHTML = match.Location;
        currentNode.querySelector(".home").innerHTML     = score.Home.Name;
        currentNode.querySelector(".visitor").innerHTML  = score.Visitor.Name;
        currentNode.querySelector(".score").innerHTML    = score.HomeScore + " - " + score.VisitorScore;
        results.appendChild(currentNode);
    }

}
xmlhttp.send();