const url_students = "/student";

async function getapi(url) {
  const response = await fetch(url)
  var data = await response.json();
  show(data);
}
getapi(url_students);
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
        <td><button type="button" id="btn_update" class="btn btn-outline-primary" data-id="${student.id}" data-target="#exampleModal2" data-toggle="modal"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-pencil" viewBox="0 0 16 16">
        <path d="M12.146.146a.5.5 0 0 1 .708 0l3 3a.5.5 0 0 1 0 .708l-10 10a.5.5 0 0 1-.168.11l-5 2a.5.5 0 0 1-.65-.65l2-5a.5.5 0 0 1 .11-.168l10-10zM11.207 2.5 13.5 4.793 14.793 3.5 12.5 1.207 11.207 2.5zm1.586 3L10.5 3.207 4 9.707V10h.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.5h.293l6.5-6.5zm-9.761 5.175-.106.106-1.528 3.821 3.821-1.528.106-.106A.5.5 0 0 1 5 12.5V12h-.5a.5.5 0 0 1-.5-.5V11h-.5a.5.5 0 0 1-.468-.325z"/>
      </svg>Edit</button></td>
        <td><button type="button" id="btn_delete" class="btn btn-outline-danger" data-id="${student.id}" data-target="#exampleModal3" data-toggle="modal"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-trash" viewBox="0 0 16 16">
        <path d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"/>
        <path fill-rule="evenodd" d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z"/>
      </svg>Delete</button></td>
        </tr>`;
  }
  document.querySelector('.first').innerHTML = tab;
}
$(document).on('submit', '#myForm', function (event) {
  event.preventDefault();
  var formData = {
    firstname: $("#firstname").val(),
    lastname: $("#lastname").val(),
    age: $("#age").val(),
    grade: $("#grade").val(),
  };
  if(formData.firstname == "" || formData.lastname == "" || formData.age == 0 || formData.grade == 0) {
      $('.info').empty();
      $('.info').append('<div class="alert alert-danger">Please enter student data to form</div>');
  } else {
    $.ajax({
      type: "POST",
      url: "/student",
      data: formData,
      dataType: "JSON",
      encode: true,
    }).done(function () {
      $('.info').empty();
      getapi(url_students);
      $('input#firstname').val('');
      $('input#lastname').val('');
      $('input#age').val('');
      $('input#grade').val('');
      $('#exampleModal').hide();
      $('.close').click();
      $('.info').addClass("success");
      $('.info').append('<div class="alert alert-success">You have succesffully added a student!</div>');
    }).fail(function() {
      $('.info').append('<div class="alert alert-danger">Invalid connection to the server!!!</div>');
    });
  }
})
$(document).on('click', '#btn_delete', function (event) {
  event.preventDefault();
  var id = $(this).attr('data-id');
  $(document).on('click', '#delete-modal', function (event) {
    event.preventDefault();
    $('.info').empty();
    $.ajax({
      type: "DELETE",
      url: "/student/" + id,
      dataType: "JSON",
      encode: true,
    }).done(function () {
      $('.info').empty();
      getapi(url_students);
      $('#exampleModal3').hide();
      $('.close').click();
      $('.info').addClass("success");
      $('.info').append('<div class="alert alert-danger">You have successfully deleted a student!</div>');
    }).fail(function () {
      $('.info').append('<div class="alert alert-danger">Invalid connection to the server!!!</div>');
    });
  })
})
$(document).on('click', '#btn_update', function (event) {
  event.preventDefault();
  var id = $(this).attr('data-id');
  $(document).on('click', '#update-modal', function (event) {
    event.preventDefault();
    $('.info').empty();
    var formData = {
      "firstname": $('#myForm2').find('#firstname').val(),
      "lastname": $('#myForm2').find('#lastname').val(),
      "age": $('#myForm2').find('#age').val(),
      "grade": $('#myForm2').find('#grade').val(),
    };
    if(formData.firstname == "" || formData.lastname == "" || formData.age == 0 || formData.grade == 0) {
      $('.info').empty();
      $('.info').append('<div class="alert alert-danger">Please enter student data to form</div>');
    } else {
      $.ajax({
        type: "PUT",
        url: "/student/" + id,
        dataType: "JSON",
        encode: true,
        data: formData,
      }).done(function (data) {
        $('.info').empty();
        getapi(url_students);
        $('input#firstname').val('');
        $('input#lastname').val('');
        $('input#age').val('');
        $('input#grade').val('');
        $('#exampleModal2').hide();
        $('.close').click();
        $('.info').addClass("success");
        $('.info').append('<div class="alert alert-primary">You have successfully updated a student!</div>');
      }).fail(function () {
        $('.info').append('<div class="alert alert-danger">Invalid connection to the server!!!</div>');
      }).catch(e => {
        console.log(e);
      });
    }
  })
})