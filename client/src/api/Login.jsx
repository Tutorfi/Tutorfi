// verify
function Login(email, password) {
    fetch("/api/account/verify", {
      method: "POST",
      body: JSON.stringify({
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