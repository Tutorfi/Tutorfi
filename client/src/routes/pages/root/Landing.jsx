import NavbarLanding from '../../../components/nav/root/NavbarLanding'

function Landing () {
  document.title = 'Tutorfi'

  return (
    <>
      <NavbarLanding />
      <div class="mt-16">
        <p>This is a landing page</p>
      </div>
    </>
  )
}

export default Landing
