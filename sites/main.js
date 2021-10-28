const url_students = "/student";

// GET data
async function getapi(url)  {
    // store response
    const response = await fetch(url)
    // store data in json
    var data = await response.json();
    console.log(data);
    show(data);
}
getapi(url_students);
function show(data) {
    let tab = ``;
    for (let student of data) {
        tab += `
        <tr>
        <th scope="row">${student.id}</th>
        <td>${student.firstname} </td>
        <td>${student.lastname} </td>
        <td>${student.age} </td>
        <td>${student.grade} </td>
        <td><button type="button" class="btn btn-outline-primary">Edit</button></td>
      <td><button type="button" id="btn_delete" class="btn btn-outline-danger">Delete</button></td>
        </tr>`;
    }
    document.querySelector('.first').innerHTML = tab;
}
//POST data example by JSON data
fetch('/student', {
    method: 'POST',
    headers: {
        'Accept': 'application/json, text/plain, */*',
        'Content-type': 'application/json'
    },
    body: JSON.stringify({firstname: "Jan", lastname: "Kowalski", age: 24, grade: 5})
}).then(res => res.json())
.then(res => console.log(res))
//UPDATE


//DELETE