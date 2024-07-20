import NavbarLanding from '../../../components/nav/root/NavbarLanding'
import { register } from '../../../api_calls/authentication/register'

function Landing () {
  document.title = 'Tutorfi'

    // Change to click button to register

  return (
    <>
      <NavbarLanding />
      <div class="mt-16">
        <p>This is a landing page</p>
      </div>
      <button class="btn" onClick={() => register('fname', 'lname', '123@email.com', 'password')}>register</button>
    </>
  )
}

export default Landing
