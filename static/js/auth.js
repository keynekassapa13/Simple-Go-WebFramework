$(document).ready(function() {

  $('#LoginSubmit').click(function(event) {
    event.preventDefault();

    data = {
      "username": $('#InputUsername').val(),
      "password": $('#InputPassword').val()
    }

    fetch( URL + '/backend/login', {
      method: "POST",
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    })
    .then(function(response) {
      return response.json();
    })
    .then(function(result) {
      console.log(result);
      if (result == "OK") {

      } else {

      }
    });
  })

  $('#TestSubmit').click(function(event) {
    event.preventDefault();

    fetch( URL + '/backend/isLoggedIn', {
      method: "GET"
    })
    .then(function(response) {
      return response.json();
    })
    .then(function(result) {
      console.log(result);
      if (result == "OK") {

      } else {

      }
    });

  })

  $('#SignUpSubmit').click(function(event) {
    event.preventDefault();

    data = {
      "username": $('#InputUsername').val(),
      "password": $('#InputPassword').val()
    }

    fetch( URL + '/backend/addUser', {
      method: "POST",
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    })
    .then(function(response) {
      return response.json();
    })
    .then(function(result) {
      console.log(result);
      if (result == "OK") {

      } else {

      }
    });

  })

})
