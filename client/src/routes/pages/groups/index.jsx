import RouteGuard from '../../RouteGuard'
import NavbarUser from '../../../components/nav/user/NavbarUser'

function Group () {
  RouteGuard()

  return (
    <>
      <NavbarUser type="group" />
      <div>
        <p>This is a group page</p>
      </div>
    </>
  )
}

export default Group
