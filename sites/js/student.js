const url_students = "/student";
const formCreateStudent = document.getElementById('formCreateStudent');
const formUpdateStudent = document.getElementById('formUpdateStudent');
const btnDelete = document.getElementById('btnDelete');
const btnUpdate = document.getElementById('btnUpdate');
let currentID;

async function getapi(url) {
  const response = await fetch(url)
  var data = await response.json();
  show(data);
}
getapi(url_students);
// Get Students
function show(data) {
  let tab = ``;
  for (let student of data) {
    tab += `
        <tr data-id='${student.id}'>
        <th scope="row">${student.id}</th>
        <td class="firstname">${student.firstname} </td>
        <td class="lastname">${student.lastname} </td>
        <td class="age">${student.age} </td>
        <td class="grade">${student.grade} </td>
        <td><button type="button" id="btnUpdate" class="btn btn-outline-primary" data-target="#exampleModal2" data-toggle="modal" onclick="getID(${student.id})">Edit</button></td>
        <td><button type="button" id="btnDelete" class="btn btn-outline-danger" data-target="#exampleModal3" data-toggle="modal" onclick="getID(${student.id})">Delete</button></td>
        </tr>`;
  }

  if (tab != null) {
    document.querySelector('.table__content').innerHTML = tab;
  }
}
function getID(id) {
  currentID = id;
}
// POST Student
formCreateStudent.addEventListener('submit', async (event) => {
  event.preventDefault();
  let studentData = {
    firstname: document.getElementById("fname").value,
    lastname: document.getElementById("lname").value,
    age: document.getElementById("fage").value,
    grade: document.getElementById("fgrade").value
  }
  let studentAge = parseInt(studentData.age);
  let studentGrade = parseInt(studentData.grade);
  // let token = JSON.parse(localStorage.getItem('token'));
  if (!studentData.firstname || !studentData.lastname || studentAge == 0 || studentGrade == 0) { alert('Please enter data to form'); return; }
  fetch(`/student`, {
    method: "POST",
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json',
      'Authorization': 'Bearer ' + localStorage.getItem('access_token')
    },
    body: JSON.stringify({
      FirstName: studentData.firstname,
      LastName: studentData.lastname,
      Age: studentAge,
      Grade: studentGrade
    }),
  }).then(async (result) => {
    const content = await result.json();
    if (result) {
      getapi(url_students);
      document.getElementById("fname").value = "";
      document.getElementById("lname").value = "";
      document.getElementById('fage').value = "";
      document.getElementById("fgrade").value = null;
      // hide modal window
      let modalWindow = document.querySelector('#exampleModal');
      modalWindow.style.display = 'none';
      let divBackground = document.querySelector('.modal-backdrop');
      divBackground.remove();
      // alert
      console.log(content);
    } else {
      alert('You not have permissions to do that action!')
    }
  }).catch((err) => {
    console.log(err)
  });
})
// PUT Student
formUpdateStudent.addEventListener('submit', (event) => {
  event.preventDefault();
  let studentData = {
    firstname: document.getElementById("firstname").value,
    lastname: document.getElementById("lastname").value,
    age: document.getElementById("age").value,
    grade: document.getElementById("grade").value
  }
  let studentAge = parseInt(studentData.age);
  let studentGrade = parseInt(studentData.grade);
  if (!studentData.firstname || !studentData.lastname || studentAge == 0 || studentGrade == 0) { alert('Please enter data to form'); return; }
  fetch(`/student/` + currentID, {
    method: 'PUT',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json',
      'Authorization': 'Bearer ' + localStorage.getItem('access_token')
    },
    body: JSON.stringify({
      FirstName: studentData.firstname,
      LastName: studentData.lastname,
      Age: studentAge,
      Grade: studentGrade
    })
  }).then(async (result) => {
    const content = await result.json();
    if (result) {
      getapi(url_students);
      document.getElementById("firstname").value = "";
      document.getElementById("lastname").value = "";
      document.getElementById('age').value = "";
      document.getElementById("grade").value = null;
      // hide modal window
      let modalWindow = document.querySelector('#exampleModal2');
      modalWindow.style.display = 'none';
      let divBackground = document.querySelector('.modal-backdrop');
      divBackground.remove();
      // alert
      console.log(content);
    } else {
      alert('You not have permissions to do that action!')
    }
  }).catch((err) => {
    console.log(err)
  });
})
// DELETE Student
btnDelete.addEventListener('click', (event) => {
  event.preventDefault();
  fetch(`/student/` + currentID, {
    method: 'DELETE',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json',
      'Authorization': 'Bearer ' + localStorage.getItem('access_token')
    },
  }).then(async (result) => {
    if (result) {
      getapi(url_students);
      // hide modal window
      let modalWindow = document.querySelector('#exampleModal3');
      modalWindow.style.display = 'none';
      let divBackground = document.querySelector('.modal-backdrop');
      divBackground.remove();
      // alert
    } else {
      alert('You not have permissions to do that action!')
    }
  }).catch((err) => {
    console.log(err)
  });
})
function logout() {
  if (confirm("Are you sure to logout from your account?")) {
    localStorage.removeItem("access_token");
    localStorage.removeItem("refresh_token");
    window.location.href = '/';
  } else {
    return
  }
}
function refresh() {
  if (confirm("Are you sure to refresh your tokens?")) {
    fetch(`/refresh`, {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + localStorage.getItem('refresh_token')
      },
    }).then(async (result) => {
      const content = await result.json();
      if (result) {
        console.log(content)
        localStorage.setItem("access_token", JSON.stringify(content.AccessToken));
        localStorage.setItem("refresh_token", JSON.stringify(content.RefreshToken));
        alert('Successfully refreshed your accesstoken and refreshtoken!')
      } else {
        alert("error!");
      }
    })
  } else {
    return
  }
}
function tokenCheck() {
  let accessToken = localStorage.getItem('access_token');
  if (!accessToken) {
    window.location.replace('/');
  }
}

