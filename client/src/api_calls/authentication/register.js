// create account
export function register (fname, lname, email, password) {
  fetch('/api/account/create', {
    method: 'POST',
    body: JSON.stringify({
      fname,
      lname,
      email,
      password
    }),
    headers: {
      'Content-type': 'application/json; charset=UTF-8'
    }
  })
    .then((response) => response.json())
    .then((json) => { console.log(json) })
}
