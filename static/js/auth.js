$(document).ready(function() {


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
    }).then((res) => {
      console.log(res)
    }).catch((err) => {
      console.log(err)
    })

  })

})
