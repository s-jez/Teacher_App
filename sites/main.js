const url_students = "/student";
const btn_delete = document.querySelector('#delete');
const btn_update = document.querySelector('#update');

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
        <td><button type="button" id="btn_update" class="btn btn-outline-primary" data-target="#exampleModal2" data-toggle="modal">Edit</button></td>
        <td><button type="button" id="btn_delete" class="btn btn-outline-danger" data-target="#exampleModal3" data-toggle="modal">Delete</button></td>
        </tr>`;
    }
    document.querySelector('.first').innerHTML = tab;
}
$(document).ready(() => {
    $("#myForm").submit(function (event) {
      $('.info').empty();
      event.preventDefault();
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
  });
if(btn_delete) {
  btn_delete.addEventListener('click', function(event) {
    event.preventDefault();
    $('.info').empty();
    var formData = {
      id: $('#id').val(),
    };
    $.ajax({
      type: "DELETE",
      url: "/student/" + formData.id,
      data: formData,
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
}
if(btn_update) {
  btn_update.addEventListener('click', function(event) {
    event.preventDefault();
    $('.info').empty();
    var formData = {
      id: $('#id').val(),
      firstname: $('#firstname').val(),
      lastname: $('#lastname').val(),
      age: $('#age').val(),
      grade: $('#grade').val(),
    };
    $.ajax({
      type: "PUT",
      url: "/student/" + formData.id,
      data: formData,
      dataType: "JSON",
      encode: true,
    }).done(function (data) {
      console.log(data);
      getapi(url_students);
      $('#exampleModal2').hide();
      $('.close').click();
      $('.info').addClass("success");
      $('.info').append('<div class="alert alert-primary">Pomyślnie przebiegła aktualizacja studenta!</div>');
    }).fail(function (data) {
        $('.info').append('<div class="alert alert-danger">Błąd z połączeniem z serwerem!!</div>');
    });
  })
}