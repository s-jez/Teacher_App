const url_students = "/student";
const formCreateStudent = document.getElementById('formCreateStudent');
const formUpdateStudent = document.getElementById('formUpdateStudent');
const formRegisterUser = document.getElementById('formRegisterUser');
const formLoginUser = document.getElementById('formLoginUser');
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
  try {
    event.preventDefault();
    let studentData = {
      firstname: document.getElementById("fname").value,
      lastname: document.getElementById("lname").value,
      age: document.getElementById("fage").value,
      grade: document.getElementById("fgrade").value
    }
    console.log(studentData.age)
    let studentAge = parseInt(studentData.age);
    let studentGrade = parseInt(studentData.grade);
    // let token = JSON.parse(localStorage.getItem('token'));
    if (!studentData.firstname || !studentData.lastname || studentAge == 0 || studentGrade == 0) { alert('Please enter data to form'); return; }
    fetch(`/student`, {
      method: "POST",
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        'Authorization': 'Bearer' + localStorage.getItem('access_token')
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
        document.getElementById("firstname").value = "";
        document.getElementById("lastname").value = "";
        document.getElementById('age').value = "";
        document.getElementById("grade").value = null;
        // hide modal window
        let modalWindow = document.querySelector('#exampleModal');
        modalWindow.style.display = 'none';
        let divBackground = document.querySelector('.modal-backdrop');
        divBackground.remove();
        // alert
        console.log(content);
      }
    })
  }
  catch (err) {
    alert("Invalid connection to the server!!!");
    document.querySelector('.info').innerHTML += `<div class="alert alert-danger>Invalid connection to the server</div>`;
    throw (err)
  }
})
// PUT Student
formUpdateStudent.addEventListener('submit', (event) => {
  try {
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
        'Authorization': 'Bearer' + localStorage.getItem('access_token')
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
      }
    })
  }
  catch (err) {
    alert("Invalid connection to the server!!!");
    document.querySelector('.info').innerHTML += `<div class="alert alert-danger>Invalid connection to the server</div>`;
    throw (err)
  }
})
// DELETE Student
btnDelete.addEventListener('click', (event) => {
  event.preventDefault();
  try {
    fetch(`/student/` + currentID, {
      method: 'DELETE',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        'Authorization': 'Bearer' + localStorage.getItem('access_token')
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
      }
    })
  }
  catch (err) {
    alert("Invalid connection to the server!!!");
    document.querySelector('.info').innerHTML += `<div class="alert alert-danger>Invalid connection to the server</div>`;
    throw (err)
  }
})
// POST UserRegister
formRegisterUser.addEventListener('submit', async (event) => {
  try {
    event.preventDefault();
    let userData = {
      UserName: document.getElementById("fusername").value,
      Password: document.getElementById("fpassword").value,
      Email: document.getElementById('femail').value,
      RoleID: document.getElementById('frole').value
    }
    let userRole = parseInt(userData.RoleID);
    if (!userData.UserName || !userData.Password || !userData.Email || userRole == 0) { alert('Please enter data to form'); return; }
    const validateEmail = (email) => {
      return String(email)
        .toLowerCase()
        .match(
          /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
        );
    }
    if (!validateEmail(userData.Email)) {
      alert('Email incorrect!');
      return
    }
    fetch(`/register`, {
      method: "POST",
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        'Authorization': 'Bearer'
      },
      body: JSON.stringify({
        "UserName": userData.UserName,
        "Password": userData.Password,
        "Email": userData.Email,
        "RoleID": userRole
      })
    }).then(async (result) => {
      const content = await result.json();
      if (result) {
        getapi(url_students);
        document.getElementById("fusername").value = "";
        document.getElementById("fpassword").value = "";
        document.getElementById('femail').value = "";
        document.getElementById('frole').value = null;
        // hide modal window
        let modalWindow = document.querySelector('#exampleModal4');
        modalWindow.style.display = 'none';
        let divBackground = document.querySelector('.modal-backdrop');
        divBackground.remove();
        // alert
        console.log('Created user successfully!');
        console.log(content);
      } else {
        alert('Error deleting user');
      }
    }).catch(e => {
      console.log(e);
    });
  }
  catch (err) {
    alert("Invalid connection to the server!!!");
    document.querySelector('.info').innerHTML += `<div class="alert alert-danger>Invalid connection to the server</div>`;
    throw (err)
  }
})
// POST UserLogin
formLoginUser.addEventListener('submit', async (event) => {
  try {
    event.preventDefault();
    let userData = {
      UserName: document.getElementById("username").value,
      Password: document.getElementById("password").value,
      Email: document.getElementById('email').value
    }
    if (!userData.UserName || !userData.Password || !userData.Email) { alert('Please enter data to form'); return; }
    fetch(`/login`, {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        'Authorization': 'Bearer'
      },
      body: JSON.stringify(userData)
    }).then(async (result) => {
      const content = await result.json();
      if (result) {
        getapi(url_students);
        document.getElementById("username").value = "";
        document.getElementById("password").value = "";
        document.getElementById('email').value = "";
        // hide modal window
        let modalWindow = document.querySelector('#exampleModal5');
        modalWindow.style.display = 'none';
        let divBackground = document.querySelector('.modal-backdrop');
        divBackground.remove();
        // alert
        console.log('User logged in successfully!');
        localStorage.setItem("access_token", JSON.stringify(content.AccessToken));
        localStorage.setItem("refresh_token", JSON.stringify(content.RefreshToken));
      } else {
        alert('Error in logging user!');
      }
    })
  }
  catch (err) {
    alert("Invalid connection to the server!!!");
    document.querySelector('.info').innerHTML += `<div class="alert alert-danger>Invalid connection to the server</div>`;
    throw (err)
  }
})