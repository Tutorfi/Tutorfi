import RouteGuard from '../../RouteGuard'
import NavbarUser from '../../../components/nav/user/NavbarUser'

function Organization () {

  return (
    <>
      <NavbarUser type="organization" />
      <div>
        <p>This is an organization page</p>
        <p>Show the calendar in the overview tab</p>
      </div>
    </>
  )
}

export default Organization
