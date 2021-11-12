const url_students = "/student";

// GET data
async function getapi(url) {
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
        <td><button type="button" id="btn_update" class="btn btn-outline-primary" data-id="${student.id}" data-target="#exampleModal2" data-toggle="modal">Edit</button></td>
        <td><button type="button" id="btn_delete" class="btn btn-outline-danger" data-id="${student.id}" data-target="#exampleModal3" data-toggle="modal">Delete</button></td>
        </tr>`;
  }
  document.querySelector('.first').innerHTML = tab;
}
$(document).on('submit', '#myForm', function (event) {
  event.preventDefault();
  $('.info').empty();
  var formData = {
    firstname: $("#firstname").val(),
    lastname: $("#lastname").val(),
    age: $("#age").val(),
    grade: $("#grade").val(),
  };
  $.ajax({
    type: "POST",
    url: "/student",
    data: formData,
    dataType: "JSON",
    encode: true,
  }).done(function (data) {
    console.log(data);
    getapi(url_students);
    $('#exampleModal').hide();
    $('.close').click();
    $('.info').addClass("success");
    $('.info').append('<div class="alert alert-success">Pomyślnie dodałeś studenta!</div>');
  }).fail(function (data) {
    $('.info').append('<div class="alert alert-danger">Błąd z połączeniem z serwerem!!</div>');
  });
})
$(document).on('click', '#btn_delete', function (event) {
  event.preventDefault();
  var id = $(this).attr('data-id');
  $('.info').empty();
  $(document).on('click', '#delete-modal', function (event) {
    event.preventDefault();
    $('.info').empty();
    $.ajax({
      type: "DELETE",
      url: "/student/" + id,
      dataType: "JSON",
      encode: true,
    }).done(function (data) {
      console.log(data);
      getapi(url_students);
      $('#exampleModal3').hide();
      $('.close').click();
      $('.info').addClass("success");
      $('.info').append('<div class="alert alert-danger">Pomyślnie usunales studenta!</div>');
    }).fail(function (data) {
      $('.info').append('<div class="alert alert-danger">Błąd z połączeniem z serwerem!!</div>');
    });
  })
})
$(document).on('click', '#btn_update', function (event) {
  event.preventDefault();
  var id = $(this).attr('data-id');
  $('.info').empty();
  $(document).on('click', '#update-modal', function (event) {
    event.preventDefault();
    $('.info').empty();
    var formData = {
      firstname: $('#myForm2').find('#firstname').val(),
      lastname: $('#myForm2').find('#lastname').val(),
      age: $('#myForm2').find('#age').val(),
      grade: $('#myForm2').find('#grade').val(),
    };
    $.ajax({
      type: "PUT",
      url: "/student/" + id,
      data: JSON.stringify({
        firstname: formData.firstname,
        lastname: formData.lastname,
        age: formData.age,
        grade: formData.grade,
      }),
      dataType: "JSON",
      encode: true,
    }).done(function (data) {
      console.log(data);
      getapi(url_students);
      $('#exampleModal2').hide();
      $('.close').click();
      $('.info').addClass("success");
      $('.info').append('<div class="alert alert-primary">Pomyślnie przebiegła aktualizacja studenta!</div>');
    }).fail(function () {
      $('.info').append('<div class="alert alert-danger">Błąd z połączeniem z serwerem!!</div>');
    }).catch(e => {
      console.log(e);
    });
  })
})