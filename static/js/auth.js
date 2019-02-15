$(document).ready(function() {

  $('#LogoutButton').click(function(event) {
    event.preventDefault();
    console.log('hey');

    addLoader();

    fetch( URL + '/backend/logout', {
      method: "GET",
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      }
    })
    .then(function(response) {
      return response.json();
    })
    .then(function(response) {
      stopLoader();
      if (response["Result"] == "OK") {
        // redirect
        console.log("OK");
        window.location.href = URL + "";

      } else {
        displayPopUp('Cannot logout');
      }
    });
  })

  $('#LoginSubmit').click(function(event) {
    event.preventDefault();

    data = {
      "username": $('#InputUsername').val(),
      "password": $('#InputPassword').val()
    }
    addLoader();

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
    .then(function(response) {
      stopLoader();
      if (response["Result"] == "OK") {
        // redirect
        window.location.href = URL + "/auth/hi";
      } else {
        displayPopUp('Please try different username');
      }
    });
  })

  $('#SignUpSubmit').click(function(event) {
    event.preventDefault();

    if ($('#InputUsername').val() == "") {
      displayPopUp("Username is required");
      return false;
    }

    if ($('#InputPassword').val() == "") {
      displayPopUp("Password is required");
      return false;
    }

    data = {
      "username": $('#InputUsername').val(),
      "password": $('#InputPassword').val()
    }
    addLoader();

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
    .then(function(response) {
      stopLoader();
      if (response["Result"] == "OK") {
        $('.alert').fadeIn();
      } else {
        displayPopUp('Please try different username');
      }
    });

  })

})
