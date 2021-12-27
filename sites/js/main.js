const url_students = "/student";

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
        <td>${student.firstname} </td>
        <td>${student.lastname} </td>
        <td>${student.age} </td>
        <td>${student.grade} </td>
        <td><button type="button" id="btn_update" class="btn btn-outline-primary" data-id="${student.id}" data-target="#exampleModal2" data-toggle="modal">Edit</button></td>
        <td><button type="button" id="btn_delete" class="btn btn-outline-danger" data-id="${student.id}" data-target="#exampleModal3" data-toggle="modal">Delete</button></td>
        </tr>`;
  }
  document.querySelector('.first').innerHTML = tab;
}
// // Add Student
// $(document).on('submit', '#formCreate', function (event) {
//   event.preventDefault();
//   var formData = {
//     firstname: $("#formRegister__firstname").val(),
//     lastname: $("#formRegister__lastname").val(),
//     age: $("#formRegister__age").val(),
//     grade: $("#formRegister__grade").val(),
//   };
//   if (formData.firstname == "" || formData.lastname == "" || formData.age == 0 || formData.grade == 0) {
//     $('.info').empty();
//     $('.info').append('<div class="alert alert-danger">Please enter student data to form</div>');
//   } else {
//     $.ajax({
//       type: "POST",
//       url: "/student",
//       data: formData,
//       dataType: "JSON",
//       encode: true,
//       headers: {
//         "Authorization": "Bearer"
//       }
//     }).done(function () {
//       $('.info').empty();
//       getapi(url_students);
//       $('input#formRegister__firstname').val('');
//       $('input#formRegister__lastname').val('');
//       $('input#formRegister__age').val('');
//       $('input#formRegister__grade').val('');
//       $('#exampleModal').hide();
//       $('.close').click();
//       $('.info').addClass("success");
//       $('.info').append('<div class="alert alert-success">You have succesffully added a student!</div>');
//     }).fail(function () {
//       $('.info').append('<div class="alert alert-danger">Invalid connection to the server!!!</div>');
//     }).catch(e => {
//       console.log(e)
//     })
//   }
// })
// // Delete Student
// $(document).on('click', '#btn_delete', function (event) {
//   event.preventDefault();
//   var id = $(this).attr('data-id');
//   $(document).on('click', '#delete-modal', function (event) {
//     event.preventDefault();
//     $('.info').empty();
//     $.ajax({
//       type: "DELETE",
//       url: "/student/" + id,
//       dataType: "JSON",
//       encode: true,
//       headers: {
//         "Authorization": "Bearer"
//       }
//     }).done(function () {
//       $('.info').empty();
//       getapi(url_students);
//       $('#exampleModal3').hide();
//       $('.close').click();
//       $('.info').addClass("success");
//       $('.info').append('<div class="alert alert-danger">You have successfully deleted a student!</div>');
//     }).fail(function () {
//       $('.info').append('<div class="alert alert-danger">Invalid connection to the server!!!</div>');
//     }).catch(e => {
//       console.log(e)
//     })
//   })
// })
// // Update Student
// $(document).on('click', '#btn_update', function (event) {
//   event.preventDefault();
//   var id = $(this).attr('data-id');
//   $(document).on('click', '#update-modal', function (event) {
//     event.preventDefault();
//     $('.info').empty();
//     var formData = {
//       "firstname": $('#myForm2').find('#formUpdate__firstname').val(),
//       "lastname": $('#myForm2').find('#formUpdate__lastname').val(),
//       "age": $('#myForm2').find('#formUpdate__age').val(),
//       "grade": $('#myForm2').find('#formUpdate__grade').val(),
//     };
//     firstname.value = $('tbody.first tr').attr('data-id');
//     if (formData.firstname == "" || formData.lastname == "" || formData.age == 0 || formData.grade == 0) {
//       $('.info').empty();
//       $('.info').append('<div class="alert alert-danger">Please enter student data to form</div>');
//     } else {
//       $.ajax({
//         type: "PUT",
//         url: "/student/" + id,
//         dataType: "JSON",
//         encode: true,
//         headers: {
//           "Authorization": "Bearer"
//         },
//         data: formData,
//       }).done(function (data) {
//         $('.info').empty();
//         getapi(url_students);
//         $('input#formUpdate__firstname').val('');
//         $('input#formUpdate__lastname').val('');
//         $('input#formUpdate__age').val('');
//         $('input#formUpdate__grade').val('');
//         $('#exampleModal2').hide();
//         $('.close').click();
//         $('.info').addClass("success");
//         $('.info').append('<div class="alert alert-primary">You have successfully updated a student!</div>');
//       }).fail(function () {
//         $('.info').append('<div class="alert alert-danger">Invalid connection to the server!!!</div>');
//       }).catch(e => {
//         console.log(e);
//       });
//     }
//   })
// })
// $(document).on('click', 'btn__register', function (e) {
//   e.preventDefault();
//   $(document).on('click', 'btn-register', function (e) {
//     let user = {
//       "username": $('form-register').find('formRegister__username').val(),
//       "password": $('form-register').find('formRegister__password').val(),
//       "email": $('form-register').find('formRegister__email').val(),
//       "role": $('form-register').find('formRegister__role').val(),
//     }
//     $.ajax({
//       type: 'POST',
//       url: '/register',
//       dataType: "json",
//       data: JSON.stringify({ user }),
//       encode: true,
//       headers: {
//         "Authorization": "Bearer"
//       },
//       success: function () {

//       },
//       error: function () {

//       }
//     }).catch(e => {
//       console.log(e)
//     })
//   })
// })
// let registerUser = {
//   username: document.getElementById('fusername').value(),
//   password: document.getElementById('fpassword').value(),
//   email: document.getElementById('fpassword').value(),
//   role: document.getElementById('frole').value(),
// };