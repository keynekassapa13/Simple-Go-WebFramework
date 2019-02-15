$(document).ready(function() {
  $('#LogoutButton').click(function(event) {
    event.preventDefault();

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
        window.location.href = URL + "/";
      } else {

      }
    });

  })
})
