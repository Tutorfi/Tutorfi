import { lazy, Show } from 'solid-js'
import { A } from '@solidjs/router'
import logo from '../../../assets/favicon/favicon.ico'
const NavbarOrgs = lazy(() => import('./NavbarOrgs'))
const NavbarGroups = lazy(() => import('./NavbarGroups'))

function NavbarUser (props) {
  const logout = () => {
    sessionStorage.removeItem('token')
    location.reload()
  }

  return (
    <header>
      <nav class="bg-white border-b border-gray-200 dark:bg-gray-900">
        <div class="flex flex-wrap justify-between items-center mx-auto max-w-screen-xl p-4">
          <A href="/" class="flex items-center space-x-3 rtl:space-x-reverse">
            <img src={logo} class="h-8" alt="Tutorfi Logo" />
            <span class="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">
              {!props.type ? 'Dashboard' : props.type === 'group' ? 'Org/GroupName' : props.type === 'organization' ? 'OrganzationName' : 'TitleNotFound'}
            </span>
          </A>
          <div class="flex items-center space-x-6 rtl:space-x-reverse">
            <p class="text-sm  text-red-600 dark:text-white">Create</p>
            <p class="text-sm  text-gray-500 dark:text-white">Message</p>
            <p class="text-sm  text-gray-500 dark:text-white">Setting</p>
            <p class="text-sm  text-gray-500 dark:text-white">Profile</p>
            <button class="text-sm text-blue-600 dark:text-blue-500 hover:underline" onClick={logout}>Log out</button>
          </div>
        </div>

        <div class="ml-8 flex space-x-4">
          <A href="/orgs" class="text-blue-600 dark:text-blue-500 hover:underline">Show Organization</A>
          <A href="/groups" class="text-blue-600 dark:text-blue-500 hover:underline">Show Groups</A>
          <A href="#" class="text-red-600 hover:underline">Autograder Config</A>
        </div>

        <Show when={props.type === 'organization'}>
          <NavbarOrgs />
        </Show>
        <Show when={props.type === 'group'}>
          <NavbarGroups />
        </Show>
      </nav>
    </header>
  )
}

export default NavbarUser
