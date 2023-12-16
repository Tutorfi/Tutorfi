import { Routes, Route } from '@solidjs/router'

function Home () {
  return (
    <div>
      <h1>Home testing</h1>
      <h2>Home testing</h2>
    </div>
  )
}

function App () {
  return (
    <>
      <Routes>
        <Route path='/' element={<Home />} />
      </Routes>
    </>
  )
}

export default App
