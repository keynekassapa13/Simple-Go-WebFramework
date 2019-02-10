URL = 'http://' + window.location.hostname + ':' + window.location.port;

function handleError(jqXHR) {
  if (jqXHR.status === 0) {
      alert('Not connect.\n Verify Network.');
  } else if (jqXHR.status == 404) {
      alert('Requested page not found. [404]');
  } else if (jqXHR.status == 500) {
      alert('Internal Server Error [500].');
  } else if (exception === 'parsererror') {
      alert('Requested JSON parse failed.');
  } else if (exception === 'timeout') {
      alert('Time out error.');
  } else if (exception === 'abort') {
      alert('Ajax request aborted.');
  } else {
      alert('Uncaught Error.\n' + jqXHR.responseText);
  }
}

function addLoader() {
  $('#LoaderWindow').addClass('loader-window');
  $('#Loader').addClass('loader');
}

function stopLoader() {
  $('#LoaderWindow').removeClass('loader-window');
  $('#Loader').removeClass('loader');
}

function displayPopUp(content) {
  $.confirm({
    title: 'Error!',
    closeIcon: true,
    content: content,
    theme: 'supervan',
    buttons: {
      Ok: function () {}
    }
  });
}
