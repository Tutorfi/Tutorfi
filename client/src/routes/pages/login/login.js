import NavbarLanding from '../../../components/nav/root/NavbarLanding'
import { register } from '../../../api_calls/authentication/register'

function Landing () {
  document.title = 'Tutorfi - Login'

  // add login page here and redirect to root if logined in 
  // add register page 

  return (
    <>
      <NavbarLanding />
      <div class="mt-16">
      </div>
      <button class="btn" onClick={() => register('fname', 'lname', '123@email.com', 'password')}>register</button>
    </>
  )
}

export default Landing
