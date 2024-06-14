// create account
function SignUp(fname, lname, email, password) {
    fetch("/api/account/create", {
      method: "POST",
      body: JSON.stringify({
        "Fname": fname,
        "Lname": lname,
        "Email": email,
        "Password": password
      }),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    })
    .then((response) => response.json())
    .then((json) => { console.log(json); });
}