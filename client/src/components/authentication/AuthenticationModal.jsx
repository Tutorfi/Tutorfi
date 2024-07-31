import { createSignal, createContext, onMount } from 'solid-js'
import { login } from '../../api_calls/authentication/login'
import { register } from '../../api_calls/authentication/register'
// import { useNavigate } from '@solidjs/router'

export default function AuthenticationModal (props) {
  // const navigate = useNavigate()
  const context = createContext()

  const [loginUsername, setLoginUsername] = createSignal('')
  const [loginPassword, setLoginPassword] = createSignal('')
  const [loginRemember, setLoginRemember] = createSignal(false)

  const [signupUsername, setSignupUsername] = createSignal('')
  const [signupPassword, setSignupPassword] = createSignal('')
  const [signupRepeatPassword, setSignupRepeatPassword] = createSignal('')

  const [loginVisible, setLoginVisible] = createSignal(false)
  const [signupVisible, setSignupVisible] = createSignal(false)

  // Change this to redirect to login page
    const handlelogin = async () => {
      const res = await login('123@email.com', 'password')
      if (res.ok) {
        location.reload()
      }
    }

    const handleSignup = async () => {
      const res = await register("","","","")
      if (res.ok) {
        location.reload()
      }
    }

  onMount(() => {
    props.authRef?.({
      loginVisible: () => {
        setLoginVisible(true)
      },
      signupVisible: () => {
        setSignupVisible(true)
      }
    })
  })

  // toggle modal visibility
  const toggleModal = () => {
    setLoginVisible(false)
    setSignupVisible(false)
  }

  // entire modal including login and signup html form
  return (
    <div
      id="authentication-modal"
      class={'absolute overflow-y-auto overflow-x-hidden z-10 top-0 left-0 right-0 flex justify-center items-center w-full h-full' + ((loginVisible() || signupVisible()) ? '' : ' hidden')}
    >
      <div
        onClick={toggleModal()}
        class="fixed w-full h-full bg-black bg-opacity-50"
      ></div>
      <div class="fixed w-full h-full z-50">
        <div class="relative p-4 w-full h-full flex justify-center items-center">
          {/* Login */}
          <form
            id="login-form"
            class={'w-full md:max-w-md bg-white rounded-lg shadow dark:bg-gray-700' + (loginVisible() ? '' : ' hidden')}
            onSubmit={(e) => {
              e.preventDefault()
              // handleLogIn()
            }}
          >
            <div class="flex items-center justify-between p-4 md:p-5 border-b rounded-t dark:border-gray-600">
              <h3 class="text-xl font-semibold text-gray-900 dark:text-white">
                Login to Tutorfi
              </h3>
              <button
                onClick={toggleModal}
                type="button"
                class="end-2.5 text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white"
              >
                <svg
                  class="w-3 h-3"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 14 14"
                >
                  <path
                    stroke="currentColor"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth="2"
                    d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"
                  />
                </svg>
                <span class="sr-only">Close modal</span>
              </button>
            </div>
            <div class="p-4 md:p-5">
              <div class="mx-auto space-y-4">
                <div class="mb-5">
                  <label
                    for="loginUsername"
                    class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                  >
                    Your username
                  </label>
                  <input
                    type="username"
                    id="loginUsername"
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    placeholder="username"
                    required
                    value={loginUsername()}
                    onChange={(e) => setLoginUsername(e.target.value)}
                  />
                </div>
                <div class="mb-5">
                  <label
                    for="loginPassword"
                    class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                  >
                    Your password
                  </label>
                  <input
                    type="password"
                    id="loginPassword"
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    required
                    value={loginPassword()}
                    onChange={(e) => setLoginPassword(e.target.value)}
                  />
                </div>
                <div class="flex items-start mb-5">
                  <div class="flex items-center h-5">
                    <input
                      id="remember"
                      type="checkbox"
                      class="w-4 h-4 border border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-blue-300 dark:bg-gray-700 dark:border-gray-600 dark:focus:ring-blue-600 dark:ring-offset-gray-800 dark:focus:ring-offset-gray-800"
                      value={loginRemember ? 'on' : undefined}
                      onChange={(e) => {
                        setLoginRemember(e.target.checked)
                        context.setData({ remember: e.target.checked })
                      }}
                    />
                  </div>
                  <label
                    for="remember"
                    class="ms-2 text-sm font-medium text-gray-900 dark:text-gray-300"
                  >
                    Remember me
                  </label>
                </div>
                <button
                  type="submit"
                  class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                >
                  Login to your account
                </button>
                <div class="text-sm font-medium text-gray-500 dark:text-gray-300">
                  Not registered?{' '}
                  <button
                    onClick={() => {
                      setLoginVisible(false)
                      setSignupVisible(true)
                    }}
                    type="button"
                    class="text-blue-700 hover:underline dark:text-blue-500"
                  >
                    Create account
                  </button>
                </div>
              </div>
            </div>
          </form>

          {/* Sign Up */}
          <form
            id="signup-form"
            class={'w-full md:max-w-md bg-white rounded-lg shadow dark:bg-gray-700' + (signupVisible() ? '' : ' hidden')}
            onSubmit={(e) => {
              e.preventDefault()
              // handleSignUp()
            }}
          >
            <div class="flex items-center justify-between p-4 md:p-5 border-b rounded-t dark:border-gray-600">
              <h3 class="text-xl font-semibold text-gray-900 dark:text-white">
                Sign Up to Tutorfi
              </h3>
              <button
                onClick={toggleModal}
                type="button"
                class="end-2.5 text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white"
              >
                <svg
                  class="w-3 h-3"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 14 14"
                >
                  <path
                    stroke="currentColor"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth="2"
                    d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"
                  />
                </svg>
                <span class="sr-only">Close modal</span>
              </button>
            </div>
            <div class="p-4 md:p-5">
              <div class="mx-auto space-y-4">
                <div class="relative z-0 w-full mb-5 group">
                  <input
                    type="username"
                    name="floating_username"
                    id="floating_username"
                    class="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-blue-500 focus:outline-none focus:ring-0 focus:border-blue-600 peer"
                    placeholder=" "
                    required
                    value={signupUsername()}
                    onChange={(e) => setSignupUsername(e.target.value)}
                  />
                  <label
                    for="floating_username"
                    class="peer-focus:font-medium absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto peer-focus:text-blue-600 peer-focus:dark:text-blue-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6"
                  >
                    Username
                  </label>
                </div>
                <div class="relative z-0 w-full mb-5 group">
                  <input
                    type="password"
                    name="floating_password"
                    id="floating_password"
                    class="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-blue-500 focus:outline-none focus:ring-0 focus:border-blue-600 peer"
                    placeholder=" "
                    required
                    value={signupPassword()}
                    onChange={(e) => setSignupPassword(e.target.value)}
                  />
                  <label
                    for="floating_password"
                    class="peer-focus:font-medium absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 rtl:peer-focus:translate-x-1/4 peer-focus:text-blue-600 peer-focus:dark:text-blue-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6"
                  >
                    Password
                  </label>
                </div>
                <div class="relative z-0 w-full mb-5 group">
                  <input
                    type="password"
                    name="repeat_password"
                    id="floating_repeat_password"
                    class="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-blue-500 focus:outline-none focus:ring-0 focus:border-blue-600 peer"
                    placeholder=" "
                    required
                    value={signupRepeatPassword()}
                    onChange={(e) => setSignupRepeatPassword(e.target.value)}
                  />
                  <label
                    for="floating_repeat_password"
                    class="peer-focus:font-medium absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 rtl:peer-focus:translate-x-1/4 peer-focus:text-blue-600 peer-focus:dark:text-blue-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6"
                  >
                    Confirm password
                  </label>
                </div>
                <div class="flex items-start mb-5">
                  <div class="flex items-center h-5">
                    <input
                      id="terms"
                      type="checkbox"
                      value=""
                      class="w-4 h-4 border border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-blue-300 dark:bg-gray-700 dark:border-gray-600 dark:focus:ring-blue-600 dark:ring-offset-gray-800 dark:focus:ring-offset-gray-800"
                      required
                    />
                  </div>
                  <label
                    for="terms"
                    class="ms-2 text-sm font-medium text-gray-900 dark:text-gray-300"
                  >
                    I agree with the{' '}
                    <a
                      href="#"
                      class="text-blue-600 hover:underline dark:text-blue-500"
                    >
                      terms and conditions
                    </a>
                  </label>
                </div>
                <button
                  type="submit"
                  class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                >
                  Create your account
                </button>
                <div class="text-sm font-medium text-gray-500 dark:text-gray-300">
                  Already have an account?{' '}
                  <button
                    onClick={() => {
                      setLoginVisible(true)
                      setSignupVisible(false)
                    }} type="button"
                    class="text-blue-700 hover:underline dark:text-blue-500"
                  >
                    Login to your account
                  </button>
                </div>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
  )
}
