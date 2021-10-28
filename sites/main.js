const url_students = "/student";

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
      <td><button type="button" class="btn btn-outline-danger">Delete</button></td>
        </tr>`;
    }
    document.querySelector('.first').innerHTML = tab;
}