const xhttp = new XMLHttpRequest();
xhttp.onload = function() {
    var json = JSON.parse(xhr.responseText);
    console.log(json);
    document.querySelector('.first').innerHTML = json.students.FirstName;
}
xhttp.open("GET", "/student")
xhttp.send()